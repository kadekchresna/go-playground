// Code generated by protoc-gen-go. DO NOT EDIT.
// source: address.proto

package tutorial

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0, 0}
}

// [START messages]
type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Username             string                `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Addess               string                `protobuf:"bytes,6,opt,name=addess,proto3" json:"addess,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func (m *Person) GetAddess() string {
	if m != nil {
		return m.Addess
	}
	return ""
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
}

func init() {
	proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e)
}

var fileDescriptor_982c640dad8fe78e = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xd1, 0xaa, 0xd3, 0x40,
	0x10, 0x86, 0x4d, 0x9a, 0xc6, 0x76, 0xa2, 0x25, 0x2e, 0x45, 0x42, 0x10, 0x0c, 0xc5, 0x8b, 0x80,
	0xb0, 0x85, 0x2a, 0x78, 0xe5, 0x85, 0x81, 0xa2, 0xa2, 0xb5, 0x21, 0xb4, 0x78, 0x29, 0x5b, 0x77,
	0xac, 0xc1, 0x24, 0xbb, 0x64, 0x37, 0x60, 0x5f, 0xc9, 0x97, 0xf3, 0x15, 0x0e, 0xd9, 0x4d, 0x4a,
	0x39, 0x9c, 0xbb, 0xf9, 0x67, 0xbe, 0x64, 0xe6, 0xff, 0x17, 0x9e, 0x32, 0xce, 0x5b, 0x54, 0x8a,
	0xca, 0x56, 0x68, 0x41, 0x66, 0xba, 0xd3, 0xa2, 0x2d, 0x59, 0x15, 0xbf, 0x3c, 0x0b, 0x71, 0xae,
	0x70, 0x6d, 0xfa, 0xa7, 0xee, 0xd7, 0x5a, 0x97, 0x35, 0x2a, 0xcd, 0x6a, 0x69, 0xd1, 0xd5, 0x7f,
	0x17, 0xfc, 0x1c, 0x5b, 0x25, 0x1a, 0x42, 0xc0, 0x6b, 0x58, 0x8d, 0x91, 0x93, 0x38, 0xe9, 0xbc,
	0x30, 0x35, 0x59, 0x80, 0x5b, 0xf2, 0xc8, 0x4d, 0x9c, 0x74, 0x5a, 0xb8, 0x25, 0x27, 0x31, 0xcc,
	0x3a, 0x85, 0xad, 0xe1, 0x1e, 0x1b, 0xee, 0xaa, 0xc9, 0x12, 0xa6, 0x58, 0xb3, 0xb2, 0x8a, 0x26,
	0x66, 0x60, 0x05, 0x79, 0x0b, 0xbe, 0xfc, 0x2d, 0x1a, 0x54, 0x91, 0x97, 0x4c, 0xd2, 0x60, 0xf3,
	0x82, 0x8e, 0xc7, 0x51, 0xbb, 0x97, 0xe6, 0xfd, 0xf8, 0x5b, 0x57, 0x9f, 0xb0, 0x2d, 0x06, 0x96,
	0xbc, 0x87, 0x27, 0x15, 0x53, 0xfa, 0x47, 0x27, 0x39, 0xd3, 0xc8, 0xa3, 0x69, 0xe2, 0xa4, 0xc1,
	0x26, 0xa6, 0xd6, 0x0e, 0x1d, 0xed, 0xd0, 0xc3, 0x68, 0xa7, 0x08, 0x7a, 0xfe, 0x68, 0x71, 0xf2,
	0x1c, 0x7c, 0xc6, 0x39, 0x2a, 0x15, 0xf9, 0xe6, 0x96, 0x41, 0xc5, 0x47, 0x08, 0x6e, 0xb6, 0xf5,
	0x58, 0x63, 0xaa, 0xc1, 0xf3, 0xa0, 0x08, 0x05, 0x4f, 0x5f, 0x24, 0x1a, 0xdf, 0x8b, 0x4d, 0xfc,
	0xf0, 0xc5, 0x87, 0x8b, 0xc4, 0xc2, 0x70, 0xab, 0xd7, 0x30, 0xbf, 0xb6, 0x08, 0x80, 0xbf, 0xdb,
	0x67, 0x9f, 0xbf, 0x6e, 0xc3, 0x47, 0x64, 0x06, 0xde, 0xa7, 0xfd, 0x6e, 0x1b, 0x3a, 0x7d, 0xf5,
	0x7d, 0x5f, 0x7c, 0x09, 0xdd, 0xd5, 0x3b, 0x08, 0x3e, 0xd8, 0xd7, 0xca, 0x84, 0xf8, 0x43, 0x52,
	0xf0, 0x25, 0x0a, 0x59, 0xf5, 0xb9, 0xf7, 0xf9, 0x84, 0xf7, 0xb7, 0x15, 0xc3, 0x3c, 0xcb, 0x61,
	0xf9, 0x53, 0xd4, 0x14, 0xff, 0xb2, 0x5a, 0x56, 0x78, 0xc5, 0xb2, 0x67, 0x37, 0xbf, 0xcb, 0xfb,
	0x60, 0xd4, 0x3f, 0xf7, 0xd5, 0x47, 0x1b, 0x54, 0x3e, 0x06, 0xb5, 0xb5, 0x5f, 0x29, 0x7a, 0x03,
	0x9f, 0x7c, 0x93, 0xe3, 0x9b, 0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x42, 0x11, 0x71, 0x3f,
	0x02, 0x00, 0x00,
}
