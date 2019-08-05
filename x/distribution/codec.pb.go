// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/distribution/codec.proto

package distribution

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_iov_one_weave "github.com/iov-one/weave"
	weave "github.com/iov-one/weave"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Revenue represents an account with funds collected from the fees. This is a
// temporary account used for storing fees that are later distributed between
// the owners.
type Revenue struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Admin key belongs to the governance entities. It can be used to transfer
	// stored amount to an another account.
	// While not enforced it is best to use a multisig contract here.
	Admin github_com_iov_one_weave.Address `protobuf:"bytes,2,opt,name=admin,proto3,casttype=github.com/iov-one/weave.Address" json:"admin,omitempty"`
	// Destinations holds any number of addresses that the collected revenue is
	// distributed to. Must be at least one.
	Destinations []*Destination `protobuf:"bytes,3,rep,name=destinations,proto3" json:"destinations,omitempty"`
	// Address of this entity. Set during creation and does not change.
	Address github_com_iov_one_weave.Address `protobuf:"bytes,4,opt,name=address,proto3,casttype=github.com/iov-one/weave.Address" json:"address,omitempty"`
}

func (m *Revenue) Reset()         { *m = Revenue{} }
func (m *Revenue) String() string { return proto.CompactTextString(m) }
func (*Revenue) ProtoMessage()    {}
func (*Revenue) Descriptor() ([]byte, []int) {
	return fileDescriptor_186299c22854933b, []int{0}
}
func (m *Revenue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Revenue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Revenue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Revenue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Revenue.Merge(m, src)
}
func (m *Revenue) XXX_Size() int {
	return m.Size()
}
func (m *Revenue) XXX_DiscardUnknown() {
	xxx_messageInfo_Revenue.DiscardUnknown(m)
}

var xxx_messageInfo_Revenue proto.InternalMessageInfo

func (m *Revenue) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Revenue) GetAdmin() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Admin
	}
	return nil
}

func (m *Revenue) GetDestinations() []*Destination {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *Revenue) GetAddress() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type Destination struct {
	// An address that the funds should be transferred to.
	// This should not be the validator addresses, as the keys used to sign
	// blocks should never be in a wallet. This can be the wallets of the admins
	// of the validators.
	Address github_com_iov_one_weave.Address `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/iov-one/weave.Address" json:"address,omitempty"`
	// Weight defines what part of the total revenue goes to this destination.
	// Each destination receives part of the total revenue amount proportional to
	// the weight. For example, if there are two destinations with weights 1 and 2
	// accordingly, distribution will be 1/3 to the first address and 2/3 to the
	// second one.
	Weight int32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (m *Destination) Reset()         { *m = Destination{} }
func (m *Destination) String() string { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()    {}
func (*Destination) Descriptor() ([]byte, []int) {
	return fileDescriptor_186299c22854933b, []int{1}
}
func (m *Destination) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Destination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Destination.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Destination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Destination.Merge(m, src)
}
func (m *Destination) XXX_Size() int {
	return m.Size()
}
func (m *Destination) XXX_DiscardUnknown() {
	xxx_messageInfo_Destination.DiscardUnknown(m)
}

var xxx_messageInfo_Destination proto.InternalMessageInfo

func (m *Destination) GetAddress() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Destination) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

// CreateMsg is issuing the creation of a new revenue stream instance.
type CreateMsg struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Admin key belongs to the governance entities. It can be used to transfer
	// stored amount to an another account.
	// While not enforced it is best to use a multisig contract here.
	Admin github_com_iov_one_weave.Address `protobuf:"bytes,2,opt,name=admin,proto3,casttype=github.com/iov-one/weave.Address" json:"admin,omitempty"`
	// Destinations holds any number of addresses that the collected revenue is
	// distributed to. Must be at least one.
	Destinations []*Destination `protobuf:"bytes,3,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (m *CreateMsg) Reset()         { *m = CreateMsg{} }
func (m *CreateMsg) String() string { return proto.CompactTextString(m) }
func (*CreateMsg) ProtoMessage()    {}
func (*CreateMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_186299c22854933b, []int{2}
}
func (m *CreateMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMsg.Merge(m, src)
}
func (m *CreateMsg) XXX_Size() int {
	return m.Size()
}
func (m *CreateMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMsg proto.InternalMessageInfo

func (m *CreateMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateMsg) GetAdmin() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Admin
	}
	return nil
}

func (m *CreateMsg) GetDestinations() []*Destination {
	if m != nil {
		return m.Destinations
	}
	return nil
}

// DistributeMsg is a request to distribute all funds collected within a single
// revenue instance. Revenue is distributed between destinations. Request must be
// signed using admin key.
type DistributeMsg struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Revenue ID reference an ID of a revenue instance that the collected fees
	// should be distributed between destinations.
	RevenueID []byte `protobuf:"bytes,2,opt,name=revenue_id,json=revenueId,proto3" json:"revenue_id,omitempty"`
}

func (m *DistributeMsg) Reset()         { *m = DistributeMsg{} }
func (m *DistributeMsg) String() string { return proto.CompactTextString(m) }
func (*DistributeMsg) ProtoMessage()    {}
func (*DistributeMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_186299c22854933b, []int{3}
}
func (m *DistributeMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributeMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributeMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributeMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributeMsg.Merge(m, src)
}
func (m *DistributeMsg) XXX_Size() int {
	return m.Size()
}
func (m *DistributeMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributeMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DistributeMsg proto.InternalMessageInfo

func (m *DistributeMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *DistributeMsg) GetRevenueID() []byte {
	if m != nil {
		return m.RevenueID
	}
	return nil
}

// ResetMsg change the configuration of a revenue instance.
// To assure destinations that they will receive money, every revenue update is
// forcing funds distribution. Before applying any change all funds stored by
// the revenue account are distributed using old configuration. Only when the
// collected revenue amount is equal to zero the change is applied.
type ResetMsg struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Revenue ID reference an ID of a revenue instance that is updated.
	RevenueID []byte `protobuf:"bytes,2,opt,name=revenue_id,json=revenueId,proto3" json:"revenue_id,omitempty"`
	// Destinations holds any number of addresses that the collected revenue is
	// distributed to. Must be at least one.
	Destinations []*Destination `protobuf:"bytes,3,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (m *ResetMsg) Reset()         { *m = ResetMsg{} }
func (m *ResetMsg) String() string { return proto.CompactTextString(m) }
func (*ResetMsg) ProtoMessage()    {}
func (*ResetMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_186299c22854933b, []int{4}
}
func (m *ResetMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResetMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResetMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResetMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetMsg.Merge(m, src)
}
func (m *ResetMsg) XXX_Size() int {
	return m.Size()
}
func (m *ResetMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ResetMsg proto.InternalMessageInfo

func (m *ResetMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ResetMsg) GetRevenueID() []byte {
	if m != nil {
		return m.RevenueID
	}
	return nil
}

func (m *ResetMsg) GetDestinations() []*Destination {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func init() {
	proto.RegisterType((*Revenue)(nil), "distribution.Revenue")
	proto.RegisterType((*Destination)(nil), "distribution.Destination")
	proto.RegisterType((*CreateMsg)(nil), "distribution.CreateMsg")
	proto.RegisterType((*DistributeMsg)(nil), "distribution.DistributeMsg")
	proto.RegisterType((*ResetMsg)(nil), "distribution.ResetMsg")
}

func init() { proto.RegisterFile("x/distribution/codec.proto", fileDescriptor_186299c22854933b) }

var fileDescriptor_186299c22854933b = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xee, 0x5a, 0xfb, 0x37, 0x69, 0x11, 0x82, 0x48, 0xec, 0x21, 0x0d, 0xc1, 0x43, 0x41, 0x4d,
	0xa0, 0xde, 0x04, 0x05, 0x6b, 0x2f, 0x3d, 0xf4, 0xb2, 0x2f, 0x20, 0xdb, 0xee, 0x90, 0xae, 0xd0,
	0xac, 0x64, 0xb7, 0xad, 0x8f, 0xe1, 0x43, 0xf8, 0x06, 0xbe, 0x84, 0xc7, 0x1e, 0x3d, 0x15, 0x49,
	0x9f, 0xc0, 0xab, 0x27, 0x69, 0x92, 0x6a, 0x3c, 0x56, 0xf1, 0xe0, 0x6d, 0x76, 0xe6, 0xfb, 0xbe,
	0xf9, 0xf8, 0x98, 0x85, 0xe6, 0xbd, 0xcf, 0x85, 0xd2, 0x91, 0x18, 0x4e, 0xb5, 0x90, 0xa1, 0x3f,
	0x92, 0x1c, 0x47, 0xde, 0x5d, 0x24, 0xb5, 0x34, 0xeb, 0xf9, 0x49, 0xd3, 0xc8, 0x8d, 0x9a, 0xfb,
	0x81, 0x0c, 0x64, 0x52, 0xfa, 0xeb, 0x2a, 0xed, 0xba, 0x6f, 0x04, 0x2a, 0x14, 0x67, 0x18, 0x4e,
	0xd1, 0x3c, 0x86, 0xea, 0x04, 0x35, 0xe3, 0x4c, 0x33, 0x8b, 0x38, 0xa4, 0x6d, 0x74, 0xf6, 0xbc,
	0x39, 0xb2, 0x19, 0x7a, 0x83, 0xac, 0x4d, 0x3f, 0x01, 0xe6, 0x39, 0x94, 0x18, 0x9f, 0x88, 0xd0,
	0xda, 0x71, 0x48, 0xbb, 0xde, 0x3d, 0x7a, 0x5f, 0xb6, 0x9c, 0x40, 0xe8, 0xf1, 0x74, 0xe8, 0x8d,
	0xe4, 0xc4, 0x17, 0x72, 0x76, 0x2a, 0x43, 0xf4, 0x53, 0xfe, 0x15, 0xe7, 0x11, 0x2a, 0x45, 0x53,
	0x8a, 0x79, 0x01, 0x75, 0x8e, 0x4a, 0x8b, 0x90, 0xad, 0x6d, 0x2a, 0xab, 0xe8, 0x14, 0xdb, 0x46,
	0xe7, 0xd0, 0xcb, 0x9b, 0xf7, 0x7a, 0x5f, 0x08, 0xfa, 0x0d, 0x6e, 0x5e, 0x42, 0x85, 0xa5, 0x82,
	0xd6, 0xee, 0x16, 0xcb, 0x37, 0x24, 0x17, 0xc1, 0xc8, 0x89, 0xe7, 0xe5, 0xc8, 0x0f, 0xe4, 0xcc,
	0x03, 0x28, 0xcf, 0x51, 0x04, 0x63, 0x9d, 0x44, 0x51, 0xa2, 0xd9, 0xcb, 0x7d, 0x22, 0x50, 0xbb,
	0x8e, 0x90, 0x69, 0x1c, 0xa8, 0xe0, 0xbf, 0x84, 0xeb, 0xde, 0x42, 0xa3, 0xb7, 0x41, 0x6e, 0x6f,
	0xfc, 0x04, 0x20, 0x4a, 0xaf, 0xe9, 0x46, 0xf0, 0xcc, 0x7d, 0x23, 0x5e, 0xb6, 0x6a, 0xd9, 0x8d,
	0xf5, 0x7b, 0xb4, 0x96, 0x01, 0xfa, 0xdc, 0x7d, 0x24, 0x50, 0xa5, 0xa8, 0x50, 0xff, 0xed, 0x9e,
	0x5f, 0x46, 0xd2, 0xb5, 0x9e, 0x63, 0x9b, 0x2c, 0x62, 0x9b, 0xbc, 0xc6, 0x36, 0x79, 0x58, 0xd9,
	0x85, 0xc5, 0xca, 0x2e, 0xbc, 0xac, 0xec, 0xc2, 0xb0, 0x9c, 0x7c, 0xa2, 0xb3, 0x8f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xad, 0xdc, 0xe2, 0x23, 0x93, 0x03, 0x00, 0x00,
}

func (m *Revenue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Revenue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Admin) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Admin)))
		i += copy(dAtA[i:], m.Admin)
	}
	if len(m.Destinations) > 0 {
		for _, msg := range m.Destinations {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	return i, nil
}

func (m *Destination) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Destination) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if m.Weight != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Weight))
	}
	return i, nil
}

func (m *CreateMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n2, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Admin) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Admin)))
		i += copy(dAtA[i:], m.Admin)
	}
	if len(m.Destinations) > 0 {
		for _, msg := range m.Destinations {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *DistributeMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributeMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n3, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.RevenueID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.RevenueID)))
		i += copy(dAtA[i:], m.RevenueID)
	}
	return i, nil
}

func (m *ResetMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResetMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n4, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.RevenueID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.RevenueID)))
		i += copy(dAtA[i:], m.RevenueID)
	}
	if len(m.Destinations) > 0 {
		for _, msg := range m.Destinations {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Revenue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Destinations) > 0 {
		for _, e := range m.Destinations {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *Destination) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Weight != 0 {
		n += 1 + sovCodec(uint64(m.Weight))
	}
	return n
}

func (m *CreateMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Destinations) > 0 {
		for _, e := range m.Destinations {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *DistributeMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.RevenueID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *ResetMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.RevenueID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Destinations) > 0 {
		for _, e := range m.Destinations {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Revenue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Revenue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Revenue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = append(m.Admin[:0], dAtA[iNdEx:postIndex]...)
			if m.Admin == nil {
				m.Admin = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destinations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destinations = append(m.Destinations, &Destination{})
			if err := m.Destinations[len(m.Destinations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Destination) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Destination: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Destination: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = append(m.Admin[:0], dAtA[iNdEx:postIndex]...)
			if m.Admin == nil {
				m.Admin = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destinations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destinations = append(m.Destinations, &Destination{})
			if err := m.Destinations[len(m.Destinations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DistributeMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DistributeMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributeMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevenueID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevenueID = append(m.RevenueID[:0], dAtA[iNdEx:postIndex]...)
			if m.RevenueID == nil {
				m.RevenueID = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResetMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResetMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResetMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevenueID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevenueID = append(m.RevenueID[:0], dAtA[iNdEx:postIndex]...)
			if m.RevenueID == nil {
				m.RevenueID = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destinations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destinations = append(m.Destinations, &Destination{})
			if err := m.Destinations[len(m.Destinations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCodec
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)
