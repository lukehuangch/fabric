// Code generated by protoc-gen-go.
// source: ledger/rwset/rwset.proto
// DO NOT EDIT!

/*
Package rwset is a generated protocol buffer package.

It is generated from these files:
	ledger/rwset/rwset.proto

It has these top-level messages:
	TxReadWriteSet
	NsReadWriteSet
*/
package rwset

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

type TxReadWriteSet_DataModel int32

const (
	TxReadWriteSet_KV TxReadWriteSet_DataModel = 0
)

var TxReadWriteSet_DataModel_name = map[int32]string{
	0: "KV",
}
var TxReadWriteSet_DataModel_value = map[string]int32{
	"KV": 0,
}

func (x TxReadWriteSet_DataModel) String() string {
	return proto.EnumName(TxReadWriteSet_DataModel_name, int32(x))
}
func (TxReadWriteSet_DataModel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// TxReadWriteSet encapsulates a read-write set for a transaction
// DataModel specifies the enum value of the data model
// ns_rwset field specifies a list of chaincode specific read-write set (one for each chaincode)
type TxReadWriteSet struct {
	DataModel TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsRwset   []*NsReadWriteSet        `protobuf:"bytes,2,rep,name=ns_rwset,json=nsRwset" json:"ns_rwset,omitempty"`
}

func (m *TxReadWriteSet) Reset()                    { *m = TxReadWriteSet{} }
func (m *TxReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*TxReadWriteSet) ProtoMessage()               {}
func (*TxReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TxReadWriteSet) GetNsRwset() []*NsReadWriteSet {
	if m != nil {
		return m.NsRwset
	}
	return nil
}

// NsReadWriteSet encapsulates the read-write set for a chaincode
type NsReadWriteSet struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Rwset     []byte `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
}

func (m *NsReadWriteSet) Reset()                    { *m = NsReadWriteSet{} }
func (m *NsReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*NsReadWriteSet) ProtoMessage()               {}
func (*NsReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*TxReadWriteSet)(nil), "rwset.TxReadWriteSet")
	proto.RegisterType((*NsReadWriteSet)(nil), "rwset.NsReadWriteSet")
	proto.RegisterEnum("rwset.TxReadWriteSet_DataModel", TxReadWriteSet_DataModel_name, TxReadWriteSet_DataModel_value)
}

func init() { proto.RegisterFile("ledger/rwset/rwset.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0x64, 0xd3, 0x3c, 0xa5, 0x48, 0x54, 0xe8, 0x41, 0x70, 0xec, 0x54, 0x3c, 0x24,
	0x3a, 0xef, 0x1e, 0x64, 0x37, 0xd1, 0x43, 0x14, 0x05, 0x2f, 0xe5, 0xb5, 0x79, 0x76, 0x85, 0xb5,
	0x29, 0x49, 0x44, 0xfd, 0x24, 0x7e, 0x5d, 0x31, 0xd9, 0xb4, 0xbb, 0x84, 0xbc, 0x3f, 0xbf, 0x5f,
	0xf2, 0xf8, 0x43, 0xb6, 0x22, 0x5d, 0x93, 0x95, 0xf6, 0xc3, 0x91, 0x8f, 0xa7, 0xe8, 0xad, 0xf1,
	0x86, 0x8f, 0xc3, 0x30, 0xfb, 0x4e, 0x20, 0x7d, 0xfa, 0x54, 0x84, 0xfa, 0xc5, 0x36, 0x9e, 0x1e,
	0xc9, 0xf3, 0x1b, 0x00, 0x8d, 0x1e, 0x8b, 0xd6, 0x68, 0x5a, 0x65, 0xc9, 0x34, 0xc9, 0xd3, 0xf9,
	0xb9, 0x88, 0xee, 0x36, 0x2a, 0x16, 0xe8, 0xf1, 0xfe, 0x17, 0x53, 0x4c, 0x6f, 0xae, 0xfc, 0x12,
	0xf6, 0x3b, 0x57, 0x04, 0x3e, 0x1b, 0x4d, 0x77, 0xf3, 0x83, 0xf9, 0xe9, 0xda, 0x7e, 0x70, 0x43,
	0x5b, 0xed, 0x75, 0x4e, 0x85, 0x25, 0x8e, 0x81, 0xfd, 0xbd, 0xc4, 0x27, 0x30, 0xba, 0x7b, 0x3e,
	0xda, 0x99, 0x2d, 0x20, 0xdd, 0xe6, 0xf9, 0x19, 0xb0, 0x0e, 0x5b, 0x72, 0x3d, 0x56, 0x14, 0xf6,
	0x62, 0xea, 0x3f, 0xe0, 0x27, 0x30, 0xde, 0xfc, 0x99, 0xe4, 0x87, 0x2a, 0x0e, 0xb7, 0x05, 0x5c,
	0x18, 0x5b, 0x8b, 0xe5, 0x57, 0x4f, 0x36, 0x76, 0x21, 0xde, 0xb0, 0xb4, 0x4d, 0x15, 0x6b, 0x70,
	0x62, 0x1d, 0x06, 0xfa, 0xf5, 0xaa, 0x6e, 0xfc, 0xf2, 0xbd, 0x14, 0x95, 0x69, 0xe5, 0x40, 0x91,
	0x51, 0x91, 0x51, 0x91, 0xc3, 0x4e, 0xcb, 0x49, 0x08, 0xaf, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0x3b, 0xef, 0x59, 0x6a, 0x01, 0x00, 0x00,
}
