// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.10.1
// source: file.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UnixfsData_DataType int32

const (
	UnixfsData_Raw       UnixfsData_DataType = 0
	UnixfsData_Directory UnixfsData_DataType = 1
	UnixfsData_File      UnixfsData_DataType = 2
	UnixfsData_Metadata  UnixfsData_DataType = 3
	UnixfsData_Symlink   UnixfsData_DataType = 4
	UnixfsData_HAMTShard UnixfsData_DataType = 5
)

// Enum value maps for UnixfsData_DataType.
var (
	UnixfsData_DataType_name = map[int32]string{
		0: "Raw",
		1: "Directory",
		2: "File",
		3: "Metadata",
		4: "Symlink",
		5: "HAMTShard",
	}
	UnixfsData_DataType_value = map[string]int32{
		"Raw":       0,
		"Directory": 1,
		"File":      2,
		"Metadata":  3,
		"Symlink":   4,
		"HAMTShard": 5,
	}
)

func (x UnixfsData_DataType) Enum() *UnixfsData_DataType {
	p := new(UnixfsData_DataType)
	*p = x
	return p
}

func (x UnixfsData_DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UnixfsData_DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_file_proto_enumTypes[0].Descriptor()
}

func (UnixfsData_DataType) Type() protoreflect.EnumType {
	return &file_file_proto_enumTypes[0]
}

func (x UnixfsData_DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *UnixfsData_DataType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = UnixfsData_DataType(num)
	return nil
}

// Deprecated: Use UnixfsData_DataType.Descriptor instead.
func (UnixfsData_DataType) EnumDescriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0, 0}
}

type UnixfsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       *UnixfsData_DataType `protobuf:"varint,1,req,name=Type,enum=pb.UnixfsData_DataType" json:"Type,omitempty"`
	Data       []byte               `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Filesize   *uint64              `protobuf:"varint,3,opt,name=filesize" json:"filesize,omitempty"`
	Blocksizes []uint64             `protobuf:"varint,4,rep,name=blocksizes" json:"blocksizes,omitempty"`
	HashType   *uint64              `protobuf:"varint,5,opt,name=hashType" json:"hashType,omitempty"`
	Fanout     *uint64              `protobuf:"varint,6,opt,name=fanout" json:"fanout,omitempty"`
}

func (x *UnixfsData) Reset() {
	*x = UnixfsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnixfsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnixfsData) ProtoMessage() {}

func (x *UnixfsData) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnixfsData.ProtoReflect.Descriptor instead.
func (*UnixfsData) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *UnixfsData) GetType() UnixfsData_DataType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return UnixfsData_Raw
}

func (x *UnixfsData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UnixfsData) GetFilesize() uint64 {
	if x != nil && x.Filesize != nil {
		return *x.Filesize
	}
	return 0
}

func (x *UnixfsData) GetBlocksizes() []uint64 {
	if x != nil {
		return x.Blocksizes
	}
	return nil
}

func (x *UnixfsData) GetHashType() uint64 {
	if x != nil && x.HashType != nil {
		return *x.HashType
	}
	return 0
}

func (x *UnixfsData) GetFanout() uint64 {
	if x != nil && x.Fanout != nil {
		return *x.Fanout
	}
	return 0
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MimeType *string `protobuf:"bytes,1,opt,name=MimeType" json:"MimeType,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetMimeType() string {
	if x != nil && x.MimeType != nil {
		return *x.MimeType
	}
	return ""
}

type PBLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// binary CID (with no multibase prefix) of the target object
	Hash []byte `protobuf:"bytes,1,opt,name=Hash" json:"Hash,omitempty"`
	// UTF-8 string name
	Name *string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	// cumulative size of target object
	Tsize *uint64 `protobuf:"varint,3,opt,name=Tsize" json:"Tsize,omitempty"`
}

func (x *PBLink) Reset() {
	*x = PBLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBLink) ProtoMessage() {}

func (x *PBLink) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBLink.ProtoReflect.Descriptor instead.
func (*PBLink) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *PBLink) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *PBLink) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *PBLink) GetTsize() uint64 {
	if x != nil && x.Tsize != nil {
		return *x.Tsize
	}
	return 0
}

type PBNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// refs to other objects
	Links []*PBLink `protobuf:"bytes,2,rep,name=Links" json:"Links,omitempty"`
	// opaque user data
	Data []byte `protobuf:"bytes,1,opt,name=Data" json:"Data,omitempty"`
}

func (x *PBNode) Reset() {
	*x = PBNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBNode) ProtoMessage() {}

func (x *PBNode) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBNode.ProtoReflect.Descriptor instead.
func (*PBNode) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *PBNode) GetLinks() []*PBLink {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *PBNode) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x95, 0x02, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x78, 0x66, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x75, 0x6e, 0x69, 0x78, 0x66, 0x73, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x68, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6e, 0x6f,
	0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x61, 0x6e, 0x6f, 0x75, 0x74,
	0x22, 0x56, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x52, 0x61, 0x77, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x79, 0x6d, 0x6c, 0x69, 0x6e, 0x6b, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x41, 0x4d,
	0x54, 0x53, 0x68, 0x61, 0x72, 0x64, 0x10, 0x05, 0x22, 0x26, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x46, 0x0a, 0x06, 0x50, 0x42, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x54, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x50, 0x42, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x42, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x03, 0x5a, 0x01, 0x2e,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_file_proto_goTypes = []interface{}{
	(UnixfsData_DataType)(0), // 0: pb.unixfsData.DataType
	(*UnixfsData)(nil),       // 1: pb.unixfsData
	(*Metadata)(nil),         // 2: pb.Metadata
	(*PBLink)(nil),           // 3: pb.PBLink
	(*PBNode)(nil),           // 4: pb.PBNode
}
var file_file_proto_depIdxs = []int32{
	0, // 0: pb.unixfsData.Type:type_name -> pb.unixfsData.DataType
	3, // 1: pb.PBNode.Links:type_name -> pb.PBLink
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnixfsData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBLink); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		EnumInfos:         file_file_proto_enumTypes,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
