// idl/hello/hello.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.2
// source: dashboard.proto

package dashboard

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "jobor/kitex_gen/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DashboardQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64    `protobuf:"varint,1,opt,name=id,proto3" query:"id" form:"id" json:"id"`
	Name  string   `protobuf:"bytes,2,opt,name=name,proto3" form:"name" json:"name" query:"name"`
	Range []string `protobuf:"bytes,3,rep,name=range,proto3" form:"range" json:"range" query:"range"`
}

func (x *DashboardQuery) Reset() {
	*x = DashboardQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DashboardQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardQuery) ProtoMessage() {}

func (x *DashboardQuery) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardQuery.ProtoReflect.Descriptor instead.
func (*DashboardQuery) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *DashboardQuery) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DashboardQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DashboardQuery) GetRange() []string {
	if x != nil {
		return x.Range
	}
	return nil
}

type DashboardResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" query:"id" form:"id" json:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" query:"name" form:"name" json:"name"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" form:"description" json:"description" query:"description"`
}

func (x *DashboardResp) Reset() {
	*x = DashboardResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DashboardResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardResp) ProtoMessage() {}

func (x *DashboardResp) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardResp.ProtoReflect.Descriptor instead.
func (*DashboardResp) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *DashboardResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DashboardResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DashboardResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_dashboard_proto protoreflect.FileDescriptor

var file_dashboard_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x09, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0c, 0xb2, 0xbb, 0x18, 0x02, 0x69, 0x64, 0xca, 0xbb, 0x18,
	0x02, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xb2, 0xbb, 0x18, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xca,
	0xbb, 0x18, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x12, 0xb2, 0xbb,
	0x18, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0xca, 0xbb, 0x18, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0d, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0c, 0xb2, 0xbb, 0x18, 0x02, 0x69, 0x64, 0xca, 0xbb, 0x18,
	0x02, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xb2, 0xbb, 0x18, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xca,
	0xbb, 0x18, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1e, 0xb2, 0xbb, 0x18, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0xca, 0xbb, 0x18, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0x74, 0x0a, 0x10, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x19, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x18,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0xca, 0xc1, 0x18, 0x17, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x6f, 0x72, 0x2f, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x1b, 0x5a, 0x19, 0x6a, 0x6f, 0x62, 0x6f, 0x72, 0x2f, 0x6b,
	0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dashboard_proto_rawDescOnce sync.Once
	file_dashboard_proto_rawDescData = file_dashboard_proto_rawDesc
)

func file_dashboard_proto_rawDescGZIP() []byte {
	file_dashboard_proto_rawDescOnce.Do(func() {
		file_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_dashboard_proto_rawDescData)
	})
	return file_dashboard_proto_rawDescData
}

var file_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dashboard_proto_goTypes = []interface{}{
	(*DashboardQuery)(nil), // 0: dashboard.dashboardQuery
	(*DashboardResp)(nil),  // 1: dashboard.dashboardResp
}
var file_dashboard_proto_depIdxs = []int32{
	0, // 0: dashboard.DashboardService.GetDashboard:input_type -> dashboard.dashboardQuery
	1, // 1: dashboard.DashboardService.GetDashboard:output_type -> dashboard.dashboardResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dashboard_proto_init() }
func file_dashboard_proto_init() {
	if File_dashboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DashboardQuery); i {
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
		file_dashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DashboardResp); i {
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
			RawDescriptor: file_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dashboard_proto_goTypes,
		DependencyIndexes: file_dashboard_proto_depIdxs,
		MessageInfos:      file_dashboard_proto_msgTypes,
	}.Build()
	File_dashboard_proto = out.File
	file_dashboard_proto_rawDesc = nil
	file_dashboard_proto_goTypes = nil
	file_dashboard_proto_depIdxs = nil
}
