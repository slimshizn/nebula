// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.5
// source: cert_v1.proto

package cert

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Curve int32

const (
	Curve_CURVE25519 Curve = 0
	Curve_P256       Curve = 1
)

// Enum value maps for Curve.
var (
	Curve_name = map[int32]string{
		0: "CURVE25519",
		1: "P256",
	}
	Curve_value = map[string]int32{
		"CURVE25519": 0,
		"P256":       1,
	}
)

func (x Curve) Enum() *Curve {
	p := new(Curve)
	*p = x
	return p
}

func (x Curve) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Curve) Descriptor() protoreflect.EnumDescriptor {
	return file_cert_v1_proto_enumTypes[0].Descriptor()
}

func (Curve) Type() protoreflect.EnumType {
	return &file_cert_v1_proto_enumTypes[0]
}

func (x Curve) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Curve.Descriptor instead.
func (Curve) EnumDescriptor() ([]byte, []int) {
	return file_cert_v1_proto_rawDescGZIP(), []int{0}
}

type RawNebulaCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details   *RawNebulaCertificateDetails `protobuf:"bytes,1,opt,name=Details,proto3" json:"Details,omitempty"`
	Signature []byte                       `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *RawNebulaCertificate) Reset() {
	*x = RawNebulaCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cert_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawNebulaCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawNebulaCertificate) ProtoMessage() {}

func (x *RawNebulaCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_cert_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawNebulaCertificate.ProtoReflect.Descriptor instead.
func (*RawNebulaCertificate) Descriptor() ([]byte, []int) {
	return file_cert_v1_proto_rawDescGZIP(), []int{0}
}

func (x *RawNebulaCertificate) GetDetails() *RawNebulaCertificateDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *RawNebulaCertificate) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type RawNebulaCertificateDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Ips and Subnets are in big endian 32 bit pairs, 1st the ip, 2nd the mask
	Ips       []uint32 `protobuf:"varint,2,rep,packed,name=Ips,proto3" json:"Ips,omitempty"`
	Subnets   []uint32 `protobuf:"varint,3,rep,packed,name=Subnets,proto3" json:"Subnets,omitempty"`
	Groups    []string `protobuf:"bytes,4,rep,name=Groups,proto3" json:"Groups,omitempty"`
	NotBefore int64    `protobuf:"varint,5,opt,name=NotBefore,proto3" json:"NotBefore,omitempty"`
	NotAfter  int64    `protobuf:"varint,6,opt,name=NotAfter,proto3" json:"NotAfter,omitempty"`
	PublicKey []byte   `protobuf:"bytes,7,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	IsCA      bool     `protobuf:"varint,8,opt,name=IsCA,proto3" json:"IsCA,omitempty"`
	// sha-256 of the issuer certificate, if this field is blank the cert is self-signed
	Issuer []byte `protobuf:"bytes,9,opt,name=Issuer,proto3" json:"Issuer,omitempty"`
	Curve  Curve  `protobuf:"varint,100,opt,name=curve,proto3,enum=cert.Curve" json:"curve,omitempty"`
}

func (x *RawNebulaCertificateDetails) Reset() {
	*x = RawNebulaCertificateDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cert_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawNebulaCertificateDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawNebulaCertificateDetails) ProtoMessage() {}

func (x *RawNebulaCertificateDetails) ProtoReflect() protoreflect.Message {
	mi := &file_cert_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawNebulaCertificateDetails.ProtoReflect.Descriptor instead.
func (*RawNebulaCertificateDetails) Descriptor() ([]byte, []int) {
	return file_cert_v1_proto_rawDescGZIP(), []int{1}
}

func (x *RawNebulaCertificateDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RawNebulaCertificateDetails) GetIps() []uint32 {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *RawNebulaCertificateDetails) GetSubnets() []uint32 {
	if x != nil {
		return x.Subnets
	}
	return nil
}

func (x *RawNebulaCertificateDetails) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *RawNebulaCertificateDetails) GetNotBefore() int64 {
	if x != nil {
		return x.NotBefore
	}
	return 0
}

func (x *RawNebulaCertificateDetails) GetNotAfter() int64 {
	if x != nil {
		return x.NotAfter
	}
	return 0
}

func (x *RawNebulaCertificateDetails) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *RawNebulaCertificateDetails) GetIsCA() bool {
	if x != nil {
		return x.IsCA
	}
	return false
}

func (x *RawNebulaCertificateDetails) GetIssuer() []byte {
	if x != nil {
		return x.Issuer
	}
	return nil
}

func (x *RawNebulaCertificateDetails) GetCurve() Curve {
	if x != nil {
		return x.Curve
	}
	return Curve_CURVE25519
}

type RawNebulaEncryptedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptionMetadata *RawNebulaEncryptionMetadata `protobuf:"bytes,1,opt,name=EncryptionMetadata,proto3" json:"EncryptionMetadata,omitempty"`
	Ciphertext         []byte                       `protobuf:"bytes,2,opt,name=Ciphertext,proto3" json:"Ciphertext,omitempty"`
}

func (x *RawNebulaEncryptedData) Reset() {
	*x = RawNebulaEncryptedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cert_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawNebulaEncryptedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawNebulaEncryptedData) ProtoMessage() {}

func (x *RawNebulaEncryptedData) ProtoReflect() protoreflect.Message {
	mi := &file_cert_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawNebulaEncryptedData.ProtoReflect.Descriptor instead.
func (*RawNebulaEncryptedData) Descriptor() ([]byte, []int) {
	return file_cert_v1_proto_rawDescGZIP(), []int{2}
}

func (x *RawNebulaEncryptedData) GetEncryptionMetadata() *RawNebulaEncryptionMetadata {
	if x != nil {
		return x.EncryptionMetadata
	}
	return nil
}

func (x *RawNebulaEncryptedData) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

type RawNebulaEncryptionMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptionAlgorithm string                     `protobuf:"bytes,1,opt,name=EncryptionAlgorithm,proto3" json:"EncryptionAlgorithm,omitempty"`
	Argon2Parameters    *RawNebulaArgon2Parameters `protobuf:"bytes,2,opt,name=Argon2Parameters,proto3" json:"Argon2Parameters,omitempty"`
}

func (x *RawNebulaEncryptionMetadata) Reset() {
	*x = RawNebulaEncryptionMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cert_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawNebulaEncryptionMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawNebulaEncryptionMetadata) ProtoMessage() {}

func (x *RawNebulaEncryptionMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_cert_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawNebulaEncryptionMetadata.ProtoReflect.Descriptor instead.
func (*RawNebulaEncryptionMetadata) Descriptor() ([]byte, []int) {
	return file_cert_v1_proto_rawDescGZIP(), []int{3}
}

func (x *RawNebulaEncryptionMetadata) GetEncryptionAlgorithm() string {
	if x != nil {
		return x.EncryptionAlgorithm
	}
	return ""
}

func (x *RawNebulaEncryptionMetadata) GetArgon2Parameters() *RawNebulaArgon2Parameters {
	if x != nil {
		return x.Argon2Parameters
	}
	return nil
}

type RawNebulaArgon2Parameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"` // rune in Go
	Memory      uint32 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Parallelism uint32 `protobuf:"varint,4,opt,name=parallelism,proto3" json:"parallelism,omitempty"` // uint8 in Go
	Iterations  uint32 `protobuf:"varint,3,opt,name=iterations,proto3" json:"iterations,omitempty"`
	Salt        []byte `protobuf:"bytes,5,opt,name=salt,proto3" json:"salt,omitempty"`
}

func (x *RawNebulaArgon2Parameters) Reset() {
	*x = RawNebulaArgon2Parameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cert_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawNebulaArgon2Parameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawNebulaArgon2Parameters) ProtoMessage() {}

func (x *RawNebulaArgon2Parameters) ProtoReflect() protoreflect.Message {
	mi := &file_cert_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawNebulaArgon2Parameters.ProtoReflect.Descriptor instead.
func (*RawNebulaArgon2Parameters) Descriptor() ([]byte, []int) {
	return file_cert_v1_proto_rawDescGZIP(), []int{4}
}

func (x *RawNebulaArgon2Parameters) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *RawNebulaArgon2Parameters) GetMemory() uint32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *RawNebulaArgon2Parameters) GetParallelism() uint32 {
	if x != nil {
		return x.Parallelism
	}
	return 0
}

func (x *RawNebulaArgon2Parameters) GetIterations() uint32 {
	if x != nil {
		return x.Iterations
	}
	return 0
}

func (x *RawNebulaArgon2Parameters) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

var File_cert_v1_proto protoreflect.FileDescriptor

var file_cert_v1_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x63, 0x65, 0x72, 0x74, 0x22, 0x71, 0x0a, 0x14, 0x52, 0x61, 0x77, 0x4e, 0x65, 0x62, 0x75,
	0x6c, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a,
	0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x52, 0x61, 0x77, 0x4e, 0x65, 0x62, 0x75, 0x6c, 0x61, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x9c, 0x02, 0x0a, 0x1b, 0x52, 0x61, 0x77,
	0x4e, 0x65, 0x62, 0x75, 0x6c, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x49, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x49, 0x70, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x07, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x4e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x4e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x73, 0x43, 0x41,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x49, 0x73, 0x43, 0x41, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x76, 0x65,
	0x52, 0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x16, 0x52, 0x61, 0x77, 0x4e,
	0x65, 0x62, 0x75, 0x6c, 0x61, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x51, 0x0a, 0x12, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x52, 0x61, 0x77, 0x4e, 0x65, 0x62, 0x75, 0x6c, 0x61, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x12, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x43, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x74, 0x65, 0x78, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x1b, 0x52, 0x61, 0x77, 0x4e, 0x65, 0x62,
	0x75, 0x6c, 0x61, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x13, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x4b, 0x0a, 0x10, 0x41, 0x72, 0x67, 0x6f, 0x6e,
	0x32, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x52, 0x61, 0x77, 0x4e, 0x65, 0x62, 0x75,
	0x6c, 0x61, 0x41, 0x72, 0x67, 0x6f, 0x6e, 0x32, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x10, 0x41, 0x72, 0x67, 0x6f, 0x6e, 0x32, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x19, 0x52, 0x61, 0x77, 0x4e, 0x65, 0x62, 0x75,
	0x6c, 0x61, 0x41, 0x72, 0x67, 0x6f, 0x6e, 0x32, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c,
	0x69, 0x73, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6c,
	0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x2a, 0x21, 0x0a, 0x05, 0x43, 0x75,
	0x72, 0x76, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x55, 0x52, 0x56, 0x45, 0x32, 0x35, 0x35, 0x31,
	0x39, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x32, 0x35, 0x36, 0x10, 0x01, 0x42, 0x20, 0x5a,
	0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x61, 0x63,
	0x6b, 0x68, 0x71, 0x2f, 0x6e, 0x65, 0x62, 0x75, 0x6c, 0x61, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cert_v1_proto_rawDescOnce sync.Once
	file_cert_v1_proto_rawDescData = file_cert_v1_proto_rawDesc
)

func file_cert_v1_proto_rawDescGZIP() []byte {
	file_cert_v1_proto_rawDescOnce.Do(func() {
		file_cert_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_cert_v1_proto_rawDescData)
	})
	return file_cert_v1_proto_rawDescData
}

var file_cert_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cert_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cert_v1_proto_goTypes = []any{
	(Curve)(0),                          // 0: cert.Curve
	(*RawNebulaCertificate)(nil),        // 1: cert.RawNebulaCertificate
	(*RawNebulaCertificateDetails)(nil), // 2: cert.RawNebulaCertificateDetails
	(*RawNebulaEncryptedData)(nil),      // 3: cert.RawNebulaEncryptedData
	(*RawNebulaEncryptionMetadata)(nil), // 4: cert.RawNebulaEncryptionMetadata
	(*RawNebulaArgon2Parameters)(nil),   // 5: cert.RawNebulaArgon2Parameters
}
var file_cert_v1_proto_depIdxs = []int32{
	2, // 0: cert.RawNebulaCertificate.Details:type_name -> cert.RawNebulaCertificateDetails
	0, // 1: cert.RawNebulaCertificateDetails.curve:type_name -> cert.Curve
	4, // 2: cert.RawNebulaEncryptedData.EncryptionMetadata:type_name -> cert.RawNebulaEncryptionMetadata
	5, // 3: cert.RawNebulaEncryptionMetadata.Argon2Parameters:type_name -> cert.RawNebulaArgon2Parameters
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cert_v1_proto_init() }
func file_cert_v1_proto_init() {
	if File_cert_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cert_v1_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RawNebulaCertificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cert_v1_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RawNebulaCertificateDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cert_v1_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RawNebulaEncryptedData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cert_v1_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RawNebulaEncryptionMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cert_v1_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RawNebulaArgon2Parameters); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cert_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cert_v1_proto_goTypes,
		DependencyIndexes: file_cert_v1_proto_depIdxs,
		EnumInfos:         file_cert_v1_proto_enumTypes,
		MessageInfos:      file_cert_v1_proto_msgTypes,
	}.Build()
	File_cert_v1_proto = out.File
	file_cert_v1_proto_rawDesc = nil
	file_cert_v1_proto_goTypes = nil
	file_cert_v1_proto_depIdxs = nil
}
