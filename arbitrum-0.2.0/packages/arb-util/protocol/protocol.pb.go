// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package protocol

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	value "github.com/offchainlabs/arbitrum/packages/arb-util/value"
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

type TokenTypeBuf struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenTypeBuf) Reset()         { *m = TokenTypeBuf{} }
func (m *TokenTypeBuf) String() string { return proto.CompactTextString(m) }
func (*TokenTypeBuf) ProtoMessage()    {}
func (*TokenTypeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *TokenTypeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenTypeBuf.Unmarshal(m, b)
}

func (m *TokenTypeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenTypeBuf.Marshal(b, m, deterministic)
}

func (m *TokenTypeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenTypeBuf.Merge(m, src)
}

func (m *TokenTypeBuf) XXX_Size() int {
	return xxx_messageInfo_TokenTypeBuf.Size(m)
}

func (m *TokenTypeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenTypeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_TokenTypeBuf proto.InternalMessageInfo

func (m *TokenTypeBuf) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type MessageBuf struct {
	Value                *value.ValueBuf      `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	TokenType            *TokenTypeBuf        `protobuf:"bytes,2,opt,name=tokenType,proto3" json:"tokenType,omitempty"`
	Amount               *value.BigIntegerBuf `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Sender               *value.HashBuf       `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MessageBuf) Reset()         { *m = MessageBuf{} }
func (m *MessageBuf) String() string { return proto.CompactTextString(m) }
func (*MessageBuf) ProtoMessage()    {}
func (*MessageBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

func (m *MessageBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageBuf.Unmarshal(m, b)
}

func (m *MessageBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageBuf.Marshal(b, m, deterministic)
}

func (m *MessageBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageBuf.Merge(m, src)
}

func (m *MessageBuf) XXX_Size() int {
	return xxx_messageInfo_MessageBuf.Size(m)
}

func (m *MessageBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageBuf.DiscardUnknown(m)
}

var xxx_messageInfo_MessageBuf proto.InternalMessageInfo

func (m *MessageBuf) GetValue() *value.ValueBuf {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MessageBuf) GetTokenType() *TokenTypeBuf {
	if m != nil {
		return m.TokenType
	}
	return nil
}

func (m *MessageBuf) GetAmount() *value.BigIntegerBuf {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MessageBuf) GetSender() *value.HashBuf {
	if m != nil {
		return m.Sender
	}
	return nil
}

type TimeBoundsBuf struct {
	StartTime            uint64   `protobuf:"varint,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              uint64   `protobuf:"varint,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeBoundsBuf) Reset()         { *m = TimeBoundsBuf{} }
func (m *TimeBoundsBuf) String() string { return proto.CompactTextString(m) }
func (*TimeBoundsBuf) ProtoMessage()    {}
func (*TimeBoundsBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{2}
}

func (m *TimeBoundsBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeBoundsBuf.Unmarshal(m, b)
}

func (m *TimeBoundsBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeBoundsBuf.Marshal(b, m, deterministic)
}

func (m *TimeBoundsBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeBoundsBuf.Merge(m, src)
}

func (m *TimeBoundsBuf) XXX_Size() int {
	return xxx_messageInfo_TimeBoundsBuf.Size(m)
}

func (m *TimeBoundsBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeBoundsBuf.DiscardUnknown(m)
}

var xxx_messageInfo_TimeBoundsBuf proto.InternalMessageInfo

func (m *TimeBoundsBuf) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TimeBoundsBuf) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type BalanceTrackerBuf struct {
	Types                []*TokenTypeBuf        `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
	Amounts              []*value.BigIntegerBuf `protobuf:"bytes,2,rep,name=amounts,proto3" json:"amounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BalanceTrackerBuf) Reset()         { *m = BalanceTrackerBuf{} }
func (m *BalanceTrackerBuf) String() string { return proto.CompactTextString(m) }
func (*BalanceTrackerBuf) ProtoMessage()    {}
func (*BalanceTrackerBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{3}
}

func (m *BalanceTrackerBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceTrackerBuf.Unmarshal(m, b)
}

func (m *BalanceTrackerBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceTrackerBuf.Marshal(b, m, deterministic)
}

func (m *BalanceTrackerBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceTrackerBuf.Merge(m, src)
}

func (m *BalanceTrackerBuf) XXX_Size() int {
	return xxx_messageInfo_BalanceTrackerBuf.Size(m)
}

func (m *BalanceTrackerBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceTrackerBuf.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceTrackerBuf proto.InternalMessageInfo

func (m *BalanceTrackerBuf) GetTypes() []*TokenTypeBuf {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *BalanceTrackerBuf) GetAmounts() []*value.BigIntegerBuf {
	if m != nil {
		return m.Amounts
	}
	return nil
}

type PreconditionBuf struct {
	BeforeHash           *value.HashBuf     `protobuf:"bytes,1,opt,name=beforeHash,proto3" json:"beforeHash,omitempty"`
	TimeBounds           *TimeBoundsBuf     `protobuf:"bytes,2,opt,name=timeBounds,proto3" json:"timeBounds,omitempty"`
	BalanceTracker       *BalanceTrackerBuf `protobuf:"bytes,3,opt,name=balanceTracker,proto3" json:"balanceTracker,omitempty"`
	BeforeInbox          *value.HashBuf     `protobuf:"bytes,4,opt,name=beforeInbox,proto3" json:"beforeInbox,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PreconditionBuf) Reset()         { *m = PreconditionBuf{} }
func (m *PreconditionBuf) String() string { return proto.CompactTextString(m) }
func (*PreconditionBuf) ProtoMessage()    {}
func (*PreconditionBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{4}
}

func (m *PreconditionBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreconditionBuf.Unmarshal(m, b)
}

func (m *PreconditionBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreconditionBuf.Marshal(b, m, deterministic)
}

func (m *PreconditionBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreconditionBuf.Merge(m, src)
}

func (m *PreconditionBuf) XXX_Size() int {
	return xxx_messageInfo_PreconditionBuf.Size(m)
}

func (m *PreconditionBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_PreconditionBuf.DiscardUnknown(m)
}

var xxx_messageInfo_PreconditionBuf proto.InternalMessageInfo

func (m *PreconditionBuf) GetBeforeHash() *value.HashBuf {
	if m != nil {
		return m.BeforeHash
	}
	return nil
}

func (m *PreconditionBuf) GetTimeBounds() *TimeBoundsBuf {
	if m != nil {
		return m.TimeBounds
	}
	return nil
}

func (m *PreconditionBuf) GetBalanceTracker() *BalanceTrackerBuf {
	if m != nil {
		return m.BalanceTracker
	}
	return nil
}

func (m *PreconditionBuf) GetBeforeInbox() *value.HashBuf {
	if m != nil {
		return m.BeforeInbox
	}
	return nil
}

type AssertionBuf struct {
	AfterHash            *value.HashBuf    `protobuf:"bytes,1,opt,name=afterHash,proto3" json:"afterHash,omitempty"`
	NumSteps             uint32            `protobuf:"varint,2,opt,name=numSteps,proto3" json:"numSteps,omitempty"`
	Messages             []*MessageBuf     `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	Logs                 []*value.ValueBuf `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AssertionBuf) Reset()         { *m = AssertionBuf{} }
func (m *AssertionBuf) String() string { return proto.CompactTextString(m) }
func (*AssertionBuf) ProtoMessage()    {}
func (*AssertionBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{5}
}

func (m *AssertionBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssertionBuf.Unmarshal(m, b)
}

func (m *AssertionBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssertionBuf.Marshal(b, m, deterministic)
}

func (m *AssertionBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssertionBuf.Merge(m, src)
}

func (m *AssertionBuf) XXX_Size() int {
	return xxx_messageInfo_AssertionBuf.Size(m)
}

func (m *AssertionBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_AssertionBuf.DiscardUnknown(m)
}

var xxx_messageInfo_AssertionBuf proto.InternalMessageInfo

func (m *AssertionBuf) GetAfterHash() *value.HashBuf {
	if m != nil {
		return m.AfterHash
	}
	return nil
}

func (m *AssertionBuf) GetNumSteps() uint32 {
	if m != nil {
		return m.NumSteps
	}
	return 0
}

func (m *AssertionBuf) GetMessages() []*MessageBuf {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *AssertionBuf) GetLogs() []*value.ValueBuf {
	if m != nil {
		return m.Logs
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenTypeBuf)(nil), "protocol.TokenTypeBuf")
	proto.RegisterType((*MessageBuf)(nil), "protocol.MessageBuf")
	proto.RegisterType((*TimeBoundsBuf)(nil), "protocol.TimeBoundsBuf")
	proto.RegisterType((*BalanceTrackerBuf)(nil), "protocol.BalanceTrackerBuf")
	proto.RegisterType((*PreconditionBuf)(nil), "protocol.PreconditionBuf")
	proto.RegisterType((*AssertionBuf)(nil), "protocol.AssertionBuf")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xc7, 0x95, 0xdd, 0xed, 0x76, 0x77, 0xba, 0xdd, 0xaa, 0x56, 0x05, 0x56, 0xe1, 0x50, 0x85,
	0x0f, 0x55, 0x68, 0xc9, 0x56, 0x80, 0xc4, 0x99, 0x70, 0x80, 0x1e, 0x90, 0x50, 0x58, 0x71, 0xe0,
	0xe6, 0x64, 0x27, 0xd9, 0xa8, 0x89, 0x1d, 0x6c, 0xa7, 0xa2, 0x4f, 0xc5, 0x23, 0xf0, 0x4c, 0xbc,
	0x01, 0x8a, 0xed, 0x7c, 0xd0, 0x42, 0x2f, 0x96, 0xc7, 0xff, 0xdf, 0xd8, 0xf3, 0xf7, 0x0c, 0x2c,
	0x2b, 0x29, 0xb4, 0x48, 0x44, 0x11, 0x98, 0x0d, 0x99, 0xb5, 0xf1, 0xe9, 0xf1, 0x35, 0x2b, 0x6a,
	0x5c, 0x9b, 0xd5, 0x8a, 0xfe, 0x53, 0x58, 0x6c, 0xc4, 0x15, 0xf2, 0xcd, 0x4d, 0x85, 0x61, 0x9d,
	0x92, 0x13, 0xd8, 0x33, 0x32, 0xf5, 0xce, 0xbc, 0xf3, 0x45, 0x64, 0x03, 0xff, 0x97, 0x07, 0xf0,
	0x09, 0x95, 0x62, 0x99, 0x81, 0x9e, 0x0d, 0xa1, 0x83, 0x57, 0x47, 0x81, 0xbd, 0xf1, 0x6b, 0xb3,
	0x86, 0x75, 0xea, 0xb2, 0xc8, 0x1b, 0x98, 0xeb, 0xf6, 0x6e, 0x3a, 0x32, 0xe8, 0x83, 0xa0, 0x2b,
	0x6e, 0xf8, 0x6c, 0xd4, 0x83, 0x64, 0x05, 0x53, 0x56, 0x8a, 0x9a, 0x6b, 0x3a, 0x36, 0x29, 0x27,
	0xee, 0xf6, 0x30, 0xcf, 0x2e, 0xb9, 0xc6, 0x0c, 0x65, 0x93, 0xe0, 0x18, 0xf2, 0x1c, 0xa6, 0x0a,
	0xf9, 0x16, 0x25, 0x9d, 0x18, 0x7a, 0xe9, 0xe8, 0x8f, 0x4c, 0xed, 0x0c, 0x67, 0x55, 0xff, 0x03,
	0x1c, 0x6e, 0xf2, 0x12, 0x43, 0x51, 0xf3, 0xad, 0x6a, 0x3c, 0x3c, 0x86, 0xb9, 0xd2, 0x4c, 0xea,
	0xe6, 0xd4, 0xf8, 0x98, 0x44, 0xfd, 0x01, 0xa1, 0xb0, 0x8f, 0x7c, 0x6b, 0xb4, 0x91, 0xd1, 0xda,
	0xd0, 0xff, 0x0e, 0xc7, 0x21, 0x2b, 0x18, 0x4f, 0x70, 0x23, 0x59, 0x72, 0x65, 0xaa, 0x21, 0x2b,
	0xd8, 0xd3, 0x37, 0x15, 0x2a, 0xea, 0x9d, 0x8d, 0xef, 0x71, 0x69, 0x21, 0x12, 0xc0, 0xbe, 0xad,
	0x5e, 0xd1, 0x91, 0xe1, 0xff, 0x6d, 0xb1, 0x85, 0xfc, 0xdf, 0x1e, 0x1c, 0x7d, 0x96, 0x98, 0x08,
	0xbe, 0xcd, 0x75, 0x2e, 0x78, 0xf3, 0x62, 0x00, 0x10, 0x63, 0x2a, 0x24, 0x36, 0x46, 0x5d, 0x1f,
	0x6e, 0x7b, 0x1f, 0x10, 0xe4, 0x2d, 0x80, 0xee, 0xfc, 0xbb, 0x66, 0x3c, 0x1c, 0x94, 0x39, 0xfc,
	0x9b, 0x68, 0x80, 0x92, 0xf7, 0xb0, 0x8c, 0xff, 0xf2, 0xeb, 0xda, 0xf2, 0xa8, 0x4f, 0xbe, 0xf3,
	0x1f, 0xd1, 0xad, 0x14, 0x72, 0x01, 0x07, 0xb6, 0x96, 0x4b, 0x1e, 0x8b, 0x1f, 0xff, 0x69, 0xd5,
	0x10, 0xf1, 0x7f, 0x7a, 0xb0, 0x78, 0xa7, 0x14, 0xca, 0xd6, 0xf0, 0x0a, 0xe6, 0x2c, 0xd5, 0x28,
	0xef, 0xf1, 0xdb, 0x03, 0xe4, 0x14, 0x66, 0xbc, 0x2e, 0xbf, 0x68, 0xac, 0xac, 0xd9, 0xc3, 0xa8,
	0x8b, 0xc9, 0x05, 0xcc, 0x4a, 0x3b, 0xcb, 0x8a, 0x8e, 0xdd, 0xff, 0x77, 0x5e, 0xfa, 0x29, 0x8f,
	0x3a, 0x8a, 0x3c, 0x81, 0x49, 0x21, 0x32, 0x45, 0x27, 0x86, 0xbe, 0x33, 0xee, 0x46, 0x0c, 0x5f,
	0x7c, 0x3b, 0xcf, 0x72, 0xbd, 0xab, 0xe3, 0x20, 0x11, 0xe5, 0x5a, 0xa4, 0x69, 0xb2, 0x63, 0x39,
	0x2f, 0x58, 0xac, 0xd6, 0x4c, 0xc6, 0x2f, 0xd9, 0x75, 0xb9, 0x6e, 0x5f, 0x89, 0xa7, 0x66, 0xf7,
	0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xed, 0x33, 0x36, 0xab, 0x03, 0x00, 0x00,
}
