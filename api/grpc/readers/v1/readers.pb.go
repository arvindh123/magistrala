// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: readers/v1/readers.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Aggregation defines supported data aggregations.
type Aggregation int32

const (
	Aggregation_AGGREGATION_UNSPECIFIED Aggregation = 0
	Aggregation_MAX                     Aggregation = 1
	Aggregation_MIN                     Aggregation = 2
	Aggregation_SUM                     Aggregation = 3
	Aggregation_COUNT                   Aggregation = 4
	Aggregation_AVG                     Aggregation = 5
)

// Enum value maps for Aggregation.
var (
	Aggregation_name = map[int32]string{
		0: "AGGREGATION_UNSPECIFIED",
		1: "MAX",
		2: "MIN",
		3: "SUM",
		4: "COUNT",
		5: "AVG",
	}
	Aggregation_value = map[string]int32{
		"AGGREGATION_UNSPECIFIED": 0,
		"MAX":                     1,
		"MIN":                     2,
		"SUM":                     3,
		"COUNT":                   4,
		"AVG":                     5,
	}
)

func (x Aggregation) Enum() *Aggregation {
	p := new(Aggregation)
	*p = x
	return p
}

func (x Aggregation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Aggregation) Descriptor() protoreflect.EnumDescriptor {
	return file_readers_v1_readers_proto_enumTypes[0].Descriptor()
}

func (Aggregation) Type() protoreflect.EnumType {
	return &file_readers_v1_readers_proto_enumTypes[0]
}

func (x Aggregation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Aggregation.Descriptor instead.
func (Aggregation) EnumDescriptor() ([]byte, []int) {
	return file_readers_v1_readers_proto_rawDescGZIP(), []int{0}
}

type PageMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         uint64                 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        uint64                 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Protocol      string                 `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Value         float64                `protobuf:"fixed64,5,opt,name=value,proto3" json:"value,omitempty"`
	Publisher     string                 `protobuf:"bytes,6,opt,name=publisher,proto3" json:"publisher,omitempty"`
	BoolValue     bool                   `protobuf:"varint,7,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue   string                 `protobuf:"bytes,8,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	DataValue     string                 `protobuf:"bytes,9,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
	From          float64                `protobuf:"fixed64,10,opt,name=from,proto3" json:"from,omitempty"`
	To            float64                `protobuf:"fixed64,11,opt,name=to,proto3" json:"to,omitempty"`
	Subtopic      string                 `protobuf:"bytes,12,opt,name=subtopic,proto3" json:"subtopic,omitempty"`
	Interval      string                 `protobuf:"bytes,13,opt,name=interval,proto3" json:"interval,omitempty"`
	Read          bool                   `protobuf:"varint,14,opt,name=read,proto3" json:"read,omitempty"`
	Aggregation   Aggregation            `protobuf:"varint,15,opt,name=aggregation,proto3,enum=readers.v1.Aggregation" json:"aggregation,omitempty"`
	Comparator    string                 `protobuf:"bytes,16,opt,name=comparator,proto3" json:"comparator,omitempty"`
	Format        string                 `protobuf:"bytes,17,opt,name=format,proto3" json:"format,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PageMetadata) Reset() {
	*x = PageMetadata{}
	mi := &file_readers_v1_readers_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageMetadata) ProtoMessage() {}

func (x *PageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_readers_v1_readers_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageMetadata.ProtoReflect.Descriptor instead.
func (*PageMetadata) Descriptor() ([]byte, []int) {
	return file_readers_v1_readers_proto_rawDescGZIP(), []int{0}
}

func (x *PageMetadata) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PageMetadata) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PageMetadata) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *PageMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PageMetadata) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PageMetadata) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *PageMetadata) GetBoolValue() bool {
	if x != nil {
		return x.BoolValue
	}
	return false
}

func (x *PageMetadata) GetStringValue() string {
	if x != nil {
		return x.StringValue
	}
	return ""
}

func (x *PageMetadata) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *PageMetadata) GetFrom() float64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *PageMetadata) GetTo() float64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *PageMetadata) GetSubtopic() string {
	if x != nil {
		return x.Subtopic
	}
	return ""
}

func (x *PageMetadata) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

func (x *PageMetadata) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *PageMetadata) GetAggregation() Aggregation {
	if x != nil {
		return x.Aggregation
	}
	return Aggregation_AGGREGATION_UNSPECIFIED
}

func (x *PageMetadata) GetComparator() string {
	if x != nil {
		return x.Comparator
	}
	return ""
}

func (x *PageMetadata) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

type ReadMessagesRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         uint64                 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageMetadata  *PageMetadata          `protobuf:"bytes,2,opt,name=page_metadata,json=pageMetadata,proto3" json:"page_metadata,omitempty"`
	Messages      []*Message             `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadMessagesRes) Reset() {
	*x = ReadMessagesRes{}
	mi := &file_readers_v1_readers_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadMessagesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadMessagesRes) ProtoMessage() {}

func (x *ReadMessagesRes) ProtoReflect() protoreflect.Message {
	mi := &file_readers_v1_readers_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadMessagesRes.ProtoReflect.Descriptor instead.
func (*ReadMessagesRes) Descriptor() ([]byte, []int) {
	return file_readers_v1_readers_proto_rawDescGZIP(), []int{1}
}

func (x *ReadMessagesRes) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ReadMessagesRes) GetPageMetadata() *PageMetadata {
	if x != nil {
		return x.PageMetadata
	}
	return nil
}

func (x *ReadMessagesRes) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*Message_Senml
	//	*Message_Json
	Payload       isMessage_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_readers_v1_readers_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_readers_v1_readers_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_readers_v1_readers_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetPayload() isMessage_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Message) GetSenml() *SenMLMessage {
	if x != nil {
		if x, ok := x.Payload.(*Message_Senml); ok {
			return x.Senml
		}
	}
	return nil
}

func (x *Message) GetJson() *JsonMessage {
	if x != nil {
		if x, ok := x.Payload.(*Message_Json); ok {
			return x.Json
		}
	}
	return nil
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Senml struct {
	Senml *SenMLMessage `protobuf:"bytes,1,opt,name=senml,proto3,oneof"`
}

type Message_Json struct {
	Json *JsonMessage `protobuf:"bytes,2,opt,name=json,proto3,oneof"`
}

func (*Message_Senml) isMessage_Payload() {}

func (*Message_Json) isMessage_Payload() {}

type BaseMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Channel       string                 `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Subtopic      string                 `protobuf:"bytes,2,opt,name=subtopic,proto3" json:"subtopic,omitempty"`
	Publisher     string                 `protobuf:"bytes,3,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Protocol      string                 `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseMessage) Reset() {
	*x = BaseMessage{}
	mi := &file_readers_v1_readers_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseMessage) ProtoMessage() {}

func (x *BaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_readers_v1_readers_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseMessage.ProtoReflect.Descriptor instead.
func (*BaseMessage) Descriptor() ([]byte, []int) {
	return file_readers_v1_readers_proto_rawDescGZIP(), []int{3}
}

func (x *BaseMessage) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *BaseMessage) GetSubtopic() string {
	if x != nil {
		return x.Subtopic
	}
	return ""
}

func (x *BaseMessage) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *BaseMessage) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

type SenMLMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Base          *BaseMessage           `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Unit          string                 `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	Time          float64                `protobuf:"fixed64,4,opt,name=time,proto3" json:"time,omitempty"`
	UpdateTime    float64                `protobuf:"fixed64,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Value         float64                `protobuf:"fixed64,6,opt,name=value,proto3" json:"value,omitempty"`
	StringValue   string                 `protobuf:"bytes,7,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	DataValue     string                 `protobuf:"bytes,8,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
	BoolValue     bool                   `protobuf:"varint,9,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	Sum           float64                `protobuf:"fixed64,10,opt,name=sum,proto3" json:"sum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SenMLMessage) Reset() {
	*x = SenMLMessage{}
	mi := &file_readers_v1_readers_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SenMLMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SenMLMessage) ProtoMessage() {}

func (x *SenMLMessage) ProtoReflect() protoreflect.Message {
	mi := &file_readers_v1_readers_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SenMLMessage.ProtoReflect.Descriptor instead.
func (*SenMLMessage) Descriptor() ([]byte, []int) {
	return file_readers_v1_readers_proto_rawDescGZIP(), []int{4}
}

func (x *SenMLMessage) GetBase() *BaseMessage {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *SenMLMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SenMLMessage) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *SenMLMessage) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *SenMLMessage) GetUpdateTime() float64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *SenMLMessage) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *SenMLMessage) GetStringValue() string {
	if x != nil {
		return x.StringValue
	}
	return ""
}

func (x *SenMLMessage) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *SenMLMessage) GetBoolValue() bool {
	if x != nil {
		return x.BoolValue
	}
	return false
}

func (x *SenMLMessage) GetSum() float64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type JsonMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Base          *BaseMessage           `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Created       int64                  `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Payload       []byte                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JsonMessage) Reset() {
	*x = JsonMessage{}
	mi := &file_readers_v1_readers_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JsonMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonMessage) ProtoMessage() {}

func (x *JsonMessage) ProtoReflect() protoreflect.Message {
	mi := &file_readers_v1_readers_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonMessage.ProtoReflect.Descriptor instead.
func (*JsonMessage) Descriptor() ([]byte, []int) {
	return file_readers_v1_readers_proto_rawDescGZIP(), []int{5}
}

func (x *JsonMessage) GetBase() *BaseMessage {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *JsonMessage) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *JsonMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ReadMessagesReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChannelId     string                 `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	DomainId      string                 `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	PageMetadata  *PageMetadata          `protobuf:"bytes,3,opt,name=page_metadata,json=pageMetadata,proto3" json:"page_metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadMessagesReq) Reset() {
	*x = ReadMessagesReq{}
	mi := &file_readers_v1_readers_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadMessagesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadMessagesReq) ProtoMessage() {}

func (x *ReadMessagesReq) ProtoReflect() protoreflect.Message {
	mi := &file_readers_v1_readers_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadMessagesReq.ProtoReflect.Descriptor instead.
func (*ReadMessagesReq) Descriptor() ([]byte, []int) {
	return file_readers_v1_readers_proto_rawDescGZIP(), []int{6}
}

func (x *ReadMessagesReq) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ReadMessagesReq) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *ReadMessagesReq) GetPageMetadata() *PageMetadata {
	if x != nil {
		return x.PageMetadata
	}
	return nil
}

var File_readers_v1_readers_proto protoreflect.FileDescriptor

const file_readers_v1_readers_proto_rawDesc = "" +
	"\n" +
	"\x18readers/v1/readers.proto\x12\n" +
	"readers.v1\"\xe4\x03\n" +
	"\fPageMetadata\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x04R\x05limit\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x04R\x06offset\x12\x1a\n" +
	"\bprotocol\x18\x03 \x01(\tR\bprotocol\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12\x14\n" +
	"\x05value\x18\x05 \x01(\x01R\x05value\x12\x1c\n" +
	"\tpublisher\x18\x06 \x01(\tR\tpublisher\x12\x1d\n" +
	"\n" +
	"bool_value\x18\a \x01(\bR\tboolValue\x12!\n" +
	"\fstring_value\x18\b \x01(\tR\vstringValue\x12\x1d\n" +
	"\n" +
	"data_value\x18\t \x01(\tR\tdataValue\x12\x12\n" +
	"\x04from\x18\n" +
	" \x01(\x01R\x04from\x12\x0e\n" +
	"\x02to\x18\v \x01(\x01R\x02to\x12\x1a\n" +
	"\bsubtopic\x18\f \x01(\tR\bsubtopic\x12\x1a\n" +
	"\binterval\x18\r \x01(\tR\binterval\x12\x12\n" +
	"\x04read\x18\x0e \x01(\bR\x04read\x129\n" +
	"\vaggregation\x18\x0f \x01(\x0e2\x17.readers.v1.AggregationR\vaggregation\x12\x1e\n" +
	"\n" +
	"comparator\x18\x10 \x01(\tR\n" +
	"comparator\x12\x16\n" +
	"\x06format\x18\x11 \x01(\tR\x06format\"\x97\x01\n" +
	"\x0fReadMessagesRes\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x04R\x05total\x12=\n" +
	"\rpage_metadata\x18\x02 \x01(\v2\x18.readers.v1.PageMetadataR\fpageMetadata\x12/\n" +
	"\bmessages\x18\x03 \x03(\v2\x13.readers.v1.MessageR\bmessages\"u\n" +
	"\aMessage\x120\n" +
	"\x05senml\x18\x01 \x01(\v2\x18.readers.v1.SenMLMessageH\x00R\x05senml\x12-\n" +
	"\x04json\x18\x02 \x01(\v2\x17.readers.v1.JsonMessageH\x00R\x04jsonB\t\n" +
	"\apayload\"}\n" +
	"\vBaseMessage\x12\x18\n" +
	"\achannel\x18\x01 \x01(\tR\achannel\x12\x1a\n" +
	"\bsubtopic\x18\x02 \x01(\tR\bsubtopic\x12\x1c\n" +
	"\tpublisher\x18\x03 \x01(\tR\tpublisher\x12\x1a\n" +
	"\bprotocol\x18\x04 \x01(\tR\bprotocol\"\xa1\x02\n" +
	"\fSenMLMessage\x12+\n" +
	"\x04base\x18\x01 \x01(\v2\x17.readers.v1.BaseMessageR\x04base\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x12\n" +
	"\x04unit\x18\x03 \x01(\tR\x04unit\x12\x12\n" +
	"\x04time\x18\x04 \x01(\x01R\x04time\x12\x1f\n" +
	"\vupdate_time\x18\x05 \x01(\x01R\n" +
	"updateTime\x12\x14\n" +
	"\x05value\x18\x06 \x01(\x01R\x05value\x12!\n" +
	"\fstring_value\x18\a \x01(\tR\vstringValue\x12\x1d\n" +
	"\n" +
	"data_value\x18\b \x01(\tR\tdataValue\x12\x1d\n" +
	"\n" +
	"bool_value\x18\t \x01(\bR\tboolValue\x12\x10\n" +
	"\x03sum\x18\n" +
	" \x01(\x01R\x03sum\"n\n" +
	"\vJsonMessage\x12+\n" +
	"\x04base\x18\x01 \x01(\v2\x17.readers.v1.BaseMessageR\x04base\x12\x18\n" +
	"\acreated\x18\x02 \x01(\x03R\acreated\x12\x18\n" +
	"\apayload\x18\x03 \x01(\fR\apayload\"\x8c\x01\n" +
	"\x0fReadMessagesReq\x12\x1d\n" +
	"\n" +
	"channel_id\x18\x01 \x01(\tR\tchannelId\x12\x1b\n" +
	"\tdomain_id\x18\x02 \x01(\tR\bdomainId\x12=\n" +
	"\rpage_metadata\x18\x03 \x01(\v2\x18.readers.v1.PageMetadataR\fpageMetadata*Y\n" +
	"\vAggregation\x12\x1b\n" +
	"\x17AGGREGATION_UNSPECIFIED\x10\x00\x12\a\n" +
	"\x03MAX\x10\x01\x12\a\n" +
	"\x03MIN\x10\x02\x12\a\n" +
	"\x03SUM\x10\x03\x12\t\n" +
	"\x05COUNT\x10\x04\x12\a\n" +
	"\x03AVG\x10\x052\\\n" +
	"\x0eReadersService\x12J\n" +
	"\fReadMessages\x12\x1b.readers.v1.ReadMessagesReq\x1a\x1b.readers.v1.ReadMessagesRes\"\x00B3Z1github.com/absmach/magistrala/api/grpc/readers/v1b\x06proto3"

var (
	file_readers_v1_readers_proto_rawDescOnce sync.Once
	file_readers_v1_readers_proto_rawDescData []byte
)

func file_readers_v1_readers_proto_rawDescGZIP() []byte {
	file_readers_v1_readers_proto_rawDescOnce.Do(func() {
		file_readers_v1_readers_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_readers_v1_readers_proto_rawDesc), len(file_readers_v1_readers_proto_rawDesc)))
	})
	return file_readers_v1_readers_proto_rawDescData
}

var file_readers_v1_readers_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_readers_v1_readers_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_readers_v1_readers_proto_goTypes = []any{
	(Aggregation)(0),        // 0: readers.v1.Aggregation
	(*PageMetadata)(nil),    // 1: readers.v1.PageMetadata
	(*ReadMessagesRes)(nil), // 2: readers.v1.ReadMessagesRes
	(*Message)(nil),         // 3: readers.v1.Message
	(*BaseMessage)(nil),     // 4: readers.v1.BaseMessage
	(*SenMLMessage)(nil),    // 5: readers.v1.SenMLMessage
	(*JsonMessage)(nil),     // 6: readers.v1.JsonMessage
	(*ReadMessagesReq)(nil), // 7: readers.v1.ReadMessagesReq
}
var file_readers_v1_readers_proto_depIdxs = []int32{
	0, // 0: readers.v1.PageMetadata.aggregation:type_name -> readers.v1.Aggregation
	1, // 1: readers.v1.ReadMessagesRes.page_metadata:type_name -> readers.v1.PageMetadata
	3, // 2: readers.v1.ReadMessagesRes.messages:type_name -> readers.v1.Message
	5, // 3: readers.v1.Message.senml:type_name -> readers.v1.SenMLMessage
	6, // 4: readers.v1.Message.json:type_name -> readers.v1.JsonMessage
	4, // 5: readers.v1.SenMLMessage.base:type_name -> readers.v1.BaseMessage
	4, // 6: readers.v1.JsonMessage.base:type_name -> readers.v1.BaseMessage
	1, // 7: readers.v1.ReadMessagesReq.page_metadata:type_name -> readers.v1.PageMetadata
	7, // 8: readers.v1.ReadersService.ReadMessages:input_type -> readers.v1.ReadMessagesReq
	2, // 9: readers.v1.ReadersService.ReadMessages:output_type -> readers.v1.ReadMessagesRes
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_readers_v1_readers_proto_init() }
func file_readers_v1_readers_proto_init() {
	if File_readers_v1_readers_proto != nil {
		return
	}
	file_readers_v1_readers_proto_msgTypes[2].OneofWrappers = []any{
		(*Message_Senml)(nil),
		(*Message_Json)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_readers_v1_readers_proto_rawDesc), len(file_readers_v1_readers_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_readers_v1_readers_proto_goTypes,
		DependencyIndexes: file_readers_v1_readers_proto_depIdxs,
		EnumInfos:         file_readers_v1_readers_proto_enumTypes,
		MessageInfos:      file_readers_v1_readers_proto_msgTypes,
	}.Build()
	File_readers_v1_readers_proto = out.File
	file_readers_v1_readers_proto_goTypes = nil
	file_readers_v1_readers_proto_depIdxs = nil
}
