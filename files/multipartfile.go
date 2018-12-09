package files

import (
	"errors"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/url"
	"path"
	"strings"
)

const (
	multipartFormdataType = "multipart/form-data"
	multipartMixedType    = "multipart/mixed"

	applicationDirectory = "application/x-directory"
	applicationSymlink   = "application/symlink"
	applicationFile      = "application/octet-stream"

	contentTypeHeader = "Content-Type"
)

var ErrPartOutsideParent = errors.New("file outside parent dir")
var ErrPartInChildTree = errors.New("file in child tree")

// MultipartFile implements Node, and is created from a `multipart.Part`.
//
// Note: iterating entries can be done only once and must be done in order,
// meaning that when iterator returns a directory, you MUST read all it's
// children before calling Next again
type MultipartFile struct {
	Node

	Part      *multipart.Part
	Reader    PartReader
	Mediatype string
}

func NewFileFromPartReader(reader *multipart.Reader, mediatype string) (Directory, error) {
	if !isDirectory(mediatype) {
		return nil, ErrNotDirectory
	}

	f := &MultipartFile{
		Reader:    &peekReader{r: reader},
		Mediatype: mediatype,
	}

	return f, nil
}

func newFileFromPart(parent string, part *multipart.Part, reader PartReader) (string, Node, error) {
	f := &MultipartFile{
		Part:   part,
		Reader: reader,
	}

	dir, base := path.Split(f.fileName())
	dir = path.Clean(dir)
	parent = path.Clean(parent)
	if dir == "." {
		dir = ""
	}
	if parent == "." {
		parent = ""
	}

	if dir != parent {
		if strings.HasPrefix(dir, parent) {
			return "", nil, ErrPartInChildTree
		}
		return "", nil, ErrPartOutsideParent
	}

	contentType := part.Header.Get(contentTypeHeader)
	switch contentType {
	case applicationSymlink:
		out, err := ioutil.ReadAll(part)
		if err != nil {
			return "", nil, err
		}

		return base, NewLinkFile(string(out), nil), nil
	case "": // default to application/octet-stream
		fallthrough
	case applicationFile:
		return base, &ReaderFile{
			reader:  part,
			abspath: part.Header.Get("abspath"),
		}, nil
	}

	var err error
	f.Mediatype, _, err = mime.ParseMediaType(contentType)
	if err != nil {
		return "", nil, err
	}

	if !isDirectory(f.Mediatype) {
		return base, &ReaderFile{
			reader:  part,
			abspath: part.Header.Get("abspath"),
		}, nil
	}

	return base, f, nil
}

func isDirectory(mediatype string) bool {
	return mediatype == multipartFormdataType || mediatype == applicationDirectory
}

type multipartIterator struct {
	f *MultipartFile

	curFile Node
	curName string
	err     error
}

func (it *multipartIterator) Name() string {
	return it.curName
}

func (it *multipartIterator) Node() Node {
	return it.curFile
}

func (it *multipartIterator) Next() bool {
	if it.f.Reader == nil {
		return false
	}
	var part *multipart.Part
	for {
		var err error
		part, err = it.f.Reader.NextPart()
		if err != nil {
			if err == io.EOF {
				return false
			}
			it.err = err
			return false
		}

		name, cf, err := newFileFromPart(it.f.fileName(), part, it.f.Reader)
		if err == ErrPartOutsideParent {
			break
		}
		if err != ErrPartInChildTree {
			it.curFile = cf
			it.curName = name
			it.err = err
			return err == nil
		}
	}

	// we read too much, try to fix this
	pr, ok := it.f.Reader.(*peekReader)
	if !ok {
		it.err = errors.New("cannot undo NextPart")
		return false
	}

	it.err = pr.put(part)
	return false
}

func (it *multipartIterator) Err() error {
	return it.err
}

func (f *MultipartFile) Entries() DirIterator {
	return &multipartIterator{f: f}
}

func (f *MultipartFile) fileName() string {
	if f == nil || f.Part == nil {
		return ""
	}

	filename, err := url.QueryUnescape(f.Part.FileName())
	if err != nil {
		// if there is a unescape error, just treat the name as unescaped
		return f.Part.FileName()
	}
	return filename
}

func (f *MultipartFile) Close() error {
	if f.Part != nil {
		return f.Part.Close()
	}
	return nil
}

func (f *MultipartFile) Size() (int64, error) {
	return 0, ErrNotSupported
}

type PartReader interface {
	NextPart() (*multipart.Part, error)
}

type peekReader struct {
	r    PartReader
	next *multipart.Part
}

func (pr *peekReader) NextPart() (*multipart.Part, error) {
	if pr.next != nil {
		p := pr.next
		pr.next = nil
		return p, nil
	}

	if pr.r == nil {
		return nil, io.EOF
	}

	p, err := pr.r.NextPart()
	if err == io.EOF {
		pr.r = nil
	}
	return p, err
}

func (pr *peekReader) put(p *multipart.Part) error {
	if pr.next != nil {
		return errors.New("cannot put multiple parts")
	}
	pr.next = p
	return nil
}

var _ Directory = &MultipartFile{}
