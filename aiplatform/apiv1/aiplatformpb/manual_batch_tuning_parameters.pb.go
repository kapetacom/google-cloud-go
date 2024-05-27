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
// source: google/cloud/aiplatform/v1/manual_batch_tuning_parameters.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Manual batch tuning parameters.
type ManualBatchTuningParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The number of the records (e.g. instances) of the operation
	// given in each batch to a machine replica. Machine type, and size of a
	// single record should be considered when setting this parameter, higher
	// value speeds up the batch operation's execution, but too high value will
	// result in a whole batch not fitting in a machine's memory, and the whole
	// operation will fail.
	// The default value is 64.
	BatchSize int32 `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
}

func (x *ManualBatchTuningParameters) Reset() {
	*x = ManualBatchTuningParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManualBatchTuningParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManualBatchTuningParameters) ProtoMessage() {}

func (x *ManualBatchTuningParameters) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManualBatchTuningParameters.ProtoReflect.Descriptor instead.
func (*ManualBatchTuningParameters) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDescGZIP(), []int{0}
}

func (x *ManualBatchTuningParameters) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

var File_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41,
	0x0a, 0x1b, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x54, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a,
	0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a,
	0x65, 0x42, 0xde, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x42, 0x20, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDescData = file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_goTypes = []interface{}{
	(*ManualBatchTuningParameters)(nil), // 0: google.cloud.aiplatform.v1.ManualBatchTuningParameters
}
var file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_init() }
func file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_init() {
	if File_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManualBatchTuningParameters); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto = out.File
	file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_manual_batch_tuning_parameters_proto_depIdxs = nil
}
