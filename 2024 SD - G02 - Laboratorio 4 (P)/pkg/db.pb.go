// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: pkg/db.proto

package pkg

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

// TODO: completar
type ParametroPut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clave string `protobuf:"bytes,1,opt,name=clave,proto3" json:"clave,omitempty"`
	Valor []byte `protobuf:"bytes,2,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *ParametroPut) Reset() {
	*x = ParametroPut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParametroPut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParametroPut) ProtoMessage() {}

func (x *ParametroPut) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParametroPut.ProtoReflect.Descriptor instead.
func (*ParametroPut) Descriptor() ([]byte, []int) {
	return file_pkg_db_proto_rawDescGZIP(), []int{0}
}

func (x *ParametroPut) GetClave() string {
	if x != nil {
		return x.Clave
	}
	return ""
}

func (x *ParametroPut) GetValor() []byte {
	if x != nil {
		return x.Valor
	}
	return nil
}

// TODO: completar
type ParametroGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clave string `protobuf:"bytes,1,opt,name=clave,proto3" json:"clave,omitempty"`
}

func (x *ParametroGet) Reset() {
	*x = ParametroGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParametroGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParametroGet) ProtoMessage() {}

func (x *ParametroGet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParametroGet.ProtoReflect.Descriptor instead.
func (*ParametroGet) Descriptor() ([]byte, []int) {
	return file_pkg_db_proto_rawDescGZIP(), []int{1}
}

func (x *ParametroGet) GetClave() string {
	if x != nil {
		return x.Clave
	}
	return ""
}

// TODO: completar
type ParametroGetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ParametroGetAll) Reset() {
	*x = ParametroGetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParametroGetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParametroGetAll) ProtoMessage() {}

func (x *ParametroGetAll) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParametroGetAll.ProtoReflect.Descriptor instead.
func (*ParametroGetAll) Descriptor() ([]byte, []int) {
	return file_pkg_db_proto_rawDescGZIP(), []int{2}
}

// TODO: completar
type ResultadoPut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResultadoPut) Reset() {
	*x = ResultadoPut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultadoPut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultadoPut) ProtoMessage() {}

func (x *ResultadoPut) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultadoPut.ProtoReflect.Descriptor instead.
func (*ResultadoPut) Descriptor() ([]byte, []int) {
	return file_pkg_db_proto_rawDescGZIP(), []int{3}
}

// TODO: completar
type ResultadoGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valor []byte `protobuf:"bytes,1,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *ResultadoGet) Reset() {
	*x = ResultadoGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultadoGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultadoGet) ProtoMessage() {}

func (x *ResultadoGet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultadoGet.ProtoReflect.Descriptor instead.
func (*ResultadoGet) Descriptor() ([]byte, []int) {
	return file_pkg_db_proto_rawDescGZIP(), []int{4}
}

func (x *ResultadoGet) GetValor() []byte {
	if x != nil {
		return x.Valor
	}
	return nil
}

// TODO: completar
type ResultadoGetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClaveValor []*ClaveValor `protobuf:"bytes,1,rep,name=claveValor,proto3" json:"claveValor,omitempty"`
}

func (x *ResultadoGetAll) Reset() {
	*x = ResultadoGetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultadoGetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultadoGetAll) ProtoMessage() {}

func (x *ResultadoGetAll) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultadoGetAll.ProtoReflect.Descriptor instead.
func (*ResultadoGetAll) Descriptor() ([]byte, []int) {
	return file_pkg_db_proto_rawDescGZIP(), []int{5}
}

func (x *ResultadoGetAll) GetClaveValor() []*ClaveValor {
	if x != nil {
		return x.ClaveValor
	}
	return nil
}

// Resultado para la operación GetAll
type ClaveValor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clave string `protobuf:"bytes,1,opt,name=clave,proto3" json:"clave,omitempty"`
	Valor []byte `protobuf:"bytes,2,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *ClaveValor) Reset() {
	*x = ClaveValor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaveValor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaveValor) ProtoMessage() {}

func (x *ClaveValor) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_db_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaveValor.ProtoReflect.Descriptor instead.
func (*ClaveValor) Descriptor() ([]byte, []int) {
	return file_pkg_db_proto_rawDescGZIP(), []int{6}
}

func (x *ClaveValor) GetClave() string {
	if x != nil {
		return x.Clave
	}
	return ""
}

func (x *ClaveValor) GetValor() []byte {
	if x != nil {
		return x.Valor
	}
	return nil
}

var File_pkg_db_proto protoreflect.FileDescriptor

var file_pkg_db_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x64, 0x62, 0x22, 0x3a, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x50,
	0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x24,
	0x0a, 0x0c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x47, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6c, 0x61, 0x76, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x72,
	0x6f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x61, 0x64, 0x6f, 0x50, 0x75, 0x74, 0x22, 0x24, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x61, 0x64, 0x6f, 0x47, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x41, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x64, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x6c, 0x61, 0x76, 0x65, 0x56,
	0x61, 0x6c, 0x6f, 0x72, 0x52, 0x0a, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x56, 0x61, 0x6c, 0x6f, 0x72,
	0x22, 0x38, 0x0a, 0x0a, 0x43, 0x6c, 0x61, 0x76, 0x65, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6c, 0x61, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x32, 0x90, 0x01, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x70, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x64, 0x62, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x50, 0x75, 0x74, 0x1a, 0x10, 0x2e, 0x64,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x64, 0x6f, 0x50, 0x75, 0x74, 0x12, 0x29,
	0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x64, 0x62, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x72, 0x6f, 0x47, 0x65, 0x74, 0x1a, 0x10, 0x2e, 0x64, 0x62, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x61, 0x64, 0x6f, 0x47, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x64, 0x62, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x72, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x1a, 0x13, 0x2e, 0x64, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x61, 0x64, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x08, 0x5a,
	0x06, 0x64, 0x62, 0x2f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_db_proto_rawDescOnce sync.Once
	file_pkg_db_proto_rawDescData = file_pkg_db_proto_rawDesc
)

func file_pkg_db_proto_rawDescGZIP() []byte {
	file_pkg_db_proto_rawDescOnce.Do(func() {
		file_pkg_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_db_proto_rawDescData)
	})
	return file_pkg_db_proto_rawDescData
}

var file_pkg_db_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_db_proto_goTypes = []interface{}{
	(*ParametroPut)(nil),    // 0: db.ParametroPut
	(*ParametroGet)(nil),    // 1: db.ParametroGet
	(*ParametroGetAll)(nil), // 2: db.ParametroGetAll
	(*ResultadoPut)(nil),    // 3: db.ResultadoPut
	(*ResultadoGet)(nil),    // 4: db.ResultadoGet
	(*ResultadoGetAll)(nil), // 5: db.ResultadoGetAll
	(*ClaveValor)(nil),      // 6: db.ClaveValor
}
var file_pkg_db_proto_depIdxs = []int32{
	6, // 0: db.ResultadoGetAll.claveValor:type_name -> db.ClaveValor
	0, // 1: db.base.put:input_type -> db.ParametroPut
	1, // 2: db.base.get:input_type -> db.ParametroGet
	2, // 3: db.base.getAll:input_type -> db.ParametroGetAll
	3, // 4: db.base.put:output_type -> db.ResultadoPut
	4, // 5: db.base.get:output_type -> db.ResultadoGet
	5, // 6: db.base.getAll:output_type -> db.ResultadoGetAll
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_db_proto_init() }
func file_pkg_db_proto_init() {
	if File_pkg_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParametroPut); i {
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
		file_pkg_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParametroGet); i {
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
		file_pkg_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParametroGetAll); i {
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
		file_pkg_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultadoPut); i {
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
		file_pkg_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultadoGet); i {
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
		file_pkg_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultadoGetAll); i {
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
		file_pkg_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaveValor); i {
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
			RawDescriptor: file_pkg_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_db_proto_goTypes,
		DependencyIndexes: file_pkg_db_proto_depIdxs,
		MessageInfos:      file_pkg_db_proto_msgTypes,
	}.Build()
	File_pkg_db_proto = out.File
	file_pkg_db_proto_rawDesc = nil
	file_pkg_db_proto_goTypes = nil
	file_pkg_db_proto_depIdxs = nil
}
