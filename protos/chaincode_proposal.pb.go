// Code generated by protoc-gen-go.
// source: chaincode_proposal.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ChaincodeHeaderExtension is the Header's extentions message to be used when
// the Header's type is CHAINCODE.  This extensions is used to specify which
// chaincode to invoke and what should appear on the ledger.
type ChaincodeHeaderExtension struct {
	// The PayloadVisibility field controls to what extent the Proposal's payload
	// (recall that for the type CHAINCODE, it is ChaincodeProposalPayload
	// message) field will be visible in the final transaction and in the ledger.
	// Ideally, it would be configurable, supporting at least 3 main “visibility
	// modes”:
	// 1. all bytes of the payload are visible;
	// 2. only a hash of the payload is visible;
	// 3. nothing is visible.
	// Notice that the visibility function may be potentially part of the ESCC.
	// In that case it overrides PayloadVisibility field.  Finally notice that
	// this field impacts the content of ProposalResponsePayload.proposalHash.
	PayloadVisibility []byte `protobuf:"bytes,1,opt,name=payloadVisibility,proto3" json:"payloadVisibility,omitempty"`
	// The ID of the chaincode to target.
	ChaincodeID *ChaincodeID `protobuf:"bytes,2,opt,name=chaincodeID" json:"chaincodeID,omitempty"`
}

func (m *ChaincodeHeaderExtension) Reset()                    { *m = ChaincodeHeaderExtension{} }
func (m *ChaincodeHeaderExtension) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeHeaderExtension) ProtoMessage()               {}
func (*ChaincodeHeaderExtension) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ChaincodeHeaderExtension) GetChaincodeID() *ChaincodeID {
	if m != nil {
		return m.ChaincodeID
	}
	return nil
}

// ChaincodeProposalPayload is the Proposal's payload message to be used when
// the Header's type is CHAINCODE.  It contains the arguments for this
// invocation.
type ChaincodeProposalPayload struct {
	// Input contains the arguments for this invocation. If this invocation
	// deploys a new chaincode, ESCC/VSCC are part of this field.
	Input []byte `protobuf:"bytes,1,opt,name=Input,proto3" json:"Input,omitempty"`
	// Transient contains data (e.g. cryptographic material) that might be used
	// to implement some form of application-level confidentiality. The contents
	// of this field are supposed to always be omitted from the transaction and
	// excluded from the ledger.
	Transient []byte `protobuf:"bytes,2,opt,name=Transient,proto3" json:"Transient,omitempty"`
}

func (m *ChaincodeProposalPayload) Reset()                    { *m = ChaincodeProposalPayload{} }
func (m *ChaincodeProposalPayload) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeProposalPayload) ProtoMessage()               {}
func (*ChaincodeProposalPayload) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// ChaincodeAction contains the actions the events generated by the execution
// of the chaincode.
type ChaincodeAction struct {
	// This field contains the read set and the write set produced by the
	// chaincode executing this invocation.
	Results []byte `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
	// This field contains the events generated by the chaincode executing this
	// invocation.
	Events []byte `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty"`
}

func (m *ChaincodeAction) Reset()                    { *m = ChaincodeAction{} }
func (m *ChaincodeAction) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeAction) ProtoMessage()               {}
func (*ChaincodeAction) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func init() {
	proto.RegisterType((*ChaincodeHeaderExtension)(nil), "protos.ChaincodeHeaderExtension")
	proto.RegisterType((*ChaincodeProposalPayload)(nil), "protos.ChaincodeProposalPayload")
	proto.RegisterType((*ChaincodeAction)(nil), "protos.ChaincodeAction")
}

func init() { proto.RegisterFile("chaincode_proposal.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0x48, 0xcc,
	0xcc, 0x4b, 0xce, 0x4f, 0x49, 0x8d, 0x2f, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x52, 0xfc, 0x70, 0x15, 0x10, 0x09, 0xa5,
	0x7a, 0x2e, 0x09, 0x67, 0x98, 0x90, 0x47, 0x6a, 0x62, 0x4a, 0x6a, 0x91, 0x6b, 0x45, 0x49, 0x6a,
	0x5e, 0x71, 0x66, 0x7e, 0x9e, 0x90, 0x0e, 0x97, 0x60, 0x41, 0x62, 0x65, 0x4e, 0x7e, 0x62, 0x4a,
	0x58, 0x66, 0x71, 0x66, 0x52, 0x66, 0x4e, 0x66, 0x49, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f,
	0x10, 0xa6, 0x84, 0x90, 0x29, 0x17, 0x37, 0xdc, 0x70, 0x4f, 0x17, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0x6e, 0x23, 0x61, 0x88, 0x35, 0xc5, 0x7a, 0xce, 0x08, 0xa9, 0x20, 0x64, 0x75, 0x4a, 0x7e, 0x48,
	0x0e, 0x08, 0x80, 0x3a, 0x3a, 0x00, 0x62, 0xb8, 0x90, 0x08, 0x17, 0xab, 0x67, 0x5e, 0x41, 0x69,
	0x09, 0xd4, 0x52, 0x08, 0x47, 0x48, 0x86, 0x8b, 0x33, 0xa4, 0x28, 0x31, 0xaf, 0x38, 0x33, 0x35,
	0xaf, 0x04, 0x6c, 0x0d, 0x4f, 0x10, 0x42, 0x40, 0xc9, 0x99, 0x8b, 0x1f, 0x6e, 0x9e, 0x63, 0x72,
	0x09, 0xc8, 0x1f, 0x12, 0x5c, 0xec, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0xc5, 0x50, 0x83, 0x60,
	0x5c, 0x21, 0x31, 0x2e, 0xb6, 0xd4, 0xb2, 0xd4, 0xbc, 0x92, 0x62, 0xa8, 0x39, 0x50, 0x5e, 0x12,
	0x24, 0xb8, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0x10, 0x2e, 0xce, 0x51, 0x01, 0x00,
	0x00,
}
