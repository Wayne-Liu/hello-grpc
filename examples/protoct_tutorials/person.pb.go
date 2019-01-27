// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

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
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
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
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 0}
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
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
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

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x80, 0xdf, 0xa4, 0x69, 0x68, 0x27, 0xa5, 0x84, 0xe1, 0x45, 0x42, 0x10, 0x0c, 0x3d, 0x05,
	0x84, 0x2d, 0x54, 0xc1, 0x93, 0x07, 0x0b, 0x05, 0x45, 0x6b, 0xcb, 0xd2, 0xe2, 0x51, 0x52, 0x32,
	0xd6, 0x60, 0x92, 0x5d, 0xb2, 0x9b, 0x43, 0x7f, 0x9b, 0x7f, 0x4e, 0xb2, 0x49, 0x44, 0xc4, 0xdb,
	0x7c, 0x3c, 0x3b, 0xfb, 0xcc, 0xc0, 0x44, 0x52, 0xa5, 0x44, 0xc9, 0x64, 0x25, 0xb4, 0xc0, 0x91,
	0xae, 0xb5, 0xa8, 0xb2, 0x24, 0x0f, 0x2f, 0x8e, 0x42, 0x1c, 0x73, 0x9a, 0x9b, 0xfa, 0xa1, 0x7e,
	0x9b, 0xeb, 0xac, 0x20, 0xa5, 0x93, 0x42, 0xb6, 0xe8, 0xec, 0xd3, 0x06, 0x77, 0x6b, 0xde, 0x22,
	0x82, 0x53, 0x26, 0x05, 0x05, 0x56, 0x64, 0xc5, 0x63, 0x6e, 0x62, 0x9c, 0x82, 0x9d, 0xa5, 0x81,
	0x1d, 0x59, 0xf1, 0x90, 0xdb, 0x59, 0x8a, 0xff, 0x61, 0x48, 0x45, 0x92, 0xe5, 0xc1, 0xc0, 0x40,
	0x6d, 0x82, 0xd7, 0xe0, 0xca, 0x77, 0x51, 0x92, 0x0a, 0x9c, 0x68, 0x10, 0x7b, 0x8b, 0x73, 0xd6,
	0x0b, 0xb0, 0x76, 0x36, 0xdb, 0x36, 0xed, 0xe7, 0xba, 0x38, 0x50, 0xc5, 0x3b, 0x16, 0x6f, 0x61,
	0x92, 0x27, 0x4a, 0xbf, 0xd6, 0x32, 0x4d, 0x34, 0xa5, 0xc1, 0x30, 0xb2, 0x62, 0x6f, 0x11, 0xb2,
	0x56, 0x99, 0xf5, 0xca, 0x6c, 0xd7, 0x2b, 0x73, 0xaf, 0xe1, 0xf7, 0x2d, 0x1e, 0xee, 0xc1, 0xfb,
	0x31, 0x15, 0xcf, 0xc0, 0x2d, 0x4d, 0xd4, 0xf9, 0x77, 0x19, 0x32, 0x70, 0xf4, 0x49, 0x92, 0xd9,
	0x61, 0xba, 0x08, 0xff, 0x36, 0xdb, 0x9d, 0x24, 0x71, 0xc3, 0xcd, 0x2e, 0x61, 0xfc, 0x5d, 0x42,
	0x00, 0x77, 0xbd, 0x59, 0x3e, 0x3c, 0xad, 0xfc, 0x7f, 0x38, 0x02, 0xe7, 0x7e, 0xb3, 0x5e, 0xf9,
	0x56, 0x13, 0xbd, 0x6c, 0xf8, 0xa3, 0x6f, 0xcf, 0x6e, 0xc0, 0xbb, 0x4b, 0xd3, 0x8a, 0x94, 0x5a,
	0x0a, 0xf1, 0x81, 0x31, 0xb8, 0x92, 0x84, 0xcc, 0x9b, 0x1b, 0x36, 0x77, 0xf0, 0x7f, 0xff, 0xc6,
	0xbb, 0xfe, 0xc1, 0x35, 0xdb, 0x5d, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x96, 0x3a, 0x18,
	0xb8, 0x01, 0x00, 0x00,
}