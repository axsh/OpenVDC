// Code generated by protoc-gen-go.
// source: cluster.proto
// DO NOT EDIT!

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Console_Transport int32

const (
	Console_SSH Console_Transport = 0
)

var Console_Transport_name = map[int32]string{
	0: "SSH",
}
var Console_Transport_value = map[string]int32{
	"SSH": 0,
}

func (x Console_Transport) String() string {
	return proto.EnumName(Console_Transport_name, int32(x))
}
func (Console_Transport) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type NodeState_State int32

const (
	NodeState_REGISTERED NodeState_State = 0
)

var NodeState_State_name = map[int32]string{
	0: "REGISTERED",
}
var NodeState_State_value = map[string]int32{
	"REGISTERED": 0,
}

func (x NodeState_State) String() string {
	return proto.EnumName(NodeState_State_name, int32(x))
}
func (NodeState_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{5, 0} }

type Console struct {
	Type     Console_Transport `protobuf:"varint,1,opt,name=type,enum=model.Console_Transport" json:"type,omitempty"`
	BindAddr string            `protobuf:"bytes,2,opt,name=bind_addr" json:"bind_addr,omitempty"`
}

func (m *Console) Reset()                    { *m = Console{} }
func (m *Console) String() string            { return proto.CompactTextString(m) }
func (*Console) ProtoMessage()               {}
func (*Console) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Console) GetType() Console_Transport {
	if m != nil {
		return m.Type
	}
	return Console_SSH
}

func (m *Console) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

type ExecutorNode struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
	Console   *Console                   `protobuf:"bytes,3,opt,name=console" json:"console,omitempty"`
	GrpcAddr  string                     `protobuf:"bytes,4,opt,name=grpc_addr" json:"grpc_addr,omitempty"`
	LastState *NodeState                 `protobuf:"bytes,5,opt,name=last_state" json:"last_state,omitempty"`
}

func (m *ExecutorNode) Reset()                    { *m = ExecutorNode{} }
func (m *ExecutorNode) String() string            { return proto.CompactTextString(m) }
func (*ExecutorNode) ProtoMessage()               {}
func (*ExecutorNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ExecutorNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExecutorNode) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ExecutorNode) GetConsole() *Console {
	if m != nil {
		return m.Console
	}
	return nil
}

func (m *ExecutorNode) GetGrpcAddr() string {
	if m != nil {
		return m.GrpcAddr
	}
	return ""
}

func (m *ExecutorNode) GetLastState() *NodeState {
	if m != nil {
		return m.LastState
	}
	return nil
}

type SchedulerNode struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *SchedulerNode) Reset()                    { *m = SchedulerNode{} }
func (m *SchedulerNode) String() string            { return proto.CompactTextString(m) }
func (*SchedulerNode) ProtoMessage()               {}
func (*SchedulerNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SchedulerNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SchedulerNode) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type AgentNode struct {
	AgentMesosId string                     `protobuf:"bytes,1,opt,name=agent_mesos_id,json=agentMesosId" json:"agent_mesos_id,omitempty"`
	AgentId      string                     `protobuf:"bytes,2,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	CreatedAt    *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *AgentNode) Reset()                    { *m = AgentNode{} }
func (m *AgentNode) String() string            { return proto.CompactTextString(m) }
func (*AgentNode) ProtoMessage()               {}
func (*AgentNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *AgentNode) GetAgentMesosId() string {
	if m != nil {
		return m.AgentMesosId
	}
	return ""
}

func (m *AgentNode) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *AgentNode) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type CrashedNode struct {
	Uuid          string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	AgentId       string                     `protobuf:"bytes,2,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	AgentMesosId  string                     `protobuf:"bytes,3,opt,name=agent_mesos_id,json=agentMesosId" json:"agent_mesos_id,omitempty"`
	Reconnected   bool                       `protobuf:"varint,4,opt,name=reconnected" json:"reconnected,omitempty"`
	CreatedAt     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	ReconnectedAt *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=reconnected_at,json=reconnectedAt" json:"reconnected_at,omitempty"`
}

func (m *CrashedNode) Reset()                    { *m = CrashedNode{} }
func (m *CrashedNode) String() string            { return proto.CompactTextString(m) }
func (*CrashedNode) ProtoMessage()               {}
func (*CrashedNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *CrashedNode) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CrashedNode) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *CrashedNode) GetAgentMesosId() string {
	if m != nil {
		return m.AgentMesosId
	}
	return ""
}

func (m *CrashedNode) GetReconnected() bool {
	if m != nil {
		return m.Reconnected
	}
	return false
}

func (m *CrashedNode) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *CrashedNode) GetReconnectedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.ReconnectedAt
	}
	return nil
}

type NodeState struct {
	State     NodeState_State            `protobuf:"varint,1,opt,name=state,enum=model.NodeState_State" json:"state,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *NodeState) Reset()                    { *m = NodeState{} }
func (m *NodeState) String() string            { return proto.CompactTextString(m) }
func (*NodeState) ProtoMessage()               {}
func (*NodeState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *NodeState) GetState() NodeState_State {
	if m != nil {
		return m.State
	}
	return NodeState_REGISTERED
}

func (m *NodeState) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Console)(nil), "model.Console")
	proto.RegisterType((*ExecutorNode)(nil), "model.ExecutorNode")
	proto.RegisterType((*SchedulerNode)(nil), "model.SchedulerNode")
	proto.RegisterType((*AgentNode)(nil), "model.AgentNode")
	proto.RegisterType((*CrashedNode)(nil), "model.CrashedNode")
	proto.RegisterType((*NodeState)(nil), "model.NodeState")
	proto.RegisterEnum("model.Console_Transport", Console_Transport_name, Console_Transport_value)
	proto.RegisterEnum("model.NodeState_State", NodeState_State_name, NodeState_State_value)
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x76, 0x5d, 0x97, 0xd7, 0x35, 0xaa, 0x2c, 0x04, 0x65, 0x02, 0x51, 0x45, 0x1c,
	0x7a, 0x98, 0x12, 0x34, 0x4e, 0xc0, 0xa9, 0x8c, 0x0a, 0x7a, 0x80, 0x43, 0xd2, 0x13, 0x1c, 0x22,
	0xd7, 0x7e, 0xa4, 0x11, 0x49, 0x1c, 0xd9, 0x0e, 0x1a, 0x57, 0x2e, 0xdc, 0xf8, 0x74, 0x7c, 0x20,
	0x14, 0xbb, 0xcd, 0xb2, 0x81, 0x98, 0x10, 0xdc, 0xec, 0xf7, 0xfe, 0xef, 0xff, 0x7e, 0x4f, 0x7e,
	0x86, 0x31, 0xcb, 0x6b, 0xa5, 0x51, 0x06, 0x95, 0x14, 0x5a, 0x90, 0x41, 0x21, 0x38, 0xe6, 0xa7,
	0x2f, 0xd2, 0x4c, 0x6f, 0xeb, 0x4d, 0xc0, 0x44, 0x11, 0xa6, 0x22, 0xa7, 0x65, 0x1a, 0x9a, 0xfc,
	0xa6, 0xfe, 0x18, 0x56, 0xfa, 0x4b, 0x85, 0x2a, 0xd4, 0x59, 0x81, 0x4a, 0xd3, 0xa2, 0xba, 0x3a,
	0x59, 0x0f, 0xff, 0x13, 0x0c, 0x2f, 0x44, 0xa9, 0x44, 0x8e, 0xe4, 0x0c, 0x0e, 0x1b, 0xf5, 0xd4,
	0x99, 0x39, 0x73, 0xef, 0x7c, 0x1a, 0x18, 0xf7, 0x60, 0x97, 0x0d, 0xd6, 0x92, 0x96, 0xaa, 0x12,
	0x52, 0x47, 0x46, 0x45, 0x1e, 0x80, 0xbb, 0xc9, 0x4a, 0x9e, 0x50, 0xce, 0xe5, 0xb4, 0x37, 0x73,
	0xe6, 0x6e, 0x74, 0x15, 0xf0, 0xef, 0x80, 0xdb, 0x16, 0x90, 0x21, 0xf4, 0xe3, 0xf8, 0xcd, 0xe4,
	0xc0, 0xff, 0xe1, 0xc0, 0xc9, 0xf2, 0x12, 0x59, 0xad, 0x85, 0x7c, 0x27, 0x38, 0x12, 0x0f, 0x7a,
	0x19, 0x37, 0x0d, 0xdd, 0xa8, 0x97, 0x71, 0xf2, 0x1c, 0x80, 0x49, 0xa4, 0x1a, 0x79, 0x42, 0xb5,
	0x71, 0x1d, 0x9d, 0x9f, 0x06, 0xa9, 0x10, 0x69, 0x8e, 0xc1, 0x7e, 0xa8, 0x60, 0xbd, 0x9f, 0x21,
	0xea, 0xa8, 0xc9, 0x1c, 0x86, 0xcc, 0xb2, 0x4e, 0xfb, 0xa6, 0xd0, 0xbb, 0x3e, 0x41, 0xb4, 0x4f,
	0x37, 0xe8, 0xa9, 0xac, 0x98, 0x45, 0x3f, 0xb4, 0xe8, 0x6d, 0x80, 0x3c, 0x01, 0xc8, 0xa9, 0xd2,
	0x89, 0xd2, 0x54, 0xe3, 0x74, 0x60, 0xac, 0x26, 0x3b, 0xab, 0x06, 0x3a, 0x6e, 0xe2, 0x51, 0x47,
	0xe3, 0x7f, 0x80, 0x71, 0xcc, 0xb6, 0xc8, 0xeb, 0x1c, 0xff, 0xfb, 0x58, 0xfe, 0x37, 0x07, 0xdc,
	0x45, 0x8a, 0xa5, 0x36, 0xce, 0x8f, 0xc1, 0xa3, 0xcd, 0x25, 0x29, 0x50, 0x09, 0x95, 0xb4, 0x5d,
	0x4e, 0x4c, 0xf4, 0x6d, 0x13, 0x5c, 0x71, 0x72, 0x1f, 0x8e, 0xad, 0x2a, 0xe3, 0xbb, 0xa7, 0x19,
	0x9a, 0xfb, 0x8a, 0x93, 0x67, 0xd7, 0x50, 0xfa, 0xb7, 0xa2, 0xb8, 0x3b, 0xf5, 0x42, 0xfb, 0x5f,
	0x7b, 0x30, 0xba, 0x90, 0x54, 0x6d, 0x91, 0x1b, 0x16, 0x02, 0x87, 0x75, 0xdd, 0x12, 0x98, 0xf3,
	0x9f, 0x3a, 0xff, 0x8a, 0xde, 0xff, 0x0d, 0xfa, 0x0c, 0x46, 0x12, 0x99, 0x28, 0x4b, 0x64, 0x1a,
	0xb9, 0x79, 0x9d, 0xe3, 0xa8, 0x1b, 0xba, 0x31, 0xc1, 0xe0, 0x2f, 0x26, 0x20, 0x0b, 0xf0, 0x3a,
	0x4e, 0x4d, 0xf9, 0xd1, 0xad, 0xe5, 0xe3, 0x4e, 0xc5, 0x42, 0xfb, 0xdf, 0x1d, 0x70, 0xdb, 0x2d,
	0x20, 0x67, 0x30, 0xb0, 0x6b, 0x62, 0xff, 0xcc, 0xdd, 0x9b, 0x6b, 0x12, 0xd8, 0x65, 0xb1, 0xa2,
	0x7f, 0x5a, 0x83, 0x7b, 0x30, 0xb0, 0x2d, 0x3d, 0x80, 0x68, 0xf9, 0x7a, 0x15, 0xaf, 0x97, 0xd1,
	0xf2, 0xd5, 0xe4, 0xe0, 0xe5, 0xa3, 0xf7, 0x0f, 0x3b, 0xff, 0x9f, 0x5e, 0xaa, 0x6d, 0x28, 0x2a,
	0x2c, 0x3f, 0x73, 0x16, 0x1a, 0x98, 0xcd, 0x91, 0x71, 0x7e, 0xfa, 0x33, 0x00, 0x00, 0xff, 0xff,
	0x38, 0x6c, 0x9a, 0x4e, 0x3d, 0x04, 0x00, 0x00,
}
