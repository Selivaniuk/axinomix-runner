// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axinomixrunner/axinomixrunner/leaderboard.proto

package types

import (
	fmt "fmt"
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

type Leaderboard struct {
	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PlayerAddress string `protobuf:"bytes,2,opt,name=playerAddress,proto3" json:"playerAddress,omitempty"`
	Score         uint64 `protobuf:"varint,3,opt,name=Score,proto3" json:"Score,omitempty"`
	PlayerRank    uint64 `protobuf:"varint,4,opt,name=playerRank,proto3" json:"playerRank,omitempty"`
}

func (m *Leaderboard) Reset()         { *m = Leaderboard{} }
func (m *Leaderboard) String() string { return proto.CompactTextString(m) }
func (*Leaderboard) ProtoMessage()    {}
func (*Leaderboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a94dd5c055c0a4a, []int{0}
}
func (m *Leaderboard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Leaderboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Leaderboard.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Leaderboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Leaderboard.Merge(m, src)
}
func (m *Leaderboard) XXX_Size() int {
	return m.Size()
}
func (m *Leaderboard) XXX_DiscardUnknown() {
	xxx_messageInfo_Leaderboard.DiscardUnknown(m)
}

var xxx_messageInfo_Leaderboard proto.InternalMessageInfo

func (m *Leaderboard) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Leaderboard) GetPlayerAddress() string {
	if m != nil {
		return m.PlayerAddress
	}
	return ""
}

func (m *Leaderboard) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Leaderboard) GetPlayerRank() uint64 {
	if m != nil {
		return m.PlayerRank
	}
	return 0
}

func init() {
	proto.RegisterType((*Leaderboard)(nil), "axinomixrunner.axinomixrunner.Leaderboard")
}

func init() {
	proto.RegisterFile("axinomixrunner/axinomixrunner/leaderboard.proto", fileDescriptor_8a94dd5c055c0a4a)
}

var fileDescriptor_8a94dd5c055c0a4a = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xac, 0xc8, 0xcc,
	0xcb, 0xcf, 0xcd, 0xac, 0x28, 0x2a, 0xcd, 0xcb, 0x4b, 0x2d, 0x42, 0xe7, 0xe6, 0xa4, 0x26, 0xa6,
	0xa4, 0x16, 0x25, 0xe5, 0x27, 0x16, 0xa5, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xc9, 0xa2,
	0xaa, 0xd0, 0x43, 0xe5, 0x2a, 0x55, 0x72, 0x71, 0xfb, 0x20, 0xf4, 0x08, 0xf1, 0x71, 0x31, 0x65,
	0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x31, 0x65, 0xa6, 0x08, 0xa9, 0x70, 0xf1, 0x16,
	0xe4, 0x24, 0x56, 0xa6, 0x16, 0x39, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x29, 0x30,
	0x6a, 0x70, 0x06, 0xa1, 0x0a, 0x0a, 0x89, 0x70, 0xb1, 0x06, 0x27, 0xe7, 0x17, 0xa5, 0x4a, 0x30,
	0x83, 0x35, 0x42, 0x38, 0x42, 0x72, 0x5c, 0x5c, 0x10, 0x65, 0x41, 0x89, 0x79, 0xd9, 0x12, 0x2c,
	0x60, 0x29, 0x24, 0x11, 0x27, 0x87, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0,
	0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88,
	0x52, 0x83, 0x39, 0x52, 0x17, 0xea, 0xad, 0x0a, 0x74, 0x7f, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27,
	0xb1, 0x81, 0xbd, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xc4, 0xce, 0xe1, 0x15, 0x01,
	0x00, 0x00,
}

func (m *Leaderboard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Leaderboard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Leaderboard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlayerRank != 0 {
		i = encodeVarintLeaderboard(dAtA, i, uint64(m.PlayerRank))
		i--
		dAtA[i] = 0x20
	}
	if m.Score != 0 {
		i = encodeVarintLeaderboard(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PlayerAddress) > 0 {
		i -= len(m.PlayerAddress)
		copy(dAtA[i:], m.PlayerAddress)
		i = encodeVarintLeaderboard(dAtA, i, uint64(len(m.PlayerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLeaderboard(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLeaderboard(dAtA []byte, offset int, v uint64) int {
	offset -= sovLeaderboard(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Leaderboard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLeaderboard(uint64(m.Id))
	}
	l = len(m.PlayerAddress)
	if l > 0 {
		n += 1 + l + sovLeaderboard(uint64(l))
	}
	if m.Score != 0 {
		n += 1 + sovLeaderboard(uint64(m.Score))
	}
	if m.PlayerRank != 0 {
		n += 1 + sovLeaderboard(uint64(m.PlayerRank))
	}
	return n
}

func sovLeaderboard(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLeaderboard(x uint64) (n int) {
	return sovLeaderboard(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Leaderboard) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLeaderboard
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
			return fmt.Errorf("proto: Leaderboard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Leaderboard: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaderboard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaderboard
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
				return ErrInvalidLengthLeaderboard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLeaderboard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaderboard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerRank", wireType)
			}
			m.PlayerRank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaderboard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlayerRank |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLeaderboard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLeaderboard
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
func skipLeaderboard(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLeaderboard
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
					return 0, ErrIntOverflowLeaderboard
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
					return 0, ErrIntOverflowLeaderboard
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
				return 0, ErrInvalidLengthLeaderboard
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLeaderboard
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLeaderboard
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLeaderboard        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLeaderboard          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLeaderboard = fmt.Errorf("proto: unexpected end of group")
)
