// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/config.proto

package config // import "github.com/linkerd/linkerd2/controller/gen/config"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GlobalConfig struct {
	LinkerdNamespace string `protobuf:"bytes,1,opt,name=linkerd_namespace,json=linkerdNamespace,proto3" json:"linkerd_namespace,omitempty"`
	CniEnabled       bool   `protobuf:"varint,2,opt,name=cni_enabled,json=cniEnabled,proto3" json:"cni_enabled,omitempty"`
	Registry         string `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
	// null indicates TLS is disabled.
	// Otherwise, a non-null struct indicates the equivalence
	// of --tls=optional.
	IdentityContext      *IdentityContext `protobuf:"bytes,4,opt,name=identity_context,json=identityContext,proto3" json:"identity_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GlobalConfig) Reset()         { *m = GlobalConfig{} }
func (m *GlobalConfig) String() string { return proto.CompactTextString(m) }
func (*GlobalConfig) ProtoMessage()    {}
func (*GlobalConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f7da3f9300481bd1, []int{0}
}
func (m *GlobalConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalConfig.Unmarshal(m, b)
}
func (m *GlobalConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalConfig.Marshal(b, m, deterministic)
}
func (dst *GlobalConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalConfig.Merge(dst, src)
}
func (m *GlobalConfig) XXX_Size() int {
	return xxx_messageInfo_GlobalConfig.Size(m)
}
func (m *GlobalConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalConfig proto.InternalMessageInfo

func (m *GlobalConfig) GetLinkerdNamespace() string {
	if m != nil {
		return m.LinkerdNamespace
	}
	return ""
}

func (m *GlobalConfig) GetCniEnabled() bool {
	if m != nil {
		return m.CniEnabled
	}
	return false
}

func (m *GlobalConfig) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *GlobalConfig) GetIdentityContext() *IdentityContext {
	if m != nil {
		return m.IdentityContext
	}
	return nil
}

type ProxyConfig struct {
	ProxyImage              *Image                `protobuf:"bytes,1,opt,name=proxy_image,json=proxyImage,proto3" json:"proxy_image,omitempty"`
	ProxyInitImage          *Image                `protobuf:"bytes,2,opt,name=proxy_init_image,json=proxyInitImage,proto3" json:"proxy_init_image,omitempty"`
	DestinationApiPort      *Port                 `protobuf:"bytes,3,opt,name=destination_api_port,json=destinationApiPort,proto3" json:"destination_api_port,omitempty"`
	ControlPort             *Port                 `protobuf:"bytes,4,opt,name=control_port,json=controlPort,proto3" json:"control_port,omitempty"`
	IgnoreInboundPorts      []*Port               `protobuf:"bytes,5,rep,name=ignore_inbound_ports,json=ignoreInboundPorts,proto3" json:"ignore_inbound_ports,omitempty"`
	IgnoreOutboundPorts     []*Port               `protobuf:"bytes,6,rep,name=ignore_outbound_ports,json=ignoreOutboundPorts,proto3" json:"ignore_outbound_ports,omitempty"`
	InboundPort             *Port                 `protobuf:"bytes,7,opt,name=inbound_port,json=inboundPort,proto3" json:"inbound_port,omitempty"`
	MetricsPort             *Port                 `protobuf:"bytes,8,opt,name=metrics_port,json=metricsPort,proto3" json:"metrics_port,omitempty"`
	OutboundPort            *Port                 `protobuf:"bytes,9,opt,name=outbound_port,json=outboundPort,proto3" json:"outbound_port,omitempty"`
	Resource                *ResourceRequirements `protobuf:"bytes,10,opt,name=resource,proto3" json:"resource,omitempty"`
	ProxyUid                int64                 `protobuf:"varint,11,opt,name=proxy_uid,json=proxyUid,proto3" json:"proxy_uid,omitempty"`
	LogLevel                *LogLevel             `protobuf:"bytes,12,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	DisableExternalProfiles bool                  `protobuf:"varint,13,opt,name=disable_external_profiles,json=disableExternalProfiles,proto3" json:"disable_external_profiles,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *ProxyConfig) Reset()         { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()    {}
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f7da3f9300481bd1, []int{1}
}
func (m *ProxyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyConfig.Unmarshal(m, b)
}
func (m *ProxyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyConfig.Marshal(b, m, deterministic)
}
func (dst *ProxyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyConfig.Merge(dst, src)
}
func (m *ProxyConfig) XXX_Size() int {
	return xxx_messageInfo_ProxyConfig.Size(m)
}
func (m *ProxyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyConfig proto.InternalMessageInfo

func (m *ProxyConfig) GetProxyImage() *Image {
	if m != nil {
		return m.ProxyImage
	}
	return nil
}

func (m *ProxyConfig) GetProxyInitImage() *Image {
	if m != nil {
		return m.ProxyInitImage
	}
	return nil
}

func (m *ProxyConfig) GetDestinationApiPort() *Port {
	if m != nil {
		return m.DestinationApiPort
	}
	return nil
}

func (m *ProxyConfig) GetControlPort() *Port {
	if m != nil {
		return m.ControlPort
	}
	return nil
}

func (m *ProxyConfig) GetIgnoreInboundPorts() []*Port {
	if m != nil {
		return m.IgnoreInboundPorts
	}
	return nil
}

func (m *ProxyConfig) GetIgnoreOutboundPorts() []*Port {
	if m != nil {
		return m.IgnoreOutboundPorts
	}
	return nil
}

func (m *ProxyConfig) GetInboundPort() *Port {
	if m != nil {
		return m.InboundPort
	}
	return nil
}

func (m *ProxyConfig) GetMetricsPort() *Port {
	if m != nil {
		return m.MetricsPort
	}
	return nil
}

func (m *ProxyConfig) GetOutboundPort() *Port {
	if m != nil {
		return m.OutboundPort
	}
	return nil
}

func (m *ProxyConfig) GetResource() *ResourceRequirements {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ProxyConfig) GetProxyUid() int64 {
	if m != nil {
		return m.ProxyUid
	}
	return 0
}

func (m *ProxyConfig) GetLogLevel() *LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return nil
}

func (m *ProxyConfig) GetDisableExternalProfiles() bool {
	if m != nil {
		return m.DisableExternalProfiles
	}
	return false
}

type Image struct {
	ImageName            string   `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	PullPolicy           string   `protobuf:"bytes,2,opt,name=pull_policy,json=pullPolicy,proto3" json:"pull_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f7da3f9300481bd1, []int{2}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *Image) GetPullPolicy() string {
	if m != nil {
		return m.PullPolicy
	}
	return ""
}

type Port struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f7da3f9300481bd1, []int{3}
}
func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (dst *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(dst, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ResourceRequirements struct {
	RequestCpu           string   `protobuf:"bytes,1,opt,name=request_cpu,json=requestCpu,proto3" json:"request_cpu,omitempty"`
	RequestMemory        string   `protobuf:"bytes,2,opt,name=request_memory,json=requestMemory,proto3" json:"request_memory,omitempty"`
	LimitCpu             string   `protobuf:"bytes,3,opt,name=limit_cpu,json=limitCpu,proto3" json:"limit_cpu,omitempty"`
	LimitMemory          string   `protobuf:"bytes,4,opt,name=limit_memory,json=limitMemory,proto3" json:"limit_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceRequirements) Reset()         { *m = ResourceRequirements{} }
func (m *ResourceRequirements) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirements) ProtoMessage()    {}
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f7da3f9300481bd1, []int{4}
}
func (m *ResourceRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirements.Unmarshal(m, b)
}
func (m *ResourceRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirements.Marshal(b, m, deterministic)
}
func (dst *ResourceRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirements.Merge(dst, src)
}
func (m *ResourceRequirements) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirements.Size(m)
}
func (m *ResourceRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirements proto.InternalMessageInfo

func (m *ResourceRequirements) GetRequestCpu() string {
	if m != nil {
		return m.RequestCpu
	}
	return ""
}

func (m *ResourceRequirements) GetRequestMemory() string {
	if m != nil {
		return m.RequestMemory
	}
	return ""
}

func (m *ResourceRequirements) GetLimitCpu() string {
	if m != nil {
		return m.LimitCpu
	}
	return ""
}

func (m *ResourceRequirements) GetLimitMemory() string {
	if m != nil {
		return m.LimitMemory
	}
	return ""
}

type IdentityContext struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityContext) Reset()         { *m = IdentityContext{} }
func (m *IdentityContext) String() string { return proto.CompactTextString(m) }
func (*IdentityContext) ProtoMessage()    {}
func (*IdentityContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f7da3f9300481bd1, []int{5}
}
func (m *IdentityContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityContext.Unmarshal(m, b)
}
func (m *IdentityContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityContext.Marshal(b, m, deterministic)
}
func (dst *IdentityContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityContext.Merge(dst, src)
}
func (m *IdentityContext) XXX_Size() int {
	return xxx_messageInfo_IdentityContext.Size(m)
}
func (m *IdentityContext) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityContext.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityContext proto.InternalMessageInfo

type LogLevel struct {
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevel) Reset()         { *m = LogLevel{} }
func (m *LogLevel) String() string { return proto.CompactTextString(m) }
func (*LogLevel) ProtoMessage()    {}
func (*LogLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f7da3f9300481bd1, []int{6}
}
func (m *LogLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevel.Unmarshal(m, b)
}
func (m *LogLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevel.Marshal(b, m, deterministic)
}
func (dst *LogLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevel.Merge(dst, src)
}
func (m *LogLevel) XXX_Size() int {
	return xxx_messageInfo_LogLevel.Size(m)
}
func (m *LogLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevel.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevel proto.InternalMessageInfo

func (m *LogLevel) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func init() {
	proto.RegisterType((*GlobalConfig)(nil), "linkerd2.config.GlobalConfig")
	proto.RegisterType((*ProxyConfig)(nil), "linkerd2.config.ProxyConfig")
	proto.RegisterType((*Image)(nil), "linkerd2.config.Image")
	proto.RegisterType((*Port)(nil), "linkerd2.config.Port")
	proto.RegisterType((*ResourceRequirements)(nil), "linkerd2.config.ResourceRequirements")
	proto.RegisterType((*IdentityContext)(nil), "linkerd2.config.IdentityContext")
	proto.RegisterType((*LogLevel)(nil), "linkerd2.config.LogLevel")
}

func init() { proto.RegisterFile("config/config.proto", fileDescriptor_config_f7da3f9300481bd1) }

var fileDescriptor_config_f7da3f9300481bd1 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5d, 0x6b, 0xdb, 0x4a,
	0x10, 0xc5, 0x89, 0x93, 0x6b, 0x8f, 0xec, 0x7c, 0x6c, 0x9c, 0x7b, 0x95, 0x5c, 0x4a, 0x55, 0x41,
	0xc0, 0x50, 0xb0, 0xa9, 0x03, 0x6d, 0xc8, 0x53, 0xd3, 0x10, 0x4c, 0x68, 0xda, 0x1a, 0x41, 0x5f,
	0xfa, 0x22, 0x64, 0x69, 0xa3, 0x0e, 0x5d, 0xed, 0x2a, 0xab, 0x55, 0x89, 0xff, 0x4c, 0x5f, 0xfa,
	0x5b, 0xfa, 0xbf, 0xca, 0x7e, 0x38, 0x75, 0xe2, 0xa2, 0x27, 0xed, 0x9e, 0x39, 0xe7, 0xec, 0xec,
	0x68, 0x66, 0xe1, 0x20, 0x15, 0xfc, 0x16, 0xf3, 0xb1, 0xfd, 0x8c, 0x4a, 0x29, 0x94, 0x20, 0xbb,
	0x0c, 0xf9, 0x37, 0x2a, 0xb3, 0xc9, 0xc8, 0xc2, 0xe1, 0xaf, 0x16, 0xf4, 0xa6, 0x4c, 0xcc, 0x13,
	0x76, 0x69, 0x00, 0xf2, 0x12, 0xf6, 0x1d, 0x27, 0xe6, 0x49, 0x41, 0xab, 0x32, 0x49, 0xa9, 0xdf,
	0x0a, 0x5a, 0xc3, 0x6e, 0xb4, 0xe7, 0x02, 0x1f, 0x97, 0x38, 0x79, 0x0e, 0x5e, 0xca, 0x31, 0xa6,
	0x3c, 0x99, 0x33, 0x9a, 0xf9, 0x1b, 0x41, 0x6b, 0xd8, 0x89, 0x20, 0xe5, 0x78, 0x65, 0x11, 0x72,
	0x0c, 0x1d, 0x49, 0x73, 0xac, 0x94, 0x5c, 0xf8, 0x9b, 0xc6, 0xe4, 0x61, 0x4f, 0xde, 0xc3, 0x1e,
	0x66, 0x94, 0x2b, 0x54, 0x8b, 0x38, 0x15, 0x5c, 0xd1, 0x7b, 0xe5, 0xb7, 0x83, 0xd6, 0xd0, 0x9b,
	0x04, 0xa3, 0x27, 0x69, 0x8e, 0xae, 0x1d, 0xf1, 0xd2, 0xf2, 0xa2, 0x5d, 0x7c, 0x0c, 0x84, 0x3f,
	0xb7, 0xc1, 0x9b, 0x49, 0x71, 0xbf, 0x70, 0xd7, 0x78, 0x03, 0x5e, 0xa9, 0xb7, 0x31, 0x16, 0x49,
	0x6e, 0x2f, 0xe0, 0x4d, 0xfe, 0x5d, 0xf7, 0xd5, 0xd1, 0x08, 0x0c, 0xd5, 0xac, 0xc9, 0x5b, 0xd8,
	0x73, 0x42, 0x8e, 0xca, 0xa9, 0x37, 0x1a, 0xd5, 0x3b, 0x56, 0xcd, 0x51, 0x59, 0x87, 0x29, 0x0c,
	0x32, 0x5a, 0x29, 0xe4, 0x89, 0x42, 0xc1, 0xe3, 0xa4, 0xc4, 0xb8, 0x14, 0x52, 0x99, 0xfb, 0x7b,
	0x93, 0xc3, 0x35, 0x97, 0x99, 0x90, 0x2a, 0x22, 0x2b, 0x92, 0x8b, 0x12, 0x35, 0x46, 0xce, 0xa0,
	0xa7, 0xeb, 0x22, 0x05, 0xb3, 0x06, 0xed, 0x26, 0x03, 0xcf, 0x51, 0x8d, 0x72, 0x0a, 0x03, 0xcc,
	0xb9, 0x90, 0x34, 0x46, 0x3e, 0x17, 0x35, 0xcf, 0x8c, 0x41, 0xe5, 0x6f, 0x05, 0x9b, 0x0d, 0x29,
	0x58, 0xc9, 0xb5, 0x55, 0x68, 0xa8, 0x22, 0xd7, 0x70, 0xe8, 0x8c, 0x44, 0xad, 0x56, 0x9d, 0xb6,
	0x9b, 0x9c, 0x0e, 0xac, 0xe6, 0x93, 0x93, 0x58, 0xab, 0x33, 0xe8, 0xad, 0x26, 0xe3, 0xff, 0xd3,
	0x78, 0x1b, 0xfc, 0x93, 0x85, 0x56, 0x16, 0x54, 0x49, 0x4c, 0x2b, 0xab, 0xec, 0x34, 0x2a, 0x1d,
	0xd5, 0x28, 0xcf, 0xa1, 0xff, 0x28, 0x6f, 0xbf, 0xdb, 0x24, 0xed, 0x89, 0x95, 0x84, 0xc9, 0x85,
	0x6e, 0xdd, 0x4a, 0xd4, 0x32, 0xa5, 0x3e, 0x18, 0xd9, 0xc9, 0x9a, 0x2c, 0x72, 0x84, 0x88, 0xde,
	0xd5, 0x28, 0x69, 0x41, 0xb9, 0xaa, 0xa2, 0x07, 0x19, 0xf9, 0x1f, 0xba, 0xb6, 0x97, 0x6a, 0xcc,
	0x7c, 0x2f, 0x68, 0x0d, 0x37, 0xa3, 0x8e, 0x01, 0x3e, 0x63, 0x46, 0x5e, 0x43, 0x97, 0x89, 0x3c,
	0x66, 0xf4, 0x3b, 0x65, 0x7e, 0xcf, 0x1c, 0x70, 0xb4, 0x76, 0xc0, 0x8d, 0xc8, 0x6f, 0x34, 0x21,
	0xea, 0x30, 0xb7, 0x22, 0xe7, 0x70, 0x94, 0x61, 0xa5, 0xc7, 0x2b, 0xa6, 0xf7, 0x8a, 0x4a, 0x9e,
	0xb0, 0xb8, 0x94, 0xe2, 0x16, 0x19, 0xad, 0xfc, 0xbe, 0x99, 0xc0, 0xff, 0x1c, 0xe1, 0xca, 0xc5,
	0x67, 0x2e, 0x1c, 0x4e, 0x61, 0xcb, 0xf6, 0xe8, 0x33, 0x00, 0xd3, 0xda, 0x66, 0xc6, 0xdd, 0x78,
	0x77, 0x0d, 0xa2, 0x87, 0x5b, 0xcf, 0x75, 0x59, 0x33, 0xdd, 0x76, 0x0c, 0xd3, 0x85, 0xe9, 0xff,
	0x6e, 0x04, 0x1a, 0x9a, 0x19, 0x24, 0x3c, 0x86, 0xb6, 0x29, 0x12, 0x81, 0xb6, 0xa9, 0xab, 0x76,
	0xe8, 0x47, 0x66, 0x1d, 0xfe, 0x68, 0xc1, 0xe0, 0x6f, 0x85, 0xd1, 0xae, 0x92, 0xde, 0xd5, 0xb4,
	0x52, 0x71, 0x5a, 0xd6, 0xee, 0x54, 0x70, 0xd0, 0x65, 0x59, 0x93, 0x13, 0xd8, 0x59, 0x12, 0x0a,
	0x5a, 0x08, 0xb9, 0x3c, 0xb9, 0xef, 0xd0, 0x0f, 0x06, 0xd4, 0x65, 0x65, 0x58, 0xa0, 0x75, 0x71,
	0xaf, 0x8a, 0x01, 0xb4, 0xc7, 0x0b, 0xe8, 0xd9, 0xa0, 0x73, 0x68, 0x9b, 0xb8, 0x67, 0x30, 0xab,
	0x0f, 0xf7, 0x61, 0xf7, 0xc9, 0x7b, 0x12, 0x06, 0xd0, 0x59, 0x96, 0x9a, 0x0c, 0x60, 0xcb, 0xfe,
	0x14, 0x9b, 0xa0, 0xdd, 0xbc, 0x3b, 0xfd, 0xf2, 0x2a, 0x47, 0xf5, 0xb5, 0x9e, 0x8f, 0x52, 0x51,
	0x8c, 0xdd, 0x7f, 0x5a, 0x7e, 0x27, 0x63, 0x37, 0x7d, 0x8c, 0xca, 0x71, 0x4e, 0xb9, 0x7b, 0x74,
	0xe7, 0xdb, 0xe6, 0xd5, 0x3d, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x87, 0x01, 0xe5, 0x8c,
	0x05, 0x00, 0x00,
}
