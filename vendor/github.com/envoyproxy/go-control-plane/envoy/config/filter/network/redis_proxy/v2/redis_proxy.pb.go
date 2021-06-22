// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/redis_proxy/v2/redis_proxy.proto

package envoy_config_filter_network_redis_proxy_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type RedisProxy_ConnPoolSettings_ReadPolicy int32

const (
	RedisProxy_ConnPoolSettings_MASTER         RedisProxy_ConnPoolSettings_ReadPolicy = 0
	RedisProxy_ConnPoolSettings_PREFER_MASTER  RedisProxy_ConnPoolSettings_ReadPolicy = 1
	RedisProxy_ConnPoolSettings_REPLICA        RedisProxy_ConnPoolSettings_ReadPolicy = 2
	RedisProxy_ConnPoolSettings_PREFER_REPLICA RedisProxy_ConnPoolSettings_ReadPolicy = 3
	RedisProxy_ConnPoolSettings_ANY            RedisProxy_ConnPoolSettings_ReadPolicy = 4
)

var RedisProxy_ConnPoolSettings_ReadPolicy_name = map[int32]string{
	0: "MASTER",
	1: "PREFER_MASTER",
	2: "REPLICA",
	3: "PREFER_REPLICA",
	4: "ANY",
}

var RedisProxy_ConnPoolSettings_ReadPolicy_value = map[string]int32{
	"MASTER":         0,
	"PREFER_MASTER":  1,
	"REPLICA":        2,
	"PREFER_REPLICA": 3,
	"ANY":            4,
}

func (x RedisProxy_ConnPoolSettings_ReadPolicy) String() string {
	return proto.EnumName(RedisProxy_ConnPoolSettings_ReadPolicy_name, int32(x))
}

func (RedisProxy_ConnPoolSettings_ReadPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67e7179f1292d5ae, []int{0, 0, 0}
}

type RedisProxy struct {
	StatPrefix             string                       `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	Cluster                string                       `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"` // Deprecated: Do not use.
	Settings               *RedisProxy_ConnPoolSettings `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	LatencyInMicros        bool                         `protobuf:"varint,4,opt,name=latency_in_micros,json=latencyInMicros,proto3" json:"latency_in_micros,omitempty"`
	PrefixRoutes           *RedisProxy_PrefixRoutes     `protobuf:"bytes,5,opt,name=prefix_routes,json=prefixRoutes,proto3" json:"prefix_routes,omitempty"`
	DownstreamAuthPassword *core.DataSource             `protobuf:"bytes,6,opt,name=downstream_auth_password,json=downstreamAuthPassword,proto3" json:"downstream_auth_password,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                     `json:"-"`
	XXX_unrecognized       []byte                       `json:"-"`
	XXX_sizecache          int32                        `json:"-"`
}

func (m *RedisProxy) Reset()         { *m = RedisProxy{} }
func (m *RedisProxy) String() string { return proto.CompactTextString(m) }
func (*RedisProxy) ProtoMessage()    {}
func (*RedisProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e7179f1292d5ae, []int{0}
}

func (m *RedisProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy.Unmarshal(m, b)
}
func (m *RedisProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy.Marshal(b, m, deterministic)
}
func (m *RedisProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy.Merge(m, src)
}
func (m *RedisProxy) XXX_Size() int {
	return xxx_messageInfo_RedisProxy.Size(m)
}
func (m *RedisProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy proto.InternalMessageInfo

func (m *RedisProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

// Deprecated: Do not use.
func (m *RedisProxy) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy) GetSettings() *RedisProxy_ConnPoolSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *RedisProxy) GetLatencyInMicros() bool {
	if m != nil {
		return m.LatencyInMicros
	}
	return false
}

func (m *RedisProxy) GetPrefixRoutes() *RedisProxy_PrefixRoutes {
	if m != nil {
		return m.PrefixRoutes
	}
	return nil
}

func (m *RedisProxy) GetDownstreamAuthPassword() *core.DataSource {
	if m != nil {
		return m.DownstreamAuthPassword
	}
	return nil
}

type RedisProxy_ConnPoolSettings struct {
	OpTimeout                     *duration.Duration                     `protobuf:"bytes,1,opt,name=op_timeout,json=opTimeout,proto3" json:"op_timeout,omitempty"`
	EnableHashtagging             bool                                   `protobuf:"varint,2,opt,name=enable_hashtagging,json=enableHashtagging,proto3" json:"enable_hashtagging,omitempty"`
	EnableRedirection             bool                                   `protobuf:"varint,3,opt,name=enable_redirection,json=enableRedirection,proto3" json:"enable_redirection,omitempty"`
	MaxBufferSizeBeforeFlush      uint32                                 `protobuf:"varint,4,opt,name=max_buffer_size_before_flush,json=maxBufferSizeBeforeFlush,proto3" json:"max_buffer_size_before_flush,omitempty"`
	BufferFlushTimeout            *duration.Duration                     `protobuf:"bytes,5,opt,name=buffer_flush_timeout,json=bufferFlushTimeout,proto3" json:"buffer_flush_timeout,omitempty"`
	MaxUpstreamUnknownConnections *wrappers.UInt32Value                  `protobuf:"bytes,6,opt,name=max_upstream_unknown_connections,json=maxUpstreamUnknownConnections,proto3" json:"max_upstream_unknown_connections,omitempty"`
	EnableCommandStats            bool                                   `protobuf:"varint,8,opt,name=enable_command_stats,json=enableCommandStats,proto3" json:"enable_command_stats,omitempty"`
	ReadPolicy                    RedisProxy_ConnPoolSettings_ReadPolicy `protobuf:"varint,7,opt,name=read_policy,json=readPolicy,proto3,enum=envoy.config.filter.network.redis_proxy.v2.RedisProxy_ConnPoolSettings_ReadPolicy" json:"read_policy,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                               `json:"-"`
	XXX_unrecognized              []byte                                 `json:"-"`
	XXX_sizecache                 int32                                  `json:"-"`
}

func (m *RedisProxy_ConnPoolSettings) Reset()         { *m = RedisProxy_ConnPoolSettings{} }
func (m *RedisProxy_ConnPoolSettings) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_ConnPoolSettings) ProtoMessage()    {}
func (*RedisProxy_ConnPoolSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e7179f1292d5ae, []int{0, 0}
}

func (m *RedisProxy_ConnPoolSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Unmarshal(m, b)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Marshal(b, m, deterministic)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_ConnPoolSettings.Merge(m, src)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Size(m)
}
func (m *RedisProxy_ConnPoolSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_ConnPoolSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_ConnPoolSettings proto.InternalMessageInfo

func (m *RedisProxy_ConnPoolSettings) GetOpTimeout() *duration.Duration {
	if m != nil {
		return m.OpTimeout
	}
	return nil
}

func (m *RedisProxy_ConnPoolSettings) GetEnableHashtagging() bool {
	if m != nil {
		return m.EnableHashtagging
	}
	return false
}

func (m *RedisProxy_ConnPoolSettings) GetEnableRedirection() bool {
	if m != nil {
		return m.EnableRedirection
	}
	return false
}

func (m *RedisProxy_ConnPoolSettings) GetMaxBufferSizeBeforeFlush() uint32 {
	if m != nil {
		return m.MaxBufferSizeBeforeFlush
	}
	return 0
}

func (m *RedisProxy_ConnPoolSettings) GetBufferFlushTimeout() *duration.Duration {
	if m != nil {
		return m.BufferFlushTimeout
	}
	return nil
}

func (m *RedisProxy_ConnPoolSettings) GetMaxUpstreamUnknownConnections() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxUpstreamUnknownConnections
	}
	return nil
}

func (m *RedisProxy_ConnPoolSettings) GetEnableCommandStats() bool {
	if m != nil {
		return m.EnableCommandStats
	}
	return false
}

func (m *RedisProxy_ConnPoolSettings) GetReadPolicy() RedisProxy_ConnPoolSettings_ReadPolicy {
	if m != nil {
		return m.ReadPolicy
	}
	return RedisProxy_ConnPoolSettings_MASTER
}

type RedisProxy_PrefixRoutes struct {
	Routes               []*RedisProxy_PrefixRoutes_Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	CaseInsensitive      bool                             `protobuf:"varint,2,opt,name=case_insensitive,json=caseInsensitive,proto3" json:"case_insensitive,omitempty"`
	CatchAllCluster      string                           `protobuf:"bytes,3,opt,name=catch_all_cluster,json=catchAllCluster,proto3" json:"catch_all_cluster,omitempty"` // Deprecated: Do not use.
	CatchAllRoute        *RedisProxy_PrefixRoutes_Route   `protobuf:"bytes,4,opt,name=catch_all_route,json=catchAllRoute,proto3" json:"catch_all_route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RedisProxy_PrefixRoutes) Reset()         { *m = RedisProxy_PrefixRoutes{} }
func (m *RedisProxy_PrefixRoutes) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_PrefixRoutes) ProtoMessage()    {}
func (*RedisProxy_PrefixRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e7179f1292d5ae, []int{0, 1}
}

func (m *RedisProxy_PrefixRoutes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_PrefixRoutes.Unmarshal(m, b)
}
func (m *RedisProxy_PrefixRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_PrefixRoutes.Marshal(b, m, deterministic)
}
func (m *RedisProxy_PrefixRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_PrefixRoutes.Merge(m, src)
}
func (m *RedisProxy_PrefixRoutes) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_PrefixRoutes.Size(m)
}
func (m *RedisProxy_PrefixRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_PrefixRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_PrefixRoutes proto.InternalMessageInfo

func (m *RedisProxy_PrefixRoutes) GetRoutes() []*RedisProxy_PrefixRoutes_Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *RedisProxy_PrefixRoutes) GetCaseInsensitive() bool {
	if m != nil {
		return m.CaseInsensitive
	}
	return false
}

// Deprecated: Do not use.
func (m *RedisProxy_PrefixRoutes) GetCatchAllCluster() string {
	if m != nil {
		return m.CatchAllCluster
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes) GetCatchAllRoute() *RedisProxy_PrefixRoutes_Route {
	if m != nil {
		return m.CatchAllRoute
	}
	return nil
}

type RedisProxy_PrefixRoutes_Route struct {
	Prefix               string                                               `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	RemovePrefix         bool                                                 `protobuf:"varint,2,opt,name=remove_prefix,json=removePrefix,proto3" json:"remove_prefix,omitempty"`
	Cluster              string                                               `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	RequestMirrorPolicy  []*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy `protobuf:"bytes,4,rep,name=request_mirror_policy,json=requestMirrorPolicy,proto3" json:"request_mirror_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                             `json:"-"`
	XXX_unrecognized     []byte                                               `json:"-"`
	XXX_sizecache        int32                                                `json:"-"`
}

func (m *RedisProxy_PrefixRoutes_Route) Reset()         { *m = RedisProxy_PrefixRoutes_Route{} }
func (m *RedisProxy_PrefixRoutes_Route) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_PrefixRoutes_Route) ProtoMessage()    {}
func (*RedisProxy_PrefixRoutes_Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e7179f1292d5ae, []int{0, 1, 0}
}

func (m *RedisProxy_PrefixRoutes_Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Unmarshal(m, b)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Marshal(b, m, deterministic)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Merge(m, src)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Size(m)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_PrefixRoutes_Route proto.InternalMessageInfo

func (m *RedisProxy_PrefixRoutes_Route) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes_Route) GetRemovePrefix() bool {
	if m != nil {
		return m.RemovePrefix
	}
	return false
}

func (m *RedisProxy_PrefixRoutes_Route) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes_Route) GetRequestMirrorPolicy() []*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy {
	if m != nil {
		return m.RequestMirrorPolicy
	}
	return nil
}

type RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy struct {
	Cluster              string                         `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	RuntimeFraction      *core.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=runtime_fraction,json=runtimeFraction,proto3" json:"runtime_fraction,omitempty"`
	ExcludeReadCommands  bool                           `protobuf:"varint,3,opt,name=exclude_read_commands,json=excludeReadCommands,proto3" json:"exclude_read_commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) Reset() {
	*m = RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy{}
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) String() string {
	return proto.CompactTextString(m)
}
func (*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) ProtoMessage() {}
func (*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e7179f1292d5ae, []int{0, 1, 0, 0}
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Unmarshal(m, b)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Marshal(b, m, deterministic)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Merge(m, src)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Size(m)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy proto.InternalMessageInfo

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) GetRuntimeFraction() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.RuntimeFraction
	}
	return nil
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) GetExcludeReadCommands() bool {
	if m != nil {
		return m.ExcludeReadCommands
	}
	return false
}

type RedisProtocolOptions struct {
	AuthPassword         *core.DataSource `protobuf:"bytes,1,opt,name=auth_password,json=authPassword,proto3" json:"auth_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RedisProtocolOptions) Reset()         { *m = RedisProtocolOptions{} }
func (m *RedisProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*RedisProtocolOptions) ProtoMessage()    {}
func (*RedisProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e7179f1292d5ae, []int{1}
}

func (m *RedisProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProtocolOptions.Unmarshal(m, b)
}
func (m *RedisProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProtocolOptions.Marshal(b, m, deterministic)
}
func (m *RedisProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProtocolOptions.Merge(m, src)
}
func (m *RedisProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_RedisProtocolOptions.Size(m)
}
func (m *RedisProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProtocolOptions proto.InternalMessageInfo

func (m *RedisProtocolOptions) GetAuthPassword() *core.DataSource {
	if m != nil {
		return m.AuthPassword
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.filter.network.redis_proxy.v2.RedisProxy_ConnPoolSettings_ReadPolicy", RedisProxy_ConnPoolSettings_ReadPolicy_name, RedisProxy_ConnPoolSettings_ReadPolicy_value)
	proto.RegisterType((*RedisProxy)(nil), "envoy.config.filter.network.redis_proxy.v2.RedisProxy")
	proto.RegisterType((*RedisProxy_ConnPoolSettings)(nil), "envoy.config.filter.network.redis_proxy.v2.RedisProxy.ConnPoolSettings")
	proto.RegisterType((*RedisProxy_PrefixRoutes)(nil), "envoy.config.filter.network.redis_proxy.v2.RedisProxy.PrefixRoutes")
	proto.RegisterType((*RedisProxy_PrefixRoutes_Route)(nil), "envoy.config.filter.network.redis_proxy.v2.RedisProxy.PrefixRoutes.Route")
	proto.RegisterType((*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy)(nil), "envoy.config.filter.network.redis_proxy.v2.RedisProxy.PrefixRoutes.Route.RequestMirrorPolicy")
	proto.RegisterType((*RedisProtocolOptions)(nil), "envoy.config.filter.network.redis_proxy.v2.RedisProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/redis_proxy/v2/redis_proxy.proto", fileDescriptor_67e7179f1292d5ae)
}

var fileDescriptor_67e7179f1292d5ae = []byte{
	// 1096 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0x66, 0xed, 0xd8, 0x71, 0xc7, 0x71, 0xed, 0x4c, 0xd3, 0x62, 0xac, 0x04, 0x99, 0xf4, 0x62,
	0x8a, 0xd8, 0x05, 0x87, 0x03, 0x07, 0x84, 0x14, 0xbb, 0x09, 0x04, 0x08, 0x58, 0x93, 0xa6, 0x12,
	0x07, 0x34, 0x1a, 0xef, 0x8e, 0xed, 0x55, 0x77, 0x67, 0xb6, 0x33, 0xb3, 0x8e, 0x53, 0x09, 0xa9,
	0xe2, 0x58, 0x90, 0x10, 0x37, 0xfe, 0x00, 0x4e, 0x9c, 0xb8, 0x51, 0x71, 0x44, 0x1c, 0xb8, 0xc2,
	0x1f, 0xc2, 0xa1, 0xc7, 0x1e, 0x10, 0xda, 0x99, 0x59, 0xdb, 0x69, 0x52, 0x51, 0x50, 0x4e, 0xf6,
	0xcc, 0xf7, 0xde, 0xf7, 0xe6, 0xfd, 0xfa, 0x16, 0xbc, 0x47, 0xd9, 0x94, 0x9f, 0x7a, 0x3e, 0x67,
	0xa3, 0x70, 0xec, 0x8d, 0xc2, 0x48, 0x51, 0xe1, 0x31, 0xaa, 0x4e, 0xb8, 0xb8, 0xe7, 0x09, 0x1a,
	0x84, 0x12, 0x27, 0x82, 0xcf, 0x4e, 0xbd, 0x69, 0x77, 0xf9, 0xe8, 0x26, 0x82, 0x2b, 0x0e, 0x6f,
	0x69, 0x6f, 0xd7, 0x78, 0xbb, 0xc6, 0xdb, 0xb5, 0xde, 0xee, 0xb2, 0xf9, 0xb4, 0xdb, 0xda, 0x34,
	0x91, 0x48, 0x12, 0x66, 0x5c, 0x3e, 0x17, 0xd4, 0x1b, 0x12, 0x49, 0x0d, 0x53, 0xeb, 0xd5, 0x31,
	0xe7, 0xe3, 0x88, 0x7a, 0xfa, 0x34, 0x4c, 0x47, 0x5e, 0x90, 0x0a, 0xa2, 0x42, 0xce, 0x9e, 0x87,
	0x9f, 0x08, 0x92, 0x24, 0x54, 0x48, 0x8b, 0xdf, 0xb4, 0xec, 0x8c, 0x71, 0xa5, 0xfd, 0xa4, 0x17,
	0xd0, 0x44, 0x50, 0xff, 0x0c, 0x49, 0x1a, 0x24, 0xe4, 0x8c, 0x4d, 0x1c, 0x8e, 0x05, 0x51, 0xf9,
	0x23, 0xda, 0xe7, 0x70, 0x49, 0x99, 0x0c, 0x55, 0x38, 0xcd, 0x2d, 0xb6, 0xce, 0x5b, 0x28, 0xa2,
	0xd2, 0xfc, 0x15, 0x2f, 0x4f, 0x49, 0x14, 0x06, 0x44, 0x51, 0x2f, 0xff, 0x63, 0x80, 0xed, 0x87,
	0x75, 0x00, 0x50, 0x56, 0x8f, 0x41, 0x56, 0x0e, 0xd8, 0x01, 0xd5, 0xcc, 0x0f, 0x27, 0x82, 0x8e,
	0xc2, 0x59, 0xd3, 0x69, 0x3b, 0x9d, 0x2b, 0xbd, 0xd5, 0xa7, 0xbd, 0x15, 0x51, 0x68, 0x3b, 0x08,
	0x64, 0xd8, 0x40, 0x43, 0x70, 0x1b, 0xac, 0xfa, 0x51, 0x2a, 0x15, 0x15, 0xcd, 0x82, 0xb6, 0xaa,
	0x3c, 0xfe, 0xeb, 0xc9, 0x9f, 0x25, 0xa7, 0xe9, 0xa0, 0x1c, 0x80, 0x31, 0xa8, 0x48, 0xaa, 0x54,
	0xc8, 0xc6, 0xb2, 0x59, 0x6c, 0x3b, 0x9d, 0x6a, 0xf7, 0x03, 0xf7, 0xc5, 0x1b, 0xe3, 0x2e, 0xde,
	0xe5, 0xf6, 0x39, 0x63, 0x03, 0xce, 0xa3, 0x23, 0x4b, 0xd7, 0xab, 0x3c, 0xed, 0x95, 0x1e, 0x39,
	0x85, 0x86, 0x83, 0xe6, 0x21, 0xe0, 0x2d, 0xb0, 0x1e, 0x11, 0x45, 0x99, 0x7f, 0x8a, 0x43, 0x86,
	0xe3, 0xd0, 0x17, 0x5c, 0x36, 0x57, 0xda, 0x4e, 0xa7, 0x82, 0xea, 0x16, 0x38, 0x60, 0x87, 0xfa,
	0x1a, 0x4e, 0x40, 0xcd, 0xe4, 0x88, 0x05, 0x4f, 0x15, 0x95, 0xcd, 0x92, 0x7e, 0x5f, 0xff, 0x7f,
	0xbe, 0xcf, 0x14, 0x05, 0x69, 0x2a, 0xb4, 0x96, 0x2c, 0x9d, 0x20, 0x06, 0xcd, 0x80, 0x9f, 0x30,
	0xa9, 0x04, 0x25, 0x31, 0x26, 0xa9, 0x9a, 0xe0, 0x84, 0x48, 0x79, 0xc2, 0x45, 0xd0, 0x2c, 0xeb,
	0xa0, 0x5b, 0x36, 0x28, 0x49, 0xc2, 0x8c, 0x36, 0x9b, 0x40, 0xf7, 0x36, 0x51, 0xe4, 0x88, 0xa7,
	0xc2, 0xa7, 0xbd, 0xf2, 0xe3, 0x9f, 0xbf, 0xfe, 0xa1, 0xe0, 0xa0, 0x1b, 0x0b, 0x9a, 0xdd, 0x54,
	0x4d, 0x06, 0x96, 0xa4, 0xf5, 0x53, 0x09, 0x34, 0x9e, 0xad, 0x0f, 0xec, 0x01, 0xc0, 0x13, 0xac,
	0xc2, 0x98, 0xf2, 0x54, 0xe9, 0x3e, 0x56, 0xbb, 0xaf, 0xb8, 0x66, 0x56, 0xdd, 0x7c, 0x56, 0xdd,
	0xdb, 0x76, 0x96, 0x75, 0x39, 0x7f, 0x74, 0x0a, 0x15, 0x07, 0x5d, 0xe1, 0xc9, 0x1d, 0xe3, 0x05,
	0xdf, 0x04, 0x90, 0x32, 0x32, 0x8c, 0x28, 0x9e, 0x10, 0x39, 0x51, 0x64, 0x3c, 0x0e, 0xd9, 0x58,
	0x77, 0xbb, 0x82, 0xd6, 0x0d, 0xf2, 0xe1, 0x02, 0x58, 0x32, 0xcf, 0xea, 0x24, 0xa8, 0x9f, 0x31,
	0xeb, 0xbe, 0xcf, 0xcd, 0xd1, 0x02, 0x80, 0xef, 0x83, 0xcd, 0x98, 0xcc, 0xf0, 0x30, 0x1d, 0x8d,
	0xa8, 0xc0, 0x32, 0x7c, 0x40, 0xf1, 0x90, 0x8e, 0xb8, 0xa0, 0x78, 0x14, 0xa5, 0x72, 0xa2, 0x1b,
	0x57, 0x43, 0xcd, 0x98, 0xcc, 0x7a, 0xda, 0xe4, 0x28, 0x7c, 0x40, 0x7b, 0xda, 0x60, 0x3f, 0xc3,
	0xe1, 0xc7, 0x60, 0xc3, 0xfa, 0x6a, 0xfb, 0x79, 0xae, 0xa5, 0x7f, 0xc9, 0x15, 0x41, 0xe3, 0xa6,
	0x59, 0xf2, 0x54, 0x29, 0x68, 0x67, 0x8f, 0x49, 0x13, 0xdb, 0xa6, 0x94, 0xdd, 0x63, 0xfc, 0x84,
	0x61, 0x9f, 0x33, 0x66, 0xde, 0x2b, 0x6d, 0xb3, 0x36, 0xcf, 0x11, 0x1f, 0x1f, 0x30, 0xb5, 0xd3,
	0xbd, 0x4b, 0xa2, 0x94, 0xa2, 0xad, 0x98, 0xcc, 0x8e, 0x2d, 0xc9, 0xb1, 0xe1, 0xe8, 0x2f, 0x28,
	0xe0, 0x5b, 0x60, 0xc3, 0x96, 0xc8, 0xe7, 0x71, 0x4c, 0x58, 0x80, 0xb3, 0x8d, 0x92, 0xcd, 0x8a,
	0x2e, 0x92, 0x2d, 0x5f, 0xdf, 0x40, 0x47, 0x19, 0x02, 0xbf, 0x04, 0x55, 0x41, 0x49, 0x80, 0x13,
	0x1e, 0x85, 0xfe, 0x69, 0x73, 0xb5, 0xed, 0x74, 0xae, 0x76, 0xd1, 0x25, 0x6d, 0x91, 0x8b, 0x28,
	0x09, 0x06, 0x9a, 0x59, 0x4f, 0xc0, 0x57, 0x7a, 0xa1, 0x80, 0x98, 0xdf, 0x6e, 0x1f, 0x67, 0xea,
	0x90, 0x9f, 0x20, 0x00, 0xe5, 0xc3, 0xdd, 0xa3, 0x3b, 0x7b, 0xa8, 0xf1, 0x12, 0x5c, 0x07, 0xb5,
	0x01, 0xda, 0xdb, 0xdf, 0x43, 0xd8, 0x5e, 0x39, 0xb0, 0x0a, 0x56, 0xd1, 0xde, 0xe0, 0x93, 0x83,
	0xfe, 0x6e, 0xa3, 0x00, 0x21, 0xb8, 0x6a, 0xf1, 0xfc, 0xae, 0x08, 0x57, 0x41, 0x71, 0xf7, 0xd3,
	0xcf, 0x1b, 0x2b, 0xad, 0x47, 0x65, 0xb0, 0xb6, 0xbc, 0x32, 0x90, 0x80, 0xb2, 0xdd, 0x43, 0xa7,
	0x5d, 0xec, 0x54, 0xbb, 0x07, 0x97, 0xb0, 0x87, 0xae, 0xfe, 0x41, 0x96, 0x18, 0xbe, 0x0e, 0x1a,
	0x3e, 0x91, 0x14, 0x87, 0x6c, 0xae, 0x9d, 0x76, 0x96, 0xeb, 0xd9, 0xfd, 0xc1, 0xe2, 0x1a, 0xbe,
	0x03, 0xd6, 0x7d, 0xa2, 0xfc, 0x09, 0x26, 0x51, 0x84, 0x73, 0x95, 0x2b, 0x3e, 0xa3, 0x72, 0x75,
	0x6d, 0xb2, 0x1b, 0x45, 0x7d, 0xab, 0x76, 0xf7, 0x41, 0x7d, 0xe1, 0xa5, 0x83, 0xea, 0x19, 0xbe,
	0xd4, 0x64, 0x6a, 0x79, 0x50, 0x7d, 0x6c, 0xfd, 0x5a, 0x04, 0x25, 0xfd, 0x0f, 0xde, 0x00, 0xe5,
	0x65, 0xcd, 0x46, 0xf6, 0x04, 0x6f, 0x82, 0x9a, 0xa0, 0x31, 0x9f, 0xd2, 0x5c, 0xd2, 0x4d, 0xca,
	0x6b, 0xe6, 0xd2, 0x6a, 0xf9, 0x6b, 0x0b, 0x2d, 0x2f, 0x9e, 0x55, 0xfc, 0xb9, 0x94, 0x7f, 0xe7,
	0x80, 0xeb, 0x82, 0xde, 0x4f, 0xa9, 0x54, 0x38, 0x0e, 0x85, 0xe0, 0x22, 0x1f, 0xc9, 0x15, 0xdd,
	0xb0, 0x2f, 0x2e, 0x2d, 0x47, 0x17, 0x99, 0x30, 0x87, 0x3a, 0x8a, 0x99, 0x3c, 0x74, 0x4d, 0x9c,
	0xbf, 0x6c, 0xfd, 0xe6, 0x80, 0x6b, 0x17, 0x18, 0x2f, 0xa7, 0xe3, 0x3c, 0x27, 0x9d, 0xbb, 0xa0,
	0x21, 0x52, 0x96, 0x49, 0x06, 0x1e, 0x09, 0x62, 0x94, 0xaa, 0xa0, 0x9b, 0xf5, 0xc6, 0x05, 0x62,
	0x8c, 0x8c, 0xe9, 0xbe, 0xb5, 0x24, 0xd1, 0x80, 0x0a, 0x9f, 0x32, 0x85, 0xea, 0xe2, 0x2c, 0x02,
	0xbb, 0xe0, 0x3a, 0x9d, 0xf9, 0x51, 0x1a, 0x64, 0x22, 0x48, 0x82, 0x7c, 0xcd, 0xa5, 0x95, 0xc1,
	0x6b, 0x16, 0xcc, 0x76, 0xca, 0xae, 0xb9, 0xdc, 0x1e, 0x82, 0x8d, 0xbc, 0x20, 0x8a, 0xfb, 0x3c,
	0xfa, 0x2c, 0x31, 0x62, 0xf1, 0x11, 0xa8, 0x9d, 0xfd, 0x5a, 0x38, 0xff, 0xe5, 0x6b, 0xb1, 0x46,
	0x96, 0xbe, 0x11, 0xbd, 0x6f, 0x9c, 0x27, 0xdf, 0xff, 0xfd, 0x6d, 0xe9, 0x6d, 0xe8, 0x19, 0x67,
	0x3a, 0x53, 0xd9, 0xb4, 0x73, 0x26, 0x6d, 0xab, 0xe4, 0xc5, 0xbd, 0xda, 0xf9, 0xe5, 0xe1, 0xef,
	0x7f, 0x94, 0x0b, 0x0d, 0x07, 0xbc, 0x1b, 0x72, 0x13, 0xd8, 0x20, 0x2f, 0xde, 0xed, 0x5e, 0x7d,
	0xd1, 0x6e, 0x9d, 0xe2, 0xc0, 0x19, 0x96, 0xb5, 0x78, 0xee, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x6d, 0x76, 0xa5, 0x4f, 0xe5, 0x09, 0x00, 0x00,
}