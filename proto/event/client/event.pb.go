// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/event/client/event.proto

package client

import (
	fmt "fmt"
	feature "github.com/ca-dp/bucketeer-go-server-sdk/proto/feature"
	user "github.com/ca-dp/bucketeer-go-server-sdk/proto/user"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Event struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event                *any.Any `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	EnvironmentNamespace string   `protobuf:"bytes,3,opt,name=environment_namespace,json=environmentNamespace,proto3" json:"environment_namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetEvent() *any.Any {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetEnvironmentNamespace() string {
	if m != nil {
		return m.EnvironmentNamespace
	}
	return ""
}

type EvaluationEvent struct {
	Timestamp            int64           `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	FeatureId            string          `protobuf:"bytes,2,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	FeatureVersion       int32           `protobuf:"varint,3,opt,name=feature_version,json=featureVersion,proto3" json:"feature_version,omitempty"`
	UserId               string          `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VariationId          string          `protobuf:"bytes,5,opt,name=variation_id,json=variationId,proto3" json:"variation_id,omitempty"`
	User                 *user.User      `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	Reason               *feature.Reason `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EvaluationEvent) Reset()         { *m = EvaluationEvent{} }
func (m *EvaluationEvent) String() string { return proto.CompactTextString(m) }
func (*EvaluationEvent) ProtoMessage()    {}
func (*EvaluationEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{1}
}

func (m *EvaluationEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluationEvent.Unmarshal(m, b)
}
func (m *EvaluationEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluationEvent.Marshal(b, m, deterministic)
}
func (m *EvaluationEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluationEvent.Merge(m, src)
}
func (m *EvaluationEvent) XXX_Size() int {
	return xxx_messageInfo_EvaluationEvent.Size(m)
}
func (m *EvaluationEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluationEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluationEvent proto.InternalMessageInfo

func (m *EvaluationEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *EvaluationEvent) GetFeatureId() string {
	if m != nil {
		return m.FeatureId
	}
	return ""
}

func (m *EvaluationEvent) GetFeatureVersion() int32 {
	if m != nil {
		return m.FeatureVersion
	}
	return 0
}

func (m *EvaluationEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *EvaluationEvent) GetVariationId() string {
	if m != nil {
		return m.VariationId
	}
	return ""
}

func (m *EvaluationEvent) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *EvaluationEvent) GetReason() *feature.Reason {
	if m != nil {
		return m.Reason
	}
	return nil
}

type GoalEvent struct {
	Timestamp            int64                 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GoalId               string                `protobuf:"bytes,2,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
	UserId               string                `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Value                float64               `protobuf:"fixed64,4,opt,name=value,proto3" json:"value,omitempty"`
	User                 *user.User            `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Evaluations          []*feature.Evaluation `protobuf:"bytes,6,rep,name=evaluations,proto3" json:"evaluations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GoalEvent) Reset()         { *m = GoalEvent{} }
func (m *GoalEvent) String() string { return proto.CompactTextString(m) }
func (*GoalEvent) ProtoMessage()    {}
func (*GoalEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{2}
}

func (m *GoalEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoalEvent.Unmarshal(m, b)
}
func (m *GoalEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoalEvent.Marshal(b, m, deterministic)
}
func (m *GoalEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoalEvent.Merge(m, src)
}
func (m *GoalEvent) XXX_Size() int {
	return xxx_messageInfo_GoalEvent.Size(m)
}
func (m *GoalEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GoalEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GoalEvent proto.InternalMessageInfo

func (m *GoalEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *GoalEvent) GetGoalId() string {
	if m != nil {
		return m.GoalId
	}
	return ""
}

func (m *GoalEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GoalEvent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *GoalEvent) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GoalEvent) GetEvaluations() []*feature.Evaluation {
	if m != nil {
		return m.Evaluations
	}
	return nil
}

type ExperimentEvent struct {
	Timestamp            int64      `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ExperimentId         string     `protobuf:"bytes,2,opt,name=experiment_id,json=experimentId,proto3" json:"experiment_id,omitempty"`
	FeatureId            string     `protobuf:"bytes,3,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	FeatureVersion       int32      `protobuf:"varint,4,opt,name=feature_version,json=featureVersion,proto3" json:"feature_version,omitempty"`
	GoalId               string     `protobuf:"bytes,5,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
	UserId               string     `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VariationId          string     `protobuf:"bytes,7,opt,name=variation_id,json=variationId,proto3" json:"variation_id,omitempty"`
	Value                float64    `protobuf:"fixed64,8,opt,name=value,proto3" json:"value,omitempty"`
	User                 *user.User `protobuf:"bytes,9,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ExperimentEvent) Reset()         { *m = ExperimentEvent{} }
func (m *ExperimentEvent) String() string { return proto.CompactTextString(m) }
func (*ExperimentEvent) ProtoMessage()    {}
func (*ExperimentEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{3}
}

func (m *ExperimentEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExperimentEvent.Unmarshal(m, b)
}
func (m *ExperimentEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExperimentEvent.Marshal(b, m, deterministic)
}
func (m *ExperimentEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExperimentEvent.Merge(m, src)
}
func (m *ExperimentEvent) XXX_Size() int {
	return xxx_messageInfo_ExperimentEvent.Size(m)
}
func (m *ExperimentEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ExperimentEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ExperimentEvent proto.InternalMessageInfo

func (m *ExperimentEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ExperimentEvent) GetExperimentId() string {
	if m != nil {
		return m.ExperimentId
	}
	return ""
}

func (m *ExperimentEvent) GetFeatureId() string {
	if m != nil {
		return m.FeatureId
	}
	return ""
}

func (m *ExperimentEvent) GetFeatureVersion() int32 {
	if m != nil {
		return m.FeatureVersion
	}
	return 0
}

func (m *ExperimentEvent) GetGoalId() string {
	if m != nil {
		return m.GoalId
	}
	return ""
}

func (m *ExperimentEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ExperimentEvent) GetVariationId() string {
	if m != nil {
		return m.VariationId
	}
	return ""
}

func (m *ExperimentEvent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *ExperimentEvent) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

type MetricsEvent struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Event                *any.Any `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsEvent) Reset()         { *m = MetricsEvent{} }
func (m *MetricsEvent) String() string { return proto.CompactTextString(m) }
func (*MetricsEvent) ProtoMessage()    {}
func (*MetricsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{4}
}

func (m *MetricsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsEvent.Unmarshal(m, b)
}
func (m *MetricsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsEvent.Marshal(b, m, deterministic)
}
func (m *MetricsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsEvent.Merge(m, src)
}
func (m *MetricsEvent) XXX_Size() int {
	return xxx_messageInfo_MetricsEvent.Size(m)
}
func (m *MetricsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsEvent proto.InternalMessageInfo

func (m *MetricsEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MetricsEvent) GetEvent() *any.Any {
	if m != nil {
		return m.Event
	}
	return nil
}

type GetEvaluationLatencyMetricsEvent struct {
	Labels               map[string]string  `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Duration             *duration.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetEvaluationLatencyMetricsEvent) Reset()         { *m = GetEvaluationLatencyMetricsEvent{} }
func (m *GetEvaluationLatencyMetricsEvent) String() string { return proto.CompactTextString(m) }
func (*GetEvaluationLatencyMetricsEvent) ProtoMessage()    {}
func (*GetEvaluationLatencyMetricsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{5}
}

func (m *GetEvaluationLatencyMetricsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEvaluationLatencyMetricsEvent.Unmarshal(m, b)
}
func (m *GetEvaluationLatencyMetricsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEvaluationLatencyMetricsEvent.Marshal(b, m, deterministic)
}
func (m *GetEvaluationLatencyMetricsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEvaluationLatencyMetricsEvent.Merge(m, src)
}
func (m *GetEvaluationLatencyMetricsEvent) XXX_Size() int {
	return xxx_messageInfo_GetEvaluationLatencyMetricsEvent.Size(m)
}
func (m *GetEvaluationLatencyMetricsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEvaluationLatencyMetricsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GetEvaluationLatencyMetricsEvent proto.InternalMessageInfo

func (m *GetEvaluationLatencyMetricsEvent) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *GetEvaluationLatencyMetricsEvent) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type GetEvaluationSizeMetricsEvent struct {
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SizeByte             int32             `protobuf:"varint,2,opt,name=size_byte,json=sizeByte,proto3" json:"size_byte,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetEvaluationSizeMetricsEvent) Reset()         { *m = GetEvaluationSizeMetricsEvent{} }
func (m *GetEvaluationSizeMetricsEvent) String() string { return proto.CompactTextString(m) }
func (*GetEvaluationSizeMetricsEvent) ProtoMessage()    {}
func (*GetEvaluationSizeMetricsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{6}
}

func (m *GetEvaluationSizeMetricsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEvaluationSizeMetricsEvent.Unmarshal(m, b)
}
func (m *GetEvaluationSizeMetricsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEvaluationSizeMetricsEvent.Marshal(b, m, deterministic)
}
func (m *GetEvaluationSizeMetricsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEvaluationSizeMetricsEvent.Merge(m, src)
}
func (m *GetEvaluationSizeMetricsEvent) XXX_Size() int {
	return xxx_messageInfo_GetEvaluationSizeMetricsEvent.Size(m)
}
func (m *GetEvaluationSizeMetricsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEvaluationSizeMetricsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GetEvaluationSizeMetricsEvent proto.InternalMessageInfo

func (m *GetEvaluationSizeMetricsEvent) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *GetEvaluationSizeMetricsEvent) GetSizeByte() int32 {
	if m != nil {
		return m.SizeByte
	}
	return 0
}

type TimeoutErrorCountMetricsEvent struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeoutErrorCountMetricsEvent) Reset()         { *m = TimeoutErrorCountMetricsEvent{} }
func (m *TimeoutErrorCountMetricsEvent) String() string { return proto.CompactTextString(m) }
func (*TimeoutErrorCountMetricsEvent) ProtoMessage()    {}
func (*TimeoutErrorCountMetricsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{7}
}

func (m *TimeoutErrorCountMetricsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeoutErrorCountMetricsEvent.Unmarshal(m, b)
}
func (m *TimeoutErrorCountMetricsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeoutErrorCountMetricsEvent.Marshal(b, m, deterministic)
}
func (m *TimeoutErrorCountMetricsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeoutErrorCountMetricsEvent.Merge(m, src)
}
func (m *TimeoutErrorCountMetricsEvent) XXX_Size() int {
	return xxx_messageInfo_TimeoutErrorCountMetricsEvent.Size(m)
}
func (m *TimeoutErrorCountMetricsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeoutErrorCountMetricsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TimeoutErrorCountMetricsEvent proto.InternalMessageInfo

func (m *TimeoutErrorCountMetricsEvent) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type InternalErrorCountMetricsEvent struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalErrorCountMetricsEvent) Reset()         { *m = InternalErrorCountMetricsEvent{} }
func (m *InternalErrorCountMetricsEvent) String() string { return proto.CompactTextString(m) }
func (*InternalErrorCountMetricsEvent) ProtoMessage()    {}
func (*InternalErrorCountMetricsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{8}
}

func (m *InternalErrorCountMetricsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalErrorCountMetricsEvent.Unmarshal(m, b)
}
func (m *InternalErrorCountMetricsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalErrorCountMetricsEvent.Marshal(b, m, deterministic)
}
func (m *InternalErrorCountMetricsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalErrorCountMetricsEvent.Merge(m, src)
}
func (m *InternalErrorCountMetricsEvent) XXX_Size() int {
	return xxx_messageInfo_InternalErrorCountMetricsEvent.Size(m)
}
func (m *InternalErrorCountMetricsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalErrorCountMetricsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_InternalErrorCountMetricsEvent proto.InternalMessageInfo

func (m *InternalErrorCountMetricsEvent) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type OpsEvent struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	FeatureId            string   `protobuf:"bytes,2,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	FeatureVersion       int32    `protobuf:"varint,3,opt,name=feature_version,json=featureVersion,proto3" json:"feature_version,omitempty"`
	VariationId          string   `protobuf:"bytes,4,opt,name=variation_id,json=variationId,proto3" json:"variation_id,omitempty"`
	GoalId               string   `protobuf:"bytes,5,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
	UserId               string   `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpsEvent) Reset()         { *m = OpsEvent{} }
func (m *OpsEvent) String() string { return proto.CompactTextString(m) }
func (*OpsEvent) ProtoMessage()    {}
func (*OpsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{9}
}

func (m *OpsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpsEvent.Unmarshal(m, b)
}
func (m *OpsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpsEvent.Marshal(b, m, deterministic)
}
func (m *OpsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpsEvent.Merge(m, src)
}
func (m *OpsEvent) XXX_Size() int {
	return xxx_messageInfo_OpsEvent.Size(m)
}
func (m *OpsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_OpsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_OpsEvent proto.InternalMessageInfo

func (m *OpsEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *OpsEvent) GetFeatureId() string {
	if m != nil {
		return m.FeatureId
	}
	return ""
}

func (m *OpsEvent) GetFeatureVersion() int32 {
	if m != nil {
		return m.FeatureVersion
	}
	return 0
}

func (m *OpsEvent) GetVariationId() string {
	if m != nil {
		return m.VariationId
	}
	return ""
}

func (m *OpsEvent) GetGoalId() string {
	if m != nil {
		return m.GoalId
	}
	return ""
}

func (m *OpsEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GoalBatchEvent struct {
	UserId                 string                   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserGoalEventsOverTags []*UserGoalEventsOverTag `protobuf:"bytes,2,rep,name=user_goal_events_over_tags,json=userGoalEventsOverTags,proto3" json:"user_goal_events_over_tags,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                 `json:"-"`
	XXX_unrecognized       []byte                   `json:"-"`
	XXX_sizecache          int32                    `json:"-"`
}

func (m *GoalBatchEvent) Reset()         { *m = GoalBatchEvent{} }
func (m *GoalBatchEvent) String() string { return proto.CompactTextString(m) }
func (*GoalBatchEvent) ProtoMessage()    {}
func (*GoalBatchEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{10}
}

func (m *GoalBatchEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoalBatchEvent.Unmarshal(m, b)
}
func (m *GoalBatchEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoalBatchEvent.Marshal(b, m, deterministic)
}
func (m *GoalBatchEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoalBatchEvent.Merge(m, src)
}
func (m *GoalBatchEvent) XXX_Size() int {
	return xxx_messageInfo_GoalBatchEvent.Size(m)
}
func (m *GoalBatchEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GoalBatchEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GoalBatchEvent proto.InternalMessageInfo

func (m *GoalBatchEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GoalBatchEvent) GetUserGoalEventsOverTags() []*UserGoalEventsOverTag {
	if m != nil {
		return m.UserGoalEventsOverTags
	}
	return nil
}

type UserGoalEventsOverTag struct {
	Tag                  string           `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	UserGoalEvents       []*UserGoalEvent `protobuf:"bytes,2,rep,name=user_goal_events,json=userGoalEvents,proto3" json:"user_goal_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UserGoalEventsOverTag) Reset()         { *m = UserGoalEventsOverTag{} }
func (m *UserGoalEventsOverTag) String() string { return proto.CompactTextString(m) }
func (*UserGoalEventsOverTag) ProtoMessage()    {}
func (*UserGoalEventsOverTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{11}
}

func (m *UserGoalEventsOverTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGoalEventsOverTag.Unmarshal(m, b)
}
func (m *UserGoalEventsOverTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGoalEventsOverTag.Marshal(b, m, deterministic)
}
func (m *UserGoalEventsOverTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGoalEventsOverTag.Merge(m, src)
}
func (m *UserGoalEventsOverTag) XXX_Size() int {
	return xxx_messageInfo_UserGoalEventsOverTag.Size(m)
}
func (m *UserGoalEventsOverTag) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGoalEventsOverTag.DiscardUnknown(m)
}

var xxx_messageInfo_UserGoalEventsOverTag proto.InternalMessageInfo

func (m *UserGoalEventsOverTag) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *UserGoalEventsOverTag) GetUserGoalEvents() []*UserGoalEvent {
	if m != nil {
		return m.UserGoalEvents
	}
	return nil
}

type UserGoalEvent struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GoalId               string   `protobuf:"bytes,2,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
	Value                float64  `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGoalEvent) Reset()         { *m = UserGoalEvent{} }
func (m *UserGoalEvent) String() string { return proto.CompactTextString(m) }
func (*UserGoalEvent) ProtoMessage()    {}
func (*UserGoalEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead1ae0565f33ac0, []int{12}
}

func (m *UserGoalEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGoalEvent.Unmarshal(m, b)
}
func (m *UserGoalEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGoalEvent.Marshal(b, m, deterministic)
}
func (m *UserGoalEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGoalEvent.Merge(m, src)
}
func (m *UserGoalEvent) XXX_Size() int {
	return xxx_messageInfo_UserGoalEvent.Size(m)
}
func (m *UserGoalEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGoalEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UserGoalEvent proto.InternalMessageInfo

func (m *UserGoalEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *UserGoalEvent) GetGoalId() string {
	if m != nil {
		return m.GoalId
	}
	return ""
}

func (m *UserGoalEvent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "bucketeer.event.client.Event")
	proto.RegisterType((*EvaluationEvent)(nil), "bucketeer.event.client.EvaluationEvent")
	proto.RegisterType((*GoalEvent)(nil), "bucketeer.event.client.GoalEvent")
	proto.RegisterType((*ExperimentEvent)(nil), "bucketeer.event.client.ExperimentEvent")
	proto.RegisterType((*MetricsEvent)(nil), "bucketeer.event.client.MetricsEvent")
	proto.RegisterType((*GetEvaluationLatencyMetricsEvent)(nil), "bucketeer.event.client.GetEvaluationLatencyMetricsEvent")
	proto.RegisterMapType((map[string]string)(nil), "bucketeer.event.client.GetEvaluationLatencyMetricsEvent.LabelsEntry")
	proto.RegisterType((*GetEvaluationSizeMetricsEvent)(nil), "bucketeer.event.client.GetEvaluationSizeMetricsEvent")
	proto.RegisterMapType((map[string]string)(nil), "bucketeer.event.client.GetEvaluationSizeMetricsEvent.LabelsEntry")
	proto.RegisterType((*TimeoutErrorCountMetricsEvent)(nil), "bucketeer.event.client.TimeoutErrorCountMetricsEvent")
	proto.RegisterType((*InternalErrorCountMetricsEvent)(nil), "bucketeer.event.client.InternalErrorCountMetricsEvent")
	proto.RegisterType((*OpsEvent)(nil), "bucketeer.event.client.OpsEvent")
	proto.RegisterType((*GoalBatchEvent)(nil), "bucketeer.event.client.GoalBatchEvent")
	proto.RegisterType((*UserGoalEventsOverTag)(nil), "bucketeer.event.client.UserGoalEventsOverTag")
	proto.RegisterType((*UserGoalEvent)(nil), "bucketeer.event.client.UserGoalEvent")
}

func init() { proto.RegisterFile("proto/event/client/event.proto", fileDescriptor_ead1ae0565f33ac0) }

var fileDescriptor_ead1ae0565f33ac0 = []byte{
	// 816 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0xe3, 0xd8, 0x49, 0x4e, 0xba, 0xd9, 0xca, 0x6a, 0x77, 0xdd, 0x40, 0xaa, 0x60, 0x84,
	0x88, 0x90, 0x6a, 0x6b, 0xbb, 0x42, 0x02, 0x6e, 0xd0, 0x96, 0x8d, 0x56, 0x91, 0x16, 0x2a, 0x99,
	0x05, 0x01, 0x5a, 0x29, 0x9a, 0xd8, 0x67, 0xdd, 0x51, 0x1d, 0x3b, 0x1a, 0x8f, 0xa3, 0xa6, 0x4f,
	0xc1, 0x05, 0x2f, 0xc2, 0x43, 0xf0, 0x02, 0x5c, 0xf3, 0x14, 0x3c, 0x01, 0xf2, 0x8c, 0xe3, 0x9f,
	0x24, 0xa5, 0x29, 0x48, 0x7b, 0xd3, 0x66, 0xce, 0xf9, 0xe6, 0xcc, 0xf7, 0x7d, 0xe7, 0xcc, 0x24,
	0x70, 0xba, 0x60, 0x31, 0x8f, 0x1d, 0x5c, 0x62, 0xc4, 0x1d, 0x2f, 0xa4, 0xd9, 0x3f, 0xb1, 0xb0,
	0x45, 0xc2, 0x78, 0x32, 0x4b, 0xbd, 0x6b, 0xe4, 0x88, 0xcc, 0x96, 0x61, 0x89, 0xe9, 0x9f, 0x04,
	0x71, 0x1c, 0x84, 0xe8, 0x08, 0xd4, 0x2c, 0x7d, 0xe7, 0x90, 0x68, 0x25, 0xb7, 0xf4, 0x4f, 0x37,
	0x53, 0x7e, 0xca, 0x08, 0xa7, 0x71, 0xb4, 0xce, 0xcb, 0x23, 0xdf, 0x21, 0xe1, 0x29, 0x43, 0x07,
	0x97, 0x24, 0x4c, 0xab, 0xf9, 0x7e, 0x3d, 0xcf, 0x90, 0x24, 0x45, 0xee, 0x58, 0xe6, 0xd2, 0x04,
	0x99, 0xf8, 0x23, 0xc3, 0xd6, 0x0d, 0x68, 0xe3, 0x8c, 0x9d, 0xd1, 0x83, 0x06, 0xf5, 0x4d, 0x65,
	0xa8, 0x8c, 0x3a, 0x6e, 0x83, 0xfa, 0xc6, 0x67, 0xa0, 0x09, 0xda, 0x66, 0x63, 0xa8, 0x8c, 0xba,
	0xe7, 0x47, 0xb6, 0xe4, 0x66, 0xaf, 0xb9, 0xd9, 0x2f, 0xa2, 0x95, 0x2b, 0x21, 0xc6, 0x73, 0x38,
	0xc6, 0x68, 0x49, 0x59, 0x1c, 0xcd, 0x31, 0xe2, 0xd3, 0x88, 0xcc, 0x31, 0x59, 0x10, 0x0f, 0x4d,
	0x55, 0x94, 0x3b, 0xaa, 0x24, 0xbf, 0x5b, 0xe7, 0xac, 0x5f, 0x1b, 0xf0, 0x78, 0x5c, 0x28, 0x90,
	0x24, 0x3e, 0x84, 0x0e, 0xa7, 0x73, 0x4c, 0x38, 0x99, 0x2f, 0x04, 0x17, 0xd5, 0x2d, 0x03, 0xc6,
	0x00, 0x20, 0x97, 0x36, 0xa5, 0xbe, 0xe0, 0xd5, 0x71, 0x3b, 0x79, 0x64, 0xe2, 0x1b, 0x9f, 0xc2,
	0xe3, 0x75, 0x7a, 0x89, 0x2c, 0xa1, 0x71, 0x24, 0xce, 0xd7, 0xdc, 0x5e, 0x1e, 0xfe, 0x51, 0x46,
	0x8d, 0xa7, 0xd0, 0xca, 0x1c, 0xc8, 0x8a, 0x34, 0x45, 0x11, 0x3d, 0x5b, 0x4e, 0x7c, 0xe3, 0x23,
	0x38, 0x58, 0x12, 0x46, 0x05, 0xa1, 0x2c, 0xab, 0x89, 0x6c, 0xb7, 0x88, 0x4d, 0x7c, 0x63, 0x04,
	0xcd, 0x0c, 0x6c, 0xea, 0xb9, 0x2b, 0x65, 0x93, 0x85, 0xa9, 0x3f, 0x24, 0xc8, 0x5c, 0x81, 0x30,
	0x9e, 0x81, 0x2e, 0x1b, 0x60, 0xb6, 0x04, 0xf6, 0xa4, 0x82, 0xcd, 0x09, 0xd9, 0xae, 0x00, 0xb8,
	0x39, 0xd0, 0xfa, 0x4b, 0x81, 0xce, 0xab, 0x98, 0x84, 0xfb, 0x98, 0xf1, 0x14, 0x5a, 0x41, 0x4c,
	0xc2, 0xd2, 0x09, 0x3d, 0x5b, 0x4e, 0xfc, 0xaa, 0x3a, 0xb5, 0xa6, 0xee, 0x08, 0xb4, 0xcc, 0x6e,
	0x14, 0xa2, 0x15, 0x57, 0x2e, 0x0a, 0x41, 0xda, 0xbd, 0x82, 0xbe, 0x86, 0x6e, 0x39, 0x71, 0x89,
	0xa9, 0x0f, 0xd5, 0x51, 0xf7, 0x7c, 0xb0, 0x43, 0x55, 0xd9, 0x55, 0xb7, 0xba, 0xc3, 0xfa, 0x3d,
	0xeb, 0xf8, 0xcd, 0x02, 0x19, 0xcd, 0x26, 0x61, 0x1f, 0x91, 0x1f, 0xc3, 0x23, 0x2c, 0x36, 0x94,
	0x52, 0x0f, 0xca, 0xe0, 0xc4, 0xdf, 0x18, 0x0b, 0x75, 0x8f, 0xb1, 0x68, 0xde, 0x35, 0x16, 0x6b,
	0x47, 0xb5, 0xbb, 0x1c, 0xd5, 0xff, 0x75, 0x5e, 0x5a, 0xdb, 0xf3, 0x52, 0x98, 0xde, 0xde, 0x65,
	0x7a, 0xe7, 0x3e, 0xd3, 0xad, 0x9f, 0xe0, 0xe0, 0x5b, 0xe4, 0x8c, 0x7a, 0xc9, 0x3e, 0x7e, 0x3d,
	0xe0, 0xd2, 0x5a, 0x7f, 0x2b, 0x30, 0x7c, 0x85, 0xbc, 0x6c, 0xd6, 0x6b, 0xc2, 0x31, 0xf2, 0x56,
	0xb5, 0xe3, 0xde, 0x82, 0x1e, 0x92, 0x19, 0x86, 0x89, 0xa9, 0x88, 0x76, 0xbf, 0xb4, 0x77, 0xbf,
	0x6a, 0xf6, 0x7d, 0x95, 0xec, 0xd7, 0xa2, 0xcc, 0x38, 0xe2, 0x6c, 0xe5, 0xe6, 0x35, 0x8d, 0xcf,
	0xa1, 0xbd, 0x7e, 0xe1, 0x72, 0xc6, 0x27, 0x5b, 0x8c, 0x5f, 0xe6, 0x00, 0xb7, 0x80, 0xf6, 0xbf,
	0x84, 0x6e, 0xa5, 0x9a, 0x71, 0x08, 0xea, 0x35, 0xae, 0xf2, 0xa7, 0x2b, 0xfb, 0x58, 0x9a, 0x2e,
	0xc7, 0x45, 0x2e, 0xbe, 0x6a, 0x7c, 0xa1, 0x58, 0x7f, 0x2a, 0x30, 0xa8, 0x51, 0xfd, 0x9e, 0xde,
	0x62, 0x4d, 0xf1, 0xcf, 0x1b, 0x8a, 0x5f, 0xec, 0xa5, 0x78, 0xb3, 0xcc, 0x4e, 0xb9, 0x1f, 0x40,
	0x27, 0xa1, 0xb7, 0x38, 0x9d, 0xad, 0xb8, 0xa4, 0xa6, 0xb9, 0xed, 0x2c, 0x70, 0xb1, 0xe2, 0xf8,
	0x7f, 0x44, 0x3d, 0x83, 0xc1, 0x1b, 0x3a, 0xc7, 0x38, 0xe5, 0x63, 0xc6, 0x62, 0xf6, 0x4d, 0x9c,
	0x46, 0xbc, 0xa6, 0xe9, 0x10, 0x54, 0x4e, 0x82, 0x75, 0x31, 0x4e, 0x02, 0xeb, 0x1c, 0x4e, 0x27,
	0x11, 0x47, 0x16, 0x91, 0x70, 0xef, 0x3d, 0x7f, 0x28, 0xd0, 0xbe, 0x5c, 0x24, 0xef, 0xf3, 0xa5,
	0xde, 0xbc, 0x60, 0xcd, 0xed, 0x0b, 0xf6, 0xe0, 0x5b, 0x6b, 0xfd, 0xa6, 0x40, 0x2f, 0x7b, 0x65,
	0x2f, 0x08, 0xf7, 0xae, 0xa4, 0x9a, 0x0a, 0x56, 0xa9, 0xdd, 0x70, 0x0a, 0x7d, 0x91, 0x10, 0x47,
	0x88, 0xf6, 0x27, 0xd3, 0x78, 0x89, 0x6c, 0xca, 0x49, 0x90, 0x98, 0x0d, 0x31, 0x21, 0x67, 0x77,
	0x4d, 0x48, 0x76, 0x8d, 0x8b, 0xe7, 0x3c, 0xb9, 0x5c, 0x22, 0x7b, 0x43, 0x02, 0xf7, 0x49, 0xba,
	0x2b, 0x9c, 0x58, 0xb7, 0x70, 0xbc, 0x73, 0xc3, 0x76, 0x27, 0x8c, 0x4b, 0x38, 0xdc, 0x64, 0x95,
	0x73, 0xf9, 0x64, 0x2f, 0x2e, 0x6e, 0xaf, 0xce, 0xc1, 0x7a, 0x0b, 0x8f, 0x6a, 0x80, 0xff, 0xfa,
	0xdd, 0x53, 0xcc, 0xa8, 0x5a, 0x79, 0xed, 0x2e, 0x9c, 0x5f, 0xce, 0x02, 0xca, 0xaf, 0xd2, 0x99,
	0xed, 0xc5, 0x73, 0xc7, 0x23, 0x67, 0xfe, 0xc2, 0x29, 0x68, 0x3a, 0xdb, 0x3f, 0xa3, 0x66, 0xba,
	0x88, 0x3d, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0x39, 0xc7, 0xe0, 0x68, 0x63, 0x09, 0x00, 0x00,
}
