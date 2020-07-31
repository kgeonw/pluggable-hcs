// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/TransactionGetReceipt.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Get the receipt of a transaction, given its transaction ID. Once a transaction reaches consensus, then information about whether it succeeded or failed will be available until the end of the receipt period.  Before and after the receipt period, and for a transaction that was never submitted, the receipt is unknown.  This query is free (the payment field is left empty). No State proof is available for this response
type TransactionGetReceiptQuery struct {
	Header               *QueryHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TransactionID        *TransactionID `protobuf:"bytes,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransactionGetReceiptQuery) Reset()         { *m = TransactionGetReceiptQuery{} }
func (m *TransactionGetReceiptQuery) String() string { return proto.CompactTextString(m) }
func (*TransactionGetReceiptQuery) ProtoMessage()    {}
func (*TransactionGetReceiptQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_ded36cff96af940d, []int{0}
}

func (m *TransactionGetReceiptQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionGetReceiptQuery.Unmarshal(m, b)
}
func (m *TransactionGetReceiptQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionGetReceiptQuery.Marshal(b, m, deterministic)
}
func (m *TransactionGetReceiptQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionGetReceiptQuery.Merge(m, src)
}
func (m *TransactionGetReceiptQuery) XXX_Size() int {
	return xxx_messageInfo_TransactionGetReceiptQuery.Size(m)
}
func (m *TransactionGetReceiptQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionGetReceiptQuery.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionGetReceiptQuery proto.InternalMessageInfo

func (m *TransactionGetReceiptQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TransactionGetReceiptQuery) GetTransactionID() *TransactionID {
	if m != nil {
		return m.TransactionID
	}
	return nil
}

// Response when the client sends the node TransactionGetReceiptQuery. If it created a new entity (account, file, or smart contract instance) then one of the three ID fields will be filled in with the ID of the new entity. Sometimes a single transaction will create more than one new entity, such as when a new contract instance is created, and this also creates the new account that it owned by that instance. No State proof is available for this response
type TransactionGetReceiptResponse struct {
	Header               *ResponseHeader     `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Receipt              *TransactionReceipt `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TransactionGetReceiptResponse) Reset()         { *m = TransactionGetReceiptResponse{} }
func (m *TransactionGetReceiptResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionGetReceiptResponse) ProtoMessage()    {}
func (*TransactionGetReceiptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ded36cff96af940d, []int{1}
}

func (m *TransactionGetReceiptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionGetReceiptResponse.Unmarshal(m, b)
}
func (m *TransactionGetReceiptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionGetReceiptResponse.Marshal(b, m, deterministic)
}
func (m *TransactionGetReceiptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionGetReceiptResponse.Merge(m, src)
}
func (m *TransactionGetReceiptResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionGetReceiptResponse.Size(m)
}
func (m *TransactionGetReceiptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionGetReceiptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionGetReceiptResponse proto.InternalMessageInfo

func (m *TransactionGetReceiptResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TransactionGetReceiptResponse) GetReceipt() *TransactionReceipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionGetReceiptQuery)(nil), "proto.TransactionGetReceiptQuery")
	proto.RegisterType((*TransactionGetReceiptResponse)(nil), "proto.TransactionGetReceiptResponse")
}

func init() {
	proto.RegisterFile("proto/TransactionGetReceipt.proto", fileDescriptor_ded36cff96af940d)
}

var fileDescriptor_ded36cff96af940d = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xbf, 0x4b, 0x03, 0x31,
	0x1c, 0xc5, 0x89, 0x60, 0x85, 0x88, 0x4b, 0xf0, 0x47, 0x0d, 0x28, 0xda, 0xa9, 0x08, 0x97, 0x03,
	0xbb, 0x39, 0x16, 0xc1, 0xba, 0x69, 0xb8, 0xc9, 0x2d, 0xcd, 0x7d, 0xb9, 0x04, 0xe9, 0x25, 0x24,
	0xe9, 0xd0, 0x59, 0xff, 0x70, 0x21, 0xc9, 0xd5, 0xbb, 0xf6, 0xa6, 0x40, 0xde, 0xe7, 0xbd, 0x7c,
	0x82, 0x1f, 0xad, 0x33, 0xc1, 0x94, 0x95, 0x13, 0xad, 0x17, 0x32, 0x68, 0xd3, 0xbe, 0x41, 0xe0,
	0x20, 0x41, 0xdb, 0xc0, 0x62, 0x46, 0x4e, 0xe3, 0x41, 0xef, 0x8f, 0xc8, 0x01, 0x46, 0xaf, 0x53,
	0xbe, 0x14, 0x5e, 0xcb, 0x6a, 0x67, 0xc1, 0xe7, 0xfb, 0x9b, 0x74, 0xff, 0xb9, 0x05, 0xb7, 0x5b,
	0x81, 0xa8, 0xc1, 0xe5, 0x80, 0xa6, 0x80, 0x83, 0xb7, 0xa6, 0xf5, 0xd0, 0xcf, 0x66, 0xbf, 0x08,
	0xd3, 0x51, 0xa7, 0x38, 0x43, 0x9e, 0xf0, 0x44, 0x45, 0x7c, 0x8a, 0x1e, 0xd0, 0xfc, 0xfc, 0x99,
	0xa4, 0x1a, 0xeb, 0x3d, 0xc2, 0x33, 0x41, 0x5e, 0xf0, 0x45, 0xf8, 0x5f, 0x7a, 0x7f, 0x9d, 0x9e,
	0xc4, 0xca, 0x65, 0xae, 0x54, 0xfd, 0x8c, 0x0f, 0xd1, 0xd9, 0x0f, 0xc2, 0x77, 0xa3, 0x1a, 0x9d,
	0x34, 0x29, 0x0e, 0x4c, 0xae, 0xf2, 0xec, 0xf0, 0x57, 0x7b, 0x99, 0x05, 0x3e, 0x73, 0x69, 0x21,
	0x6b, 0xdc, 0x1e, 0x6b, 0x74, 0x4f, 0x74, 0xe4, 0x72, 0x85, 0xa9, 0x34, 0x1b, 0xa6, 0xa0, 0x06,
	0x27, 0x98, 0x12, 0x5e, 0x35, 0x4e, 0x58, 0x95, 0x9a, 0x1f, 0xe8, 0x6b, 0xde, 0xe8, 0xa0, 0xb6,
	0x6b, 0x26, 0xcd, 0xa6, 0xdc, 0xa7, 0x65, 0xc2, 0x0b, 0x5f, 0x7f, 0x17, 0x8d, 0x29, 0x23, 0xbb,
	0x9e, 0xc4, 0x63, 0xf1, 0x17, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xce, 0x2d, 0x3f, 0xf6, 0x01, 0x00,
	0x00,
}