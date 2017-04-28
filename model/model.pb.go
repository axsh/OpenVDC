// Code generated by protoc-gen-go.
// source: model.proto
// DO NOT EDIT!

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	model.proto
	cluster.proto

It has these top-level messages:
	Instance
	InstanceState
	FailureMessage
	Template
	NoneTemplate
	LxcTemplate
	NullTemplate
	Console
	ExecutorNode
	SchedulerNode
	NodeState
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InstanceState_State int32

const (
	InstanceState_REGISTERED   InstanceState_State = 0
	InstanceState_QUEUED       InstanceState_State = 1
	InstanceState_STARTING     InstanceState_State = 2
	InstanceState_RUNNING      InstanceState_State = 3
	InstanceState_STOPPING     InstanceState_State = 4
	InstanceState_STOPPED      InstanceState_State = 5
	InstanceState_REBOOTING    InstanceState_State = 6
	InstanceState_SHUTTINGDOWN InstanceState_State = 7
	InstanceState_TERMINATED   InstanceState_State = 8
)

var InstanceState_State_name = map[int32]string{
	0: "REGISTERED",
	1: "QUEUED",
	2: "STARTING",
	3: "RUNNING",
	4: "STOPPING",
	5: "STOPPED",
	6: "REBOOTING",
	7: "SHUTTINGDOWN",
	8: "TERMINATED",
}
var InstanceState_State_value = map[string]int32{
	"REGISTERED":   0,
	"QUEUED":       1,
	"STARTING":     2,
	"RUNNING":      3,
	"STOPPING":     4,
	"STOPPED":      5,
	"REBOOTING":    6,
	"SHUTTINGDOWN": 7,
	"TERMINATED":   8,
}

func (x InstanceState_State) String() string {
	return proto.EnumName(InstanceState_State_name, int32(x))
}
func (InstanceState_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type FailureMessage_ErrorType int32

const (
	FailureMessage_FAILED_BOOT      FailureMessage_ErrorType = 0
	FailureMessage_FAILED_START     FailureMessage_ErrorType = 1
	FailureMessage_FAILED_STOP      FailureMessage_ErrorType = 2
	FailureMessage_FAILED_REBOOT    FailureMessage_ErrorType = 3
	FailureMessage_FAILED_TERMINATE FailureMessage_ErrorType = 4
)

var FailureMessage_ErrorType_name = map[int32]string{
	0: "FAILED_BOOT",
	1: "FAILED_START",
	2: "FAILED_STOP",
	3: "FAILED_REBOOT",
	4: "FAILED_TERMINATE",
}
var FailureMessage_ErrorType_value = map[string]int32{
	"FAILED_BOOT":      0,
	"FAILED_START":     1,
	"FAILED_STOP":      2,
	"FAILED_REBOOT":    3,
	"FAILED_TERMINATE": 4,
}

func (x FailureMessage_ErrorType) String() string {
	return proto.EnumName(FailureMessage_ErrorType_name, int32(x))
}
func (FailureMessage_ErrorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Instance struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	SlaveId string `protobuf:"bytes,2,opt,name=slave_id" json:"slave_id,omitempty"`
	// string resource_id = 3; // Obsolete
	LastState     *InstanceState             `protobuf:"bytes,4,opt,name=last_state" json:"last_state,omitempty"`
	CreatedAt     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=created_at" json:"created_at,omitempty"`
	Template      *Template                  `protobuf:"bytes,6,opt,name=template" json:"template,omitempty"`
	LatestFailure *FailureMessage            `protobuf:"bytes,7,opt,name=latest_failure,json=latestFailure" json:"latest_failure,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instance) GetSlaveId() string {
	if m != nil {
		return m.SlaveId
	}
	return ""
}

func (m *Instance) GetLastState() *InstanceState {
	if m != nil {
		return m.LastState
	}
	return nil
}

func (m *Instance) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Instance) GetTemplate() *Template {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *Instance) GetLatestFailure() *FailureMessage {
	if m != nil {
		return m.LatestFailure
	}
	return nil
}

type InstanceState struct {
	State     InstanceState_State        `protobuf:"varint,1,opt,name=state,enum=model.InstanceState_State" json:"state,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *InstanceState) Reset()                    { *m = InstanceState{} }
func (m *InstanceState) String() string            { return proto.CompactTextString(m) }
func (*InstanceState) ProtoMessage()               {}
func (*InstanceState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InstanceState) GetState() InstanceState_State {
	if m != nil {
		return m.State
	}
	return InstanceState_REGISTERED
}

func (m *InstanceState) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type FailureMessage struct {
	ErrorType FailureMessage_ErrorType   `protobuf:"varint,1,opt,name=error_type,json=errorType,enum=model.FailureMessage_ErrorType" json:"error_type,omitempty"`
	FailedAt  *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=failed_at,json=failedAt" json:"failed_at,omitempty"`
}

func (m *FailureMessage) Reset()                    { *m = FailureMessage{} }
func (m *FailureMessage) String() string            { return proto.CompactTextString(m) }
func (*FailureMessage) ProtoMessage()               {}
func (*FailureMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FailureMessage) GetErrorType() FailureMessage_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return FailureMessage_FAILED_BOOT
}

func (m *FailureMessage) GetFailedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.FailedAt
	}
	return nil
}

type Template struct {
	TemplateUri string `protobuf:"bytes,1,opt,name=template_uri" json:"template_uri,omitempty"`
	// Types that are valid to be assigned to Item:
	//	*Template_None
	//	*Template_Lxc
	//	*Template_Null
	Item      isTemplate_Item            `protobuf_oneof:"Item"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *Template) Reset()                    { *m = Template{} }
func (m *Template) String() string            { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()               {}
func (*Template) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isTemplate_Item interface {
	isTemplate_Item()
}

type Template_None struct {
	None *NoneTemplate `protobuf:"bytes,500,opt,name=none,oneof"`
}
type Template_Lxc struct {
	Lxc *LxcTemplate `protobuf:"bytes,501,opt,name=lxc,oneof"`
}
type Template_Null struct {
	Null *NullTemplate `protobuf:"bytes,502,opt,name=null,oneof"`
}

func (*Template_None) isTemplate_Item() {}
func (*Template_Lxc) isTemplate_Item()  {}
func (*Template_Null) isTemplate_Item() {}

func (m *Template) GetItem() isTemplate_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Template) GetTemplateUri() string {
	if m != nil {
		return m.TemplateUri
	}
	return ""
}

func (m *Template) GetNone() *NoneTemplate {
	if x, ok := m.GetItem().(*Template_None); ok {
		return x.None
	}
	return nil
}

func (m *Template) GetLxc() *LxcTemplate {
	if x, ok := m.GetItem().(*Template_Lxc); ok {
		return x.Lxc
	}
	return nil
}

func (m *Template) GetNull() *NullTemplate {
	if x, ok := m.GetItem().(*Template_Null); ok {
		return x.Null
	}
	return nil
}

func (m *Template) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Template) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Template_OneofMarshaler, _Template_OneofUnmarshaler, _Template_OneofSizer, []interface{}{
		(*Template_None)(nil),
		(*Template_Lxc)(nil),
		(*Template_Null)(nil),
	}
}

func _Template_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Template)
	// Item
	switch x := m.Item.(type) {
	case *Template_None:
		b.EncodeVarint(500<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.None); err != nil {
			return err
		}
	case *Template_Lxc:
		b.EncodeVarint(501<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Lxc); err != nil {
			return err
		}
	case *Template_Null:
		b.EncodeVarint(502<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Null); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Template.Item has unexpected type %T", x)
	}
	return nil
}

func _Template_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Template)
	switch tag {
	case 500: // Item.none
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NoneTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_None{msg}
		return true, err
	case 501: // Item.lxc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LxcTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_Lxc{msg}
		return true, err
	case 502: // Item.null
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NullTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_Null{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Template_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Template)
	// Item
	switch x := m.Item.(type) {
	case *Template_None:
		s := proto.Size(x.None)
		n += proto.SizeVarint(500<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Template_Lxc:
		s := proto.Size(x.Lxc)
		n += proto.SizeVarint(501<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Template_Null:
		s := proto.Size(x.Null)
		n += proto.SizeVarint(502<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NoneTemplate struct {
}

func (m *NoneTemplate) Reset()                    { *m = NoneTemplate{} }
func (m *NoneTemplate) String() string            { return proto.CompactTextString(m) }
func (*NoneTemplate) ProtoMessage()               {}
func (*NoneTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type LxcTemplate struct {
	Vcpu        int32                    `protobuf:"varint,1,opt,name=vcpu" json:"vcpu,omitempty"`
	MemoryGb    int32                    `protobuf:"varint,2,opt,name=memory_gb" json:"memory_gb,omitempty"`
	MinVcpu     int32                    `protobuf:"varint,3,opt,name=min_vcpu" json:"min_vcpu,omitempty"`
	MinMemoryGb int32                    `protobuf:"varint,4,opt,name=min_memory_gb" json:"min_memory_gb,omitempty"`
	LxcImage    *LxcTemplate_Image       `protobuf:"bytes,5,opt,name=lxc_image" json:"lxc_image,omitempty"`
	Interfaces  []*LxcTemplate_Interface `protobuf:"bytes,6,rep,name=interfaces" json:"interfaces,omitempty"`
	LxcTemplate *LxcTemplate_Template    `protobuf:"bytes,7,opt,name=lxc_template" json:"lxc_template,omitempty"`
	NodeGroups  []string                 `protobuf:"bytes,8,rep,name=node_groups" json:"node_groups,omitempty"`
}

func (m *LxcTemplate) Reset()                    { *m = LxcTemplate{} }
func (m *LxcTemplate) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate) ProtoMessage()               {}
func (*LxcTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LxcTemplate) GetVcpu() int32 {
	if m != nil {
		return m.Vcpu
	}
	return 0
}

func (m *LxcTemplate) GetMemoryGb() int32 {
	if m != nil {
		return m.MemoryGb
	}
	return 0
}

func (m *LxcTemplate) GetMinVcpu() int32 {
	if m != nil {
		return m.MinVcpu
	}
	return 0
}

func (m *LxcTemplate) GetMinMemoryGb() int32 {
	if m != nil {
		return m.MinMemoryGb
	}
	return 0
}

func (m *LxcTemplate) GetLxcImage() *LxcTemplate_Image {
	if m != nil {
		return m.LxcImage
	}
	return nil
}

func (m *LxcTemplate) GetInterfaces() []*LxcTemplate_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *LxcTemplate) GetLxcTemplate() *LxcTemplate_Template {
	if m != nil {
		return m.LxcTemplate
	}
	return nil
}

func (m *LxcTemplate) GetNodeGroups() []string {
	if m != nil {
		return m.NodeGroups
	}
	return nil
}

type LxcTemplate_Image struct {
	DownloadUrl string `protobuf:"bytes,1,opt,name=download_url" json:"download_url,omitempty"`
	ChksumType  string `protobuf:"bytes,2,opt,name=chksum_type" json:"chksum_type,omitempty"`
	Chksum      string `protobuf:"bytes,3,opt,name=chksum" json:"chksum,omitempty"`
}

func (m *LxcTemplate_Image) Reset()                    { *m = LxcTemplate_Image{} }
func (m *LxcTemplate_Image) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate_Image) ProtoMessage()               {}
func (*LxcTemplate_Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *LxcTemplate_Image) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *LxcTemplate_Image) GetChksumType() string {
	if m != nil {
		return m.ChksumType
	}
	return ""
}

func (m *LxcTemplate_Image) GetChksum() string {
	if m != nil {
		return m.Chksum
	}
	return ""
}

type LxcTemplate_Interface struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Macaddr  string `protobuf:"bytes,2,opt,name=macaddr" json:"macaddr,omitempty"`
	Ipv4Addr string `protobuf:"bytes,3,opt,name=ipv4addr" json:"ipv4addr,omitempty"`
}

func (m *LxcTemplate_Interface) Reset()                    { *m = LxcTemplate_Interface{} }
func (m *LxcTemplate_Interface) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate_Interface) ProtoMessage()               {}
func (*LxcTemplate_Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 1} }

func (m *LxcTemplate_Interface) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *LxcTemplate_Interface) GetMacaddr() string {
	if m != nil {
		return m.Macaddr
	}
	return ""
}

func (m *LxcTemplate_Interface) GetIpv4Addr() string {
	if m != nil {
		return m.Ipv4Addr
	}
	return ""
}

type LxcTemplate_Template struct {
	// Template specifies the name of the template.
	Template string `protobuf:"bytes,1,opt,name=template" json:"template,omitempty"`
	// Backend specifies the type of the backend.
	Backend int32 `protobuf:"varint,2,opt,name=backend" json:"backend,omitempty"`
	// Distro specifies the name of the distribution.
	Distro string `protobuf:"bytes,3,opt,name=distro" json:"distro,omitempty"`
	// Release specifies the name/version of the distribution.
	Release string `protobuf:"bytes,4,opt,name=release" json:"release,omitempty"`
	// Arch specified the architecture of the container.
	Arch string `protobuf:"bytes,5,opt,name=arch" json:"arch,omitempty"`
	// Variant specifies the variant of the image (default: "default").
	Variant string `protobuf:"bytes,6,opt,name=variant" json:"variant,omitempty"`
	// Image server (default: "images.linuxcontainers.org").
	Server string `protobuf:"bytes,7,opt,name=server" json:"server,omitempty"`
	// GPG keyid (default: 0x...).
	KeyId string `protobuf:"bytes,8,opt,name=key_id" json:"key_id,omitempty"`
	// GPG keyserver to use.
	KeyServer string `protobuf:"bytes,9,opt,name=key_server" json:"key_server,omitempty"`
	// Disable GPG validation (not recommended).
	DisableGpgValidation bool `protobuf:"varint,10,opt,name=disable_gpg_validation" json:"disable_gpg_validation,omitempty"`
	// Flush the local copy (if present).
	FlushCache bool `protobuf:"varint,11,opt,name=flush_cache" json:"flush_cache,omitempty"`
	// Force the use of the local copy even if expired.
	ForceCache bool `protobuf:"varint,12,opt,name=force_cache" json:"force_cache,omitempty"`
	// ExtraArgs provides a way to specify template specific args.
	ExtraArgs []string `protobuf:"bytes,13,rep,name=extra_args" json:"extra_args,omitempty"`
}

func (m *LxcTemplate_Template) Reset()                    { *m = LxcTemplate_Template{} }
func (m *LxcTemplate_Template) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate_Template) ProtoMessage()               {}
func (*LxcTemplate_Template) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 2} }

func (m *LxcTemplate_Template) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *LxcTemplate_Template) GetBackend() int32 {
	if m != nil {
		return m.Backend
	}
	return 0
}

func (m *LxcTemplate_Template) GetDistro() string {
	if m != nil {
		return m.Distro
	}
	return ""
}

func (m *LxcTemplate_Template) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *LxcTemplate_Template) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

func (m *LxcTemplate_Template) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

func (m *LxcTemplate_Template) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *LxcTemplate_Template) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *LxcTemplate_Template) GetKeyServer() string {
	if m != nil {
		return m.KeyServer
	}
	return ""
}

func (m *LxcTemplate_Template) GetDisableGpgValidation() bool {
	if m != nil {
		return m.DisableGpgValidation
	}
	return false
}

func (m *LxcTemplate_Template) GetFlushCache() bool {
	if m != nil {
		return m.FlushCache
	}
	return false
}

func (m *LxcTemplate_Template) GetForceCache() bool {
	if m != nil {
		return m.ForceCache
	}
	return false
}

func (m *LxcTemplate_Template) GetExtraArgs() []string {
	if m != nil {
		return m.ExtraArgs
	}
	return nil
}

type NullTemplate struct {
	Vcpu       int32    `protobuf:"varint,1,opt,name=vcpu" json:"vcpu,omitempty"`
	MemoryGb   int32    `protobuf:"varint,2,opt,name=memory_gb" json:"memory_gb,omitempty"`
	NodeGroups []string `protobuf:"bytes,8,rep,name=node_groups" json:"node_groups,omitempty"`
}

func (m *NullTemplate) Reset()                    { *m = NullTemplate{} }
func (m *NullTemplate) String() string            { return proto.CompactTextString(m) }
func (*NullTemplate) ProtoMessage()               {}
func (*NullTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NullTemplate) GetVcpu() int32 {
	if m != nil {
		return m.Vcpu
	}
	return 0
}

func (m *NullTemplate) GetMemoryGb() int32 {
	if m != nil {
		return m.MemoryGb
	}
	return 0
}

func (m *NullTemplate) GetNodeGroups() []string {
	if m != nil {
		return m.NodeGroups
	}
	return nil
}

func init() {
	proto.RegisterType((*Instance)(nil), "model.Instance")
	proto.RegisterType((*InstanceState)(nil), "model.InstanceState")
	proto.RegisterType((*FailureMessage)(nil), "model.FailureMessage")
	proto.RegisterType((*Template)(nil), "model.Template")
	proto.RegisterType((*NoneTemplate)(nil), "model.NoneTemplate")
	proto.RegisterType((*LxcTemplate)(nil), "model.LxcTemplate")
	proto.RegisterType((*LxcTemplate_Image)(nil), "model.LxcTemplate.Image")
	proto.RegisterType((*LxcTemplate_Interface)(nil), "model.LxcTemplate.Interface")
	proto.RegisterType((*LxcTemplate_Template)(nil), "model.LxcTemplate.Template")
	proto.RegisterType((*NullTemplate)(nil), "model.NullTemplate")
	proto.RegisterEnum("model.InstanceState_State", InstanceState_State_name, InstanceState_State_value)
	proto.RegisterEnum("model.FailureMessage_ErrorType", FailureMessage_ErrorType_name, FailureMessage_ErrorType_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1002 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5b, 0x8f, 0xdb, 0x44,
	0x14, 0xde, 0xdc, 0xed, 0x93, 0x4b, 0xcd, 0x50, 0x2a, 0x2b, 0x14, 0x1a, 0x45, 0x48, 0xac, 0x40,
	0x4a, 0xd0, 0x52, 0x15, 0x09, 0x2a, 0xd0, 0xae, 0xe2, 0xb6, 0x91, 0xda, 0x6c, 0x99, 0xcd, 0x0a,
	0x89, 0x17, 0x6b, 0x62, 0xcf, 0x3a, 0xd6, 0x8e, 0x2f, 0x1a, 0x8f, 0xc3, 0xee, 0x5f, 0xa8, 0x78,
	0xe3, 0x47, 0xf0, 0x03, 0xb9, 0x3c, 0xa3, 0x99, 0xb1, 0x1d, 0x07, 0x16, 0x71, 0xe9, 0x53, 0xe6,
	0xfb, 0xe6, 0x3b, 0x67, 0xce, 0x39, 0x73, 0xe6, 0x38, 0xd0, 0x8f, 0x12, 0x9f, 0xb2, 0x59, 0xca,
	0x13, 0x91, 0xa0, 0x8e, 0x02, 0xe3, 0xaf, 0x82, 0x50, 0x6c, 0xf3, 0xcd, 0xcc, 0x4b, 0xa2, 0x79,
	0x90, 0x30, 0x12, 0x07, 0x73, 0xb5, 0xbf, 0xc9, 0xaf, 0xe6, 0xa9, 0xb8, 0x4d, 0x69, 0x36, 0x17,
	0x61, 0x44, 0x33, 0x41, 0xa2, 0x74, 0xbf, 0xd2, 0x3e, 0xa6, 0x3f, 0x35, 0xc1, 0x58, 0xc6, 0x99,
	0x20, 0xb1, 0x47, 0xd1, 0x08, 0x9a, 0xa1, 0x6f, 0x37, 0x26, 0x8d, 0x63, 0x13, 0x37, 0x43, 0x1f,
	0x8d, 0xc1, 0xc8, 0x18, 0xd9, 0x51, 0x37, 0xf4, 0xed, 0xa6, 0x62, 0x2b, 0x8c, 0x1e, 0x03, 0x30,
	0x92, 0x09, 0x37, 0x13, 0x44, 0x50, 0xbb, 0x3d, 0x69, 0x1c, 0xf7, 0x4f, 0xee, 0xcf, 0x74, 0x78,
	0xa5, 0xc3, 0x0b, 0xb9, 0x87, 0x6b, 0x3a, 0xf4, 0x25, 0x80, 0xc7, 0x29, 0x11, 0xd4, 0x77, 0x89,
	0xb0, 0x3b, 0xca, 0x6a, 0x3c, 0x0b, 0x92, 0x24, 0x60, 0x74, 0x56, 0x46, 0x3d, 0x5b, 0x97, 0x41,
	0xe2, 0x9a, 0x1a, 0x7d, 0x0a, 0x86, 0xa0, 0x51, 0xca, 0xe4, 0x79, 0x5d, 0x65, 0x79, 0xaf, 0x38,
	0x6f, 0x5d, 0xd0, 0xb8, 0x12, 0xa0, 0xa7, 0x30, 0x92, 0xbf, 0x99, 0x70, 0xaf, 0x48, 0xc8, 0x72,
	0x4e, 0xed, 0x9e, 0x32, 0x79, 0xaf, 0x30, 0x79, 0xa6, 0xd9, 0x57, 0x34, 0xcb, 0x48, 0x40, 0xf1,
	0x50, 0x8b, 0x0b, 0x76, 0xfa, 0xa6, 0x09, 0xc3, 0x83, 0x24, 0xd0, 0x67, 0xd0, 0xd1, 0x99, 0xca,
	0xea, 0x8c, 0x4e, 0xc6, 0x77, 0x65, 0x3a, 0xd3, 0xf9, 0x76, 0xee, 0x4a, 0xb5, 0xf9, 0x5f, 0x52,
	0x9d, 0xbe, 0x69, 0x40, 0x47, 0x9f, 0x3b, 0x02, 0xc0, 0xce, 0xf3, 0xe5, 0xc5, 0xda, 0xc1, 0xce,
	0xc2, 0x3a, 0x42, 0x00, 0xdd, 0x6f, 0x2f, 0x9d, 0x4b, 0x67, 0x61, 0x35, 0xd0, 0x00, 0x8c, 0x8b,
	0xf5, 0x29, 0x5e, 0x2f, 0x57, 0xcf, 0xad, 0x26, 0xea, 0x43, 0x0f, 0x5f, 0xae, 0x56, 0x12, 0xb4,
	0xf4, 0xd6, 0xf9, 0xeb, 0xd7, 0x12, 0xb5, 0xe5, 0x96, 0x42, 0xce, 0xc2, 0xea, 0xa0, 0x21, 0x98,
	0xd8, 0x39, 0x3b, 0x3f, 0x57, 0x66, 0x5d, 0x64, 0xc1, 0xe0, 0xe2, 0xc5, 0xe5, 0x5a, 0xa2, 0xc5,
	0xf9, 0x77, 0x2b, 0xab, 0x27, 0x8f, 0x5c, 0x3b, 0xf8, 0xd5, 0x72, 0x75, 0xba, 0x76, 0x16, 0x96,
	0x31, 0xfd, 0xa5, 0x01, 0xa3, 0xc3, 0x72, 0xa1, 0xaf, 0x01, 0x28, 0xe7, 0x09, 0x77, 0x65, 0x7b,
	0x15, 0x25, 0x79, 0x74, 0x67, 0x65, 0x67, 0x8e, 0xd4, 0xad, 0x6f, 0x53, 0x8a, 0x4d, 0x5a, 0x2e,
	0xd1, 0x17, 0x60, 0xca, 0x6b, 0xf9, 0xb7, 0xa5, 0x31, 0xb4, 0xf8, 0x54, 0x4c, 0xb7, 0x60, 0x56,
	0x0e, 0xd1, 0x3d, 0xe8, 0x3f, 0x3b, 0x5d, 0xbe, 0x74, 0x16, 0xae, 0x4c, 0xc7, 0x3a, 0x92, 0xb9,
	0x14, 0x84, 0xaa, 0x8b, 0xd5, 0xa8, 0x49, 0x64, 0x01, 0xac, 0x26, 0x7a, 0x07, 0x86, 0x05, 0xa1,
	0x8b, 0x60, 0xb5, 0xd0, 0x7d, 0xb0, 0x0a, 0xaa, 0x4a, 0xdb, 0x6a, 0xcb, 0xac, 0x8d, 0xb2, 0xaf,
	0xd0, 0x14, 0x06, 0x65, 0x67, 0xb9, 0x39, 0x0f, 0x8b, 0x27, 0x72, 0xc0, 0xa1, 0x4f, 0xa0, 0x1d,
	0x27, 0x31, 0xb5, 0x7f, 0x6d, 0xa9, 0x7c, 0xde, 0x2d, 0xca, 0xb1, 0x4a, 0x62, 0x5a, 0xfa, 0x79,
	0x71, 0x84, 0x95, 0x06, 0x7d, 0x0c, 0x2d, 0x76, 0xe3, 0xd9, 0xbf, 0x69, 0x29, 0x2a, 0xa4, 0x2f,
	0x6f, 0xbc, 0x9a, 0x52, 0x2a, 0x94, 0xd3, 0x9c, 0x31, 0xfb, 0xf7, 0x3f, 0x39, 0xcd, 0x19, 0x3b,
	0x70, 0x9a, 0x33, 0xf6, 0x36, 0x0d, 0x77, 0xd6, 0x85, 0xf6, 0x52, 0xd0, 0x68, 0x3a, 0x82, 0x41,
	0x3d, 0xe0, 0xe9, 0xcf, 0x3d, 0xe8, 0xd7, 0xc2, 0x42, 0x08, 0xda, 0x3b, 0x2f, 0xcd, 0x55, 0x01,
	0x3a, 0x58, 0xad, 0xd1, 0x43, 0x30, 0x23, 0x1a, 0x25, 0xfc, 0xd6, 0x0d, 0x36, 0xea, 0xd8, 0x0e,
	0xde, 0x13, 0x72, 0x86, 0x44, 0x61, 0xec, 0x2a, 0xab, 0x96, 0xda, 0xac, 0x30, 0xfa, 0x08, 0x86,
	0x72, 0xbd, 0xb7, 0x6e, 0x2b, 0xc1, 0x21, 0x89, 0x9e, 0x80, 0xc9, 0x6e, 0x3c, 0x37, 0x8c, 0x48,
	0x40, 0x8b, 0x91, 0x61, 0xff, 0xb5, 0x62, 0xb3, 0xa5, 0xdc, 0xc7, 0x7b, 0x29, 0x7a, 0x0a, 0x10,
	0xc6, 0x82, 0xf2, 0x2b, 0xe2, 0xd1, 0xcc, 0xee, 0x4e, 0x5a, 0xc7, 0xfd, 0x93, 0x87, 0x77, 0x19,
	0x96, 0x22, 0x5c, 0xd3, 0xa3, 0x6f, 0x60, 0x20, 0x5d, 0x55, 0x13, 0x47, 0x8f, 0x8f, 0xf7, 0xef,
	0xb0, 0xaf, 0xa6, 0xcf, 0x81, 0x01, 0x9a, 0x40, 0x3f, 0x4e, 0x7c, 0xea, 0x06, 0x3c, 0xc9, 0xd3,
	0xcc, 0x36, 0x26, 0xad, 0x63, 0x13, 0xd7, 0xa9, 0x31, 0x85, 0x8e, 0x0a, 0x5a, 0xb6, 0x97, 0x9f,
	0xfc, 0x10, 0xb3, 0x84, 0xf8, 0x6e, 0xce, 0x59, 0xd9, 0x5e, 0x75, 0x4e, 0xba, 0xf3, 0xb6, 0xd7,
	0x59, 0x1e, 0xe9, 0x37, 0xa7, 0xc7, 0x71, 0x9d, 0x42, 0x0f, 0xa0, 0xab, 0xa1, 0xaa, 0xb3, 0x89,
	0x0b, 0x34, 0xbe, 0x04, 0xb3, 0x4a, 0x51, 0x5e, 0x60, 0xf5, 0x66, 0x4d, 0xac, 0xd6, 0xc8, 0x86,
	0x5e, 0x44, 0x3c, 0xe2, 0xfb, 0xbc, 0x70, 0x5b, 0x42, 0x79, 0x79, 0x61, 0xba, 0x7b, 0xac, 0xb6,
	0xb4, 0xd3, 0x0a, 0x8f, 0x7f, 0x6c, 0xd5, 0x1e, 0xc8, 0xb8, 0x36, 0x9b, 0xb5, 0xeb, 0xfd, 0x28,
	0xb6, 0xa1, 0xb7, 0x21, 0xde, 0x35, 0x8d, 0xfd, 0xa2, 0x3b, 0x4a, 0x28, 0x23, 0xf6, 0xc3, 0x4c,
	0xf0, 0xa4, 0x8c, 0x58, 0x23, 0x69, 0xc1, 0x29, 0xa3, 0x24, 0xd3, 0x1f, 0x16, 0x13, 0x97, 0x50,
	0x86, 0x4f, 0xb8, 0xb7, 0x55, 0x6d, 0x60, 0x62, 0xb5, 0x96, 0xea, 0x1d, 0xe1, 0x21, 0x89, 0x85,
	0xfa, 0x2c, 0x98, 0xb8, 0x84, 0xd2, 0x7f, 0x46, 0xf9, 0x8e, 0x72, 0x75, 0x7b, 0x26, 0x2e, 0x90,
	0xe4, 0xaf, 0xe9, 0xad, 0xfc, 0xaa, 0x19, 0x9a, 0xd7, 0x08, 0x7d, 0x08, 0x20, 0x57, 0x85, 0x8d,
	0xa9, 0xf6, 0x6a, 0x0c, 0x7a, 0x02, 0x0f, 0xfc, 0x30, 0x23, 0x1b, 0x46, 0xdd, 0x20, 0x0d, 0xdc,
	0x1d, 0x61, 0xa1, 0x4f, 0x44, 0x98, 0xc4, 0x36, 0x4c, 0x1a, 0xc7, 0x06, 0xfe, 0x9b, 0x5d, 0x79,
	0x77, 0x57, 0x2c, 0xcf, 0xb6, 0xae, 0x47, 0xbc, 0x2d, 0xb5, 0xfb, 0x4a, 0x5c, 0xa7, 0x94, 0x22,
	0xe1, 0x1e, 0x2d, 0x14, 0x83, 0x42, 0xb1, 0xa7, 0x64, 0x6c, 0xf4, 0x46, 0x70, 0xe2, 0x12, 0x1e,
	0x64, 0xf6, 0x50, 0x75, 0x53, 0x8d, 0x99, 0x6e, 0x60, 0x50, 0x9f, 0x0a, 0xff, 0xe3, 0xa5, 0xfe,
	0x63, 0xc3, 0x9e, 0x3d, 0xfa, 0xfe, 0x83, 0xda, 0x7f, 0x0d, 0x72, 0x93, 0x6d, 0xe7, 0x49, 0x4a,
	0xe3, 0x9d, 0xef, 0xcd, 0xd5, 0xb3, 0xd8, 0x74, 0xd5, 0x98, 0xf9, 0xfc, 0x8f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2d, 0xf9, 0x38, 0x4e, 0xa7, 0x08, 0x00, 0x00,
}
