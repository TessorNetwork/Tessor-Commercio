// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commercionetwork/documents/documentDoSign.proto

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

type DocumentDoSign struct {
	StorageURI         string   `protobuf:"bytes,1,opt,name=storageURI,proto3" json:"storageURI,omitempty"`
	SignerInstance     string   `protobuf:"bytes,2,opt,name=signerInstance,proto3" json:"signerInstance,omitempty"`
	SdnData            []string `protobuf:"bytes,3,rep,name=sdnData,proto3" json:"sdnData,omitempty"`
	VcrID              string   `protobuf:"bytes,4,opt,name=vcrID,proto3" json:"vcrID,omitempty"`
	CertificateProfile string   `protobuf:"bytes,5,opt,name=certificateProfile,proto3" json:"certificateProfile,omitempty"`
}

func (m *DocumentDoSign) Reset()         { *m = DocumentDoSign{} }
func (m *DocumentDoSign) String() string { return proto.CompactTextString(m) }
func (*DocumentDoSign) ProtoMessage()    {}
func (*DocumentDoSign) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4f5e5c3f05aa265, []int{0}
}
func (m *DocumentDoSign) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DocumentDoSign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DocumentDoSign.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DocumentDoSign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentDoSign.Merge(m, src)
}
func (m *DocumentDoSign) XXX_Size() int {
	return m.Size()
}
func (m *DocumentDoSign) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentDoSign.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentDoSign proto.InternalMessageInfo

func (m *DocumentDoSign) GetStorageURI() string {
	if m != nil {
		return m.StorageURI
	}
	return ""
}

func (m *DocumentDoSign) GetSignerInstance() string {
	if m != nil {
		return m.SignerInstance
	}
	return ""
}

func (m *DocumentDoSign) GetSdnData() []string {
	if m != nil {
		return m.SdnData
	}
	return nil
}

func (m *DocumentDoSign) GetVcrID() string {
	if m != nil {
		return m.VcrID
	}
	return ""
}

func (m *DocumentDoSign) GetCertificateProfile() string {
	if m != nil {
		return m.CertificateProfile
	}
	return ""
}

func init() {
	proto.RegisterType((*DocumentDoSign)(nil), "commercionetwork.commercionetwork.documents.DocumentDoSign")
}

func init() {
	proto.RegisterFile("commercionetwork/documents/documentDoSign.proto", fileDescriptor_f4f5e5c3f05aa265)
}

var fileDescriptor_f4f5e5c3f05aa265 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xce, 0xcf, 0xcd,
	0x4d, 0x2d, 0x4a, 0xce, 0xcc, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x4f, 0xc9, 0x4f,
	0x2e, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0x86, 0xb3, 0x5c, 0xf2, 0x83, 0x33, 0xd3, 0xf3, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb4, 0xd1, 0x35, 0xe8, 0x61, 0x08, 0xc0, 0x4d, 0x90, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0xeb, 0xd3, 0x07, 0xb1, 0x20, 0x46, 0x28, 0xed, 0x60, 0xe4, 0xe2, 0x73,
	0x41, 0x31, 0x5b, 0x48, 0x8e, 0x8b, 0xab, 0xb8, 0x24, 0xbf, 0x28, 0x31, 0x3d, 0x35, 0x34, 0xc8,
	0x53, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x49, 0x44, 0x48, 0x8d, 0x8b, 0xaf, 0x38, 0x33,
	0x3d, 0x2f, 0xb5, 0xc8, 0x33, 0xaf, 0xb8, 0x24, 0x31, 0x2f, 0x39, 0x55, 0x82, 0x09, 0xac, 0x06,
	0x4d, 0x54, 0x48, 0x82, 0x8b, 0xbd, 0x38, 0x25, 0xcf, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x59, 0x81,
	0x59, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0x2e, 0xf2, 0x74, 0x91, 0x60,
	0x01, 0x6b, 0x84, 0x70, 0x84, 0xf4, 0xb8, 0x84, 0x92, 0x53, 0x8b, 0x4a, 0x32, 0xd3, 0x32, 0x93,
	0x13, 0x4b, 0x52, 0x03, 0x8a, 0xf2, 0xd3, 0x32, 0x73, 0x52, 0x25, 0x58, 0xc1, 0x4a, 0xb0, 0xc8,
	0x38, 0x45, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x5d, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x12, 0x28, 0x34, 0x30, 0xc3, 0x14, 0x43, 0xa0, 0x02, 0x29, 0x98, 0x4b, 0x2a,
	0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x61, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x59, 0x91,
	0xc4, 0x55, 0x91, 0x01, 0x00, 0x00,
}

func (m *DocumentDoSign) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DocumentDoSign) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DocumentDoSign) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CertificateProfile) > 0 {
		i -= len(m.CertificateProfile)
		copy(dAtA[i:], m.CertificateProfile)
		i = encodeVarintDocumentDoSign(dAtA, i, uint64(len(m.CertificateProfile)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.VcrID) > 0 {
		i -= len(m.VcrID)
		copy(dAtA[i:], m.VcrID)
		i = encodeVarintDocumentDoSign(dAtA, i, uint64(len(m.VcrID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SdnData) > 0 {
		for iNdEx := len(m.SdnData) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SdnData[iNdEx])
			copy(dAtA[i:], m.SdnData[iNdEx])
			i = encodeVarintDocumentDoSign(dAtA, i, uint64(len(m.SdnData[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SignerInstance) > 0 {
		i -= len(m.SignerInstance)
		copy(dAtA[i:], m.SignerInstance)
		i = encodeVarintDocumentDoSign(dAtA, i, uint64(len(m.SignerInstance)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StorageURI) > 0 {
		i -= len(m.StorageURI)
		copy(dAtA[i:], m.StorageURI)
		i = encodeVarintDocumentDoSign(dAtA, i, uint64(len(m.StorageURI)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDocumentDoSign(dAtA []byte, offset int, v uint64) int {
	offset -= sovDocumentDoSign(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DocumentDoSign) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StorageURI)
	if l > 0 {
		n += 1 + l + sovDocumentDoSign(uint64(l))
	}
	l = len(m.SignerInstance)
	if l > 0 {
		n += 1 + l + sovDocumentDoSign(uint64(l))
	}
	if len(m.SdnData) > 0 {
		for _, s := range m.SdnData {
			l = len(s)
			n += 1 + l + sovDocumentDoSign(uint64(l))
		}
	}
	l = len(m.VcrID)
	if l > 0 {
		n += 1 + l + sovDocumentDoSign(uint64(l))
	}
	l = len(m.CertificateProfile)
	if l > 0 {
		n += 1 + l + sovDocumentDoSign(uint64(l))
	}
	return n
}

func sovDocumentDoSign(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDocumentDoSign(x uint64) (n int) {
	return sovDocumentDoSign(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DocumentDoSign) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDocumentDoSign
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
			return fmt.Errorf("proto: DocumentDoSign: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DocumentDoSign: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentDoSign
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
				return ErrInvalidLengthDocumentDoSign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentDoSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerInstance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentDoSign
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
				return ErrInvalidLengthDocumentDoSign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentDoSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignerInstance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdnData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentDoSign
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
				return ErrInvalidLengthDocumentDoSign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentDoSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SdnData = append(m.SdnData, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VcrID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentDoSign
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
				return ErrInvalidLengthDocumentDoSign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentDoSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VcrID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificateProfile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentDoSign
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
				return ErrInvalidLengthDocumentDoSign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentDoSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificateProfile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDocumentDoSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDocumentDoSign
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
func skipDocumentDoSign(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDocumentDoSign
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
					return 0, ErrIntOverflowDocumentDoSign
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
					return 0, ErrIntOverflowDocumentDoSign
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
				return 0, ErrInvalidLengthDocumentDoSign
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDocumentDoSign
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDocumentDoSign
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDocumentDoSign        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDocumentDoSign          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDocumentDoSign = fmt.Errorf("proto: unexpected end of group")
)
