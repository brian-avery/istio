// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/server_info.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ServerInfo_State int32

const (
	ServerInfo_LIVE             ServerInfo_State = 0
	ServerInfo_DRAINING         ServerInfo_State = 1
	ServerInfo_PRE_INITIALIZING ServerInfo_State = 2
	ServerInfo_INITIALIZING     ServerInfo_State = 3
)

var ServerInfo_State_name = map[int32]string{
	0: "LIVE",
	1: "DRAINING",
	2: "PRE_INITIALIZING",
	3: "INITIALIZING",
}

var ServerInfo_State_value = map[string]int32{
	"LIVE":             0,
	"DRAINING":         1,
	"PRE_INITIALIZING": 2,
	"INITIALIZING":     3,
}

func (x ServerInfo_State) String() string {
	return proto.EnumName(ServerInfo_State_name, int32(x))
}

func (ServerInfo_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{0, 0}
}

type CommandLineOptions_IpVersion int32

const (
	CommandLineOptions_v4 CommandLineOptions_IpVersion = 0
	CommandLineOptions_v6 CommandLineOptions_IpVersion = 1
)

var CommandLineOptions_IpVersion_name = map[int32]string{
	0: "v4",
	1: "v6",
}

var CommandLineOptions_IpVersion_value = map[string]int32{
	"v4": 0,
	"v6": 1,
}

func (x CommandLineOptions_IpVersion) String() string {
	return proto.EnumName(CommandLineOptions_IpVersion_name, int32(x))
}

func (CommandLineOptions_IpVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{1, 0}
}

type CommandLineOptions_Mode int32

const (
	CommandLineOptions_Serve    CommandLineOptions_Mode = 0
	CommandLineOptions_Validate CommandLineOptions_Mode = 1
	CommandLineOptions_InitOnly CommandLineOptions_Mode = 2
)

var CommandLineOptions_Mode_name = map[int32]string{
	0: "Serve",
	1: "Validate",
	2: "InitOnly",
}

var CommandLineOptions_Mode_value = map[string]int32{
	"Serve":    0,
	"Validate": 1,
	"InitOnly": 2,
}

func (x CommandLineOptions_Mode) String() string {
	return proto.EnumName(CommandLineOptions_Mode_name, int32(x))
}

func (CommandLineOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{1, 1}
}

type ServerInfo struct {
	Version              string              `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	State                ServerInfo_State    `protobuf:"varint,2,opt,name=state,proto3,enum=envoy.admin.v2alpha.ServerInfo_State" json:"state,omitempty"`
	UptimeCurrentEpoch   *duration.Duration  `protobuf:"bytes,3,opt,name=uptime_current_epoch,json=uptimeCurrentEpoch,proto3" json:"uptime_current_epoch,omitempty"`
	UptimeAllEpochs      *duration.Duration  `protobuf:"bytes,4,opt,name=uptime_all_epochs,json=uptimeAllEpochs,proto3" json:"uptime_all_epochs,omitempty"`
	HotRestartVersion    string              `protobuf:"bytes,5,opt,name=hot_restart_version,json=hotRestartVersion,proto3" json:"hot_restart_version,omitempty"`
	CommandLineOptions   *CommandLineOptions `protobuf:"bytes,6,opt,name=command_line_options,json=commandLineOptions,proto3" json:"command_line_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetState() ServerInfo_State {
	if m != nil {
		return m.State
	}
	return ServerInfo_LIVE
}

func (m *ServerInfo) GetUptimeCurrentEpoch() *duration.Duration {
	if m != nil {
		return m.UptimeCurrentEpoch
	}
	return nil
}

func (m *ServerInfo) GetUptimeAllEpochs() *duration.Duration {
	if m != nil {
		return m.UptimeAllEpochs
	}
	return nil
}

func (m *ServerInfo) GetHotRestartVersion() string {
	if m != nil {
		return m.HotRestartVersion
	}
	return ""
}

func (m *ServerInfo) GetCommandLineOptions() *CommandLineOptions {
	if m != nil {
		return m.CommandLineOptions
	}
	return nil
}

type CommandLineOptions struct {
	BaseId                     uint64                       `protobuf:"varint,1,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	Concurrency                uint32                       `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	ConfigPath                 string                       `protobuf:"bytes,3,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	ConfigYaml                 string                       `protobuf:"bytes,4,opt,name=config_yaml,json=configYaml,proto3" json:"config_yaml,omitempty"`
	AllowUnknownStaticFields   bool                         `protobuf:"varint,5,opt,name=allow_unknown_static_fields,json=allowUnknownStaticFields,proto3" json:"allow_unknown_static_fields,omitempty"`
	RejectUnknownDynamicFields bool                         `protobuf:"varint,26,opt,name=reject_unknown_dynamic_fields,json=rejectUnknownDynamicFields,proto3" json:"reject_unknown_dynamic_fields,omitempty"`
	AdminAddressPath           string                       `protobuf:"bytes,6,opt,name=admin_address_path,json=adminAddressPath,proto3" json:"admin_address_path,omitempty"`
	LocalAddressIpVersion      CommandLineOptions_IpVersion `protobuf:"varint,7,opt,name=local_address_ip_version,json=localAddressIpVersion,proto3,enum=envoy.admin.v2alpha.CommandLineOptions_IpVersion" json:"local_address_ip_version,omitempty"`
	LogLevel                   string                       `protobuf:"bytes,8,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	ComponentLogLevel          string                       `protobuf:"bytes,9,opt,name=component_log_level,json=componentLogLevel,proto3" json:"component_log_level,omitempty"`
	LogFormat                  string                       `protobuf:"bytes,10,opt,name=log_format,json=logFormat,proto3" json:"log_format,omitempty"`
	LogPath                    string                       `protobuf:"bytes,11,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
	ServiceCluster             string                       `protobuf:"bytes,13,opt,name=service_cluster,json=serviceCluster,proto3" json:"service_cluster,omitempty"`
	ServiceNode                string                       `protobuf:"bytes,14,opt,name=service_node,json=serviceNode,proto3" json:"service_node,omitempty"`
	ServiceZone                string                       `protobuf:"bytes,15,opt,name=service_zone,json=serviceZone,proto3" json:"service_zone,omitempty"`
	FileFlushInterval          *duration.Duration           `protobuf:"bytes,16,opt,name=file_flush_interval,json=fileFlushInterval,proto3" json:"file_flush_interval,omitempty"`
	DrainTime                  *duration.Duration           `protobuf:"bytes,17,opt,name=drain_time,json=drainTime,proto3" json:"drain_time,omitempty"`
	ParentShutdownTime         *duration.Duration           `protobuf:"bytes,18,opt,name=parent_shutdown_time,json=parentShutdownTime,proto3" json:"parent_shutdown_time,omitempty"`
	Mode                       CommandLineOptions_Mode      `protobuf:"varint,19,opt,name=mode,proto3,enum=envoy.admin.v2alpha.CommandLineOptions_Mode" json:"mode,omitempty"`
	MaxStats                   uint64                       `protobuf:"varint,20,opt,name=max_stats,json=maxStats,proto3" json:"max_stats,omitempty"`                    // Deprecated: Do not use.
	MaxObjNameLen              uint64                       `protobuf:"varint,21,opt,name=max_obj_name_len,json=maxObjNameLen,proto3" json:"max_obj_name_len,omitempty"` // Deprecated: Do not use.
	DisableHotRestart          bool                         `protobuf:"varint,22,opt,name=disable_hot_restart,json=disableHotRestart,proto3" json:"disable_hot_restart,omitempty"`
	EnableMutexTracing         bool                         `protobuf:"varint,23,opt,name=enable_mutex_tracing,json=enableMutexTracing,proto3" json:"enable_mutex_tracing,omitempty"`
	RestartEpoch               uint32                       `protobuf:"varint,24,opt,name=restart_epoch,json=restartEpoch,proto3" json:"restart_epoch,omitempty"`
	CpusetThreads              bool                         `protobuf:"varint,25,opt,name=cpuset_threads,json=cpusetThreads,proto3" json:"cpuset_threads,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                     `json:"-"`
	XXX_unrecognized           []byte                       `json:"-"`
	XXX_sizecache              int32                        `json:"-"`
}

func (m *CommandLineOptions) Reset()         { *m = CommandLineOptions{} }
func (m *CommandLineOptions) String() string { return proto.CompactTextString(m) }
func (*CommandLineOptions) ProtoMessage()    {}
func (*CommandLineOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{1}
}

func (m *CommandLineOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLineOptions.Unmarshal(m, b)
}
func (m *CommandLineOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLineOptions.Marshal(b, m, deterministic)
}
func (m *CommandLineOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLineOptions.Merge(m, src)
}
func (m *CommandLineOptions) XXX_Size() int {
	return xxx_messageInfo_CommandLineOptions.Size(m)
}
func (m *CommandLineOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLineOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLineOptions proto.InternalMessageInfo

func (m *CommandLineOptions) GetBaseId() uint64 {
	if m != nil {
		return m.BaseId
	}
	return 0
}

func (m *CommandLineOptions) GetConcurrency() uint32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *CommandLineOptions) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *CommandLineOptions) GetConfigYaml() string {
	if m != nil {
		return m.ConfigYaml
	}
	return ""
}

func (m *CommandLineOptions) GetAllowUnknownStaticFields() bool {
	if m != nil {
		return m.AllowUnknownStaticFields
	}
	return false
}

func (m *CommandLineOptions) GetRejectUnknownDynamicFields() bool {
	if m != nil {
		return m.RejectUnknownDynamicFields
	}
	return false
}

func (m *CommandLineOptions) GetAdminAddressPath() string {
	if m != nil {
		return m.AdminAddressPath
	}
	return ""
}

func (m *CommandLineOptions) GetLocalAddressIpVersion() CommandLineOptions_IpVersion {
	if m != nil {
		return m.LocalAddressIpVersion
	}
	return CommandLineOptions_v4
}

func (m *CommandLineOptions) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetComponentLogLevel() string {
	if m != nil {
		return m.ComponentLogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormat() string {
	if m != nil {
		return m.LogFormat
	}
	return ""
}

func (m *CommandLineOptions) GetLogPath() string {
	if m != nil {
		return m.LogPath
	}
	return ""
}

func (m *CommandLineOptions) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

func (m *CommandLineOptions) GetServiceNode() string {
	if m != nil {
		return m.ServiceNode
	}
	return ""
}

func (m *CommandLineOptions) GetServiceZone() string {
	if m != nil {
		return m.ServiceZone
	}
	return ""
}

func (m *CommandLineOptions) GetFileFlushInterval() *duration.Duration {
	if m != nil {
		return m.FileFlushInterval
	}
	return nil
}

func (m *CommandLineOptions) GetDrainTime() *duration.Duration {
	if m != nil {
		return m.DrainTime
	}
	return nil
}

func (m *CommandLineOptions) GetParentShutdownTime() *duration.Duration {
	if m != nil {
		return m.ParentShutdownTime
	}
	return nil
}

func (m *CommandLineOptions) GetMode() CommandLineOptions_Mode {
	if m != nil {
		return m.Mode
	}
	return CommandLineOptions_Serve
}

// Deprecated: Do not use.
func (m *CommandLineOptions) GetMaxStats() uint64 {
	if m != nil {
		return m.MaxStats
	}
	return 0
}

// Deprecated: Do not use.
func (m *CommandLineOptions) GetMaxObjNameLen() uint64 {
	if m != nil {
		return m.MaxObjNameLen
	}
	return 0
}

func (m *CommandLineOptions) GetDisableHotRestart() bool {
	if m != nil {
		return m.DisableHotRestart
	}
	return false
}

func (m *CommandLineOptions) GetEnableMutexTracing() bool {
	if m != nil {
		return m.EnableMutexTracing
	}
	return false
}

func (m *CommandLineOptions) GetRestartEpoch() uint32 {
	if m != nil {
		return m.RestartEpoch
	}
	return 0
}

func (m *CommandLineOptions) GetCpusetThreads() bool {
	if m != nil {
		return m.CpusetThreads
	}
	return false
}

func init() {
	proto.RegisterEnum("envoy.admin.v2alpha.ServerInfo_State", ServerInfo_State_name, ServerInfo_State_value)
	proto.RegisterEnum("envoy.admin.v2alpha.CommandLineOptions_IpVersion", CommandLineOptions_IpVersion_name, CommandLineOptions_IpVersion_value)
	proto.RegisterEnum("envoy.admin.v2alpha.CommandLineOptions_Mode", CommandLineOptions_Mode_name, CommandLineOptions_Mode_value)
	proto.RegisterType((*ServerInfo)(nil), "envoy.admin.v2alpha.ServerInfo")
	proto.RegisterType((*CommandLineOptions)(nil), "envoy.admin.v2alpha.CommandLineOptions")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/server_info.proto", fileDescriptor_ed0f406f9d75bf97)
}

var fileDescriptor_ed0f406f9d75bf97 = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0x6d, 0x4f, 0x1b, 0x47,
	0x10, 0xc7, 0x63, 0x30, 0x60, 0x0f, 0x18, 0x8e, 0x85, 0x34, 0x0b, 0x28, 0x0d, 0xa1, 0x42, 0x41,
	0x6a, 0x7a, 0xb4, 0xb4, 0x8a, 0x2a, 0x55, 0x95, 0xca, 0x63, 0x7a, 0xad, 0x03, 0xe8, 0xa0, 0x48,
	0xc9, 0x9b, 0xd5, 0xfa, 0x6e, 0x6d, 0x2f, 0xdd, 0xdb, 0x3d, 0xdd, 0xad, 0x1d, 0xdc, 0xf7, 0xfd,
	0x74, 0xfd, 0x52, 0xd5, 0xce, 0x9e, 0x6d, 0xa2, 0x20, 0x91, 0x57, 0xa7, 0x9d, 0xf9, 0xcd, 0x7f,
	0x1f, 0xe6, 0xe1, 0x60, 0x57, 0xe8, 0xa1, 0x19, 0xed, 0xf3, 0x34, 0x93, 0x7a, 0x7f, 0x78, 0xc0,
	0x55, 0xde, 0xe7, 0xfb, 0xa5, 0x28, 0x86, 0xa2, 0x60, 0x52, 0x77, 0x4d, 0x98, 0x17, 0xc6, 0x1a,
	0xb2, 0x86, 0x58, 0x88, 0x58, 0x58, 0x61, 0x9b, 0x5f, 0xf7, 0x8c, 0xe9, 0x29, 0xb1, 0x8f, 0x48,
	0x67, 0xd0, 0xdd, 0x4f, 0x07, 0x05, 0xb7, 0xd2, 0x68, 0x1f, 0xb4, 0xf3, 0xdf, 0x2c, 0xc0, 0x15,
	0x4a, 0x45, 0xba, 0x6b, 0x08, 0x85, 0x85, 0xa1, 0x28, 0x4a, 0x69, 0x34, 0xad, 0x6d, 0xd7, 0xf6,
	0x9a, 0xf1, 0x78, 0x49, 0x7e, 0x81, 0xb9, 0xd2, 0x72, 0x2b, 0xe8, 0xcc, 0x76, 0x6d, 0x6f, 0xf9,
	0x60, 0x37, 0x7c, 0x60, 0xb7, 0x70, 0xaa, 0x14, 0x5e, 0x39, 0x38, 0xf6, 0x31, 0xe4, 0x4f, 0x58,
	0x1f, 0xe4, 0x56, 0x66, 0x82, 0x25, 0x83, 0xa2, 0x10, 0xda, 0x32, 0x91, 0x9b, 0xa4, 0x4f, 0x67,
	0xb7, 0x6b, 0x7b, 0x8b, 0x07, 0x1b, 0xa1, 0x3f, 0x64, 0x38, 0x3e, 0x64, 0x78, 0x52, 0x1d, 0x32,
	0x26, 0x3e, 0xec, 0xd8, 0x47, 0x9d, 0xba, 0x20, 0x72, 0x0a, 0xab, 0x95, 0x18, 0x57, 0xca, 0x0b,
	0x95, 0xb4, 0xfe, 0x98, 0xd2, 0x8a, 0x8f, 0x39, 0x54, 0x0a, 0x55, 0x4a, 0x12, 0xc2, 0x5a, 0xdf,
	0x58, 0x56, 0x88, 0xd2, 0xf2, 0xc2, 0xb2, 0xf1, 0xb5, 0xe7, 0xf0, 0xda, 0xab, 0x7d, 0x63, 0x63,
	0xef, 0xb9, 0xa9, 0x1e, 0xe0, 0x3d, 0xac, 0x27, 0x26, 0xcb, 0xb8, 0x4e, 0x99, 0x92, 0x5a, 0x30,
	0x93, 0x3b, 0xdd, 0x92, 0xce, 0xe3, 0xce, 0xaf, 0x1e, 0x7c, 0x8f, 0x63, 0x1f, 0xd0, 0x96, 0x5a,
	0x5c, 0x78, 0x3c, 0x26, 0xc9, 0x67, 0xb6, 0x9d, 0xb7, 0x30, 0x87, 0xcf, 0x45, 0x1a, 0x50, 0x6f,
	0x47, 0x37, 0xa7, 0xc1, 0x13, 0xb2, 0x04, 0x8d, 0x93, 0xf8, 0x30, 0x3a, 0x8f, 0xce, 0xdf, 0x06,
	0x35, 0xb2, 0x0e, 0xc1, 0x65, 0x7c, 0xca, 0xa2, 0xf3, 0xe8, 0x3a, 0x3a, 0x6c, 0x47, 0x1f, 0x9c,
	0x75, 0x86, 0x04, 0xb0, 0xf4, 0x89, 0x65, 0x76, 0xe7, 0x5f, 0x00, 0xf2, 0xf9, 0x9e, 0xe4, 0x19,
	0x2c, 0x74, 0x78, 0x29, 0x98, 0x4c, 0x31, 0xab, 0xf5, 0x78, 0xde, 0x2d, 0xa3, 0x94, 0x6c, 0xc3,
	0x62, 0x62, 0xb4, 0xcf, 0x49, 0x32, 0xc2, 0xd4, 0xb6, 0xe2, 0xfb, 0x26, 0xf2, 0x02, 0x89, 0xae,
	0xec, 0xb1, 0x9c, 0x5b, 0x9f, 0xb0, 0x66, 0x0c, 0xde, 0x74, 0xc9, 0x6d, 0xff, 0x1e, 0x30, 0xe2,
	0x99, 0xc2, 0x3c, 0x4c, 0x80, 0xf7, 0x3c, 0x53, 0xe4, 0x57, 0xd8, 0xe2, 0x4a, 0x99, 0x8f, 0x6c,
	0xa0, 0xff, 0xd6, 0xe6, 0xa3, 0x66, 0xae, 0x24, 0x64, 0xc2, 0xba, 0x52, 0xa8, 0xb4, 0xc4, 0xf7,
	0x6e, 0xc4, 0x14, 0x91, 0xbf, 0x3c, 0x71, 0x85, 0xc0, 0x19, 0xfa, 0xc9, 0x21, 0x3c, 0x2f, 0xc4,
	0xad, 0x48, 0xec, 0x24, 0x3e, 0x1d, 0x69, 0x9e, 0x4d, 0x05, 0x36, 0x51, 0x60, 0xd3, 0x43, 0x95,
	0xc2, 0x89, 0x47, 0x2a, 0x89, 0xd7, 0x40, 0x30, 0x2d, 0x8c, 0xa7, 0x69, 0x21, 0xca, 0xd2, 0x5f,
	0x65, 0x1e, 0x4f, 0x1a, 0xa0, 0xe7, 0xd0, 0x3b, 0xf0, 0x42, 0xb7, 0x40, 0x95, 0x49, 0xb8, 0x9a,
	0xd0, 0x32, 0x9f, 0x14, 0xc7, 0x02, 0xd6, 0xfe, 0x0f, 0x5f, 0x98, 0xeb, 0x30, 0xca, 0xab, 0xe2,
	0x89, 0x9f, 0xa2, 0x64, 0xb5, 0xcd, 0xc4, 0x4c, 0xb6, 0xa0, 0xa9, 0x4c, 0x8f, 0x29, 0x31, 0x14,
	0x8a, 0x36, 0xf0, 0x40, 0x0d, 0x65, 0x7a, 0x6d, 0xb7, 0x76, 0x05, 0x9a, 0x98, 0x2c, 0x37, 0xda,
	0xf5, 0xcb, 0x14, 0x6b, 0xfa, 0x02, 0x9d, 0xb8, 0xda, 0x63, 0xfe, 0x39, 0x80, 0xa3, 0xba, 0xa6,
	0xc8, 0xb8, 0xa5, 0x80, 0x98, 0x93, 0x3f, 0x43, 0x03, 0xd9, 0x00, 0x27, 0xed, 0xef, 0xbe, 0xe8,
	0x7b, 0x5b, 0x19, 0x9f, 0xc3, 0x57, 0xb0, 0xe2, 0xc6, 0x89, 0x4c, 0x04, 0x4b, 0xd4, 0xa0, 0xb4,
	0xa2, 0xa0, 0x2d, 0x24, 0x96, 0x2b, 0xf3, 0xb1, 0xb7, 0x92, 0x97, 0xb0, 0x34, 0x06, 0xb5, 0x49,
	0x05, 0x5d, 0x46, 0x6a, 0xb1, 0xb2, 0x9d, 0x9b, 0x54, 0xdc, 0x47, 0xfe, 0x31, 0x5a, 0xd0, 0x95,
	0x4f, 0x90, 0x0f, 0x46, 0x0b, 0x12, 0xc1, 0x5a, 0x57, 0x2a, 0xc1, 0xba, 0x6a, 0x50, 0xf6, 0x99,
	0xd4, 0x56, 0x14, 0x43, 0xae, 0x68, 0xf0, 0x58, 0x0b, 0xaf, 0xba, 0xa8, 0x33, 0x17, 0x14, 0x55,
	0x31, 0xe4, 0x67, 0x80, 0xb4, 0xe0, 0x52, 0x33, 0xd7, 0xdb, 0x74, 0xf5, 0x31, 0x85, 0x26, 0xc2,
	0xd7, 0x32, 0xc3, 0x91, 0x94, 0x73, 0x1c, 0x45, 0x65, 0x7f, 0x60, 0x53, 0x57, 0x58, 0xa8, 0x41,
	0x1e, 0x1d, 0x49, 0x3e, 0xec, 0xaa, 0x8a, 0x42, 0xb1, 0xdf, 0xa0, 0x9e, 0xb9, 0xf7, 0x58, 0xc3,
	0xfa, 0x78, 0xfd, 0xa5, 0xf5, 0xf1, 0xce, 0xa4, 0x22, 0xc6, 0x48, 0xf2, 0x02, 0x9a, 0x19, 0xbf,
	0xc3, 0xde, 0x28, 0xe9, 0xba, 0x6b, 0xd2, 0xa3, 0x19, 0x5a, 0x8b, 0x1b, 0x19, 0xbf, 0x73, 0xed,
	0x50, 0x92, 0x6f, 0x21, 0x70, 0x80, 0xe9, 0xdc, 0x32, 0xcd, 0x33, 0xc1, 0x94, 0xd0, 0xf4, 0xe9,
	0x84, 0x6b, 0x65, 0xfc, 0xee, 0xa2, 0x73, 0x7b, 0xce, 0x33, 0xd1, 0x16, 0xda, 0x95, 0x4e, 0x2a,
	0x4b, 0xde, 0x51, 0x82, 0xdd, 0x9b, 0x71, 0xf4, 0x2b, 0x6c, 0x95, 0xd5, 0xca, 0xf5, 0xfb, 0x64,
	0xc4, 0x91, 0xef, 0x61, 0x5d, 0x68, 0xc4, 0xb3, 0x81, 0x15, 0x77, 0xcc, 0x16, 0x3c, 0x91, 0xba,
	0x47, 0x9f, 0x61, 0x00, 0xf1, 0xbe, 0x77, 0xce, 0x75, 0xed, 0x3d, 0xe4, 0x1b, 0x68, 0x8d, 0x27,
	0xa7, 0x1f, 0xe5, 0x14, 0x67, 0xc7, 0x52, 0x65, 0xf4, 0x93, 0x7a, 0x17, 0x96, 0x93, 0x7c, 0x50,
	0x0a, 0xcb, 0x6c, 0xbf, 0x10, 0x3c, 0x2d, 0xe9, 0x06, 0x0a, 0xb6, 0xbc, 0xf5, 0xda, 0x1b, 0x77,
	0xb6, 0xa0, 0x39, 0x6d, 0x89, 0x79, 0x98, 0x19, 0xfe, 0x14, 0x3c, 0xc1, 0xef, 0x9b, 0xa0, 0xb6,
	0xf3, 0x1d, 0xd4, 0xdd, 0x33, 0x91, 0x26, 0xcc, 0xe1, 0xdf, 0xc5, 0xcf, 0xc6, 0x1b, 0xae, 0x64,
	0xca, 0xad, 0x08, 0x6a, 0x6e, 0x15, 0x69, 0x69, 0x2f, 0xb4, 0x1a, 0x05, 0x33, 0x7f, 0xd4, 0x1b,
	0x4b, 0x41, 0xeb, 0xe8, 0x0d, 0xbc, 0x94, 0xc6, 0x67, 0x21, 0x2f, 0xcc, 0xdd, 0xe8, 0xa1, 0x84,
	0x1c, 0xad, 0x4c, 0xff, 0x56, 0x97, 0x2e, 0xcb, 0x97, 0xb5, 0xce, 0x3c, 0xa6, 0xfb, 0xc7, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xe7, 0xb6, 0x90, 0x72, 0x07, 0x00, 0x00,
}
