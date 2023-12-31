// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: gobp/v1/gobp_furniture.proto

package pb_v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetFurnitureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFurnitureRequest) Reset() {
	*x = GetFurnitureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gobp_v1_gobp_furniture_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFurnitureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFurnitureRequest) ProtoMessage() {}

func (x *GetFurnitureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gobp_v1_gobp_furniture_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFurnitureRequest.ProtoReflect.Descriptor instead.
func (*GetFurnitureRequest) Descriptor() ([]byte, []int) {
	return file_gobp_v1_gobp_furniture_proto_rawDescGZIP(), []int{0}
}

func (x *GetFurnitureRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFurnitureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetFurnitureResponse) Reset() {
	*x = GetFurnitureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gobp_v1_gobp_furniture_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFurnitureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFurnitureResponse) ProtoMessage() {}

func (x *GetFurnitureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gobp_v1_gobp_furniture_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFurnitureResponse.ProtoReflect.Descriptor instead.
func (*GetFurnitureResponse) Descriptor() ([]byte, []int) {
	return file_gobp_v1_gobp_furniture_proto_rawDescGZIP(), []int{1}
}

func (x *GetFurnitureResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetFurnitureResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFurnitureResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AddFurnitureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *AddFurnitureRequest) Reset() {
	*x = AddFurnitureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gobp_v1_gobp_furniture_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFurnitureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFurnitureRequest) ProtoMessage() {}

func (x *AddFurnitureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gobp_v1_gobp_furniture_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFurnitureRequest.ProtoReflect.Descriptor instead.
func (*AddFurnitureRequest) Descriptor() ([]byte, []int) {
	return file_gobp_v1_gobp_furniture_proto_rawDescGZIP(), []int{2}
}

func (x *AddFurnitureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddFurnitureRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AddFurnitureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddFurnitureResponse) Reset() {
	*x = AddFurnitureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gobp_v1_gobp_furniture_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFurnitureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFurnitureResponse) ProtoMessage() {}

func (x *AddFurnitureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gobp_v1_gobp_furniture_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFurnitureResponse.ProtoReflect.Descriptor instead.
func (*AddFurnitureResponse) Descriptor() ([]byte, []int) {
	return file_gobp_v1_gobp_furniture_proto_rawDescGZIP(), []int{3}
}

func (x *AddFurnitureResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_gobp_v1_gobp_furniture_proto protoreflect.FileDescriptor

var file_gobp_v1_gobp_furniture_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x6f, 0x62, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x62, 0x70, 0x5f, 0x66,
	0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x67, 0x6f, 0x62, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x75,
	0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3d,
	0x0a, 0x13, 0x41, 0x64, 0x64, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x26, 0x0a,
	0x14, 0x41, 0x64, 0x64, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x5f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gobp_v1_gobp_furniture_proto_rawDescOnce sync.Once
	file_gobp_v1_gobp_furniture_proto_rawDescData = file_gobp_v1_gobp_furniture_proto_rawDesc
)

func file_gobp_v1_gobp_furniture_proto_rawDescGZIP() []byte {
	file_gobp_v1_gobp_furniture_proto_rawDescOnce.Do(func() {
		file_gobp_v1_gobp_furniture_proto_rawDescData = protoimpl.X.CompressGZIP(file_gobp_v1_gobp_furniture_proto_rawDescData)
	})
	return file_gobp_v1_gobp_furniture_proto_rawDescData
}

var file_gobp_v1_gobp_furniture_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gobp_v1_gobp_furniture_proto_goTypes = []interface{}{
	(*GetFurnitureRequest)(nil),  // 0: gobp.v1.GetFurnitureRequest
	(*GetFurnitureResponse)(nil), // 1: gobp.v1.GetFurnitureResponse
	(*AddFurnitureRequest)(nil),  // 2: gobp.v1.AddFurnitureRequest
	(*AddFurnitureResponse)(nil), // 3: gobp.v1.AddFurnitureResponse
}
var file_gobp_v1_gobp_furniture_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gobp_v1_gobp_furniture_proto_init() }
func file_gobp_v1_gobp_furniture_proto_init() {
	if File_gobp_v1_gobp_furniture_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gobp_v1_gobp_furniture_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFurnitureRequest); i {
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
		file_gobp_v1_gobp_furniture_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFurnitureResponse); i {
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
		file_gobp_v1_gobp_furniture_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFurnitureRequest); i {
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
		file_gobp_v1_gobp_furniture_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFurnitureResponse); i {
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
			RawDescriptor: file_gobp_v1_gobp_furniture_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gobp_v1_gobp_furniture_proto_goTypes,
		DependencyIndexes: file_gobp_v1_gobp_furniture_proto_depIdxs,
		MessageInfos:      file_gobp_v1_gobp_furniture_proto_msgTypes,
	}.Build()
	File_gobp_v1_gobp_furniture_proto = out.File
	file_gobp_v1_gobp_furniture_proto_rawDesc = nil
	file_gobp_v1_gobp_furniture_proto_goTypes = nil
	file_gobp_v1_gobp_furniture_proto_depIdxs = nil
}
