// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/signproxy/signproxy.proto

package signproxy

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

type TransactionType int32

const (
	TransactionType_UnKnow        TransactionType = 0
	TransactionType_CreateAccount TransactionType = 1
	TransactionType_Transaction   TransactionType = 2
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "UnKnow",
		1: "CreateAccount",
		2: "Transaction",
	}
	TransactionType_value = map[string]int32{
		"UnKnow":        0,
		"CreateAccount": 1,
		"Transaction":   2,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_signproxy_signproxy_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_npool_signproxy_signproxy_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{0}
}

type WalletBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,100,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *WalletBalanceRequest) Reset() {
	*x = WalletBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletBalanceRequest) ProtoMessage() {}

func (x *WalletBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletBalanceRequest.ProtoReflect.Descriptor instead.
func (*WalletBalanceRequest) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{0}
}

func (x *WalletBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type WalletBalanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance string `protobuf:"bytes,100,opt,name=Balance,proto3" json:"Balance,omitempty"`
}

func (x *WalletBalanceInfo) Reset() {
	*x = WalletBalanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletBalanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletBalanceInfo) ProtoMessage() {}

func (x *WalletBalanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletBalanceInfo.ProtoReflect.Descriptor instead.
func (*WalletBalanceInfo) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{1}
}

func (x *WalletBalanceInfo) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

type WalletBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *WalletBalanceInfo `protobuf:"bytes,100,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *WalletBalanceResponse) Reset() {
	*x = WalletBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletBalanceResponse) ProtoMessage() {}

func (x *WalletBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletBalanceResponse.ProtoReflect.Descriptor instead.
func (*WalletBalanceResponse) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{2}
}

func (x *WalletBalanceResponse) GetInfo() *WalletBalanceInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type MpoolGetNonceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,100,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *MpoolGetNonceRequest) Reset() {
	*x = MpoolGetNonceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpoolGetNonceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpoolGetNonceRequest) ProtoMessage() {}

func (x *MpoolGetNonceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpoolGetNonceRequest.ProtoReflect.Descriptor instead.
func (*MpoolGetNonceRequest) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{3}
}

func (x *MpoolGetNonceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type MpoolGetNonceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce uint64 `protobuf:"varint,100,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
}

func (x *MpoolGetNonceInfo) Reset() {
	*x = MpoolGetNonceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpoolGetNonceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpoolGetNonceInfo) ProtoMessage() {}

func (x *MpoolGetNonceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpoolGetNonceInfo.ProtoReflect.Descriptor instead.
func (*MpoolGetNonceInfo) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{4}
}

func (x *MpoolGetNonceInfo) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

type MpoolGetNonceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *MpoolGetNonceInfo `protobuf:"bytes,100,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *MpoolGetNonceResponse) Reset() {
	*x = MpoolGetNonceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpoolGetNonceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpoolGetNonceResponse) ProtoMessage() {}

func (x *MpoolGetNonceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpoolGetNonceResponse.ProtoReflect.Descriptor instead.
func (*MpoolGetNonceResponse) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{5}
}

func (x *MpoolGetNonceResponse) GetInfo() *MpoolGetNonceInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type RegisterCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinType string `protobuf:"bytes,100,opt,name=CoinType,proto3" json:"CoinType,omitempty"`
}

func (x *RegisterCoinRequest) Reset() {
	*x = RegisterCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCoinRequest) ProtoMessage() {}

func (x *RegisterCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCoinRequest.ProtoReflect.Descriptor instead.
func (*RegisterCoinRequest) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterCoinRequest) GetCoinType() string {
	if x != nil {
		return x.CoinType
	}
	return ""
}

type RegisterCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterCoinResponse) Reset() {
	*x = RegisterCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCoinResponse) ProtoMessage() {}

func (x *RegisterCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCoinResponse.ProtoReflect.Descriptor instead.
func (*RegisterCoinResponse) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{7}
}

type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionType TransactionType  `protobuf:"varint,100,opt,name=TransactionType,proto3,enum=sphinx.proxy.v1.TransactionType" json:"TransactionType,omitempty"`
	Message         *UnsignedMessage `protobuf:"bytes,110,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{8}
}

func (x *TransactionRequest) GetTransactionType() TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return TransactionType_UnKnow
}

func (x *TransactionRequest) GetMessage() *UnsignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type UnsignedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    uint64 `protobuf:"varint,100,opt,name=Version,proto3" json:"Version,omitempty"`
	To         string `protobuf:"bytes,110,opt,name=To,proto3" json:"To,omitempty"`
	From       string `protobuf:"bytes,120,opt,name=From,proto3" json:"From,omitempty"`
	Nonce      uint64 `protobuf:"varint,130,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Value      string `protobuf:"bytes,140,opt,name=Value,proto3" json:"Value,omitempty"`
	GasLimit   int64  `protobuf:"varint,150,opt,name=GasLimit,proto3" json:"GasLimit,omitempty"`
	GasFeeCap  string `protobuf:"bytes,160,opt,name=GasFeeCap,proto3" json:"GasFeeCap,omitempty"`
	GasPremium string `protobuf:"bytes,170,opt,name=GasPremium,proto3" json:"GasPremium,omitempty"`
	Method     uint64 `protobuf:"varint,180,opt,name=Method,proto3" json:"Method,omitempty"`
	Params     []byte `protobuf:"bytes,190,opt,name=Params,proto3" json:"Params,omitempty"`
}

func (x *UnsignedMessage) Reset() {
	*x = UnsignedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsignedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsignedMessage) ProtoMessage() {}

func (x *UnsignedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsignedMessage.ProtoReflect.Descriptor instead.
func (*UnsignedMessage) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{9}
}

func (x *UnsignedMessage) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UnsignedMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *UnsignedMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *UnsignedMessage) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *UnsignedMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *UnsignedMessage) GetGasLimit() int64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *UnsignedMessage) GetGasFeeCap() string {
	if x != nil {
		return x.GasFeeCap
	}
	return ""
}

func (x *UnsignedMessage) GetGasPremium() string {
	if x != nil {
		return x.GasPremium
	}
	return ""
}

func (x *UnsignedMessage) GetMethod() uint64 {
	if x != nil {
		return x.Method
	}
	return 0
}

func (x *UnsignedMessage) GetParams() []byte {
	if x != nil {
		return x.Params
	}
	return nil
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *TransactionResponseInfo `protobuf:"bytes,100,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{10}
}

func (x *TransactionResponse) GetInfo() *TransactionResponseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type TransactionResponseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *UnsignedMessage `protobuf:"bytes,100,opt,name=Message,proto3" json:"Message,omitempty"`
	Signature *Signature       `protobuf:"bytes,110,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *TransactionResponseInfo) Reset() {
	*x = TransactionResponseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponseInfo) ProtoMessage() {}

func (x *TransactionResponseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponseInfo.ProtoReflect.Descriptor instead.
func (*TransactionResponseInfo) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{11}
}

func (x *TransactionResponseInfo) GetMessage() *UnsignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *TransactionResponseInfo) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignType string `protobuf:"bytes,100,opt,name=SignType,proto3" json:"SignType,omitempty"` //secp256k1
	Data     []byte `protobuf:"bytes,110,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{12}
}

func (x *Signature) GetSignType() string {
	if x != nil {
		return x.SignType
	}
	return ""
}

func (x *Signature) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_npool_signproxy_signproxy_proto protoreflect.FileDescriptor

var file_npool_signproxy_signproxy_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x22, 0x30, 0x0a, 0x14, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x11, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x4f, 0x0a, 0x15, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x70, 0x68,
	0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x30, 0x0a, 0x14, 0x4d, 0x70, 0x6f, 0x6f, 0x6c, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x29, 0x0a, 0x11, 0x4d, 0x70, 0x6f, 0x6f, 0x6c, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4e, 0x6f, 0x6e, 0x63,
	0x65, 0x22, 0x4f, 0x0a, 0x15, 0x4d, 0x70, 0x6f, 0x6f, 0x6c, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x70, 0x6f, 0x6f, 0x6c,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x31, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9c, 0x01,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3a, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8c, 0x02, 0x0a,
	0x0f, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f,
	0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72,
	0x6f, 0x6d, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x15,
	0x0a, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x8c,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x08,
	0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x47, 0x61, 0x73,
	0x46, 0x65, 0x65, 0x43, 0x61, 0x70, 0x18, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47,
	0x61, 0x73, 0x46, 0x65, 0x65, 0x43, 0x61, 0x70, 0x12, 0x1f, 0x0a, 0x0a, 0x47, 0x61, 0x73, 0x50,
	0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x47,
	0x61, 0x73, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x06, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x17, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0xbe, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x53, 0x0a, 0x13, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x8f, 0x01, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70,
	0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x3b, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x2a,
	0x41, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x02, 0x32, 0x8e, 0x03, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x12, 0x5d, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e,
	0x12, 0x24, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x60, 0x0a, 0x0d, 0x4d, 0x70, 0x6f, 0x6f, 0x6c, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x25, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x70, 0x6f, 0x6f, 0x6c, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x70, 0x6f, 0x6f, 0x6c, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5e, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x24, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x60, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x68, 0x69,
	0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_signproxy_signproxy_proto_rawDescOnce sync.Once
	file_npool_signproxy_signproxy_proto_rawDescData = file_npool_signproxy_signproxy_proto_rawDesc
)

func file_npool_signproxy_signproxy_proto_rawDescGZIP() []byte {
	file_npool_signproxy_signproxy_proto_rawDescOnce.Do(func() {
		file_npool_signproxy_signproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_signproxy_signproxy_proto_rawDescData)
	})
	return file_npool_signproxy_signproxy_proto_rawDescData
}

var file_npool_signproxy_signproxy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_signproxy_signproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_npool_signproxy_signproxy_proto_goTypes = []interface{}{
	(TransactionType)(0),            // 0: sphinx.proxy.v1.TransactionType
	(*WalletBalanceRequest)(nil),    // 1: sphinx.proxy.v1.WalletBalanceRequest
	(*WalletBalanceInfo)(nil),       // 2: sphinx.proxy.v1.WalletBalanceInfo
	(*WalletBalanceResponse)(nil),   // 3: sphinx.proxy.v1.WalletBalanceResponse
	(*MpoolGetNonceRequest)(nil),    // 4: sphinx.proxy.v1.MpoolGetNonceRequest
	(*MpoolGetNonceInfo)(nil),       // 5: sphinx.proxy.v1.MpoolGetNonceInfo
	(*MpoolGetNonceResponse)(nil),   // 6: sphinx.proxy.v1.MpoolGetNonceResponse
	(*RegisterCoinRequest)(nil),     // 7: sphinx.proxy.v1.RegisterCoinRequest
	(*RegisterCoinResponse)(nil),    // 8: sphinx.proxy.v1.RegisterCoinResponse
	(*TransactionRequest)(nil),      // 9: sphinx.proxy.v1.TransactionRequest
	(*UnsignedMessage)(nil),         // 10: sphinx.proxy.v1.UnsignedMessage
	(*TransactionResponse)(nil),     // 11: sphinx.proxy.v1.TransactionResponse
	(*TransactionResponseInfo)(nil), // 12: sphinx.proxy.v1.TransactionResponseInfo
	(*Signature)(nil),               // 13: sphinx.proxy.v1.Signature
}
var file_npool_signproxy_signproxy_proto_depIdxs = []int32{
	2,  // 0: sphinx.proxy.v1.WalletBalanceResponse.Info:type_name -> sphinx.proxy.v1.WalletBalanceInfo
	5,  // 1: sphinx.proxy.v1.MpoolGetNonceResponse.Info:type_name -> sphinx.proxy.v1.MpoolGetNonceInfo
	0,  // 2: sphinx.proxy.v1.TransactionRequest.TransactionType:type_name -> sphinx.proxy.v1.TransactionType
	10, // 3: sphinx.proxy.v1.TransactionRequest.Message:type_name -> sphinx.proxy.v1.UnsignedMessage
	12, // 4: sphinx.proxy.v1.TransactionResponse.Info:type_name -> sphinx.proxy.v1.TransactionResponseInfo
	10, // 5: sphinx.proxy.v1.TransactionResponseInfo.Message:type_name -> sphinx.proxy.v1.UnsignedMessage
	13, // 6: sphinx.proxy.v1.TransactionResponseInfo.Signature:type_name -> sphinx.proxy.v1.Signature
	7,  // 7: sphinx.proxy.v1.SignProxy.RegisterCoin:input_type -> sphinx.proxy.v1.RegisterCoinRequest
	4,  // 8: sphinx.proxy.v1.SignProxy.MpoolGetNonce:input_type -> sphinx.proxy.v1.MpoolGetNonceRequest
	11, // 9: sphinx.proxy.v1.SignProxy.Transaction:input_type -> sphinx.proxy.v1.TransactionResponse
	1,  // 10: sphinx.proxy.v1.SignProxy.WalletBalance:input_type -> sphinx.proxy.v1.WalletBalanceRequest
	8,  // 11: sphinx.proxy.v1.SignProxy.RegisterCoin:output_type -> sphinx.proxy.v1.RegisterCoinResponse
	6,  // 12: sphinx.proxy.v1.SignProxy.MpoolGetNonce:output_type -> sphinx.proxy.v1.MpoolGetNonceResponse
	9,  // 13: sphinx.proxy.v1.SignProxy.Transaction:output_type -> sphinx.proxy.v1.TransactionRequest
	3,  // 14: sphinx.proxy.v1.SignProxy.WalletBalance:output_type -> sphinx.proxy.v1.WalletBalanceResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_npool_signproxy_signproxy_proto_init() }
func file_npool_signproxy_signproxy_proto_init() {
	if File_npool_signproxy_signproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_signproxy_signproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletBalanceRequest); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletBalanceInfo); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletBalanceResponse); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpoolGetNonceRequest); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpoolGetNonceInfo); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpoolGetNonceResponse); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterCoinRequest); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterCoinResponse); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsignedMessage); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponseInfo); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
			RawDescriptor: file_npool_signproxy_signproxy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_signproxy_signproxy_proto_goTypes,
		DependencyIndexes: file_npool_signproxy_signproxy_proto_depIdxs,
		EnumInfos:         file_npool_signproxy_signproxy_proto_enumTypes,
		MessageInfos:      file_npool_signproxy_signproxy_proto_msgTypes,
	}.Build()
	File_npool_signproxy_signproxy_proto = out.File
	file_npool_signproxy_signproxy_proto_rawDesc = nil
	file_npool_signproxy_signproxy_proto_goTypes = nil
	file_npool_signproxy_signproxy_proto_depIdxs = nil
}