// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: side/gmm/pool_params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type PoolType int32

const (
	PoolType_WEIGHT PoolType = 0
	PoolType_STABLE PoolType = 1
)

var PoolType_name = map[int32]string{
	0: "WEIGHT",
	1: "STABLE",
}

var PoolType_value = map[string]int32{
	"WEIGHT": 0,
	"STABLE": 1,
}

func (x PoolType) String() string {
	return proto.EnumName(PoolType_name, int32(x))
}

func (PoolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1be1bba6d1ee211c, []int{0}
}

type PoolParams struct {
	Type      PoolType                               `protobuf:"varint,1,opt,name=type,proto3,enum=side.gmm.PoolType" json:"type,omitempty"`
	SwapFee   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swapFee"`
	ExitFee   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=exitFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"exitFee"`
	UseOracle bool                                   `protobuf:"varint,4,opt,name=useOracle,proto3" json:"useOracle,omitempty"`
	// Amplifier parameters for stable pool.
	Amp github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=amp,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amp"`
}

func (m *PoolParams) Reset()         { *m = PoolParams{} }
func (m *PoolParams) String() string { return proto.CompactTextString(m) }
func (*PoolParams) ProtoMessage()    {}
func (*PoolParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be1bba6d1ee211c, []int{0}
}
func (m *PoolParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolParams.Merge(m, src)
}
func (m *PoolParams) XXX_Size() int {
	return m.Size()
}
func (m *PoolParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolParams.DiscardUnknown(m)
}

var xxx_messageInfo_PoolParams proto.InternalMessageInfo

func (m *PoolParams) GetType() PoolType {
	if m != nil {
		return m.Type
	}
	return PoolType_WEIGHT
}

func (m *PoolParams) GetUseOracle() bool {
	if m != nil {
		return m.UseOracle
	}
	return false
}

func init() {
	proto.RegisterEnum("side.gmm.PoolType", PoolType_name, PoolType_value)
	proto.RegisterType((*PoolParams)(nil), "side.gmm.PoolParams")
}

func init() { proto.RegisterFile("side/gmm/pool_params.proto", fileDescriptor_1be1bba6d1ee211c) }

var fileDescriptor_1be1bba6d1ee211c = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xce, 0x4c, 0x49,
	0xd5, 0x4f, 0xcf, 0xcd, 0xd5, 0x2f, 0xc8, 0xcf, 0xcf, 0x89, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xc9, 0xe9, 0xa5, 0xe7, 0xe6, 0x4a, 0x89,
	0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x05, 0xf5, 0x41, 0x2c, 0x88, 0xbc, 0xd2, 0x0a, 0x26, 0x2e, 0xae,
	0x80, 0xfc, 0xfc, 0x9c, 0x00, 0xb0, 0x26, 0x21, 0x35, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x21, 0x3d, 0x98, 0x6e, 0x3d, 0x90, 0x9a, 0x90, 0xca, 0x82,
	0xd4, 0x20, 0xb0, 0xbc, 0x90, 0x07, 0x17, 0x7b, 0x71, 0x79, 0x62, 0x81, 0x5b, 0x6a, 0xaa, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xde, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0xab, 0xa5,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0x17, 0xe7, 0xe6, 0x17,
	0x43, 0x29, 0xdd, 0xe2, 0x94, 0x6c, 0x7d, 0x90, 0xde, 0x62, 0x3d, 0x97, 0xd4, 0xe4, 0x20, 0x98,
	0x76, 0x90, 0x49, 0xa9, 0x15, 0x99, 0x25, 0x20, 0x93, 0x98, 0xc9, 0x33, 0x09, 0xaa, 0x5d, 0x48,
	0x86, 0x8b, 0xb3, 0xb4, 0x38, 0xd5, 0xbf, 0x28, 0x31, 0x39, 0x27, 0x55, 0x82, 0x45, 0x81, 0x51,
	0x83, 0x23, 0x08, 0x21, 0x20, 0xe4, 0xc0, 0xc5, 0x9c, 0x98, 0x5b, 0x20, 0xc1, 0x4a, 0xb2, 0x1d,
	0x9e, 0x79, 0x25, 0x41, 0x20, 0xad, 0x5a, 0x4a, 0x5c, 0x1c, 0xb0, 0x50, 0x10, 0xe2, 0xe2, 0x62,
	0x0b, 0x77, 0xf5, 0x74, 0xf7, 0x08, 0x11, 0x60, 0x00, 0xb1, 0x83, 0x43, 0x1c, 0x9d, 0x7c, 0x5c,
	0x05, 0x18, 0x9d, 0x9c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x03,
	0xc9, 0x2a, 0x50, 0xa8, 0x82, 0x83, 0x3f, 0x39, 0x3f, 0x07, 0xcc, 0xd1, 0xaf, 0x00, 0x47, 0x1f,
	0xd8, 0xc2, 0x24, 0x36, 0xb0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xf1, 0x82, 0xc9,
	0xd7, 0x01, 0x00, 0x00,
}

func (m *PoolParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amp.Size()
		i -= size
		if _, err := m.Amp.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.UseOracle {
		i--
		if m.UseOracle {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.ExitFee.Size()
		i -= size
		if _, err := m.ExitFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Type != 0 {
		i = encodeVarintPoolParams(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoolParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovPoolParams(uint64(m.Type))
	}
	l = m.SwapFee.Size()
	n += 1 + l + sovPoolParams(uint64(l))
	l = m.ExitFee.Size()
	n += 1 + l + sovPoolParams(uint64(l))
	if m.UseOracle {
		n += 2
	}
	l = m.Amp.Size()
	n += 1 + l + sovPoolParams(uint64(l))
	return n
}

func sovPoolParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolParams(x uint64) (n int) {
	return sovPoolParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolParams
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
			return fmt.Errorf("proto: PoolParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= PoolType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
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
				return ErrInvalidLengthPoolParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
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
				return ErrInvalidLengthPoolParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExitFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseOracle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UseOracle = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
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
				return ErrInvalidLengthPoolParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolParams
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
func skipPoolParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolParams
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
					return 0, ErrIntOverflowPoolParams
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
					return 0, ErrIntOverflowPoolParams
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
				return 0, ErrInvalidLengthPoolParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolParams = fmt.Errorf("proto: unexpected end of group")
)
