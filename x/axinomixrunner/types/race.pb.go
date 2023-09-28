// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axinomixrunner/axinomixrunner/race.proto

package types

import (
	encoding_binary "encoding/binary"
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

type Race struct {
	Id            uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PlayerAddress string  `protobuf:"bytes,2,opt,name=playerAddress,proto3" json:"playerAddress,omitempty"`
	Bet           string  `protobuf:"bytes,3,opt,name=bet,proto3" json:"bet,omitempty"`
	Multiplier    float32 `protobuf:"fixed32,4,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
	StartTime     uint64  `protobuf:"varint,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime       uint64  `protobuf:"varint,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
	CoinsEarned   string  `protobuf:"bytes,7,opt,name=coinsEarned,proto3" json:"coinsEarned,omitempty"`
	Score         uint64  `protobuf:"varint,8,opt,name=score,proto3" json:"score,omitempty"`
	State         string  `protobuf:"bytes,9,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *Race) Reset()         { *m = Race{} }
func (m *Race) String() string { return proto.CompactTextString(m) }
func (*Race) ProtoMessage()    {}
func (*Race) Descriptor() ([]byte, []int) {
	return fileDescriptor_208fb1fbcd391512, []int{0}
}
func (m *Race) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Race) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Race.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Race) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Race.Merge(m, src)
}
func (m *Race) XXX_Size() int {
	return m.Size()
}
func (m *Race) XXX_DiscardUnknown() {
	xxx_messageInfo_Race.DiscardUnknown(m)
}

var xxx_messageInfo_Race proto.InternalMessageInfo

func (m *Race) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Race) GetPlayerAddress() string {
	if m != nil {
		return m.PlayerAddress
	}
	return ""
}

func (m *Race) GetBet() string {
	if m != nil {
		return m.Bet
	}
	return ""
}

func (m *Race) GetMultiplier() float32 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

func (m *Race) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Race) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Race) GetCoinsEarned() string {
	if m != nil {
		return m.CoinsEarned
	}
	return ""
}

func (m *Race) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Race) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*Race)(nil), "axinomixrunner.axinomixrunner.Race")
}

func init() {
	proto.RegisterFile("axinomixrunner/axinomixrunner/race.proto", fileDescriptor_208fb1fbcd391512)
}

var fileDescriptor_208fb1fbcd391512 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x33, 0x69, 0xda, 0x9a, 0x27, 0x8a, 0x0c, 0x2e, 0x66, 0xa1, 0x43, 0x10, 0x91, 0x6c,
	0xac, 0x0b, 0x2f, 0xa0, 0x82, 0x17, 0x08, 0xae, 0xdc, 0x4d, 0x33, 0x6f, 0x31, 0x90, 0x4c, 0xc2,
	0x9b, 0x29, 0xa4, 0xb7, 0xf0, 0x58, 0x2e, 0xbb, 0x74, 0x29, 0xc9, 0x15, 0x3c, 0x80, 0x38, 0xb5,
	0xb4, 0xe9, 0xee, 0x7d, 0xdf, 0xcf, 0xe3, 0x87, 0x1f, 0x72, 0xd5, 0x19, 0xdb, 0xd4, 0xa6, 0xa3,
	0x95, 0xb5, 0x48, 0x0f, 0x47, 0x48, 0xaa, 0xc4, 0x45, 0x4b, 0x8d, 0x6f, 0xf8, 0xf5, 0x38, 0x5a,
	0x8c, 0xf1, 0xe6, 0x87, 0x41, 0x52, 0xa8, 0x12, 0xf9, 0x39, 0xc4, 0x46, 0x0b, 0x96, 0xb1, 0x3c,
	0x29, 0x62, 0xa3, 0xf9, 0x2d, 0x9c, 0xb5, 0x95, 0x5a, 0x23, 0x3d, 0x6b, 0x4d, 0xe8, 0x9c, 0x88,
	0x33, 0x96, 0xa7, 0xc5, 0x58, 0xf2, 0x0b, 0x98, 0x2c, 0xd1, 0x8b, 0x49, 0xc8, 0xfe, 0x4e, 0x2e,
	0x01, 0xea, 0x55, 0xe5, 0x4d, 0x5b, 0x19, 0x24, 0x91, 0x64, 0x2c, 0x8f, 0x8b, 0x03, 0xc3, 0xaf,
	0x20, 0x75, 0x5e, 0x91, 0x7f, 0x33, 0x35, 0x8a, 0x69, 0xa8, 0xdb, 0x0b, 0x2e, 0x60, 0x8e, 0x56,
	0x87, 0x6c, 0x16, 0xb2, 0x1d, 0xf2, 0x0c, 0x4e, 0xcb, 0xc6, 0x58, 0xf7, 0xaa, 0xc8, 0xa2, 0x16,
	0xf3, 0xd0, 0x78, 0xa8, 0xf8, 0x25, 0x4c, 0x5d, 0xd9, 0x10, 0x8a, 0x93, 0xf0, 0xb9, 0x85, 0x60,
	0xbd, 0xf2, 0x28, 0xd2, 0xf0, 0xb1, 0x85, 0x97, 0xa7, 0xcf, 0x5e, 0xb2, 0x4d, 0x2f, 0xd9, 0x77,
	0x2f, 0xd9, 0xc7, 0x20, 0xa3, 0xcd, 0x20, 0xa3, 0xaf, 0x41, 0x46, 0xef, 0x77, 0xbb, 0x81, 0xee,
	0xff, 0xb7, 0xec, 0x8e, 0xc7, 0xf5, 0xeb, 0x16, 0xdd, 0x72, 0x16, 0xe6, 0x7d, 0xfc, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x4e, 0x0c, 0xbe, 0x6f, 0x8a, 0x01, 0x00, 0x00,
}

func (m *Race) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Race) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Race) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintRace(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Score != 0 {
		i = encodeVarintRace(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x40
	}
	if len(m.CoinsEarned) > 0 {
		i -= len(m.CoinsEarned)
		copy(dAtA[i:], m.CoinsEarned)
		i = encodeVarintRace(dAtA, i, uint64(len(m.CoinsEarned)))
		i--
		dAtA[i] = 0x3a
	}
	if m.EndTime != 0 {
		i = encodeVarintRace(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x30
	}
	if m.StartTime != 0 {
		i = encodeVarintRace(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x28
	}
	if m.Multiplier != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Multiplier))))
		i--
		dAtA[i] = 0x25
	}
	if len(m.Bet) > 0 {
		i -= len(m.Bet)
		copy(dAtA[i:], m.Bet)
		i = encodeVarintRace(dAtA, i, uint64(len(m.Bet)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PlayerAddress) > 0 {
		i -= len(m.PlayerAddress)
		copy(dAtA[i:], m.PlayerAddress)
		i = encodeVarintRace(dAtA, i, uint64(len(m.PlayerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintRace(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRace(dAtA []byte, offset int, v uint64) int {
	offset -= sovRace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Race) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRace(uint64(m.Id))
	}
	l = len(m.PlayerAddress)
	if l > 0 {
		n += 1 + l + sovRace(uint64(l))
	}
	l = len(m.Bet)
	if l > 0 {
		n += 1 + l + sovRace(uint64(l))
	}
	if m.Multiplier != 0 {
		n += 5
	}
	if m.StartTime != 0 {
		n += 1 + sovRace(uint64(m.StartTime))
	}
	if m.EndTime != 0 {
		n += 1 + sovRace(uint64(m.EndTime))
	}
	l = len(m.CoinsEarned)
	if l > 0 {
		n += 1 + l + sovRace(uint64(l))
	}
	if m.Score != 0 {
		n += 1 + sovRace(uint64(m.Score))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovRace(uint64(l))
	}
	return n
}

func sovRace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRace(x uint64) (n int) {
	return sovRace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Race) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRace
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
			return fmt.Errorf("proto: Race: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Race: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRace
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
					return ErrIntOverflowRace
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
				return ErrInvalidLengthRace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRace
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
				return ErrInvalidLengthRace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Multiplier = float32(math.Float32frombits(v))
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinsEarned", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRace
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
				return ErrInvalidLengthRace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoinsEarned = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRace
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
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRace
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
				return ErrInvalidLengthRace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRace
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
func skipRace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRace
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
					return 0, ErrIntOverflowRace
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
					return 0, ErrIntOverflowRace
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
				return 0, ErrInvalidLengthRace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRace = fmt.Errorf("proto: unexpected end of group")
)