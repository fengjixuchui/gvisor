// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.0
// source: pkg/metric/metric.proto

package gvisor

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MetricMetadata_Type int32

const (
	MetricMetadata_TYPE_UINT64 MetricMetadata_Type = 0
)

// Enum value maps for MetricMetadata_Type.
var (
	MetricMetadata_Type_name = map[int32]string{
		0: "TYPE_UINT64",
	}
	MetricMetadata_Type_value = map[string]int32{
		"TYPE_UINT64": 0,
	}
)

func (x MetricMetadata_Type) Enum() *MetricMetadata_Type {
	p := new(MetricMetadata_Type)
	*p = x
	return p
}

func (x MetricMetadata_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricMetadata_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_metric_metric_proto_enumTypes[0].Descriptor()
}

func (MetricMetadata_Type) Type() protoreflect.EnumType {
	return &file_pkg_metric_metric_proto_enumTypes[0]
}

func (x MetricMetadata_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricMetadata_Type.Descriptor instead.
func (MetricMetadata_Type) EnumDescriptor() ([]byte, []int) {
	return file_pkg_metric_metric_proto_rawDescGZIP(), []int{0, 0}
}

type MetricMetadata_Units int32

const (
	MetricMetadata_UNITS_NONE        MetricMetadata_Units = 0
	MetricMetadata_UNITS_NANOSECONDS MetricMetadata_Units = 1
)

// Enum value maps for MetricMetadata_Units.
var (
	MetricMetadata_Units_name = map[int32]string{
		0: "UNITS_NONE",
		1: "UNITS_NANOSECONDS",
	}
	MetricMetadata_Units_value = map[string]int32{
		"UNITS_NONE":        0,
		"UNITS_NANOSECONDS": 1,
	}
)

func (x MetricMetadata_Units) Enum() *MetricMetadata_Units {
	p := new(MetricMetadata_Units)
	*p = x
	return p
}

func (x MetricMetadata_Units) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricMetadata_Units) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_metric_metric_proto_enumTypes[1].Descriptor()
}

func (MetricMetadata_Units) Type() protoreflect.EnumType {
	return &file_pkg_metric_metric_proto_enumTypes[1]
}

func (x MetricMetadata_Units) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricMetadata_Units.Descriptor instead.
func (MetricMetadata_Units) EnumDescriptor() ([]byte, []int) {
	return file_pkg_metric_metric_proto_rawDescGZIP(), []int{0, 1}
}

type MetricMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Cumulative  bool                 `protobuf:"varint,3,opt,name=cumulative,proto3" json:"cumulative,omitempty"`
	Sync        bool                 `protobuf:"varint,4,opt,name=sync,proto3" json:"sync,omitempty"`
	Type        MetricMetadata_Type  `protobuf:"varint,5,opt,name=type,proto3,enum=gvisor.MetricMetadata_Type" json:"type,omitempty"`
	Units       MetricMetadata_Units `protobuf:"varint,6,opt,name=units,proto3,enum=gvisor.MetricMetadata_Units" json:"units,omitempty"`
}

func (x *MetricMetadata) Reset() {
	*x = MetricMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_metric_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricMetadata) ProtoMessage() {}

func (x *MetricMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metric_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricMetadata.ProtoReflect.Descriptor instead.
func (*MetricMetadata) Descriptor() ([]byte, []int) {
	return file_pkg_metric_metric_proto_rawDescGZIP(), []int{0}
}

func (x *MetricMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MetricMetadata) GetCumulative() bool {
	if x != nil {
		return x.Cumulative
	}
	return false
}

func (x *MetricMetadata) GetSync() bool {
	if x != nil {
		return x.Sync
	}
	return false
}

func (x *MetricMetadata) GetType() MetricMetadata_Type {
	if x != nil {
		return x.Type
	}
	return MetricMetadata_TYPE_UINT64
}

func (x *MetricMetadata) GetUnits() MetricMetadata_Units {
	if x != nil {
		return x.Units
	}
	return MetricMetadata_UNITS_NONE
}

type MetricRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*MetricMetadata `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *MetricRegistration) Reset() {
	*x = MetricRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_metric_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricRegistration) ProtoMessage() {}

func (x *MetricRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metric_metric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricRegistration.ProtoReflect.Descriptor instead.
func (*MetricRegistration) Descriptor() ([]byte, []int) {
	return file_pkg_metric_metric_proto_rawDescGZIP(), []int{1}
}

func (x *MetricRegistration) GetMetrics() []*MetricMetadata {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type MetricValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Value:
	//	*MetricValue_Uint64Value
	Value isMetricValue_Value `protobuf_oneof:"value"`
}

func (x *MetricValue) Reset() {
	*x = MetricValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_metric_metric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricValue) ProtoMessage() {}

func (x *MetricValue) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metric_metric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricValue.ProtoReflect.Descriptor instead.
func (*MetricValue) Descriptor() ([]byte, []int) {
	return file_pkg_metric_metric_proto_rawDescGZIP(), []int{2}
}

func (x *MetricValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *MetricValue) GetValue() isMetricValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *MetricValue) GetUint64Value() uint64 {
	if x, ok := x.GetValue().(*MetricValue_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

type isMetricValue_Value interface {
	isMetricValue_Value()
}

type MetricValue_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,2,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

func (*MetricValue_Uint64Value) isMetricValue_Value() {}

type MetricUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*MetricValue `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *MetricUpdate) Reset() {
	*x = MetricUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_metric_metric_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricUpdate) ProtoMessage() {}

func (x *MetricUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metric_metric_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricUpdate.ProtoReflect.Descriptor instead.
func (*MetricUpdate) Descriptor() ([]byte, []int) {
	return file_pkg_metric_metric_proto_rawDescGZIP(), []int{3}
}

func (x *MetricUpdate) GetMetrics() []*MetricValue {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_pkg_metric_metric_proto protoreflect.FileDescriptor

var file_pkg_metric_metric_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x22, 0xa8, 0x02, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x79,
	0x6e, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x2f,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x67, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x05, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x22, 0x17, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x00, 0x22, 0x2e, 0x0a, 0x05,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x4e,
	0x41, 0x4e, 0x4f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x10, 0x01, 0x22, 0x46, 0x0a, 0x12,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x22, 0x4f, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52,
	0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3d, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_metric_metric_proto_rawDescOnce sync.Once
	file_pkg_metric_metric_proto_rawDescData = file_pkg_metric_metric_proto_rawDesc
)

func file_pkg_metric_metric_proto_rawDescGZIP() []byte {
	file_pkg_metric_metric_proto_rawDescOnce.Do(func() {
		file_pkg_metric_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_metric_metric_proto_rawDescData)
	})
	return file_pkg_metric_metric_proto_rawDescData
}

var file_pkg_metric_metric_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_metric_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_metric_metric_proto_goTypes = []interface{}{
	(MetricMetadata_Type)(0),   // 0: gvisor.MetricMetadata.Type
	(MetricMetadata_Units)(0),  // 1: gvisor.MetricMetadata.Units
	(*MetricMetadata)(nil),     // 2: gvisor.MetricMetadata
	(*MetricRegistration)(nil), // 3: gvisor.MetricRegistration
	(*MetricValue)(nil),        // 4: gvisor.MetricValue
	(*MetricUpdate)(nil),       // 5: gvisor.MetricUpdate
}
var file_pkg_metric_metric_proto_depIdxs = []int32{
	0, // 0: gvisor.MetricMetadata.type:type_name -> gvisor.MetricMetadata.Type
	1, // 1: gvisor.MetricMetadata.units:type_name -> gvisor.MetricMetadata.Units
	2, // 2: gvisor.MetricRegistration.metrics:type_name -> gvisor.MetricMetadata
	4, // 3: gvisor.MetricUpdate.metrics:type_name -> gvisor.MetricValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_metric_metric_proto_init() }
func file_pkg_metric_metric_proto_init() {
	if File_pkg_metric_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_metric_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricMetadata); i {
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
		file_pkg_metric_metric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricRegistration); i {
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
		file_pkg_metric_metric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricValue); i {
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
		file_pkg_metric_metric_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricUpdate); i {
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
	file_pkg_metric_metric_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MetricValue_Uint64Value)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_metric_metric_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metric_metric_proto_goTypes,
		DependencyIndexes: file_pkg_metric_metric_proto_depIdxs,
		EnumInfos:         file_pkg_metric_metric_proto_enumTypes,
		MessageInfos:      file_pkg_metric_metric_proto_msgTypes,
	}.Build()
	File_pkg_metric_metric_proto = out.File
	file_pkg_metric_metric_proto_rawDesc = nil
	file_pkg_metric_metric_proto_goTypes = nil
	file_pkg_metric_metric_proto_depIdxs = nil
}
