// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/auth/module/v1/module.proto

package module

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Module struct {
	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// account_constructor is an optional AccountI constructor config object that can be provided
	// to override the default BaseAccount constructor. The provided config object must have
	// an `NewAccount() AccountI` method defined. If this is left empty, the default constructor
	// will be used
	AccountConstructor            *types.Any `protobuf:"bytes,2,opt,name=account_constructor,json=accountConstructor,proto3" json:"account_constructor,omitempty"`
	RandomGenesisAccountsProvider *types.Any `protobuf:"bytes,3,opt,name=random_genesis_accounts_provider,json=randomGenesisAccountsProvider,proto3" json:"random_genesis_accounts_provider,omitempty"`
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7f34be9f8952c0, []int{0}
}
func (m *Module) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Module) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Module.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Module) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Module.Merge(m, src)
}
func (m *Module) XXX_Size() int {
	return m.Size()
}
func (m *Module) XXX_DiscardUnknown() {
	xxx_messageInfo_Module.DiscardUnknown(m)
}

var xxx_messageInfo_Module proto.InternalMessageInfo

func (m *Module) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Module) GetAccountConstructor() *types.Any {
	if m != nil {
		return m.AccountConstructor
	}
	return nil
}

func (m *Module) GetRandomGenesisAccountsProvider() *types.Any {
	if m != nil {
		return m.RandomGenesisAccountsProvider
	}
	return nil
}

type Permission struct {
	Address     string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Permissions []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7f34be9f8952c0, []int{1}
}
func (m *Permission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return m.Size()
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Permission) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type DefaultAccountConstructor struct {
}

func (m *DefaultAccountConstructor) Reset()         { *m = DefaultAccountConstructor{} }
func (m *DefaultAccountConstructor) String() string { return proto.CompactTextString(m) }
func (*DefaultAccountConstructor) ProtoMessage()    {}
func (*DefaultAccountConstructor) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7f34be9f8952c0, []int{2}
}
func (m *DefaultAccountConstructor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DefaultAccountConstructor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefaultAccountConstructor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DefaultAccountConstructor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultAccountConstructor.Merge(m, src)
}
func (m *DefaultAccountConstructor) XXX_Size() int {
	return m.Size()
}
func (m *DefaultAccountConstructor) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultAccountConstructor.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultAccountConstructor proto.InternalMessageInfo

type DefaultRandomGenesisAccountsProvider struct {
}

func (m *DefaultRandomGenesisAccountsProvider) Reset()         { *m = DefaultRandomGenesisAccountsProvider{} }
func (m *DefaultRandomGenesisAccountsProvider) String() string { return proto.CompactTextString(m) }
func (*DefaultRandomGenesisAccountsProvider) ProtoMessage()    {}
func (*DefaultRandomGenesisAccountsProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7f34be9f8952c0, []int{3}
}
func (m *DefaultRandomGenesisAccountsProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DefaultRandomGenesisAccountsProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefaultRandomGenesisAccountsProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DefaultRandomGenesisAccountsProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultRandomGenesisAccountsProvider.Merge(m, src)
}
func (m *DefaultRandomGenesisAccountsProvider) XXX_Size() int {
	return m.Size()
}
func (m *DefaultRandomGenesisAccountsProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultRandomGenesisAccountsProvider.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultRandomGenesisAccountsProvider proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Module)(nil), "cosmos.auth.module.v1.Module")
	proto.RegisterType((*Permission)(nil), "cosmos.auth.module.v1.Permission")
	proto.RegisterType((*DefaultAccountConstructor)(nil), "cosmos.auth.module.v1.DefaultAccountConstructor")
	proto.RegisterType((*DefaultRandomGenesisAccountsProvider)(nil), "cosmos.auth.module.v1.DefaultRandomGenesisAccountsProvider")
}

func init() {
	proto.RegisterFile("cosmos/auth/module/v1/module.proto", fileDescriptor_0f7f34be9f8952c0)
}

var fileDescriptor_0f7f34be9f8952c0 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0x87, 0x1d, 0x05, 0x2f, 0x8e, 0xbb, 0xb9, 0xf7, 0x42, 0xbc, 0x97, 0x86, 0x34, 0x94, 0x12,
	0x0a, 0x9d, 0x41, 0xfb, 0x04, 0x56, 0x4b, 0xbb, 0x29, 0x48, 0x96, 0x85, 0x12, 0x62, 0x32, 0xc6,
	0x50, 0x33, 0x47, 0xe6, 0x8f, 0xd4, 0xb7, 0xe8, 0x63, 0x75, 0xe9, 0xb2, 0xcb, 0xa2, 0x8f, 0xd0,
	0x17, 0x28, 0x9a, 0xb1, 0x2a, 0x2d, 0xae, 0x92, 0xe1, 0x7c, 0xdf, 0xe1, 0xfc, 0xce, 0x0c, 0xf6,
	0x13, 0x50, 0x05, 0x28, 0x16, 0x1b, 0x3d, 0x66, 0x05, 0xa4, 0x66, 0xc2, 0xd9, 0xac, 0x6d, 0xff,
	0xe8, 0x54, 0x82, 0x06, 0xf2, 0xb7, 0x64, 0xe8, 0x9a, 0xa1, 0xb6, 0x32, 0x6b, 0xff, 0x6b, 0x65,
	0x00, 0xd9, 0x84, 0xb3, 0x0d, 0x34, 0x34, 0x23, 0x16, 0x8b, 0x79, 0x69, 0xf8, 0x1f, 0x08, 0xd7,
	0xef, 0x37, 0x20, 0xe9, 0xe1, 0xe6, 0x94, 0xcb, 0x22, 0x57, 0x2a, 0x07, 0xa1, 0x1c, 0xe4, 0xd5,
	0x82, 0x66, 0xe7, 0x94, 0xfe, 0xd8, 0x92, 0x0e, 0xbe, 0xc8, 0x70, 0xdf, 0x22, 0x37, 0xf8, 0x77,
	0x9c, 0x24, 0x60, 0x84, 0x8e, 0x12, 0x10, 0x4a, 0x4b, 0x93, 0x68, 0x90, 0x4e, 0xd5, 0x43, 0x41,
	0xb3, 0xf3, 0x87, 0x96, 0x83, 0xd0, 0xed, 0x20, 0xb4, 0x2b, 0xe6, 0x21, 0xb1, 0x42, 0x6f, 0xc7,
	0x93, 0x47, 0xec, 0xc9, 0x58, 0xa4, 0x50, 0x44, 0x19, 0x17, 0x5c, 0xe5, 0x2a, 0xb2, 0x90, 0x8a,
	0xa6, 0x12, 0x66, 0x79, 0xca, 0xa5, 0x53, 0x3b, 0xd2, 0xf3, 0xa4, 0xb4, 0x6f, 0x4b, 0xb9, 0x6b,
	0xdd, 0x81, 0x55, 0xfd, 0x3b, 0x8c, 0x77, 0x01, 0x88, 0x83, 0x7f, 0xc5, 0x69, 0x2a, 0xb9, 0x5a,
	0x87, 0x46, 0x41, 0x23, 0xdc, 0x1e, 0x89, 0x77, 0xb8, 0x92, 0xaa, 0x57, 0x0b, 0x1a, 0x07, 0x79,
	0xfd, 0xff, 0xb8, 0xd5, 0xe7, 0xa3, 0xd8, 0x4c, 0x74, 0xf7, 0x5b, 0x0a, 0xff, 0x1c, 0x9f, 0xd9,
	0x62, 0x78, 0x6c, 0x9c, 0xeb, 0xfe, 0xeb, 0xd2, 0x45, 0x8b, 0xa5, 0x8b, 0xde, 0x97, 0x2e, 0x7a,
	0x59, 0xb9, 0x95, 0xc5, 0xca, 0xad, 0xbc, 0xad, 0xdc, 0xca, 0xc3, 0x45, 0x96, 0xeb, 0xb1, 0x19,
	0xd2, 0x04, 0x0a, 0x66, 0xef, 0xbf, 0xfc, 0x5c, 0xaa, 0xf4, 0x89, 0x3d, 0xef, 0x3f, 0x86, 0x61,
	0x7d, 0xb3, 0x81, 0xab, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x8b, 0x32, 0x17, 0x29, 0x02,
	0x00, 0x00,
}

func (m *Module) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Module) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Module) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RandomGenesisAccountsProvider != nil {
		{
			size, err := m.RandomGenesisAccountsProvider.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModule(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.AccountConstructor != nil {
		{
			size, err := m.AccountConstructor.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModule(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Permissions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintModule(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Permission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Permission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Permission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Permissions[iNdEx])
			copy(dAtA[i:], m.Permissions[iNdEx])
			i = encodeVarintModule(dAtA, i, uint64(len(m.Permissions[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintModule(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DefaultAccountConstructor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefaultAccountConstructor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultAccountConstructor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *DefaultRandomGenesisAccountsProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefaultRandomGenesisAccountsProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultRandomGenesisAccountsProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintModule(dAtA []byte, offset int, v uint64) int {
	offset -= sovModule(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Module) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for _, e := range m.Permissions {
			l = e.Size()
			n += 1 + l + sovModule(uint64(l))
		}
	}
	if m.AccountConstructor != nil {
		l = m.AccountConstructor.Size()
		n += 1 + l + sovModule(uint64(l))
	}
	if m.RandomGenesisAccountsProvider != nil {
		l = m.RandomGenesisAccountsProvider.Size()
		n += 1 + l + sovModule(uint64(l))
	}
	return n
}

func (m *Permission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovModule(uint64(l))
	}
	if len(m.Permissions) > 0 {
		for _, s := range m.Permissions {
			l = len(s)
			n += 1 + l + sovModule(uint64(l))
		}
	}
	return n
}

func (m *DefaultAccountConstructor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *DefaultRandomGenesisAccountsProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovModule(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModule(x uint64) (n int) {
	return sovModule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Module) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModule
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
			return fmt.Errorf("proto: Module: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Module: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
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
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, &Permission{})
			if err := m.Permissions[len(m.Permissions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountConstructor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
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
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AccountConstructor == nil {
				m.AccountConstructor = &types.Any{}
			}
			if err := m.AccountConstructor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomGenesisAccountsProvider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
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
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RandomGenesisAccountsProvider == nil {
				m.RandomGenesisAccountsProvider = &types.Any{}
			}
			if err := m.RandomGenesisAccountsProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModule
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
func (m *Permission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModule
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
			return fmt.Errorf("proto: Permission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Permission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModule
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
func (m *DefaultAccountConstructor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModule
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
			return fmt.Errorf("proto: DefaultAccountConstructor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefaultAccountConstructor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipModule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModule
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
func (m *DefaultRandomGenesisAccountsProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModule
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
			return fmt.Errorf("proto: DefaultRandomGenesisAccountsProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefaultRandomGenesisAccountsProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipModule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModule
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
func skipModule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModule
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
					return 0, ErrIntOverflowModule
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowModule
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
				return 0, ErrInvalidLengthModule
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModule
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModule
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModule        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModule          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModule = fmt.Errorf("proto: unexpected end of group")
)
