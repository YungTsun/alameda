// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/predict.proto

package containers_ai_alameda_v1alpha1_datahub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RecommendationPolicy int32

const (
	RecommendationPolicy_RECOMMENDATIONPOLICY_UNDEFINED RecommendationPolicy = 0
	RecommendationPolicy_STABLE                         RecommendationPolicy = 1
	RecommendationPolicy_COMPACT                        RecommendationPolicy = 2
)

var RecommendationPolicy_name = map[int32]string{
	0: "RECOMMENDATIONPOLICY_UNDEFINED",
	1: "STABLE",
	2: "COMPACT",
}
var RecommendationPolicy_value = map[string]int32{
	"RECOMMENDATIONPOLICY_UNDEFINED": 0,
	"STABLE":                         1,
	"COMPACT":                        2,
}

func (x RecommendationPolicy) String() string {
	return proto.EnumName(RecommendationPolicy_name, int32(x))
}
func (RecommendationPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_predict_173e05f83af7760f, []int{0}
}

type PredictContainer struct {
	Name                          string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RowPredictData                []*MetricData `protobuf:"bytes,2,rep,name=row_predict_data,json=rowPredictData,proto3" json:"row_predict_data,omitempty"`
	LimitPredictData              []*MetricData `protobuf:"bytes,3,rep,name=limit_predict_data,json=limitPredictData,proto3" json:"limit_predict_data,omitempty"`
	RequestPredictData            []*MetricData `protobuf:"bytes,4,rep,name=request_predict_data,json=requestPredictData,proto3" json:"request_predict_data,omitempty"`
	InitialLimitPredictResource   []*MetricData `protobuf:"bytes,5,rep,name=initial_limit_predict_resource,json=initialLimitPredictResource,proto3" json:"initial_limit_predict_resource,omitempty"`
	InitialRequestPredictResource []*MetricData `protobuf:"bytes,6,rep,name=initial_request_predict_resource,json=initialRequestPredictResource,proto3" json:"initial_request_predict_resource,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}      `json:"-"`
	XXX_unrecognized              []byte        `json:"-"`
	XXX_sizecache                 int32         `json:"-"`
}

func (m *PredictContainer) Reset()         { *m = PredictContainer{} }
func (m *PredictContainer) String() string { return proto.CompactTextString(m) }
func (*PredictContainer) ProtoMessage()    {}
func (*PredictContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_predict_173e05f83af7760f, []int{0}
}
func (m *PredictContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictContainer.Unmarshal(m, b)
}
func (m *PredictContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictContainer.Marshal(b, m, deterministic)
}
func (dst *PredictContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictContainer.Merge(dst, src)
}
func (m *PredictContainer) XXX_Size() int {
	return xxx_messageInfo_PredictContainer.Size(m)
}
func (m *PredictContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictContainer.DiscardUnknown(m)
}

var xxx_messageInfo_PredictContainer proto.InternalMessageInfo

func (m *PredictContainer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PredictContainer) GetRowPredictData() []*MetricData {
	if m != nil {
		return m.RowPredictData
	}
	return nil
}

func (m *PredictContainer) GetLimitPredictData() []*MetricData {
	if m != nil {
		return m.LimitPredictData
	}
	return nil
}

func (m *PredictContainer) GetRequestPredictData() []*MetricData {
	if m != nil {
		return m.RequestPredictData
	}
	return nil
}

func (m *PredictContainer) GetInitialLimitPredictResource() []*MetricData {
	if m != nil {
		return m.InitialLimitPredictResource
	}
	return nil
}

func (m *PredictContainer) GetInitialRequestPredictResource() []*MetricData {
	if m != nil {
		return m.InitialRequestPredictResource
	}
	return nil
}

type PredictPod struct {
	Uid               string              `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	NamespacedName    *NamespacedName     `protobuf:"bytes,2,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	PredictContainers []*PredictContainer `protobuf:"bytes,3,rep,name=predict_containers,json=predictContainers,proto3" json:"predict_containers,omitempty"`
	// Types that are valid to be assigned to AssignPodPolicy:
	//	*PredictPod_NodePriority
	//	*PredictPod_NodeSelector
	//	*PredictPod_NodeName
	AssignPodPolicy      isPredictPod_AssignPodPolicy `protobuf_oneof:"assign_pod_policy"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PredictPod) Reset()         { *m = PredictPod{} }
func (m *PredictPod) String() string { return proto.CompactTextString(m) }
func (*PredictPod) ProtoMessage()    {}
func (*PredictPod) Descriptor() ([]byte, []int) {
	return fileDescriptor_predict_173e05f83af7760f, []int{1}
}
func (m *PredictPod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictPod.Unmarshal(m, b)
}
func (m *PredictPod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictPod.Marshal(b, m, deterministic)
}
func (dst *PredictPod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictPod.Merge(dst, src)
}
func (m *PredictPod) XXX_Size() int {
	return xxx_messageInfo_PredictPod.Size(m)
}
func (m *PredictPod) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictPod.DiscardUnknown(m)
}

var xxx_messageInfo_PredictPod proto.InternalMessageInfo

func (m *PredictPod) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *PredictPod) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *PredictPod) GetPredictContainers() []*PredictContainer {
	if m != nil {
		return m.PredictContainers
	}
	return nil
}

type isPredictPod_AssignPodPolicy interface {
	isPredictPod_AssignPodPolicy()
}

type PredictPod_NodePriority struct {
	NodePriority *NodePriority `protobuf:"bytes,4,opt,name=node_priority,json=nodePriority,proto3,oneof"`
}

type PredictPod_NodeSelector struct {
	NodeSelector *Selector `protobuf:"bytes,5,opt,name=node_selector,json=nodeSelector,proto3,oneof"`
}

type PredictPod_NodeName struct {
	NodeName string `protobuf:"bytes,6,opt,name=node_name,json=nodeName,proto3,oneof"`
}

func (*PredictPod_NodePriority) isPredictPod_AssignPodPolicy() {}

func (*PredictPod_NodeSelector) isPredictPod_AssignPodPolicy() {}

func (*PredictPod_NodeName) isPredictPod_AssignPodPolicy() {}

func (m *PredictPod) GetAssignPodPolicy() isPredictPod_AssignPodPolicy {
	if m != nil {
		return m.AssignPodPolicy
	}
	return nil
}

func (m *PredictPod) GetNodePriority() *NodePriority {
	if x, ok := m.GetAssignPodPolicy().(*PredictPod_NodePriority); ok {
		return x.NodePriority
	}
	return nil
}

func (m *PredictPod) GetNodeSelector() *Selector {
	if x, ok := m.GetAssignPodPolicy().(*PredictPod_NodeSelector); ok {
		return x.NodeSelector
	}
	return nil
}

func (m *PredictPod) GetNodeName() string {
	if x, ok := m.GetAssignPodPolicy().(*PredictPod_NodeName); ok {
		return x.NodeName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PredictPod) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PredictPod_OneofMarshaler, _PredictPod_OneofUnmarshaler, _PredictPod_OneofSizer, []interface{}{
		(*PredictPod_NodePriority)(nil),
		(*PredictPod_NodeSelector)(nil),
		(*PredictPod_NodeName)(nil),
	}
}

func _PredictPod_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PredictPod)
	// assign_pod_policy
	switch x := m.AssignPodPolicy.(type) {
	case *PredictPod_NodePriority:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodePriority); err != nil {
			return err
		}
	case *PredictPod_NodeSelector:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodeSelector); err != nil {
			return err
		}
	case *PredictPod_NodeName:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NodeName)
	case nil:
	default:
		return fmt.Errorf("PredictPod.AssignPodPolicy has unexpected type %T", x)
	}
	return nil
}

func _PredictPod_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PredictPod)
	switch tag {
	case 4: // assign_pod_policy.node_priority
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NodePriority)
		err := b.DecodeMessage(msg)
		m.AssignPodPolicy = &PredictPod_NodePriority{msg}
		return true, err
	case 5: // assign_pod_policy.node_selector
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Selector)
		err := b.DecodeMessage(msg)
		m.AssignPodPolicy = &PredictPod_NodeSelector{msg}
		return true, err
	case 6: // assign_pod_policy.node_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.AssignPodPolicy = &PredictPod_NodeName{x}
		return true, err
	default:
		return false, nil
	}
}

func _PredictPod_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PredictPod)
	// assign_pod_policy
	switch x := m.AssignPodPolicy.(type) {
	case *PredictPod_NodePriority:
		s := proto.Size(x.NodePriority)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PredictPod_NodeSelector:
		s := proto.Size(x.NodeSelector)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PredictPod_NodeName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.NodeName)))
		n += len(x.NodeName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PredictNode struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RowPredictData       []*MetricData `protobuf:"bytes,2,rep,name=row_predict_data,json=rowPredictData,proto3" json:"row_predict_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PredictNode) Reset()         { *m = PredictNode{} }
func (m *PredictNode) String() string { return proto.CompactTextString(m) }
func (*PredictNode) ProtoMessage()    {}
func (*PredictNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_predict_173e05f83af7760f, []int{2}
}
func (m *PredictNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictNode.Unmarshal(m, b)
}
func (m *PredictNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictNode.Marshal(b, m, deterministic)
}
func (dst *PredictNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictNode.Merge(dst, src)
}
func (m *PredictNode) XXX_Size() int {
	return xxx_messageInfo_PredictNode.Size(m)
}
func (m *PredictNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictNode.DiscardUnknown(m)
}

var xxx_messageInfo_PredictNode proto.InternalMessageInfo

func (m *PredictNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PredictNode) GetRowPredictData() []*MetricData {
	if m != nil {
		return m.RowPredictData
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictContainer)(nil), "containers_ai.alameda.v1alpha1.datahub.PredictContainer")
	proto.RegisterType((*PredictPod)(nil), "containers_ai.alameda.v1alpha1.datahub.PredictPod")
	proto.RegisterType((*PredictNode)(nil), "containers_ai.alameda.v1alpha1.datahub.PredictNode")
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.RecommendationPolicy", RecommendationPolicy_name, RecommendationPolicy_value)
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/predict.proto", fileDescriptor_predict_173e05f83af7760f)
}

var fileDescriptor_predict_173e05f83af7760f = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xfb, 0x6f, 0x85, 0xbd, 0x85, 0x91, 0x99, 0x1d, 0xaa, 0xa1, 0x4d, 0x55, 0x0f, 0xa8,
	0x0c, 0x29, 0xa5, 0x05, 0x21, 0xae, 0x5d, 0x5b, 0xb4, 0x4a, 0xfd, 0x13, 0xa5, 0x45, 0x08, 0x81,
	0x64, 0xbc, 0xd8, 0xda, 0x2c, 0x25, 0x76, 0x70, 0x5c, 0xaa, 0x89, 0x03, 0xdf, 0x8d, 0x8f, 0xc0,
	0x27, 0x42, 0x71, 0x9d, 0xd1, 0xf6, 0x50, 0x55, 0x3d, 0xec, 0x66, 0x5b, 0x79, 0x7f, 0xcf, 0xf3,
	0xd8, 0x6f, 0x5e, 0xb8, 0x20, 0x21, 0x89, 0x18, 0x25, 0x98, 0xc4, 0xbc, 0xf9, 0xb3, 0x45, 0xc2,
	0xf8, 0x96, 0xb4, 0x9a, 0x94, 0x68, 0x72, 0x3b, 0xbf, 0x6e, 0xc6, 0x8a, 0x51, 0x1e, 0x68, 0x37,
	0x56, 0x52, 0x4b, 0xf4, 0x32, 0x90, 0x42, 0x13, 0x2e, 0x98, 0x4a, 0x30, 0xe1, 0xae, 0xad, 0x74,
	0xb3, 0x2a, 0xd7, 0x56, 0x9d, 0xbe, 0xde, 0xca, 0x54, 0x2c, 0x91, 0x73, 0x15, 0xb0, 0x25, 0xf4,
	0xf4, 0xd5, 0xd6, 0x8f, 0x23, 0xa6, 0x15, 0x0f, 0xec, 0xa7, 0xad, 0xed, 0x5e, 0x25, 0xc5, 0x24,
	0x49, 0xf8, 0x8d, 0x88, 0x98, 0xb0, 0x96, 0xeb, 0x7f, 0x4a, 0xe0, 0x78, 0xcb, 0x10, 0xdd, 0xcc,
	0x3c, 0x42, 0x50, 0x12, 0x24, 0x62, 0xd5, 0x7c, 0x2d, 0xdf, 0x38, 0xf4, 0xcd, 0x1a, 0x7d, 0x03,
	0x47, 0xc9, 0x05, 0xb6, 0x81, 0x71, 0x0a, 0xad, 0x16, 0x6a, 0xc5, 0x46, 0xa5, 0xdd, 0x76, 0x77,
	0x8b, 0xed, 0x8e, 0x8c, 0xd7, 0x1e, 0xd1, 0xc4, 0x3f, 0x52, 0x72, 0x61, 0x65, 0xd3, 0x3d, 0xfa,
	0x0e, 0x28, 0xe4, 0x11, 0xd7, 0xeb, 0xfc, 0xe2, 0xde, 0x7c, 0xc7, 0xd0, 0x56, 0x15, 0x28, 0x9c,
	0x28, 0xf6, 0x63, 0xce, 0x92, 0x0d, 0x8d, 0xd2, 0xde, 0x1a, 0xc8, 0xf2, 0x56, 0x55, 0x16, 0x70,
	0xce, 0x05, 0xd7, 0x9c, 0x84, 0x78, 0x3d, 0x4f, 0xf6, 0xa8, 0xd5, 0x83, 0xbd, 0xf5, 0x5e, 0x58,
	0xf2, 0x70, 0x25, 0x9a, 0x6f, 0xb1, 0xe8, 0x17, 0xd4, 0x32, 0xe1, 0xcd, 0x98, 0xf7, 0xd2, 0xe5,
	0xbd, 0xa5, 0xcf, 0x2c, 0xdb, 0x5f, 0x4b, 0x9c, 0x89, 0xd7, 0xff, 0x16, 0x01, 0xec, 0x99, 0x27,
	0x29, 0x72, 0xa0, 0x38, 0xe7, 0xd4, 0x76, 0x4f, 0xba, 0x44, 0x18, 0x9e, 0xa5, 0x4d, 0x94, 0xc4,
	0x24, 0x60, 0x14, 0x9b, 0xde, 0x2a, 0xd4, 0xf2, 0x8d, 0x4a, 0xfb, 0xfd, 0xae, 0x66, 0xc6, 0xf7,
	0xe5, 0xe9, 0xca, 0x3f, 0x12, 0x6b, 0x7b, 0x74, 0x03, 0x28, 0x8b, 0xfb, 0x1f, 0x68, 0xfb, 0xe7,
	0xc3, 0xae, 0x1a, 0x9b, 0xff, 0x81, 0x7f, 0x1c, 0x6f, 0x9c, 0x24, 0xe8, 0x2b, 0x3c, 0x15, 0x92,
	0x32, 0x1c, 0x2b, 0x2e, 0x15, 0xd7, 0x77, 0xd5, 0x92, 0xc9, 0xf1, 0x6e, 0xe7, 0x1c, 0x92, 0x32,
	0xcf, 0xd6, 0x5e, 0xe5, 0xfc, 0x27, 0x62, 0x65, 0x8f, 0x3e, 0x5b, 0x78, 0xc2, 0x42, 0x16, 0x68,
	0xa9, 0xaa, 0x07, 0x06, 0xfe, 0x66, 0x57, 0xf8, 0xd4, 0xd6, 0x65, 0xe0, 0x6c, 0x8f, 0xce, 0xe0,
	0xd0, 0x80, 0xcd, 0xcd, 0x97, 0xd3, 0x77, 0xb9, 0xca, 0xf9, 0x8f, 0xd3, 0xa3, 0xf4, 0xf6, 0x2e,
	0x9f, 0xc3, 0xf1, 0x72, 0x30, 0xe0, 0x74, 0x46, 0xc4, 0x32, 0xe4, 0xc1, 0x5d, 0xfd, 0x37, 0x54,
	0xec, 0x85, 0xa4, 0x9e, 0x1f, 0x7e, 0x26, 0x5c, 0x4c, 0xe1, 0xc4, 0x67, 0x81, 0x8c, 0x22, 0x26,
	0x28, 0xd1, 0x5c, 0x0a, 0xcf, 0x18, 0x43, 0x75, 0x38, 0xf7, 0xfb, 0xdd, 0xc9, 0x68, 0xd4, 0x1f,
	0xf7, 0x3a, 0xb3, 0xc1, 0x64, 0xec, 0x4d, 0x86, 0x83, 0xee, 0x17, 0xfc, 0x69, 0xdc, 0xeb, 0x7f,
	0x1c, 0x8c, 0xfb, 0x3d, 0x27, 0x87, 0x00, 0xca, 0xd3, 0x59, 0xe7, 0x72, 0xd8, 0x77, 0xf2, 0xa8,
	0x02, 0x8f, 0xba, 0x93, 0x91, 0xd7, 0xe9, 0xce, 0x9c, 0xc2, 0x75, 0xd9, 0x8c, 0xbd, 0xb7, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xe3, 0x28, 0xf7, 0xd7, 0x05, 0x00, 0x00,
}
