// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger/rwset/rwset.proto

/*
Package rwset is a generated protocol buffer package.

It is generated from these files:
	ledger/rwset/rwset.proto

It has these top-level messages:
	TxReadWriteSet
	NsReadWriteSet
	CollectionHashedReadWriteSet
	TxPvtReadWriteSet
	NsPvtReadWriteSet
	CollectionPvtReadWriteSet
	CollectionCriteria
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

func (m *TxReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if m != nil {
		return m.DataModel
	}
	return TxReadWriteSet_KV
}

func (m *TxReadWriteSet) GetNsRwset() []*NsReadWriteSet {
	if m != nil {
		return m.NsRwset
	}
	return nil
}

// NsReadWriteSet encapsulates the read-write set for a chaincode
type NsReadWriteSet struct {
	Namespace             string                          `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Rwset                 []byte                          `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
	CollectionHashedRwset []*CollectionHashedReadWriteSet `protobuf:"bytes,3,rep,name=collection_hashed_rwset,json=collectionHashedRwset" json:"collection_hashed_rwset,omitempty"`
}

func (m *NsReadWriteSet) Reset()                    { *m = NsReadWriteSet{} }
func (m *NsReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*NsReadWriteSet) ProtoMessage()               {}
func (*NsReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NsReadWriteSet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NsReadWriteSet) GetRwset() []byte {
	if m != nil {
		return m.Rwset
	}
	return nil
}

func (m *NsReadWriteSet) GetCollectionHashedRwset() []*CollectionHashedReadWriteSet {
	if m != nil {
		return m.CollectionHashedRwset
	}
	return nil
}

// CollectionHashedReadWriteSet encapsulate the hashed representation for the private read-write set for a collection
type CollectionHashedReadWriteSet struct {
	CollectionName string `protobuf:"bytes,1,opt,name=collection_name,json=collectionName" json:"collection_name,omitempty"`
	HashedRwset    []byte `protobuf:"bytes,2,opt,name=hashed_rwset,json=hashedRwset,proto3" json:"hashed_rwset,omitempty"`
	PvtRwsetHash   []byte `protobuf:"bytes,3,opt,name=pvt_rwset_hash,json=pvtRwsetHash,proto3" json:"pvt_rwset_hash,omitempty"`
}

func (m *CollectionHashedReadWriteSet) Reset()                    { *m = CollectionHashedReadWriteSet{} }
func (m *CollectionHashedReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*CollectionHashedReadWriteSet) ProtoMessage()               {}
func (*CollectionHashedReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CollectionHashedReadWriteSet) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CollectionHashedReadWriteSet) GetHashedRwset() []byte {
	if m != nil {
		return m.HashedRwset
	}
	return nil
}

func (m *CollectionHashedReadWriteSet) GetPvtRwsetHash() []byte {
	if m != nil {
		return m.PvtRwsetHash
	}
	return nil
}

// TxPvtReadWriteSet encapsulate the private read-write set for a transaction
type TxPvtReadWriteSet struct {
	DataModel  TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsPvtRwset []*NsPvtReadWriteSet     `protobuf:"bytes,2,rep,name=ns_pvt_rwset,json=nsPvtRwset" json:"ns_pvt_rwset,omitempty"`
}

func (m *TxPvtReadWriteSet) Reset()                    { *m = TxPvtReadWriteSet{} }
func (m *TxPvtReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*TxPvtReadWriteSet) ProtoMessage()               {}
func (*TxPvtReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TxPvtReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if m != nil {
		return m.DataModel
	}
	return TxReadWriteSet_KV
}

func (m *TxPvtReadWriteSet) GetNsPvtRwset() []*NsPvtReadWriteSet {
	if m != nil {
		return m.NsPvtRwset
	}
	return nil
}

// NsPvtReadWriteSet encapsulates the private read-write set for a chaincode
type NsPvtReadWriteSet struct {
	Namespace          string                       `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	CollectionPvtRwset []*CollectionPvtReadWriteSet `protobuf:"bytes,2,rep,name=collection_pvt_rwset,json=collectionPvtRwset" json:"collection_pvt_rwset,omitempty"`
}

func (m *NsPvtReadWriteSet) Reset()                    { *m = NsPvtReadWriteSet{} }
func (m *NsPvtReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*NsPvtReadWriteSet) ProtoMessage()               {}
func (*NsPvtReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NsPvtReadWriteSet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NsPvtReadWriteSet) GetCollectionPvtRwset() []*CollectionPvtReadWriteSet {
	if m != nil {
		return m.CollectionPvtRwset
	}
	return nil
}

// CollectionPvtReadWriteSet encapsulates the private read-write set for a collection
type CollectionPvtReadWriteSet struct {
	CollectionName string `protobuf:"bytes,1,opt,name=collection_name,json=collectionName" json:"collection_name,omitempty"`
	Rwset          []byte `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
}

func (m *CollectionPvtReadWriteSet) Reset()                    { *m = CollectionPvtReadWriteSet{} }
func (m *CollectionPvtReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*CollectionPvtReadWriteSet) ProtoMessage()               {}
func (*CollectionPvtReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CollectionPvtReadWriteSet) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CollectionPvtReadWriteSet) GetRwset() []byte {
	if m != nil {
		return m.Rwset
	}
	return nil
}

type CollectionCriteria struct {
	Channel    string `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	TxId       string `protobuf:"bytes,2,opt,name=tx_id,json=txId" json:"tx_id,omitempty"`
	Collection string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
}

func (m *CollectionCriteria) Reset()                    { *m = CollectionCriteria{} }
func (m *CollectionCriteria) String() string            { return proto.CompactTextString(m) }
func (*CollectionCriteria) ProtoMessage()               {}
func (*CollectionCriteria) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CollectionCriteria) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *CollectionCriteria) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *CollectionCriteria) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func init() {
	proto.RegisterType((*TxReadWriteSet)(nil), "rwset.TxReadWriteSet")
	proto.RegisterType((*NsReadWriteSet)(nil), "rwset.NsReadWriteSet")
	proto.RegisterType((*CollectionHashedReadWriteSet)(nil), "rwset.CollectionHashedReadWriteSet")
	proto.RegisterType((*TxPvtReadWriteSet)(nil), "rwset.TxPvtReadWriteSet")
	proto.RegisterType((*NsPvtReadWriteSet)(nil), "rwset.NsPvtReadWriteSet")
	proto.RegisterType((*CollectionPvtReadWriteSet)(nil), "rwset.CollectionPvtReadWriteSet")
	proto.RegisterType((*CollectionCriteria)(nil), "rwset.CollectionCriteria")
	proto.RegisterEnum("rwset.TxReadWriteSet_DataModel", TxReadWriteSet_DataModel_name, TxReadWriteSet_DataModel_value)
}

func init() { proto.RegisterFile("ledger/rwset/rwset.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x2b, 0xdd, 0xc8, 0x5b, 0x15, 0x98, 0xb7, 0x89, 0x20, 0x4d, 0x50, 0x0a, 0x12, 0x15,
	0x87, 0x04, 0xc6, 0x8d, 0x03, 0x07, 0xc6, 0x01, 0x84, 0x98, 0x90, 0x99, 0x40, 0x1a, 0x87, 0xc8,
	0xb5, 0x4d, 0x13, 0xa9, 0x75, 0x22, 0xdb, 0x94, 0xf2, 0x07, 0x70, 0x86, 0x1b, 0x67, 0xfe, 0x53,
	0xd4, 0xe7, 0xb4, 0xf9, 0x31, 0x7e, 0x1d, 0xb8, 0x44, 0xf1, 0xf3, 0xf7, 0xbd, 0xef, 0xf3, 0xfb,
	0xe2, 0x40, 0x34, 0x53, 0x72, 0xaa, 0x4c, 0x62, 0x3e, 0x59, 0xe5, 0xfc, 0x33, 0x2e, 0x4d, 0xe1,
	0x0a, 0xda, 0xc7, 0xc5, 0xe8, 0x3b, 0x81, 0xf0, 0x6c, 0xc9, 0x14, 0x97, 0xef, 0x4c, 0xee, 0xd4,
	0x1b, 0xe5, 0xe8, 0x13, 0x00, 0xc9, 0x1d, 0x4f, 0xe7, 0x85, 0x54, 0xb3, 0x88, 0x0c, 0xc9, 0x38,
	0x3c, 0xbe, 0x15, 0x7b, 0x6e, 0x1b, 0x1a, 0x3f, 0xe3, 0x8e, 0xbf, 0x5a, 0xc1, 0x58, 0x20, 0xd7,
	0xaf, 0xf4, 0x01, 0x5c, 0xd1, 0x36, 0x45, 0x7c, 0xb4, 0x35, 0xec, 0x8d, 0x77, 0x8f, 0x0f, 0x2b,
	0xf6, 0xa9, 0x6d, 0xb2, 0xd9, 0x8e, 0xb6, 0x0c, 0x4d, 0xec, 0x43, 0xb0, 0xe9, 0x44, 0xb7, 0x61,
	0xeb, 0xe5, 0xdb, 0x6b, 0x97, 0x46, 0x3f, 0x08, 0x84, 0x6d, 0x02, 0x3d, 0x82, 0x40, 0xf3, 0xb9,
	0xb2, 0x25, 0x17, 0x0a, 0x8d, 0x05, 0xac, 0x2e, 0xd0, 0x03, 0xe8, 0xaf, 0x45, 0xc9, 0x78, 0xc0,
	0xfc, 0x82, 0xbe, 0x87, 0xeb, 0xa2, 0x98, 0xcd, 0x94, 0x70, 0x79, 0xa1, 0xd3, 0x8c, 0xdb, 0x4c,
	0xc9, 0xca, 0x5c, 0x0f, 0xcd, 0xdd, 0xa9, 0xcc, 0x9d, 0x6c, 0x50, 0xcf, 0x11, 0xd4, 0xb2, 0x7a,
	0x28, 0xba, 0xbb, 0x68, 0xfc, 0x1b, 0x81, 0xa3, 0x3f, 0xf1, 0xe8, 0x3d, 0xb8, 0xda, 0x50, 0x5f,
	0x79, 0xad, 0x7c, 0x87, 0x75, 0xf9, 0x94, 0xcf, 0x15, 0xbd, 0x0d, 0x83, 0x96, 0x37, 0x7f, 0x86,
	0xdd, 0xac, 0x16, 0xa3, 0x77, 0x21, 0x2c, 0x17, 0xce, 0xef, 0xe3, 0x41, 0xa2, 0x1e, 0x82, 0x06,
	0xe5, 0xc2, 0x21, 0x62, 0xa5, 0x3f, 0xfa, 0x4a, 0x60, 0xef, 0x6c, 0xf9, 0x7a, 0xe1, 0xfe, 0x6b,
	0xa6, 0x8f, 0x61, 0xa0, 0x6d, 0xba, 0x91, 0xaf, 0x72, 0x8d, 0x36, 0xb9, 0x76, 0xf4, 0x18, 0x68,
	0x2c, 0xe1, 0x90, 0xbe, 0x10, 0xd8, 0xbb, 0x80, 0xf8, 0x4b, 0x96, 0x0c, 0x0e, 0x1a, 0x73, 0xeb,
	0xea, 0x0e, 0x2f, 0x44, 0xd6, 0xd5, 0xa7, 0xa2, 0xb5, 0x85, 0x3e, 0xce, 0xe1, 0xc6, 0x6f, 0x09,
	0xff, 0x1e, 0xd4, 0x2f, 0xbf, 0xb2, 0x91, 0x00, 0x5a, 0xf7, 0x3e, 0x59, 0x35, 0x35, 0x39, 0xa7,
	0x11, 0xec, 0x88, 0x8c, 0x6b, 0x5d, 0x8d, 0x3c, 0x60, 0xeb, 0x25, 0xdd, 0x87, 0xbe, 0x5b, 0xa6,
	0xb9, 0xc4, 0x2e, 0x01, 0xbb, 0xec, 0x96, 0x2f, 0x24, 0xbd, 0x09, 0x50, 0x8b, 0x61, 0xb8, 0x01,
	0x6b, 0x54, 0x9e, 0xa6, 0x70, 0xbf, 0x30, 0xd3, 0x38, 0xfb, 0x5c, 0x2a, 0xe3, 0xef, 0x75, 0xfc,
	0x81, 0x4f, 0x4c, 0x2e, 0xfc, 0x95, 0xb6, 0x71, 0x55, 0x44, 0x4b, 0xe7, 0x0f, 0xa7, 0xb9, 0xcb,
	0x3e, 0x4e, 0x62, 0x51, 0xcc, 0x93, 0x06, 0x25, 0xf1, 0x94, 0xc4, 0x53, 0x92, 0xe6, 0xff, 0x61,
	0xb2, 0x8d, 0xc5, 0x47, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xae, 0xa7, 0x5f, 0x87, 0x36, 0x04,
	0x00, 0x00,
}
