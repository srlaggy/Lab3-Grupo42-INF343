// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: BP.proto

package princesa

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
		mi := &file_BP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reloj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reloj) ProtoMessage() {}

func (x *Reloj) ProtoReflect() protoreflect.Message {
	mi := &file_BP_proto_msgTypes[0]
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
	return file_BP_proto_rawDescGZIP(), []int{0}
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

type RebeldesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	Ciudad  string `protobuf:"bytes,2,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
	Reloj   *Reloj `protobuf:"bytes,3,opt,name=reloj,proto3" json:"reloj,omitempty"`
	Server  int64  `protobuf:"varint,4,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *RebeldesReq) Reset() {
	*x = RebeldesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BP_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebeldesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebeldesReq) ProtoMessage() {}

func (x *RebeldesReq) ProtoReflect() protoreflect.Message {
	mi := &file_BP_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebeldesReq.ProtoReflect.Descriptor instead.
func (*RebeldesReq) Descriptor() ([]byte, []int) {
	return file_BP_proto_rawDescGZIP(), []int{1}
}

func (x *RebeldesReq) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *RebeldesReq) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *RebeldesReq) GetReloj() *Reloj {
	if x != nil {
		return x.Reloj
	}
	return nil
}

func (x *RebeldesReq) GetServer() int64 {
	if x != nil {
		return x.Server
	}
	return 0
}

type RebeldesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta      string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	Ciudad       string `protobuf:"bytes,2,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
	Reloj        *Reloj `protobuf:"bytes,3,opt,name=reloj,proto3" json:"reloj,omitempty"`
	Server       int64  `protobuf:"varint,4,opt,name=server,proto3" json:"server,omitempty"`
	CantRebeldes int64  `protobuf:"varint,5,opt,name=cantRebeldes,proto3" json:"cantRebeldes,omitempty"`
}

func (x *RebeldesResp) Reset() {
	*x = RebeldesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BP_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebeldesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebeldesResp) ProtoMessage() {}

func (x *RebeldesResp) ProtoReflect() protoreflect.Message {
	mi := &file_BP_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebeldesResp.ProtoReflect.Descriptor instead.
func (*RebeldesResp) Descriptor() ([]byte, []int) {
	return file_BP_proto_rawDescGZIP(), []int{2}
}

func (x *RebeldesResp) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *RebeldesResp) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *RebeldesResp) GetReloj() *Reloj {
	if x != nil {
		return x.Reloj
	}
	return nil
}

func (x *RebeldesResp) GetServer() int64 {
	if x != nil {
		return x.Server
	}
	return 0
}

func (x *RebeldesResp) GetCantRebeldes() int64 {
	if x != nil {
		return x.CantRebeldes
	}
	return 0
}

var File_BP_proto protoreflect.FileDescriptor

var file_BP_proto_rawDesc = []byte{
	0x0a, 0x08, 0x42, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x42, 0x50, 0x22, 0x55,
	0x0a, 0x05, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x31, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x33, 0x22, 0x78, 0x0a, 0x0b, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x42, 0x50, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x6a,
	0x52, 0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22,
	0x9d, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69,
	0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x75, 0x64,
	0x61, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x42, 0x50, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x52, 0x05, 0x72, 0x65,
	0x6c, 0x6f, 0x6a, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x32,
	0x4d, 0x0a, 0x15, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x65, 0x73,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x42, 0x50,
	0x2e, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x42,
	0x50, 0x2e, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0e,
	0x5a, 0x0c, 0x6c, 0x61, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BP_proto_rawDescOnce sync.Once
	file_BP_proto_rawDescData = file_BP_proto_rawDesc
)

func file_BP_proto_rawDescGZIP() []byte {
	file_BP_proto_rawDescOnce.Do(func() {
		file_BP_proto_rawDescData = protoimpl.X.CompressGZIP(file_BP_proto_rawDescData)
	})
	return file_BP_proto_rawDescData
}

var file_BP_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_BP_proto_goTypes = []interface{}{
	(*Reloj)(nil),        // 0: BP.Reloj
	(*RebeldesReq)(nil),  // 1: BP.RebeldesReq
	(*RebeldesResp)(nil), // 2: BP.RebeldesResp
}
var file_BP_proto_depIdxs = []int32{
	0, // 0: BP.RebeldesReq.reloj:type_name -> BP.Reloj
	0, // 1: BP.RebeldesResp.reloj:type_name -> BP.Reloj
	1, // 2: BP.BrokerPrincesaService.RequestRebeldes:input_type -> BP.RebeldesReq
	2, // 3: BP.BrokerPrincesaService.RequestRebeldes:output_type -> BP.RebeldesResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_BP_proto_init() }
func file_BP_proto_init() {
	if File_BP_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_BP_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebeldesReq); i {
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
		file_BP_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebeldesResp); i {
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
			RawDescriptor: file_BP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_BP_proto_goTypes,
		DependencyIndexes: file_BP_proto_depIdxs,
		MessageInfos:      file_BP_proto_msgTypes,
	}.Build()
	File_BP_proto = out.File
	file_BP_proto_rawDesc = nil
	file_BP_proto_goTypes = nil
	file_BP_proto_depIdxs = nil
}
