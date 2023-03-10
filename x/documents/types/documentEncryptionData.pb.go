// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commercionetwork/documents/documentEncryptionData.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type DocumentEncryptionData struct {
	Keys          []*DocumentEncryptionKey `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	EncryptedData []string                 `protobuf:"bytes,2,rep,name=encryptedData,proto3" json:"encryptedData,omitempty"`
}

func (m *DocumentEncryptionData) Reset()         { *m = DocumentEncryptionData{} }
func (m *DocumentEncryptionData) String() string { return proto.CompactTextString(m) }
func (*DocumentEncryptionData) ProtoMessage()    {}
func (*DocumentEncryptionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d2b18a1bcac6a80, []int{0}
}
func (m *DocumentEncryptionData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DocumentEncryptionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DocumentEncryptionData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DocumentEncryptionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentEncryptionData.Merge(m, src)
}
func (m *DocumentEncryptionData) XXX_Size() int {
	return m.Size()
}
func (m *DocumentEncryptionData) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentEncryptionData.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentEncryptionData proto.InternalMessageInfo

func (m *DocumentEncryptionData) GetKeys() []*DocumentEncryptionKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *DocumentEncryptionData) GetEncryptedData() []string {
	if m != nil {
		return m.EncryptedData
	}
	return nil
}

func init() {
	proto.RegisterType((*DocumentEncryptionData)(nil), "commercionetwork.commercionetwork.documents.DocumentEncryptionData")
}

func init() {
	proto.RegisterFile("commercionetwork/documents/documentEncryptionData.proto", fileDescriptor_2d2b18a1bcac6a80)
}

var fileDescriptor_2d2b18a1bcac6a80 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4f, 0xce, 0xcf, 0xcd,
	0x4d, 0x2d, 0x4a, 0xce, 0xcc, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x4f, 0xc9, 0x4f,
	0x2e, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0x86, 0xb3, 0x5c, 0xf3, 0x92, 0x8b, 0x2a, 0x0b, 0x4a, 0x32,
	0xf3, 0xf3, 0x5c, 0x12, 0x4b, 0x12, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb4, 0xd1, 0x35,
	0xea, 0x61, 0x08, 0xc0, 0x4d, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xeb, 0xd3, 0x07, 0xb1,
	0x20, 0x46, 0x48, 0x99, 0x91, 0x64, 0xb7, 0x77, 0x6a, 0x25, 0x44, 0x9f, 0xd2, 0x34, 0x46, 0x2e,
	0x31, 0x17, 0xac, 0x6e, 0x13, 0x0a, 0xe3, 0x62, 0xc9, 0x4e, 0xad, 0x2c, 0x96, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x36, 0x72, 0xd2, 0x23, 0xc1, 0x91, 0x7a, 0x2e, 0xd8, 0xac, 0x0c, 0x02, 0x9b, 0x27,
	0xa4, 0xc2, 0xc5, 0x9b, 0x0a, 0x11, 0x4e, 0x4d, 0x01, 0x59, 0x24, 0xc1, 0xa4, 0xc0, 0xac, 0xc1,
	0x19, 0x84, 0x2a, 0xe8, 0x14, 0x71, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0x76, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x20, 0xeb, 0xf5, 0x31, 0x7c, 0x8d, 0x21, 0x50, 0x81,
	0x14, 0x10, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x9f, 0x1b, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf9, 0xe7, 0xeb, 0xc8, 0xaf, 0x01, 0x00, 0x00,
}

func (m *DocumentEncryptionData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DocumentEncryptionData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DocumentEncryptionData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedData) > 0 {
		for iNdEx := len(m.EncryptedData) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EncryptedData[iNdEx])
			copy(dAtA[i:], m.EncryptedData[iNdEx])
			i = encodeVarintDocumentEncryptionData(dAtA, i, uint64(len(m.EncryptedData[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Keys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDocumentEncryptionData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDocumentEncryptionData(dAtA []byte, offset int, v uint64) int {
	offset -= sovDocumentEncryptionData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DocumentEncryptionData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Keys) > 0 {
		for _, e := range m.Keys {
			l = e.Size()
			n += 1 + l + sovDocumentEncryptionData(uint64(l))
		}
	}
	if len(m.EncryptedData) > 0 {
		for _, s := range m.EncryptedData {
			l = len(s)
			n += 1 + l + sovDocumentEncryptionData(uint64(l))
		}
	}
	return n
}

func sovDocumentEncryptionData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDocumentEncryptionData(x uint64) (n int) {
	return sovDocumentEncryptionData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DocumentEncryptionData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDocumentEncryptionData
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
			return fmt.Errorf("proto: DocumentEncryptionData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DocumentEncryptionData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentEncryptionData
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
				return ErrInvalidLengthDocumentEncryptionData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentEncryptionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, &DocumentEncryptionKey{})
			if err := m.Keys[len(m.Keys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentEncryptionData
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
				return ErrInvalidLengthDocumentEncryptionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentEncryptionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedData = append(m.EncryptedData, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDocumentEncryptionData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDocumentEncryptionData
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
func skipDocumentEncryptionData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDocumentEncryptionData
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
					return 0, ErrIntOverflowDocumentEncryptionData
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
					return 0, ErrIntOverflowDocumentEncryptionData
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
				return 0, ErrInvalidLengthDocumentEncryptionData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDocumentEncryptionData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDocumentEncryptionData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDocumentEncryptionData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDocumentEncryptionData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDocumentEncryptionData = fmt.Errorf("proto: unexpected end of group")
)
