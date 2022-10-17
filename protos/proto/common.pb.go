// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: common.proto

package proto

import (
	types "github.com/GTLiSunnyi/cita-sdk-go/types"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Hash) Reset() {
	*x = Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hash) ProtoMessage() {}

func (x *Hash) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hash.ProtoReflect.Descriptor instead.
func (*Hash) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Hash) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type Hashes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes []*Hash `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *Hashes) Reset() {
	*x = Hashes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hashes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hashes) ProtoMessage() {}

func (x *Hashes) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hashes.ProtoReflect.Descriptor instead.
func (*Hashes) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Hashes) GetHashes() []*Hash {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *Proposal) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Proposal) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProposalWithProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proposal *Proposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
	Proof    []byte    `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *ProposalWithProof) Reset() {
	*x = ProposalWithProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalWithProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalWithProof) ProtoMessage() {}

func (x *ProposalWithProof) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalWithProof.ProtoReflect.Descriptor instead.
func (*ProposalWithProof) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *ProposalWithProof) GetProposal() *Proposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

func (x *ProposalWithProof) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type BFTProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreStateRoot []byte       `protobuf:"bytes,1,opt,name=pre_state_root,json=preStateRoot,proto3" json:"pre_state_root,omitempty"`
	PreProof     []byte       `protobuf:"bytes,2,opt,name=pre_proof,json=preProof,proto3" json:"pre_proof,omitempty"`
	Proposal     *types.Block `protobuf:"bytes,3,opt,name=proposal,proto3" json:"proposal,omitempty"`
}

func (x *BFTProposal) Reset() {
	*x = BFTProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFTProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFTProposal) ProtoMessage() {}

func (x *BFTProposal) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFTProposal.ProtoReflect.Descriptor instead.
func (*BFTProposal) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *BFTProposal) GetPreStateRoot() []byte {
	if x != nil {
		return x.PreStateRoot
	}
	return nil
}

func (x *BFTProposal) GetPreProof() []byte {
	if x != nil {
		return x.PreProof
	}
	return nil
}

func (x *BFTProposal) GetProposal() *types.Block {
	if x != nil {
		return x.Proposal
	}
	return nil
}

type ProposalEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Proposal:
	//	*ProposalEnum_BftProposal
	Proposal isProposalEnum_Proposal `protobuf_oneof:"proposal"`
}

func (x *ProposalEnum) Reset() {
	*x = ProposalEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalEnum) ProtoMessage() {}

func (x *ProposalEnum) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalEnum.ProtoReflect.Descriptor instead.
func (*ProposalEnum) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (m *ProposalEnum) GetProposal() isProposalEnum_Proposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func (x *ProposalEnum) GetBftProposal() *BFTProposal {
	if x, ok := x.GetProposal().(*ProposalEnum_BftProposal); ok {
		return x.BftProposal
	}
	return nil
}

type isProposalEnum_Proposal interface {
	isProposalEnum_Proposal()
}

type ProposalEnum_BftProposal struct {
	BftProposal *BFTProposal `protobuf:"bytes,1,opt,name=bft_proposal,json=bftProposal,proto3,oneof"`
}

func (*ProposalEnum_BftProposal) isProposalEnum_Proposal() {}

type ConsensusConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height        uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	BlockInterval uint32   `protobuf:"varint,2,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
	Validators    [][]byte `protobuf:"bytes,3,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (x *ConsensusConfiguration) Reset() {
	*x = ConsensusConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusConfiguration) ProtoMessage() {}

func (x *ConsensusConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusConfiguration.ProtoReflect.Descriptor instead.
func (*ConsensusConfiguration) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{8}
}

func (x *ConsensusConfiguration) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ConsensusConfiguration) GetBlockInterval() uint32 {
	if x != nil {
		return x.BlockInterval
	}
	return 0
}

func (x *ConsensusConfiguration) GetValidators() [][]byte {
	if x != nil {
		return x.Validators
	}
	return nil
}

type StatusCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *StatusCode) Reset() {
	*x = StatusCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusCode) ProtoMessage() {}

func (x *StatusCode) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusCode.ProtoReflect.Descriptor instead.
func (*StatusCode) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{9}
}

func (x *StatusCode) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type HashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *StatusCode `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Hash   *Hash       `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *HashResponse) Reset() {
	*x = HashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashResponse) ProtoMessage() {}

func (x *HashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashResponse.ProtoReflect.Descriptor instead.
func (*HashResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{10}
}

func (x *HashResponse) GetStatus() *StatusCode {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *HashResponse) GetHash() *Hash {
	if x != nil {
		return x.Hash
	}
	return nil
}

type ProposalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *StatusCode `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Proposal *Proposal   `protobuf:"bytes,2,opt,name=proposal,proto3" json:"proposal,omitempty"`
}

func (x *ProposalResponse) Reset() {
	*x = ProposalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalResponse) ProtoMessage() {}

func (x *ProposalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalResponse.ProtoReflect.Descriptor instead.
func (*ProposalResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{11}
}

func (x *ProposalResponse) GetStatus() *StatusCode {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ProposalResponse) GetProposal() *Proposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

type ConsensusConfigurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *StatusCode             `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Config *ConsensusConfiguration `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ConsensusConfigurationResponse) Reset() {
	*x = ConsensusConfigurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusConfigurationResponse) ProtoMessage() {}

func (x *ConsensusConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusConfigurationResponse.ProtoReflect.Descriptor instead.
func (*ConsensusConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{12}
}

func (x *ConsensusConfigurationResponse) GetStatus() *StatusCode {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ConsensusConfigurationResponse) GetConfig() *ConsensusConfiguration {
	if x != nil {
		return x.Config
	}
	return nil
}

type NodeNetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MultiAddress string `protobuf:"bytes,1,opt,name=multi_address,json=multiAddress,proto3" json:"multi_address,omitempty"`
	Origin       uint64 `protobuf:"varint,2,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *NodeNetInfo) Reset() {
	*x = NodeNetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeNetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeNetInfo) ProtoMessage() {}

func (x *NodeNetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeNetInfo.ProtoReflect.Descriptor instead.
func (*NodeNetInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{13}
}

func (x *NodeNetInfo) GetMultiAddress() string {
	if x != nil {
		return x.MultiAddress
	}
	return ""
}

func (x *NodeNetInfo) GetOrigin() uint64 {
	if x != nil {
		return x.Origin
	}
	return 0
}

type TotalNodeNetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*NodeNetInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *TotalNodeNetInfo) Reset() {
	*x = TotalNodeNetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalNodeNetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalNodeNetInfo) ProtoMessage() {}

func (x *TotalNodeNetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalNodeNetInfo.ProtoReflect.Descriptor instead.
func (*TotalNodeNetInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{14}
}

func (x *TotalNodeNetInfo) GetNodes() []*NodeNetInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	NetInfo *NodeNetInfo `protobuf:"bytes,2,opt,name=net_info,json=netInfo,proto3" json:"net_info,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{15}
}

func (x *NodeInfo) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *NodeInfo) GetNetInfo() *NodeNetInfo {
	if x != nil {
		return x.NetInfo
	}
	return nil
}

type TotalNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*NodeInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *TotalNodeInfo) Reset() {
	*x = TotalNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalNodeInfo) ProtoMessage() {}

func (x *TotalNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalNodeInfo.ProtoReflect.Descriptor instead.
func (*TotalNodeInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{16}
}

func (x *TotalNodeInfo) GetNodes() []*NodeInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1a, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x2e, 0x0a,
	0x06, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x23, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x36, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x11, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12,
	0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x22, 0x7f, 0x0a, 0x0b, 0x42, 0x46, 0x54, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x22, 0x54, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x46, 0x54, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x48,
	0x00, 0x52, 0x0b, 0x62, 0x66, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x42, 0x0a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x22, 0x77, 0x0a, 0x16, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x5c, 0x0a, 0x0c, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x20, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x22, 0x6c, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x22, 0x84, 0x01, 0x0a, 0x1e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4a, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65,
	0x4e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x22, 0x3d, 0x0a, 0x10, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x6f, 0x64,
	0x65, 0x4e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x6e, 0x65, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x37, 0x0a, 0x0d, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x54, 0x4c, 0x69, 0x53, 0x75, 0x6e, 0x6e, 0x79, 0x69, 0x2f, 0x63, 0x69, 0x74, 0x61,
	0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_common_proto_goTypes = []interface{}{
	(*Empty)(nil),                          // 0: common.Empty
	(*Hash)(nil),                           // 1: common.Hash
	(*Hashes)(nil),                         // 2: common.Hashes
	(*Address)(nil),                        // 3: common.Address
	(*Proposal)(nil),                       // 4: common.Proposal
	(*ProposalWithProof)(nil),              // 5: common.ProposalWithProof
	(*BFTProposal)(nil),                    // 6: common.BFTProposal
	(*ProposalEnum)(nil),                   // 7: common.ProposalEnum
	(*ConsensusConfiguration)(nil),         // 8: common.ConsensusConfiguration
	(*StatusCode)(nil),                     // 9: common.StatusCode
	(*HashResponse)(nil),                   // 10: common.HashResponse
	(*ProposalResponse)(nil),               // 11: common.ProposalResponse
	(*ConsensusConfigurationResponse)(nil), // 12: common.ConsensusConfigurationResponse
	(*NodeNetInfo)(nil),                    // 13: common.NodeNetInfo
	(*TotalNodeNetInfo)(nil),               // 14: common.TotalNodeNetInfo
	(*NodeInfo)(nil),                       // 15: common.NodeInfo
	(*TotalNodeInfo)(nil),                  // 16: common.TotalNodeInfo
	(*types.Block)(nil),                    // 17: blockchain.Block
}
var file_common_proto_depIdxs = []int32{
	1,  // 0: common.Hashes.hashes:type_name -> common.Hash
	4,  // 1: common.ProposalWithProof.proposal:type_name -> common.Proposal
	17, // 2: common.BFTProposal.proposal:type_name -> blockchain.Block
	6,  // 3: common.ProposalEnum.bft_proposal:type_name -> common.BFTProposal
	9,  // 4: common.HashResponse.status:type_name -> common.StatusCode
	1,  // 5: common.HashResponse.hash:type_name -> common.Hash
	9,  // 6: common.ProposalResponse.status:type_name -> common.StatusCode
	4,  // 7: common.ProposalResponse.proposal:type_name -> common.Proposal
	9,  // 8: common.ConsensusConfigurationResponse.status:type_name -> common.StatusCode
	8,  // 9: common.ConsensusConfigurationResponse.config:type_name -> common.ConsensusConfiguration
	13, // 10: common.TotalNodeNetInfo.nodes:type_name -> common.NodeNetInfo
	13, // 11: common.NodeInfo.net_info:type_name -> common.NodeNetInfo
	15, // 12: common.TotalNodeInfo.nodes:type_name -> common.NodeInfo
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hash); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hashes); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalWithProof); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFTProposal); i {
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
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalEnum); i {
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
		file_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusConfiguration); i {
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
		file_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusCode); i {
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
		file_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashResponse); i {
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
		file_common_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalResponse); i {
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
		file_common_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusConfigurationResponse); i {
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
		file_common_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeNetInfo); i {
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
		file_common_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalNodeNetInfo); i {
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
		file_common_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_common_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalNodeInfo); i {
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
	file_common_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*ProposalEnum_BftProposal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
