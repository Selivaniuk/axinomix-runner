// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axinomixrunner/axinomixrunner/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the axinomixrunner module's genesis state.
type GenesisState struct {
	Params           Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	RaceList         []Race        `protobuf:"bytes,2,rep,name=raceList,proto3" json:"raceList"`
	RaceCount        uint64        `protobuf:"varint,3,opt,name=raceCount,proto3" json:"raceCount,omitempty"`
	LeaderboardList  []Leaderboard `protobuf:"bytes,4,rep,name=leaderboardList,proto3" json:"leaderboardList"`
	LeaderboardCount uint64        `protobuf:"varint,5,opt,name=leaderboardCount,proto3" json:"leaderboardCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_88a648018c951631, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRaceList() []Race {
	if m != nil {
		return m.RaceList
	}
	return nil
}

func (m *GenesisState) GetRaceCount() uint64 {
	if m != nil {
		return m.RaceCount
	}
	return 0
}

func (m *GenesisState) GetLeaderboardList() []Leaderboard {
	if m != nil {
		return m.LeaderboardList
	}
	return nil
}

func (m *GenesisState) GetLeaderboardCount() uint64 {
	if m != nil {
		return m.LeaderboardCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "axinomixrunner.axinomixrunner.GenesisState")
}

func init() {
	proto.RegisterFile("axinomixrunner/axinomixrunner/genesis.proto", fileDescriptor_88a648018c951631)
}

var fileDescriptor_88a648018c951631 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xac, 0xc8, 0xcc,
	0xcb, 0xcf, 0xcd, 0xac, 0x28, 0x2a, 0xcd, 0xcb, 0x4b, 0x2d, 0xd2, 0x47, 0xe3, 0xa6, 0xa7, 0xe6,
	0xa5, 0x16, 0x67, 0x16, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xc9, 0xa2, 0xca, 0xea, 0xa1,
	0x72, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x2a, 0xf5, 0x41, 0x2c, 0x88, 0x26, 0x29, 0x2d,
	0xfc, 0x36, 0x14, 0x24, 0x16, 0x25, 0xe6, 0x42, 0x2d, 0x90, 0xd2, 0xc0, 0xaf, 0xb6, 0x28, 0x31,
	0x39, 0x15, 0xaa, 0x52, 0x1f, 0xbf, 0xca, 0x9c, 0xd4, 0xc4, 0x94, 0xd4, 0xa2, 0xa4, 0xfc, 0xc4,
	0xa2, 0x14, 0x88, 0x06, 0xa5, 0x03, 0x4c, 0x5c, 0x3c, 0xee, 0x10, 0xdf, 0x04, 0x97, 0x24, 0x96,
	0xa4, 0x0a, 0x39, 0x73, 0xb1, 0x41, 0xec, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd5,
	0xc3, 0xeb, 0x3b, 0xbd, 0x00, 0xb0, 0x62, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xa0, 0x5a,
	0x85, 0x5c, 0xb9, 0x38, 0x40, 0x8e, 0xf2, 0xc9, 0x2c, 0x2e, 0x91, 0x60, 0x52, 0x60, 0xd6, 0xe0,
	0x36, 0x52, 0x26, 0x60, 0x4c, 0x50, 0x62, 0x72, 0x2a, 0xd4, 0x10, 0xb8, 0x56, 0x21, 0x19, 0x2e,
	0x4e, 0x10, 0xdb, 0x39, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x59, 0x81, 0x51, 0x83, 0x25, 0x08, 0x21,
	0x20, 0x14, 0xc5, 0xc5, 0x8f, 0xe4, 0x1f, 0xb0, 0x5d, 0x2c, 0x60, 0xbb, 0xb4, 0x08, 0xd8, 0xe5,
	0x83, 0xd0, 0x05, 0xb5, 0x12, 0xdd, 0x20, 0x21, 0x2d, 0x2e, 0x01, 0x24, 0x21, 0x88, 0x03, 0x58,
	0xc1, 0x0e, 0xc0, 0x10, 0x77, 0x72, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0x35, 0x98, 0xc5, 0xba, 0xd0, 0xf0, 0xaf, 0x40, 0x8f, 0x90, 0x92, 0xca, 0x82, 0xd4, 0xe2,
	0x24, 0x36, 0x70, 0x5c, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xac, 0x13, 0x27, 0x0c, 0x76,
	0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LeaderboardCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LeaderboardCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.LeaderboardList) > 0 {
		for iNdEx := len(m.LeaderboardList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LeaderboardList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.RaceCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RaceCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RaceList) > 0 {
		for iNdEx := len(m.RaceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RaceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.RaceList) > 0 {
		for _, e := range m.RaceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.RaceCount != 0 {
		n += 1 + sovGenesis(uint64(m.RaceCount))
	}
	if len(m.LeaderboardList) > 0 {
		for _, e := range m.LeaderboardList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LeaderboardCount != 0 {
		n += 1 + sovGenesis(uint64(m.LeaderboardCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RaceList = append(m.RaceList, Race{})
			if err := m.RaceList[len(m.RaceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaceCount", wireType)
			}
			m.RaceCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaceCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderboardList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LeaderboardList = append(m.LeaderboardList, Leaderboard{})
			if err := m.LeaderboardList[len(m.LeaderboardList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderboardCount", wireType)
			}
			m.LeaderboardCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LeaderboardCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
