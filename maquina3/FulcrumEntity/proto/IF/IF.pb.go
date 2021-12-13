// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: IF.proto

package fulcrum

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

type Reloj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server1 int64 `protobuf:"varint,1,opt,name=Server1,proto3" json:"Server1,omitempty"`
	Server2 int64 `protobuf:"varint,2,opt,name=Server2,proto3" json:"Server2,omitempty"`
	Server3 int64 `protobuf:"varint,3,opt,name=Server3,proto3" json:"Server3,omitempty"`
}

func (x *Reloj) Reset() {
	*x = Reloj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reloj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reloj) ProtoMessage() {}

func (x *Reloj) ProtoReflect() protoreflect.Message {
	mi := &file_IF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reloj.ProtoReflect.Descriptor instead.
func (*Reloj) Descriptor() ([]byte, []int) {
	return file_IF_proto_rawDescGZIP(), []int{0}
}

func (x *Reloj) GetServer1() int64 {
	if x != nil {
		return x.Server1
	}
	return 0
}

func (x *Reloj) GetServer2() int64 {
	if x != nil {
		return x.Server2
	}
	return 0
}

func (x *Reloj) GetServer3() int64 {
	if x != nil {
		return x.Server3
	}
	return 0
}

type WriteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comando string `protobuf:"bytes,1,opt,name=comando,proto3" json:"comando,omitempty"`
	Planeta string `protobuf:"bytes,2,opt,name=planeta,proto3" json:"planeta,omitempty"`
	Ciudad  string `protobuf:"bytes,3,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
	Valor   string `protobuf:"bytes,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Caso    int64  `protobuf:"varint,5,opt,name=caso,proto3" json:"caso,omitempty"`
}

func (x *WriteReq) Reset() {
	*x = WriteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IF_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteReq) ProtoMessage() {}

func (x *WriteReq) ProtoReflect() protoreflect.Message {
	mi := &file_IF_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteReq.ProtoReflect.Descriptor instead.
func (*WriteReq) Descriptor() ([]byte, []int) {
	return file_IF_proto_rawDescGZIP(), []int{1}
}

func (x *WriteReq) GetComando() string {
	if x != nil {
		return x.Comando
	}
	return ""
}

func (x *WriteReq) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *WriteReq) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *WriteReq) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *WriteReq) GetCaso() int64 {
	if x != nil {
		return x.Caso
	}
	return 0
}

type WriteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reloj  *Reloj `protobuf:"bytes,1,opt,name=reloj,proto3" json:"reloj,omitempty"`
	Estado int64  `protobuf:"varint,2,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *WriteResp) Reset() {
	*x = WriteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IF_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResp) ProtoMessage() {}

func (x *WriteResp) ProtoReflect() protoreflect.Message {
	mi := &file_IF_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResp.ProtoReflect.Descriptor instead.
func (*WriteResp) Descriptor() ([]byte, []int) {
	return file_IF_proto_rawDescGZIP(), []int{2}
}

func (x *WriteResp) GetReloj() *Reloj {
	if x != nil {
		return x.Reloj
	}
	return nil
}

func (x *WriteResp) GetEstado() int64 {
	if x != nil {
		return x.Estado
	}
	return 0
}

var File_IF_proto protoreflect.FileDescriptor

var file_IF_proto_rawDesc = []byte{
	0x0a, 0x08, 0x49, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x49, 0x46, 0x22, 0x55,
	0x0a, 0x05, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x31, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x33, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x61, 0x73, 0x6f, 0x22, 0x44, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x46, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x52,
	0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x32, 0x47,
	0x0a, 0x18, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x63,
	0x72, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x0c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x49, 0x46, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x49, 0x46, 0x2e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0d, 0x5a, 0x0b, 0x6c, 0x61, 0x62, 0x2f, 0x66,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IF_proto_rawDescOnce sync.Once
	file_IF_proto_rawDescData = file_IF_proto_rawDesc
)

func file_IF_proto_rawDescGZIP() []byte {
	file_IF_proto_rawDescOnce.Do(func() {
		file_IF_proto_rawDescData = protoimpl.X.CompressGZIP(file_IF_proto_rawDescData)
	})
	return file_IF_proto_rawDescData
}

var file_IF_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_IF_proto_goTypes = []interface{}{
	(*Reloj)(nil),     // 0: IF.Reloj
	(*WriteReq)(nil),  // 1: IF.WriteReq
	(*WriteResp)(nil), // 2: IF.WriteResp
}
var file_IF_proto_depIdxs = []int32{
	0, // 0: IF.WriteResp.reloj:type_name -> IF.Reloj
	1, // 1: IF.InformanteFulcrumService.RequestWrite:input_type -> IF.WriteReq
	2, // 2: IF.InformanteFulcrumService.RequestWrite:output_type -> IF.WriteResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_IF_proto_init() }
func file_IF_proto_init() {
	if File_IF_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_IF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reloj); i {
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
		file_IF_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteReq); i {
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
		file_IF_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResp); i {
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
			RawDescriptor: file_IF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_IF_proto_goTypes,
		DependencyIndexes: file_IF_proto_depIdxs,
		MessageInfos:      file_IF_proto_msgTypes,
	}.Build()
	File_IF_proto = out.File
	file_IF_proto_rawDesc = nil
	file_IF_proto_goTypes = nil
	file_IF_proto_depIdxs = nil
}
