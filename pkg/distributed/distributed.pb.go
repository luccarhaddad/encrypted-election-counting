// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.0
// source: api/distributed.proto

package distributed

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

type KeyShareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyShare    []byte `protobuf:"bytes,1,opt,name=key_share,json=keyShare,proto3" json:"key_share,omitempty"`
	ShareIndex  int32  `protobuf:"varint,2,opt,name=share_index,json=shareIndex,proto3" json:"share_index,omitempty"`
	TotalShares int32  `protobuf:"varint,3,opt,name=total_shares,json=totalShares,proto3" json:"total_shares,omitempty"`
	PublicKey   []byte `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *KeyShareRequest) Reset() {
	*x = KeyShareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_distributed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyShareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyShareRequest) ProtoMessage() {}

func (x *KeyShareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_distributed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyShareRequest.ProtoReflect.Descriptor instead.
func (*KeyShareRequest) Descriptor() ([]byte, []int) {
	return file_api_distributed_proto_rawDescGZIP(), []int{0}
}

func (x *KeyShareRequest) GetKeyShare() []byte {
	if x != nil {
		return x.KeyShare
	}
	return nil
}

func (x *KeyShareRequest) GetShareIndex() int32 {
	if x != nil {
		return x.ShareIndex
	}
	return 0
}

func (x *KeyShareRequest) GetTotalShares() int32 {
	if x != nil {
		return x.TotalShares
	}
	return 0
}

func (x *KeyShareRequest) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type KeyShareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *KeyShareResponse) Reset() {
	*x = KeyShareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_distributed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyShareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyShareResponse) ProtoMessage() {}

func (x *KeyShareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_distributed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyShareResponse.ProtoReflect.Descriptor instead.
func (*KeyShareResponse) Descriptor() ([]byte, []int) {
	return file_api_distributed_proto_rawDescGZIP(), []int{1}
}

func (x *KeyShareResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidateId string `protobuf:"bytes,1,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_distributed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_distributed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_api_distributed_proto_rawDescGZIP(), []int{2}
}

func (x *VoteRequest) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

type VoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VoteResponse) Reset() {
	*x = VoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_distributed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_distributed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_api_distributed_proto_rawDescGZIP(), []int{3}
}

func (x *VoteResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AggregateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptedAggregates map[string][]byte `protobuf:"bytes,1,rep,name=encrypted_aggregates,json=encryptedAggregates,proto3" json:"encrypted_aggregates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AggregateRequest) Reset() {
	*x = AggregateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_distributed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateRequest) ProtoMessage() {}

func (x *AggregateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_distributed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateRequest.ProtoReflect.Descriptor instead.
func (*AggregateRequest) Descriptor() ([]byte, []int) {
	return file_api_distributed_proto_rawDescGZIP(), []int{4}
}

func (x *AggregateRequest) GetEncryptedAggregates() map[string][]byte {
	if x != nil {
		return x.EncryptedAggregates
	}
	return nil
}

type AckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AckResponse) Reset() {
	*x = AckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_distributed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckResponse) ProtoMessage() {}

func (x *AckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_distributed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckResponse.ProtoReflect.Descriptor instead.
func (*AckResponse) Descriptor() ([]byte, []int) {
	return file_api_distributed_proto_rawDescGZIP(), []int{5}
}

func (x *AckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DecryptRequest) Reset() {
	*x = DecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_distributed_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptRequest) ProtoMessage() {}

func (x *DecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_distributed_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptRequest.ProtoReflect.Descriptor instead.
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return file_api_distributed_proto_rawDescGZIP(), []int{6}
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinalResults map[string]int32 `protobuf:"bytes,1,rep,name=final_results,json=finalResults,proto3" json:"final_results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_distributed_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_distributed_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_api_distributed_proto_rawDescGZIP(), []int{7}
}

func (x *ResultResponse) GetFinalResults() map[string]int32 {
	if x != nil {
		return x.FinalResults
	}
	return nil
}

var File_api_distributed_proto protoreflect.FileDescriptor

var file_api_distributed_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x2c, 0x0a, 0x10, 0x4b, 0x65, 0x79, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xc5, 0x01, 0x0a, 0x10, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x64, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x1a, 0x46, 0x0a, 0x18, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x27, 0x0a, 0x0b, 0x41, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x46, 0x69,
	0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xe6, 0x01, 0x0a, 0x06,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x64, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x53, 0x65, 0x6e,
	0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4b,
	0x65, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x1c, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x64, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xaa, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x12, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_distributed_proto_rawDescOnce sync.Once
	file_api_distributed_proto_rawDescData = file_api_distributed_proto_rawDesc
)

func file_api_distributed_proto_rawDescGZIP() []byte {
	file_api_distributed_proto_rawDescOnce.Do(func() {
		file_api_distributed_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_distributed_proto_rawDescData)
	})
	return file_api_distributed_proto_rawDescData
}

var file_api_distributed_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_distributed_proto_goTypes = []any{
	(*KeyShareRequest)(nil),  // 0: distributed.KeyShareRequest
	(*KeyShareResponse)(nil), // 1: distributed.KeyShareResponse
	(*VoteRequest)(nil),      // 2: distributed.VoteRequest
	(*VoteResponse)(nil),     // 3: distributed.VoteResponse
	(*AggregateRequest)(nil), // 4: distributed.AggregateRequest
	(*AckResponse)(nil),      // 5: distributed.AckResponse
	(*DecryptRequest)(nil),   // 6: distributed.DecryptRequest
	(*ResultResponse)(nil),   // 7: distributed.ResultResponse
	nil,                      // 8: distributed.AggregateRequest.EncryptedAggregatesEntry
	nil,                      // 9: distributed.ResultResponse.FinalResultsEntry
}
var file_api_distributed_proto_depIdxs = []int32{
	8, // 0: distributed.AggregateRequest.encrypted_aggregates:type_name -> distributed.AggregateRequest.EncryptedAggregatesEntry
	9, // 1: distributed.ResultResponse.final_results:type_name -> distributed.ResultResponse.FinalResultsEntry
	2, // 2: distributed.Worker.SubmitVote:input_type -> distributed.VoteRequest
	4, // 3: distributed.Worker.SendAggregates:input_type -> distributed.AggregateRequest
	0, // 4: distributed.Worker.ReceiveKeyShare:input_type -> distributed.KeyShareRequest
	4, // 5: distributed.Controller.ReceiveAggregates:input_type -> distributed.AggregateRequest
	6, // 6: distributed.Controller.DecryptFinalResult:input_type -> distributed.DecryptRequest
	3, // 7: distributed.Worker.SubmitVote:output_type -> distributed.VoteResponse
	5, // 8: distributed.Worker.SendAggregates:output_type -> distributed.AckResponse
	1, // 9: distributed.Worker.ReceiveKeyShare:output_type -> distributed.KeyShareResponse
	5, // 10: distributed.Controller.ReceiveAggregates:output_type -> distributed.AckResponse
	7, // 11: distributed.Controller.DecryptFinalResult:output_type -> distributed.ResultResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_distributed_proto_init() }
func file_api_distributed_proto_init() {
	if File_api_distributed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_distributed_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KeyShareRequest); i {
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
		file_api_distributed_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*KeyShareResponse); i {
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
		file_api_distributed_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VoteRequest); i {
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
		file_api_distributed_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*VoteResponse); i {
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
		file_api_distributed_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AggregateRequest); i {
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
		file_api_distributed_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AckResponse); i {
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
		file_api_distributed_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DecryptRequest); i {
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
		file_api_distributed_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ResultResponse); i {
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
			RawDescriptor: file_api_distributed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_distributed_proto_goTypes,
		DependencyIndexes: file_api_distributed_proto_depIdxs,
		MessageInfos:      file_api_distributed_proto_msgTypes,
	}.Build()
	File_api_distributed_proto = out.File
	file_api_distributed_proto_rawDesc = nil
	file_api_distributed_proto_goTypes = nil
	file_api_distributed_proto_depIdxs = nil
}
