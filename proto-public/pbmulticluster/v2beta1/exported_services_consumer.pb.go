// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pbmulticluster/v2beta1/exported_services_consumer.proto

package multiclusterv2beta1

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

// +kubebuilder:validation:Type=object
// +kubebuilder:validation:Schemaless
// +kubebuilder:pruning:PreserveUnknownFields
type ExportedServicesConsumer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConsumerTenancy:
	//
	//	*ExportedServicesConsumer_Peer
	//	*ExportedServicesConsumer_Partition
	//	*ExportedServicesConsumer_SamenessGroup
	ConsumerTenancy isExportedServicesConsumer_ConsumerTenancy `protobuf_oneof:"consumer_tenancy"`
}

func (x *ExportedServicesConsumer) Reset() {
	*x = ExportedServicesConsumer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmulticluster_v2beta1_exported_services_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportedServicesConsumer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportedServicesConsumer) ProtoMessage() {}

func (x *ExportedServicesConsumer) ProtoReflect() protoreflect.Message {
	mi := &file_pbmulticluster_v2beta1_exported_services_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportedServicesConsumer.ProtoReflect.Descriptor instead.
func (*ExportedServicesConsumer) Descriptor() ([]byte, []int) {
	return file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDescGZIP(), []int{0}
}

func (m *ExportedServicesConsumer) GetConsumerTenancy() isExportedServicesConsumer_ConsumerTenancy {
	if m != nil {
		return m.ConsumerTenancy
	}
	return nil
}

func (x *ExportedServicesConsumer) GetPeer() string {
	if x, ok := x.GetConsumerTenancy().(*ExportedServicesConsumer_Peer); ok {
		return x.Peer
	}
	return ""
}

func (x *ExportedServicesConsumer) GetPartition() string {
	if x, ok := x.GetConsumerTenancy().(*ExportedServicesConsumer_Partition); ok {
		return x.Partition
	}
	return ""
}

func (x *ExportedServicesConsumer) GetSamenessGroup() string {
	if x, ok := x.GetConsumerTenancy().(*ExportedServicesConsumer_SamenessGroup); ok {
		return x.SamenessGroup
	}
	return ""
}

type isExportedServicesConsumer_ConsumerTenancy interface {
	isExportedServicesConsumer_ConsumerTenancy()
}

type ExportedServicesConsumer_Peer struct {
	Peer string `protobuf:"bytes,1,opt,name=peer,proto3,oneof"`
}

type ExportedServicesConsumer_Partition struct {
	Partition string `protobuf:"bytes,2,opt,name=partition,proto3,oneof"`
}

type ExportedServicesConsumer_SamenessGroup struct {
	SamenessGroup string `protobuf:"bytes,3,opt,name=sameness_group,json=samenessGroup,proto3,oneof"`
}

func (*ExportedServicesConsumer_Peer) isExportedServicesConsumer_ConsumerTenancy() {}

func (*ExportedServicesConsumer_Partition) isExportedServicesConsumer_ConsumerTenancy() {}

func (*ExportedServicesConsumer_SamenessGroup) isExportedServicesConsumer_ConsumerTenancy() {}

var File_pbmulticluster_v2beta1_exported_services_consumer_proto protoreflect.FileDescriptor

var file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDesc = []byte{
	0x0a, 0x37, 0x70, 0x62, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x22, 0x8d, 0x01, 0x0a, 0x18, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70,
	0x65, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0e, 0x73, 0x61, 0x6d, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x73,
	0x61, 0x6d, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x12, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79,
	0x42, 0xd6, 0x02, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x1d,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x3b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x4d, 0xaa, 0x02, 0x25, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x25, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x31, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x28, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x3a, 0x3a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDescOnce sync.Once
	file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDescData = file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDesc
)

func file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDescGZIP() []byte {
	file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDescOnce.Do(func() {
		file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDescData)
	})
	return file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDescData
}

var file_pbmulticluster_v2beta1_exported_services_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pbmulticluster_v2beta1_exported_services_consumer_proto_goTypes = []interface{}{
	(*ExportedServicesConsumer)(nil), // 0: hashicorp.consul.multicluster.v2beta1.ExportedServicesConsumer
}
var file_pbmulticluster_v2beta1_exported_services_consumer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbmulticluster_v2beta1_exported_services_consumer_proto_init() }
func file_pbmulticluster_v2beta1_exported_services_consumer_proto_init() {
	if File_pbmulticluster_v2beta1_exported_services_consumer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmulticluster_v2beta1_exported_services_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportedServicesConsumer); i {
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
	file_pbmulticluster_v2beta1_exported_services_consumer_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ExportedServicesConsumer_Peer)(nil),
		(*ExportedServicesConsumer_Partition)(nil),
		(*ExportedServicesConsumer_SamenessGroup)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmulticluster_v2beta1_exported_services_consumer_proto_goTypes,
		DependencyIndexes: file_pbmulticluster_v2beta1_exported_services_consumer_proto_depIdxs,
		MessageInfos:      file_pbmulticluster_v2beta1_exported_services_consumer_proto_msgTypes,
	}.Build()
	File_pbmulticluster_v2beta1_exported_services_consumer_proto = out.File
	file_pbmulticluster_v2beta1_exported_services_consumer_proto_rawDesc = nil
	file_pbmulticluster_v2beta1_exported_services_consumer_proto_goTypes = nil
	file_pbmulticluster_v2beta1_exported_services_consumer_proto_depIdxs = nil
}
