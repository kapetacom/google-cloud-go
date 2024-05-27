// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/cloud/retail/v2alpha/analytics_service.proto

package retailpb

import (
	context "context"
	reflect "reflect"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_cloud_retail_v2alpha_analytics_service_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2alpha_analytics_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c,
	0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x97, 0x03, 0x0a, 0x10,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xb7, 0x02, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x3a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc1, 0x01, 0xca, 0x41, 0x68, 0x0a, 0x3a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x50, 0x3a, 0x01, 0x2a, 0x22, 0x4b, 0x2f,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73,
	0x2f, 0x2a, 0x7d, 0x3a, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x49, 0xca, 0x41, 0x15, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xd9, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x15, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x70, 0x62, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0xa2, 0x02, 0x06, 0x52, 0x45,
	0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_cloud_retail_v2alpha_analytics_service_proto_goTypes = []interface{}{
	(*ExportAnalyticsMetricsRequest)(nil), // 0: google.cloud.retail.v2alpha.ExportAnalyticsMetricsRequest
	(*longrunningpb.Operation)(nil),       // 1: google.longrunning.Operation
}
var file_google_cloud_retail_v2alpha_analytics_service_proto_depIdxs = []int32{
	0, // 0: google.cloud.retail.v2alpha.AnalyticsService.ExportAnalyticsMetrics:input_type -> google.cloud.retail.v2alpha.ExportAnalyticsMetricsRequest
	1, // 1: google.cloud.retail.v2alpha.AnalyticsService.ExportAnalyticsMetrics:output_type -> google.longrunning.Operation
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2alpha_analytics_service_proto_init() }
func file_google_cloud_retail_v2alpha_analytics_service_proto_init() {
	if File_google_cloud_retail_v2alpha_analytics_service_proto != nil {
		return
	}
	file_google_cloud_retail_v2alpha_export_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_retail_v2alpha_analytics_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_retail_v2alpha_analytics_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2alpha_analytics_service_proto_depIdxs,
	}.Build()
	File_google_cloud_retail_v2alpha_analytics_service_proto = out.File
	file_google_cloud_retail_v2alpha_analytics_service_proto_rawDesc = nil
	file_google_cloud_retail_v2alpha_analytics_service_proto_goTypes = nil
	file_google_cloud_retail_v2alpha_analytics_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnalyticsServiceClient is the client API for AnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnalyticsServiceClient interface {
	// Exports analytics metrics.
	//
	// `Operation.response` is of type `ExportAnalyticsMetricsResponse`.
	// `Operation.metadata` is of type `ExportMetadata`.
	ExportAnalyticsMetrics(ctx context.Context, in *ExportAnalyticsMetricsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type analyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsServiceClient(cc grpc.ClientConnInterface) AnalyticsServiceClient {
	return &analyticsServiceClient{cc}
}

func (c *analyticsServiceClient) ExportAnalyticsMetrics(ctx context.Context, in *ExportAnalyticsMetricsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2alpha.AnalyticsService/ExportAnalyticsMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsServiceServer is the server API for AnalyticsService service.
type AnalyticsServiceServer interface {
	// Exports analytics metrics.
	//
	// `Operation.response` is of type `ExportAnalyticsMetricsResponse`.
	// `Operation.metadata` is of type `ExportMetadata`.
	ExportAnalyticsMetrics(context.Context, *ExportAnalyticsMetricsRequest) (*longrunningpb.Operation, error)
}

// UnimplementedAnalyticsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServiceServer struct {
}

func (*UnimplementedAnalyticsServiceServer) ExportAnalyticsMetrics(context.Context, *ExportAnalyticsMetricsRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportAnalyticsMetrics not implemented")
}

func RegisterAnalyticsServiceServer(s *grpc.Server, srv AnalyticsServiceServer) {
	s.RegisterService(&_AnalyticsService_serviceDesc, srv)
}

func _AnalyticsService_ExportAnalyticsMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportAnalyticsMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).ExportAnalyticsMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2alpha.AnalyticsService/ExportAnalyticsMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).ExportAnalyticsMetrics(ctx, req.(*ExportAnalyticsMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnalyticsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.retail.v2alpha.AnalyticsService",
	HandlerType: (*AnalyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportAnalyticsMetrics",
			Handler:    _AnalyticsService_ExportAnalyticsMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/retail/v2alpha/analytics_service.proto",
}
