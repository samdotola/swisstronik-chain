// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: swisstronik/did/genesis.proto

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

// DIDDocumentVersionSet contains all versions of DID Documents and their
// metadata for a given DID. The latest version of the DID Document set is
// stored in the latest_version field.
type DIDDocumentVersionSet struct {
	// Latest version of the DID Document set
	LatestVersion string `protobuf:"bytes,1,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
	// All versions of the DID Document set
	DidDocs []*DIDDocumentWithMetadata `protobuf:"bytes,2,rep,name=did_docs,json=didDocs,proto3" json:"did_docs,omitempty"`
}

func (m *DIDDocumentVersionSet) Reset()         { *m = DIDDocumentVersionSet{} }
func (m *DIDDocumentVersionSet) String() string { return proto.CompactTextString(m) }
func (*DIDDocumentVersionSet) ProtoMessage()    {}
func (*DIDDocumentVersionSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_34002a7e75866bc1, []int{0}
}
func (m *DIDDocumentVersionSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DIDDocumentVersionSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DIDDocumentVersionSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DIDDocumentVersionSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DIDDocumentVersionSet.Merge(m, src)
}
func (m *DIDDocumentVersionSet) XXX_Size() int {
	return m.Size()
}
func (m *DIDDocumentVersionSet) XXX_DiscardUnknown() {
	xxx_messageInfo_DIDDocumentVersionSet.DiscardUnknown(m)
}

var xxx_messageInfo_DIDDocumentVersionSet proto.InternalMessageInfo

func (m *DIDDocumentVersionSet) GetLatestVersion() string {
	if m != nil {
		return m.LatestVersion
	}
	return ""
}

func (m *DIDDocumentVersionSet) GetDidDocs() []*DIDDocumentWithMetadata {
	if m != nil {
		return m.DidDocs
	}
	return nil
}

// GenesisState defines the did module's genesis state.
type GenesisState struct {
	// All DID Document version sets (contains all versions of all DID Documents)
	VersionSets []*DIDDocumentVersionSet `protobuf:"bytes,1,rep,name=version_sets,json=versionSets,proto3" json:"version_sets,omitempty"`
	Params      Params                   `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_34002a7e75866bc1, []int{1}
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

func (m *GenesisState) GetVersionSets() []*DIDDocumentVersionSet {
	if m != nil {
		return m.VersionSets
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*DIDDocumentVersionSet)(nil), "swisstronik.did.DIDDocumentVersionSet")
	proto.RegisterType((*GenesisState)(nil), "swisstronik.did.GenesisState")
}

func init() { proto.RegisterFile("swisstronik/did/genesis.proto", fileDescriptor_34002a7e75866bc1) }

var fileDescriptor_34002a7e75866bc1 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x33, 0xfd, 0x7f, 0xaa, 0x4e, 0xab, 0x42, 0x50, 0x1a, 0x8a, 0x8e, 0xa5, 0xa0, 0x64,
	0x95, 0x62, 0xc5, 0x17, 0xa8, 0x01, 0xe9, 0x42, 0x90, 0x14, 0x14, 0xdc, 0x94, 0x31, 0x33, 0xc4,
	0x41, 0x9b, 0x09, 0xb9, 0xd7, 0xaa, 0x6b, 0x5f, 0xa0, 0x8f, 0xd5, 0x65, 0x97, 0xae, 0x44, 0x92,
	0x17, 0x11, 0x33, 0x41, 0x43, 0x04, 0x77, 0x97, 0x73, 0xce, 0x7c, 0xf7, 0xce, 0xa1, 0xfb, 0xf0,
	0xa4, 0x00, 0x30, 0xd5, 0xb1, 0xba, 0x1f, 0x08, 0x25, 0x06, 0x91, 0x8c, 0x25, 0x28, 0xf0, 0x92,
	0x54, 0xa3, 0xb6, 0xb7, 0x2b, 0xb6, 0x27, 0x94, 0xe8, 0xb2, 0x7a, 0x5e, 0xe8, 0xf0, 0x71, 0x26,
	0x63, 0x34, 0x0f, 0xba, 0x7b, 0x75, 0x3f, 0xe1, 0x29, 0x9f, 0x95, 0xb8, 0xee, 0x4e, 0xa4, 0x23,
	0x5d, 0x8c, 0x83, 0xaf, 0xc9, 0xa8, 0xfd, 0x57, 0x42, 0x77, 0xfd, 0xb1, 0xef, 0x97, 0xa4, 0x2b,
	0x99, 0x82, 0xd2, 0xf1, 0x44, 0xa2, 0x7d, 0x48, 0xb7, 0x1e, 0x38, 0x4a, 0xc0, 0xe9, 0xdc, 0x88,
	0x0e, 0xe9, 0x11, 0x77, 0x23, 0xd8, 0x34, 0x6a, 0x99, 0xb4, 0xcf, 0xe8, 0xba, 0x50, 0x62, 0x2a,
	0x74, 0x08, 0x4e, 0xa3, 0xf7, 0xcf, 0x6d, 0x0d, 0x5d, 0xaf, 0x76, 0xb8, 0x57, 0x59, 0x70, 0xad,
	0xf0, 0xee, 0x42, 0x22, 0x17, 0x1c, 0x79, 0xb0, 0x26, 0x94, 0xf0, 0x75, 0x08, 0xfd, 0x05, 0xa1,
	0xed, 0x73, 0xf3, 0xf9, 0x09, 0x72, 0x94, 0xf6, 0x98, 0xb6, 0xcb, 0xad, 0x53, 0x90, 0x08, 0x0e,
	0x29, 0xc8, 0x47, 0x7f, 0x91, 0x7f, 0x4e, 0x0f, 0x5a, 0xf3, 0xef, 0x19, 0xec, 0x53, 0xda, 0x34,
	0x3d, 0x38, 0x8d, 0x1e, 0x71, 0x5b, 0xc3, 0xce, 0x2f, 0xc8, 0x65, 0x61, 0x8f, 0xfe, 0x2f, 0xdf,
	0x0f, 0xac, 0xa0, 0x0c, 0x8f, 0x8e, 0x97, 0x19, 0x23, 0xab, 0x8c, 0x91, 0x8f, 0x8c, 0x91, 0x45,
	0xce, 0xac, 0x55, 0xce, 0xac, 0xb7, 0x9c, 0x59, 0x37, 0x9d, 0x6a, 0xcd, 0xcf, 0x45, 0xd1, 0xf8,
	0x92, 0x48, 0xb8, 0x6d, 0x16, 0x95, 0x9e, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x01, 0xeb,
	0x71, 0xd8, 0x01, 0x00, 0x00,
}

func (m *DIDDocumentVersionSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DIDDocumentVersionSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DIDDocumentVersionSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DidDocs) > 0 {
		for iNdEx := len(m.DidDocs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DidDocs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.LatestVersion) > 0 {
		i -= len(m.LatestVersion)
		copy(dAtA[i:], m.LatestVersion)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.LatestVersion)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.VersionSets) > 0 {
		for iNdEx := len(m.VersionSets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VersionSets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
func (m *DIDDocumentVersionSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LatestVersion)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.DidDocs) > 0 {
		for _, e := range m.DidDocs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VersionSets) > 0 {
		for _, e := range m.VersionSets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DIDDocumentVersionSet) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DIDDocumentVersionSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DIDDocumentVersionSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocs", wireType)
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
			m.DidDocs = append(m.DidDocs, &DIDDocumentWithMetadata{})
			if err := m.DidDocs[len(m.DidDocs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
				return fmt.Errorf("proto: wrong wireType = %d for field VersionSets", wireType)
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
			m.VersionSets = append(m.VersionSets, &DIDDocumentVersionSet{})
			if err := m.VersionSets[len(m.VersionSets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
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
