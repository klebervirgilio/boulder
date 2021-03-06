// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/proto/core.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Challenge struct {
	Id                   *int64              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type                 *string             `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Status               *string             `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	Uri                  *string             `protobuf:"bytes,9,opt,name=uri" json:"uri,omitempty"`
	Token                *string             `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	KeyAuthorization     *string             `protobuf:"bytes,5,opt,name=keyAuthorization" json:"keyAuthorization,omitempty"`
	Validationrecords    []*ValidationRecord `protobuf:"bytes,10,rep,name=validationrecords" json:"validationrecords,omitempty"`
	Error                *ProblemDetails     `protobuf:"bytes,7,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Challenge) Reset()         { *m = Challenge{} }
func (m *Challenge) String() string { return proto.CompactTextString(m) }
func (*Challenge) ProtoMessage()    {}
func (*Challenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{0}
}

func (m *Challenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Challenge.Unmarshal(m, b)
}
func (m *Challenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Challenge.Marshal(b, m, deterministic)
}
func (m *Challenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Challenge.Merge(m, src)
}
func (m *Challenge) XXX_Size() int {
	return xxx_messageInfo_Challenge.Size(m)
}
func (m *Challenge) XXX_DiscardUnknown() {
	xxx_messageInfo_Challenge.DiscardUnknown(m)
}

var xxx_messageInfo_Challenge proto.InternalMessageInfo

func (m *Challenge) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Challenge) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Challenge) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *Challenge) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Challenge) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Challenge) GetKeyAuthorization() string {
	if m != nil && m.KeyAuthorization != nil {
		return *m.KeyAuthorization
	}
	return ""
}

func (m *Challenge) GetValidationrecords() []*ValidationRecord {
	if m != nil {
		return m.Validationrecords
	}
	return nil
}

func (m *Challenge) GetError() *ProblemDetails {
	if m != nil {
		return m.Error
	}
	return nil
}

type ValidationRecord struct {
	Hostname          *string  `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Port              *string  `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	AddressesResolved [][]byte `protobuf:"bytes,3,rep,name=addressesResolved" json:"addressesResolved,omitempty"`
	AddressUsed       []byte   `protobuf:"bytes,4,opt,name=addressUsed" json:"addressUsed,omitempty"`
	Authorities       []string `protobuf:"bytes,5,rep,name=authorities" json:"authorities,omitempty"`
	Url               *string  `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	// A list of addresses tried before the address used (see
	// core/objects.go and the comment on the ValidationRecord structure
	// definition for more information.
	AddressesTried       [][]byte `protobuf:"bytes,7,rep,name=addressesTried" json:"addressesTried,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationRecord) Reset()         { *m = ValidationRecord{} }
func (m *ValidationRecord) String() string { return proto.CompactTextString(m) }
func (*ValidationRecord) ProtoMessage()    {}
func (*ValidationRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{1}
}

func (m *ValidationRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationRecord.Unmarshal(m, b)
}
func (m *ValidationRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationRecord.Marshal(b, m, deterministic)
}
func (m *ValidationRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationRecord.Merge(m, src)
}
func (m *ValidationRecord) XXX_Size() int {
	return xxx_messageInfo_ValidationRecord.Size(m)
}
func (m *ValidationRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationRecord proto.InternalMessageInfo

func (m *ValidationRecord) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *ValidationRecord) GetPort() string {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return ""
}

func (m *ValidationRecord) GetAddressesResolved() [][]byte {
	if m != nil {
		return m.AddressesResolved
	}
	return nil
}

func (m *ValidationRecord) GetAddressUsed() []byte {
	if m != nil {
		return m.AddressUsed
	}
	return nil
}

func (m *ValidationRecord) GetAuthorities() []string {
	if m != nil {
		return m.Authorities
	}
	return nil
}

func (m *ValidationRecord) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *ValidationRecord) GetAddressesTried() [][]byte {
	if m != nil {
		return m.AddressesTried
	}
	return nil
}

type ProblemDetails struct {
	ProblemType          *string  `protobuf:"bytes,1,opt,name=problemType" json:"problemType,omitempty"`
	Detail               *string  `protobuf:"bytes,2,opt,name=detail" json:"detail,omitempty"`
	HttpStatus           *int32   `protobuf:"varint,3,opt,name=httpStatus" json:"httpStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProblemDetails) Reset()         { *m = ProblemDetails{} }
func (m *ProblemDetails) String() string { return proto.CompactTextString(m) }
func (*ProblemDetails) ProtoMessage()    {}
func (*ProblemDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{2}
}

func (m *ProblemDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProblemDetails.Unmarshal(m, b)
}
func (m *ProblemDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProblemDetails.Marshal(b, m, deterministic)
}
func (m *ProblemDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProblemDetails.Merge(m, src)
}
func (m *ProblemDetails) XXX_Size() int {
	return xxx_messageInfo_ProblemDetails.Size(m)
}
func (m *ProblemDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ProblemDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ProblemDetails proto.InternalMessageInfo

func (m *ProblemDetails) GetProblemType() string {
	if m != nil && m.ProblemType != nil {
		return *m.ProblemType
	}
	return ""
}

func (m *ProblemDetails) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

func (m *ProblemDetails) GetHttpStatus() int32 {
	if m != nil && m.HttpStatus != nil {
		return *m.HttpStatus
	}
	return 0
}

type Certificate struct {
	RegistrationID       *int64   `protobuf:"varint,1,opt,name=registrationID" json:"registrationID,omitempty"`
	Serial               *string  `protobuf:"bytes,2,opt,name=serial" json:"serial,omitempty"`
	Digest               *string  `protobuf:"bytes,3,opt,name=digest" json:"digest,omitempty"`
	Der                  []byte   `protobuf:"bytes,4,opt,name=der" json:"der,omitempty"`
	Issued               *int64   `protobuf:"varint,5,opt,name=issued" json:"issued,omitempty"`
	Expires              *int64   `protobuf:"varint,6,opt,name=expires" json:"expires,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{3}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

func (m *Certificate) GetSerial() string {
	if m != nil && m.Serial != nil {
		return *m.Serial
	}
	return ""
}

func (m *Certificate) GetDigest() string {
	if m != nil && m.Digest != nil {
		return *m.Digest
	}
	return ""
}

func (m *Certificate) GetDer() []byte {
	if m != nil {
		return m.Der
	}
	return nil
}

func (m *Certificate) GetIssued() int64 {
	if m != nil && m.Issued != nil {
		return *m.Issued
	}
	return 0
}

func (m *Certificate) GetExpires() int64 {
	if m != nil && m.Expires != nil {
		return *m.Expires
	}
	return 0
}

type Registration struct {
	Id                   *int64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Contact              []string `protobuf:"bytes,3,rep,name=contact" json:"contact,omitempty"`
	ContactsPresent      *bool    `protobuf:"varint,4,opt,name=contactsPresent" json:"contactsPresent,omitempty"`
	Agreement            *string  `protobuf:"bytes,5,opt,name=agreement" json:"agreement,omitempty"`
	InitialIP            []byte   `protobuf:"bytes,6,opt,name=initialIP" json:"initialIP,omitempty"`
	CreatedAt            *int64   `protobuf:"varint,7,opt,name=createdAt" json:"createdAt,omitempty"`
	Status               *string  `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Registration) Reset()         { *m = Registration{} }
func (m *Registration) String() string { return proto.CompactTextString(m) }
func (*Registration) ProtoMessage()    {}
func (*Registration) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{4}
}

func (m *Registration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registration.Unmarshal(m, b)
}
func (m *Registration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registration.Marshal(b, m, deterministic)
}
func (m *Registration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registration.Merge(m, src)
}
func (m *Registration) XXX_Size() int {
	return xxx_messageInfo_Registration.Size(m)
}
func (m *Registration) XXX_DiscardUnknown() {
	xxx_messageInfo_Registration.DiscardUnknown(m)
}

var xxx_messageInfo_Registration proto.InternalMessageInfo

func (m *Registration) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Registration) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Registration) GetContact() []string {
	if m != nil {
		return m.Contact
	}
	return nil
}

func (m *Registration) GetContactsPresent() bool {
	if m != nil && m.ContactsPresent != nil {
		return *m.ContactsPresent
	}
	return false
}

func (m *Registration) GetAgreement() string {
	if m != nil && m.Agreement != nil {
		return *m.Agreement
	}
	return ""
}

func (m *Registration) GetInitialIP() []byte {
	if m != nil {
		return m.InitialIP
	}
	return nil
}

func (m *Registration) GetCreatedAt() int64 {
	if m != nil && m.CreatedAt != nil {
		return *m.CreatedAt
	}
	return 0
}

func (m *Registration) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

type Authorization struct {
	Id                   *string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Identifier           *string      `protobuf:"bytes,2,opt,name=identifier" json:"identifier,omitempty"`
	RegistrationID       *int64       `protobuf:"varint,3,opt,name=registrationID" json:"registrationID,omitempty"`
	Status               *string      `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Expires              *int64       `protobuf:"varint,5,opt,name=expires" json:"expires,omitempty"`
	Challenges           []*Challenge `protobuf:"bytes,6,rep,name=challenges" json:"challenges,omitempty"`
	Combinations         []byte       `protobuf:"bytes,7,opt,name=combinations" json:"combinations,omitempty"`
	V2                   *bool        `protobuf:"varint,8,opt,name=v2" json:"v2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Authorization) Reset()         { *m = Authorization{} }
func (m *Authorization) String() string { return proto.CompactTextString(m) }
func (*Authorization) ProtoMessage()    {}
func (*Authorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{5}
}

func (m *Authorization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authorization.Unmarshal(m, b)
}
func (m *Authorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authorization.Marshal(b, m, deterministic)
}
func (m *Authorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authorization.Merge(m, src)
}
func (m *Authorization) XXX_Size() int {
	return xxx_messageInfo_Authorization.Size(m)
}
func (m *Authorization) XXX_DiscardUnknown() {
	xxx_messageInfo_Authorization.DiscardUnknown(m)
}

var xxx_messageInfo_Authorization proto.InternalMessageInfo

func (m *Authorization) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Authorization) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *Authorization) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

func (m *Authorization) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *Authorization) GetExpires() int64 {
	if m != nil && m.Expires != nil {
		return *m.Expires
	}
	return 0
}

func (m *Authorization) GetChallenges() []*Challenge {
	if m != nil {
		return m.Challenges
	}
	return nil
}

func (m *Authorization) GetCombinations() []byte {
	if m != nil {
		return m.Combinations
	}
	return nil
}

func (m *Authorization) GetV2() bool {
	if m != nil && m.V2 != nil {
		return *m.V2
	}
	return false
}

type Order struct {
	Id                *int64          `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	RegistrationID    *int64          `protobuf:"varint,2,opt,name=registrationID" json:"registrationID,omitempty"`
	Expires           *int64          `protobuf:"varint,3,opt,name=expires" json:"expires,omitempty"`
	Error             *ProblemDetails `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	CertificateSerial *string         `protobuf:"bytes,5,opt,name=certificateSerial" json:"certificateSerial,omitempty"`
	// contains only v1 authorization IDs, should be
	// deprecated once fully switched to v2 authorizations
	// in favor of v2Authorizations.
	Authorizations  []string `protobuf:"bytes,6,rep,name=authorizations" json:"authorizations,omitempty"`
	Status          *string  `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	Names           []string `protobuf:"bytes,8,rep,name=names" json:"names,omitempty"`
	BeganProcessing *bool    `protobuf:"varint,9,opt,name=beganProcessing" json:"beganProcessing,omitempty"`
	Created         *int64   `protobuf:"varint,10,opt,name=created" json:"created,omitempty"`
	// contains only v2 authorization IDs.
	V2Authorizations     []int64  `protobuf:"varint,11,rep,name=v2Authorizations" json:"v2Authorizations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{6}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Order) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

func (m *Order) GetExpires() int64 {
	if m != nil && m.Expires != nil {
		return *m.Expires
	}
	return 0
}

func (m *Order) GetError() *ProblemDetails {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Order) GetCertificateSerial() string {
	if m != nil && m.CertificateSerial != nil {
		return *m.CertificateSerial
	}
	return ""
}

func (m *Order) GetAuthorizations() []string {
	if m != nil {
		return m.Authorizations
	}
	return nil
}

func (m *Order) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *Order) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *Order) GetBeganProcessing() bool {
	if m != nil && m.BeganProcessing != nil {
		return *m.BeganProcessing
	}
	return false
}

func (m *Order) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Order) GetV2Authorizations() []int64 {
	if m != nil {
		return m.V2Authorizations
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{7}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Challenge)(nil), "core.Challenge")
	proto.RegisterType((*ValidationRecord)(nil), "core.ValidationRecord")
	proto.RegisterType((*ProblemDetails)(nil), "core.ProblemDetails")
	proto.RegisterType((*Certificate)(nil), "core.Certificate")
	proto.RegisterType((*Registration)(nil), "core.Registration")
	proto.RegisterType((*Authorization)(nil), "core.Authorization")
	proto.RegisterType((*Order)(nil), "core.Order")
	proto.RegisterType((*Empty)(nil), "core.Empty")
}

func init() { proto.RegisterFile("core/proto/core.proto", fileDescriptor_80ea9561f1d738ba) }

var fileDescriptor_80ea9561f1d738ba = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xc1, 0x6e, 0xdb, 0x38,
	0x10, 0x85, 0x2c, 0x2b, 0xb6, 0xc6, 0xde, 0xc4, 0x21, 0xb2, 0x81, 0xb0, 0x58, 0x04, 0x82, 0x0e,
	0x0b, 0x21, 0x58, 0x24, 0x80, 0xff, 0x20, 0x4d, 0x7a, 0xc8, 0xa9, 0x06, 0x93, 0xf6, 0xd0, 0x9b,
	0x22, 0x4d, 0x6d, 0x36, 0xb2, 0x68, 0x90, 0xb4, 0x51, 0xf7, 0x1f, 0xfa, 0x21, 0xfd, 0xb1, 0x1e,
	0x7b, 0xef, 0xa5, 0x28, 0x38, 0x94, 0x6d, 0xc9, 0x4e, 0xd1, 0xdb, 0xcc, 0x23, 0x65, 0xce, 0xbc,
	0xf7, 0x66, 0x0c, 0x7f, 0xe7, 0x52, 0xe1, 0xf5, 0x42, 0x49, 0x23, 0xaf, 0x6d, 0x78, 0x45, 0x21,
	0xeb, 0xda, 0x38, 0xf9, 0xd2, 0x81, 0xf0, 0x76, 0x96, 0x95, 0x25, 0x56, 0x53, 0x64, 0xc7, 0xd0,
	0x11, 0x45, 0xe4, 0xc5, 0x5e, 0xea, 0xf3, 0x8e, 0x28, 0x18, 0x83, 0xae, 0x59, 0x2f, 0x30, 0xea,
	0xc4, 0x5e, 0x1a, 0x72, 0x8a, 0xd9, 0x39, 0x1c, 0x69, 0x93, 0x99, 0xa5, 0x8e, 0x8e, 0x08, 0xad,
	0x33, 0x36, 0x02, 0x7f, 0xa9, 0x44, 0x14, 0x12, 0x68, 0x43, 0x76, 0x06, 0x81, 0x91, 0xcf, 0x58,
	0x45, 0x3e, 0x61, 0x2e, 0x61, 0x97, 0x30, 0x7a, 0xc6, 0xf5, 0xcd, 0xd2, 0xcc, 0xa4, 0x12, 0x9f,
	0x33, 0x23, 0x64, 0x15, 0x05, 0x74, 0xe1, 0x00, 0x67, 0x77, 0x70, 0xba, 0xca, 0x4a, 0x51, 0x50,
	0xa6, 0x30, 0x97, 0xaa, 0xd0, 0x11, 0xc4, 0x7e, 0x3a, 0x18, 0x9f, 0x5f, 0x51, 0x2f, 0xef, 0xb6,
	0xc7, 0x9c, 0x8e, 0xf9, 0xe1, 0x07, 0xec, 0x12, 0x02, 0x54, 0x4a, 0xaa, 0xa8, 0x17, 0x7b, 0xe9,
	0x60, 0x7c, 0xe6, 0xbe, 0x9c, 0x28, 0xf9, 0x54, 0xe2, 0xfc, 0x0e, 0x4d, 0x26, 0x4a, 0xcd, 0xdd,
	0x95, 0xe4, 0xbb, 0x07, 0xa3, 0xfd, 0xdf, 0x64, 0xff, 0x40, 0x7f, 0x26, 0xb5, 0xa9, 0xb2, 0x39,
	0x12, 0x39, 0x21, 0xdf, 0xe6, 0x96, 0xa2, 0x85, 0x54, 0x66, 0x43, 0x91, 0x8d, 0xd9, 0xff, 0x70,
	0x9a, 0x15, 0x85, 0x42, 0xad, 0x51, 0x73, 0xd4, 0xb2, 0x5c, 0x61, 0x11, 0xf9, 0xb1, 0x9f, 0x0e,
	0xf9, 0xe1, 0x01, 0x8b, 0x61, 0x50, 0x83, 0x6f, 0x35, 0x16, 0x51, 0x37, 0xf6, 0xd2, 0x21, 0x6f,
	0x42, 0x74, 0xc3, 0xf1, 0x62, 0x04, 0xea, 0x28, 0x88, 0xfd, 0x34, 0xe4, 0x4d, 0xc8, 0x91, 0x5f,
	0xd6, 0x8a, 0xd8, 0x90, 0xfd, 0x07, 0xc7, 0xdb, 0xa7, 0x1e, 0x95, 0xc0, 0x22, 0xea, 0x51, 0x01,
	0x7b, 0x68, 0xf2, 0x11, 0x8e, 0xdb, 0x4c, 0xd8, 0xd7, 0x16, 0x0e, 0x79, 0xb4, 0xda, 0xbb, 0x86,
	0x9b, 0x90, 0xb5, 0x40, 0x41, 0x97, 0xeb, 0xae, 0xeb, 0x8c, 0x5d, 0x00, 0xcc, 0x8c, 0x59, 0x3c,
	0x38, 0x7b, 0x58, 0xd5, 0x03, 0xde, 0x40, 0x92, 0xaf, 0x1e, 0x0c, 0x6e, 0x51, 0x19, 0xf1, 0x41,
	0xe4, 0x99, 0x41, 0x5b, 0xa3, 0xc2, 0xa9, 0xd0, 0x46, 0x11, 0xdb, 0xf7, 0x77, 0xb5, 0xf5, 0xf6,
	0x50, 0xb2, 0x1c, 0x2a, 0x91, 0x6d, 0xdf, 0x73, 0x19, 0xd5, 0x21, 0xa6, 0xa8, 0x4d, 0xed, 0xb0,
	0x3a, 0xb3, 0x6c, 0x14, 0xa8, 0x6a, 0x26, 0x6d, 0x68, 0x6f, 0x0a, 0xad, 0x97, 0x58, 0x90, 0xd5,
	0x7c, 0x5e, 0x67, 0x2c, 0x82, 0x1e, 0x7e, 0x5a, 0x08, 0x85, 0xce, 0xcd, 0x3e, 0xdf, 0xa4, 0xc9,
	0x37, 0x0f, 0x86, 0xbc, 0x51, 0xc6, 0xc1, 0x6c, 0x8c, 0xc0, 0x7f, 0xc6, 0x35, 0x55, 0x34, 0xe4,
	0x36, 0xb4, 0x3f, 0x96, 0xcb, 0xca, 0x64, 0xb9, 0x21, 0xb1, 0x43, 0xbe, 0x49, 0x59, 0x0a, 0x27,
	0x75, 0xa8, 0x27, 0x0a, 0x35, 0x56, 0x86, 0x8a, 0xeb, 0xf3, 0x7d, 0x98, 0xfd, 0x0b, 0x61, 0x36,
	0x55, 0x88, 0x73, 0x7b, 0xc7, 0x8d, 0xc5, 0x0e, 0xb0, 0xa7, 0xa2, 0x12, 0x46, 0x64, 0xe5, 0xfd,
	0x84, 0x0a, 0x1e, 0xf2, 0x1d, 0x60, 0x4f, 0x73, 0x85, 0x99, 0xc1, 0xe2, 0xc6, 0x90, 0xd7, 0x7d,
	0xbe, 0x03, 0x1a, 0x73, 0xdb, 0x6f, 0xce, 0x6d, 0xf2, 0xd3, 0x83, 0xbf, 0xda, 0x53, 0xb7, 0xeb,
	0x34, 0xa4, 0x4e, 0x2f, 0x00, 0x44, 0x81, 0x95, 0x95, 0x0d, 0x55, 0x2d, 0x41, 0x03, 0x79, 0x41,
	0x46, 0xff, 0xb7, 0x32, 0xba, 0x0a, 0xba, 0xad, 0xcd, 0xd1, 0x10, 0x21, 0x68, 0x89, 0xc0, 0xae,
	0x01, 0xf2, 0xcd, 0x72, 0xb2, 0x0a, 0xd9, 0xc1, 0x3f, 0x71, 0xe3, 0xbb, 0x5d, 0x5a, 0xbc, 0x71,
	0x85, 0x25, 0x30, 0xcc, 0xe5, 0xfc, 0x49, 0x54, 0xf4, 0xa6, 0x26, 0x16, 0x86, 0xbc, 0x85, 0xd9,
	0xf6, 0x56, 0x63, 0x22, 0xa1, 0xcf, 0x3b, 0xab, 0x71, 0xf2, 0xa3, 0x03, 0xc1, 0x1b, 0x65, 0x5d,
	0xb2, 0x2f, 0xf1, 0x61, 0x63, 0x9d, 0x17, 0x1b, 0x6b, 0x34, 0xe0, 0xb7, 0x1b, 0xd8, 0xae, 0x9e,
	0xee, 0x1f, 0x57, 0x8f, 0xdd, 0x1a, 0xf9, 0x6e, 0x38, 0x1e, 0x9c, 0xe1, 0x9d, 0x05, 0x0e, 0x0f,
	0x68, 0xbe, 0x9b, 0xaa, 0x39, 0x7a, 0x42, 0xbe, 0x87, 0x36, 0x48, 0xef, 0xb5, 0x48, 0x3f, 0x83,
	0xc0, 0xee, 0x2f, 0xeb, 0x06, 0xfb, 0x99, 0x4b, 0xac, 0x51, 0x9f, 0x70, 0x9a, 0x55, 0x13, 0x25,
	0x73, 0xd4, 0x5a, 0x54, 0x53, 0x5a, 0xe8, 0x7d, 0xbe, 0x0f, 0x93, 0xd9, 0x9d, 0xb7, 0x22, 0x70,
	0x3d, 0xd7, 0xa9, 0x5d, 0xf0, 0xab, 0xf1, 0x4d, 0xbb, 0xb6, 0x41, 0xec, 0xa7, 0x3e, 0x3f, 0xc0,
	0x93, 0x1e, 0x04, 0xaf, 0xe7, 0x0b, 0xb3, 0x7e, 0xd5, 0x7b, 0x1f, 0xd0, 0xdf, 0xd2, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x0a, 0x30, 0x02, 0xc8, 0xae, 0x06, 0x00, 0x00,
}
