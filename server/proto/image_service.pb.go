// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: proto/image_service.proto

package proto

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

type ImageInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid *ImageUID `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ImageInfoRequest) Reset() {
	*x = ImageInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfoRequest) ProtoMessage() {}

func (x *ImageInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfoRequest.ProtoReflect.Descriptor instead.
func (*ImageInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_image_service_proto_rawDescGZIP(), []int{0}
}

func (x *ImageInfoRequest) GetUid() *ImageUID {
	if x != nil {
		return x.Uid
	}
	return nil
}

type ImageInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dim      *Dimension `protobuf:"bytes,1,opt,name=dim,proto3" json:"dim,omitempty"`
	Filename string     `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *ImageInfoResponse) Reset() {
	*x = ImageInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfoResponse) ProtoMessage() {}

func (x *ImageInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfoResponse.ProtoReflect.Descriptor instead.
func (*ImageInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_image_service_proto_rawDescGZIP(), []int{1}
}

func (x *ImageInfoResponse) GetDim() *Dimension {
	if x != nil {
		return x.Dim
	}
	return nil
}

func (x *ImageInfoResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type ImageUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ImageUID) Reset() {
	*x = ImageUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUID) ProtoMessage() {}

func (x *ImageUID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUID.ProtoReflect.Descriptor instead.
func (*ImageUID) Descriptor() ([]byte, []int) {
	return file_proto_image_service_proto_rawDescGZIP(), []int{2}
}

func (x *ImageUID) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type Dimension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  int32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Dimension) Reset() {
	*x = Dimension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dimension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dimension) ProtoMessage() {}

func (x *Dimension) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dimension.ProtoReflect.Descriptor instead.
func (*Dimension) Descriptor() ([]byte, []int) {
	return file_proto_image_service_proto_rawDescGZIP(), []int{3}
}

func (x *Dimension) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Dimension) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_proto_image_service_proto protoreflect.FileDescriptor

var file_proto_image_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x10, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x49, 0x44, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x11,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x03, 0x64, 0x69, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x64, 0x69, 0x6d, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x09, 0x44, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x32, 0x47, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x6f,
	0x74, 0x6c, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x64, 0x2d, 0x62, 0x65, 0x61, 0x6e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_image_service_proto_rawDescOnce sync.Once
	file_proto_image_service_proto_rawDescData = file_proto_image_service_proto_rawDesc
)

func file_proto_image_service_proto_rawDescGZIP() []byte {
	file_proto_image_service_proto_rawDescOnce.Do(func() {
		file_proto_image_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_image_service_proto_rawDescData)
	})
	return file_proto_image_service_proto_rawDescData
}

var file_proto_image_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_image_service_proto_goTypes = []interface{}{
	(*ImageInfoRequest)(nil),  // 0: ImageInfoRequest
	(*ImageInfoResponse)(nil), // 1: ImageInfoResponse
	(*ImageUID)(nil),          // 2: ImageUID
	(*Dimension)(nil),         // 3: Dimension
}
var file_proto_image_service_proto_depIdxs = []int32{
	2, // 0: ImageInfoRequest.uid:type_name -> ImageUID
	3, // 1: ImageInfoResponse.dim:type_name -> Dimension
	0, // 2: ImageService.GetImageInfo:input_type -> ImageInfoRequest
	1, // 3: ImageService.GetImageInfo:output_type -> ImageInfoResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_image_service_proto_init() }
func file_proto_image_service_proto_init() {
	if File_proto_image_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_image_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInfoRequest); i {
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
		file_proto_image_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInfoResponse); i {
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
		file_proto_image_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUID); i {
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
		file_proto_image_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dimension); i {
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
			RawDescriptor: file_proto_image_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_image_service_proto_goTypes,
		DependencyIndexes: file_proto_image_service_proto_depIdxs,
		MessageInfos:      file_proto_image_service_proto_msgTypes,
	}.Build()
	File_proto_image_service_proto = out.File
	file_proto_image_service_proto_rawDesc = nil
	file_proto_image_service_proto_goTypes = nil
	file_proto_image_service_proto_depIdxs = nil
}