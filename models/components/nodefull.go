package components

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/types"
)

// NodeFullStatus - Status of the node
type NodeFullStatus string

const (
	NodeFullStatusPending  NodeFullStatus = "pending"
	NodeFullStatusAccepted NodeFullStatus = "accepted"
	NodeFullStatusDeleted  NodeFullStatus = "deleted"
)

func (e NodeFullStatus) ToPointer() *NodeFullStatus {
	return &e
}
func (e *NodeFullStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "accepted":
		fallthrough
	case "deleted":
		*e = NodeFullStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeFullStatus: %v", v)
	}
}

// NodeFullType - Type of the machine
type NodeFullType string

const (
	NodeFullTypePhysical NodeFullType = "Physical"
	NodeFullTypeVirtual  NodeFullType = "Virtual"
)

func (e NodeFullType) ToPointer() *NodeFullType {
	return &e
}
func (e *NodeFullType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Physical":
		fallthrough
	case "Virtual":
		*e = NodeFullType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeFullType: %v", v)
	}
}

// Machine - Information about the underlying machine
type Machine struct {
	// Rudder unique identifier for the machine
	ID *string `json:"id,omitempty"`
	// Type of the machine
	Type *NodeFullType `json:"type,omitempty"`
	// In the case of VM, the VM technology
	Provider *string `json:"provider,omitempty"`
	// Information about machine manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`
	// If available, a unique identifier provided by the machine
	SerialNumber *string `json:"serialNumber,omitempty"`
}

func (o *Machine) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Machine) GetType() *NodeFullType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Machine) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *Machine) GetManufacturer() *string {
	if o == nil {
		return nil
	}
	return o.Manufacturer
}

func (o *Machine) GetSerialNumber() *string {
	if o == nil {
		return nil
	}
	return o.SerialNumber
}

// NodeFullOsType - Family of the OS
type NodeFullOsType string

const (
	NodeFullOsTypeWindows NodeFullOsType = "Windows"
	NodeFullOsTypeLinux   NodeFullOsType = "Linux"
	NodeFullOsTypeAix     NodeFullOsType = "AIX"
	NodeFullOsTypeFreeBsd NodeFullOsType = "FreeBSD"
)

func (e NodeFullOsType) ToPointer() *NodeFullOsType {
	return &e
}
func (e *NodeFullOsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Windows":
		fallthrough
	case "Linux":
		fallthrough
	case "AIX":
		fallthrough
	case "FreeBSD":
		*e = NodeFullOsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeFullOsType: %v", v)
	}
}

// NodeFullOs - Information about the operating system
type NodeFullOs struct {
	// Family of the OS
	Type NodeFullOsType `json:"type"`
	// Operating system name (distribution on Linux, etc.)
	Name string `json:"name"`
	// OS version
	Version string `json:"version"`
	// Full operating system name
	FullName string `json:"fullName"`
	// If relevant, the service pack of the OS
	ServicePack *string `json:"servicePack,omitempty"`
	// Version of the OS kernel
	KernelVersion string `json:"kernelVersion"`
}

func (o *NodeFullOs) GetType() NodeFullOsType {
	if o == nil {
		return NodeFullOsType("")
	}
	return o.Type
}

func (o *NodeFullOs) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *NodeFullOs) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}

func (o *NodeFullOs) GetFullName() string {
	if o == nil {
		return ""
	}
	return o.FullName
}

func (o *NodeFullOs) GetServicePack() *string {
	if o == nil {
		return nil
	}
	return o.ServicePack
}

func (o *NodeFullOs) GetKernelVersion() string {
	if o == nil {
		return ""
	}
	return o.KernelVersion
}

type ManagementTechnology struct {
	// Agent name
	Name string `json:"name"`
	// Agent version
	Version *string `json:"version,omitempty"`
	// List of agent capabilities
	Capabilities []string `json:"capabilities,omitempty"`
	// kind of node for the management engine, like `root`, `relay`, `node`, `root-component`
	NodeKind *string `json:"nodeKind,omitempty"`
	// Roles fulfilled by the agent
	RootComponents []string `json:"rootComponents,omitempty"`
}

func (o *ManagementTechnology) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ManagementTechnology) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *ManagementTechnology) GetCapabilities() []string {
	if o == nil {
		return nil
	}
	return o.Capabilities
}

func (o *ManagementTechnology) GetNodeKind() *string {
	if o == nil {
		return nil
	}
	return o.NodeKind
}

func (o *ManagementTechnology) GetRootComponents() []string {
	if o == nil {
		return nil
	}
	return o.RootComponents
}

type NodeFullProperties struct {
	// Property name
	Name string `json:"name"`
	// Property value (can be a string or JSON object)
	Value any `json:"value"`
}

func (o *NodeFullProperties) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *NodeFullProperties) GetValue() any {
	if o == nil {
		return nil
	}
	return o.Value
}

// NodeFullPolicyMode - Rudder policy mode for the node (`default` follows the global configuration)
type NodeFullPolicyMode string

const (
	NodeFullPolicyModeEnforce NodeFullPolicyMode = "enforce"
	NodeFullPolicyModeAudit   NodeFullPolicyMode = "audit"
	NodeFullPolicyModeDefault NodeFullPolicyMode = "default"
)

func (e NodeFullPolicyMode) ToPointer() *NodeFullPolicyMode {
	return &e
}
func (e *NodeFullPolicyMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enforce":
		fallthrough
	case "audit":
		fallthrough
	case "default":
		*e = NodeFullPolicyMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeFullPolicyMode: %v", v)
	}
}

type NodeFullTimezone struct {
	// Timezone name
	Name string `json:"name"`
	// Timezone offset to UTC
	Offset *string `json:"offset,omitempty"`
}

func (o *NodeFullTimezone) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *NodeFullTimezone) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

// Bios - BIOS information
type Bios struct {
	// BIOS name
	Name *string `json:"name,omitempty"`
	// BIOS version
	Version *string `json:"version,omitempty"`
	// BIOS editor
	Editor *string `json:"editor,omitempty"`
	// Number of BIOS on the machine
	Quantity *int64 `json:"quantity,omitempty"`
	// Release date of the BIOS
	ReleaseDate *string `json:"releaseDate,omitempty"`
	// System provided description of the BIOS (long name)
	Description *string `json:"description,omitempty"`
}

func (o *Bios) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Bios) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *Bios) GetEditor() *string {
	if o == nil {
		return nil
	}
	return o.Editor
}

func (o *Bios) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Bios) GetReleaseDate() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *Bios) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type Controllers struct {
	// Controller name
	Name *string `json:"name,omitempty"`
	// Controller type
	Type *string `json:"type,omitempty"`
	// Quantity of that controller
	Quantity *int64 `json:"quantity,omitempty"`
	// System provided description of the controller
	Description *string `json:"description,omitempty"`
	// Controller manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`
}

func (o *Controllers) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Controllers) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Controllers) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Controllers) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Controllers) GetManufacturer() *string {
	if o == nil {
		return nil
	}
	return o.Manufacturer
}

type EnvironmentVariables struct {
	// Environment variable name
	Name *string `json:"name,omitempty"`
	// Environment variable value
	Value *string `json:"value,omitempty"`
}

func (o *EnvironmentVariables) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *EnvironmentVariables) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type FileSystems struct {
	// Type of file system (`ext4`, `swap`, etc.)
	Name *string `json:"name,omitempty"`
	// Mount point
	MountPoint *string `json:"mountPoint,omitempty"`
	// Description of the file system
	Description *string `json:"description,omitempty"`
	// Number of files
	FileCount *int64 `json:"fileCount,omitempty"`
	// Free space remaining
	FreeSpace *int64 `json:"freeSpace,omitempty"`
	// Total space
	TotalSpace *int64 `json:"totalSpace,omitempty"`
}

func (o *FileSystems) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *FileSystems) GetMountPoint() *string {
	if o == nil {
		return nil
	}
	return o.MountPoint
}

func (o *FileSystems) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *FileSystems) GetFileCount() *int64 {
	if o == nil {
		return nil
	}
	return o.FileCount
}

func (o *FileSystems) GetFreeSpace() *int64 {
	if o == nil {
		return nil
	}
	return o.FreeSpace
}

func (o *FileSystems) GetTotalSpace() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalSpace
}

// ManagementTechnologyDetails - Additional information about management technology
type ManagementTechnologyDetails struct {
	// Certificates used by the agent
	CfengineKeys []string `json:"cfengineKeys,omitempty"`
	// Local user account used by the agent
	CfengineUser *string `json:"cfengineUser,omitempty"`
}

func (o *ManagementTechnologyDetails) GetCfengineKeys() []string {
	if o == nil {
		return nil
	}
	return o.CfengineKeys
}

func (o *ManagementTechnologyDetails) GetCfengineUser() *string {
	if o == nil {
		return nil
	}
	return o.CfengineUser
}

// Memories - Memory slots
type Memories struct {
	// Name of the memory slot or memory module
	Name *string `json:"name,omitempty"`
	// Memory speed (frequency)
	Speed *int64 `json:"speed,omitempty"`
	// Memory slot type
	Type *string `json:"type,omitempty"`
	// Manufacturer provided information about the module
	Caption *string `json:"caption,omitempty"`
	// Number of modules in that slot
	Quantity *int64 `json:"quantity,omitempty"`
	// Size of modules
	Capacity *int64 `json:"capacity,omitempty"`
	// Slot number
	SlotNumber *int64 `json:"slotNumber,omitempty"`
	// System provided description
	Description *string `json:"description,omitempty"`
	// Serial number of the module
	SerialNumber *string `json:"serialNumber,omitempty"`
}

func (o *Memories) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Memories) GetSpeed() *int64 {
	if o == nil {
		return nil
	}
	return o.Speed
}

func (o *Memories) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Memories) GetCaption() *string {
	if o == nil {
		return nil
	}
	return o.Caption
}

func (o *Memories) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Memories) GetCapacity() *int64 {
	if o == nil {
		return nil
	}
	return o.Capacity
}

func (o *Memories) GetSlotNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.SlotNumber
}

func (o *Memories) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Memories) GetSerialNumber() *string {
	if o == nil {
		return nil
	}
	return o.SerialNumber
}

type NetworkInterfaces struct {
	// Interface name (for ex "eth0")
	Name *string  `json:"name,omitempty"`
	Mask []string `json:"mask,omitempty"`
	// Information about the type of network interface
	Type *string `json:"type,omitempty"`
	// Information about synchronization speed
	Speed *string `json:"speed,omitempty"`
	// network interface status (enabled or not, up or down)
	Status *string `json:"status,omitempty"`
	// DHCP server managing that network interface
	DhcpServer *string `json:"dhcpServer,omitempty"`
	// MAC address of the network interface
	MacAddress *string `json:"macAddress,omitempty"`
	// IP addresses of the network interface
	IPAddresses []string `json:"ipAddresses,omitempty"`
}

func (o *NetworkInterfaces) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *NetworkInterfaces) GetMask() []string {
	if o == nil {
		return nil
	}
	return o.Mask
}

func (o *NetworkInterfaces) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *NetworkInterfaces) GetSpeed() *string {
	if o == nil {
		return nil
	}
	return o.Speed
}

func (o *NetworkInterfaces) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *NetworkInterfaces) GetDhcpServer() *string {
	if o == nil {
		return nil
	}
	return o.DhcpServer
}

func (o *NetworkInterfaces) GetMacAddress() *string {
	if o == nil {
		return nil
	}
	return o.MacAddress
}

func (o *NetworkInterfaces) GetIPAddresses() []string {
	if o == nil {
		return nil
	}
	return o.IPAddresses
}

type Ports struct {
	// Port name
	Name *string `json:"name,omitempty"`
	// Port type
	Type *string `json:"type,omitempty"`
	// Quantity of similar ports
	Quantity *int64 `json:"quantity,omitempty"`
	// System provided description of the port
	Description *string `json:"description,omitempty"`
}

func (o *Ports) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Ports) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Ports) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Ports) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

// Processes - Process information
type Processes struct {
	// PID of the process
	Pid *int64 `json:"pid,omitempty"`
	// TTY to which the process is
	Tty *string `json:"tty,omitempty"`
	// Process name
	Name *string `json:"name,omitempty"`
	// User account who started the process
	User *string `json:"user,omitempty"`
	// Started date and time of the process
	Started *types.Date `json:"started,omitempty"`
	// Memory allocated to the process (at inventory time)
	Memory *float32 `json:"memory,omitempty"`
	// Virtual memory allocated to the process (at inventory time)
	VirtualMemory *int64 `json:"virtualMemory,omitempty"`
	// CPU used by the process (at inventory time)
	CPUUsage *int64 `json:"cpuUsage,omitempty"`
	// System provided description about the process
	Description *string `json:"description,omitempty"`
}

func (p Processes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Processes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Processes) GetPid() *int64 {
	if o == nil {
		return nil
	}
	return o.Pid
}

func (o *Processes) GetTty() *string {
	if o == nil {
		return nil
	}
	return o.Tty
}

func (o *Processes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Processes) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *Processes) GetStarted() *types.Date {
	if o == nil {
		return nil
	}
	return o.Started
}

func (o *Processes) GetMemory() *float32 {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *Processes) GetVirtualMemory() *int64 {
	if o == nil {
		return nil
	}
	return o.VirtualMemory
}

func (o *Processes) GetCPUUsage() *int64 {
	if o == nil {
		return nil
	}
	return o.CPUUsage
}

func (o *Processes) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type Processors struct {
	// CPU name
	Name *string `json:"name,omitempty"`
	// CPU architecture
	Arch *string `json:"arch,omitempty"`
	// CPU model
	Model *int64 `json:"model,omitempty"`
	// CPU family
	FamilyName *string `json:"familyName,omitempty"`
	// Number of core for that CPU
	Core *int64 `json:"core,omitempty"`
	// Speed (frequency) of the CPU
	Speed *int64 `json:"speed,omitempty"`
	// Number of thread by core for the CPU
	Thread *int64 `json:"thread,omitempty"`
	// Stepping or power management information
	Stepping *int64 `json:"stepping,omitempty"`
	// CPU manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Number of CPU with these features
	Quantity *int64 `json:"quantity,omitempty"`
	// Identifier of the CPU
	Cpuid *string `json:"cpuid,omitempty"`
	// External clock used by the CPU
	ExternalClock *string `json:"externalClock,omitempty"`
	// System provided description of the CPU
	Description *string `json:"description,omitempty"`
}

func (o *Processors) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Processors) GetArch() *string {
	if o == nil {
		return nil
	}
	return o.Arch
}

func (o *Processors) GetModel() *int64 {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *Processors) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *Processors) GetCore() *int64 {
	if o == nil {
		return nil
	}
	return o.Core
}

func (o *Processors) GetSpeed() *int64 {
	if o == nil {
		return nil
	}
	return o.Speed
}

func (o *Processors) GetThread() *int64 {
	if o == nil {
		return nil
	}
	return o.Thread
}

func (o *Processors) GetStepping() *int64 {
	if o == nil {
		return nil
	}
	return o.Stepping
}

func (o *Processors) GetManufacturer() *string {
	if o == nil {
		return nil
	}
	return o.Manufacturer
}

func (o *Processors) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Processors) GetCpuid() *string {
	if o == nil {
		return nil
	}
	return o.Cpuid
}

func (o *Processors) GetExternalClock() *string {
	if o == nil {
		return nil
	}
	return o.ExternalClock
}

func (o *Processors) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type Slots struct {
	// Slot name or identifier
	Name *string `json:"name,omitempty"`
	// Slot status (used, powered, etc)
	Status *string `json:"status,omitempty"`
	// Quantity of similar slots
	Quantity *int64 `json:"quantity,omitempty"`
	// System provided description of the slots
	Description *string `json:"description,omitempty"`
}

func (o *Slots) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Slots) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Slots) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Slots) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

// License - Information about the license
type License struct {
	// Is this an OEM license (and information)
	Oem *string `json:"oem,omitempty"`
	// License name
	Name *string `json:"name,omitempty"`
	// License product identifier
	ProductID *string `json:"productId,omitempty"`
	// License key
	ProductKey *string `json:"productKey,omitempty"`
	// License description
	Description *string `json:"description,omitempty"`
	// License expiration date
	ExpirationDate *types.Date `json:"expirationDate,omitempty"`
}

func (l License) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *License) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *License) GetOem() *string {
	if o == nil {
		return nil
	}
	return o.Oem
}

func (o *License) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *License) GetProductID() *string {
	if o == nil {
		return nil
	}
	return o.ProductID
}

func (o *License) GetProductKey() *string {
	if o == nil {
		return nil
	}
	return o.ProductKey
}

func (o *License) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *License) GetExpirationDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ExpirationDate
}

type Software struct {
	// Name of the software (as reported by the node)
	Name *string `json:"name,omitempty"`
	// Version of the software
	Version *string `json:"version,omitempty"`
	// Editor of the software
	Editor *string `json:"editor,omitempty"`
	// A description of the software
	Description *string `json:"description,omitempty"`
	// Release date of the software
	ReleaseDate *types.Date `json:"releaseDate,omitempty"`
	// Information about the license
	License *License `json:"license,omitempty"`
}

func (s Software) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Software) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Software) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Software) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *Software) GetEditor() *string {
	if o == nil {
		return nil
	}
	return o.Editor
}

func (o *Software) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Software) GetReleaseDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *Software) GetLicense() *License {
	if o == nil {
		return nil
	}
	return o.License
}

// Kind - if available, kind of patch provided by that update, else none
type Kind string

const (
	KindNone        Kind = "none"
	KindSecurity    Kind = "security"
	KindDefect      Kind = "defect"
	KindEnhancement Kind = "enhancement"
	KindOther       Kind = "other"
)

func (e Kind) ToPointer() *Kind {
	return &e
}
func (e *Kind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "security":
		fallthrough
	case "defect":
		fallthrough
	case "enhancement":
		fallthrough
	case "other":
		*e = Kind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Kind: %v", v)
	}
}

// Severity - if available, the severity of the update
type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityHigh     Severity = "high"
	SeverityModerate Severity = "moderate"
	SeverityLow      Severity = "low"
	SeverityOther    Severity = "other"
)

func (e Severity) ToPointer() *Severity {
	return &e
}
func (e *Severity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "critical":
		fallthrough
	case "high":
		fallthrough
	case "moderate":
		fallthrough
	case "low":
		fallthrough
	case "other":
		*e = Severity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Severity: %v", v)
	}
}

type SoftwareUpdate struct {
	// name of software that can be updated
	Name *string `json:"name,omitempty"`
	// available version for software
	Version *string `json:"version,omitempty"`
	// CPU architecture of the update
	Arch *string `json:"arch,omitempty"`
	// tool that discovered that update
	From *string `json:"from,omitempty"`
	// if available, kind of patch provided by that update, else none
	Kind *Kind `json:"kind,omitempty"`
	// information about the source providing that update
	Source *string `json:"source,omitempty"`
	// details about the content of the update, if available
	Description *string `json:"description,omitempty"`
	// if available, the severity of the update
	Severity *Severity `json:"severity,omitempty"`
	Ids      []string  `json:"ids,omitempty"`
}

func (o *SoftwareUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SoftwareUpdate) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *SoftwareUpdate) GetArch() *string {
	if o == nil {
		return nil
	}
	return o.Arch
}

func (o *SoftwareUpdate) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *SoftwareUpdate) GetKind() *Kind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *SoftwareUpdate) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *SoftwareUpdate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SoftwareUpdate) GetSeverity() *Severity {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *SoftwareUpdate) GetIds() []string {
	if o == nil {
		return nil
	}
	return o.Ids
}

type Sound struct {
	// Sound card name
	Name *string `json:"name,omitempty"`
	// Quantity of similar sound cards
	Quantity *int64 `json:"quantity,omitempty"`
	// System provided description of the sound card
	Description *string `json:"description,omitempty"`
}

func (o *Sound) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Sound) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Sound) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type Storage struct {
	// Storage name
	Name *string `json:"name,omitempty"`
	// Storage type
	Type *string `json:"type,omitempty"`
	// Storage size in MB
	Size *int64 `json:"size,omitempty"`
	// Storage model
	Model *string `json:"model,omitempty"`
	// Storage firmware information
	Firmware *string `json:"firmware,omitempty"`
	// Quantity of similar storage
	Quantity *int64 `json:"quantity,omitempty"`
	// System provided information about the storage
	Description *string `json:"description,omitempty"`
	// Storage manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Storage serial number
	SerialNumber *string `json:"serialNumber,omitempty"`
}

func (o *Storage) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Storage) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Storage) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *Storage) GetModel() *string {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *Storage) GetFirmware() *string {
	if o == nil {
		return nil
	}
	return o.Firmware
}

func (o *Storage) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Storage) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Storage) GetManufacturer() *string {
	if o == nil {
		return nil
	}
	return o.Manufacturer
}

func (o *Storage) GetSerialNumber() *string {
	if o == nil {
		return nil
	}
	return o.SerialNumber
}

type Videos struct {
	// Video card name
	Name *string `json:"name,omitempty"`
	// Quantity of memory for that video card
	Memory *string `json:"memory,omitempty"`
	// information about video card chipset
	Chipset *string `json:"chipset,omitempty"`
	// Quantity of similar video cards
	Quantity *int64 `json:"quantity,omitempty"`
	// Resolution used by that video card at inventory time
	Resolution *string `json:"resolution,omitempty"`
	// System provided description for that video card
	Description *string `json:"description,omitempty"`
}

func (o *Videos) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Videos) GetMemory() *string {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *Videos) GetChipset() *string {
	if o == nil {
		return nil
	}
	return o.Chipset
}

func (o *Videos) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Videos) GetResolution() *string {
	if o == nil {
		return nil
	}
	return o.Resolution
}

func (o *Videos) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type VirtualMachines struct {
	// Name of the hosted virtual machine
	Name *string `json:"name,omitempty"`
	// Type of the hosted virtual machine
	Type *string `json:"type,omitempty"`
	// Unique identifier of the hosted virtual machine
	UUID *string `json:"uuid,omitempty"`
	// Number of virtual CPU allocated to the hosted virtual machine
	Vcpu *string `json:"vcpu,omitempty"`
	// Owner of the hosted virtual machine
	Owner *string `json:"owner,omitempty"`
	// Status (up, starting, etc) of the hosted virtual machine
	Status *string `json:"status,omitempty"`
	// Memory allocated to the hosted virtual machine
	Memory *string `json:"memory,omitempty"`
	// Technology of the hosted virtual machine
	Subsystem *string `json:"subsystem,omitempty"`
	// System provided description of the hosted virtual machine
	Description *string `json:"description,omitempty"`
}

func (o *VirtualMachines) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *VirtualMachines) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *VirtualMachines) GetUUID() *string {
	if o == nil {
		return nil
	}
	return o.UUID
}

func (o *VirtualMachines) GetVcpu() *string {
	if o == nil {
		return nil
	}
	return o.Vcpu
}

func (o *VirtualMachines) GetOwner() *string {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *VirtualMachines) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *VirtualMachines) GetMemory() *string {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *VirtualMachines) GetSubsystem() *string {
	if o == nil {
		return nil
	}
	return o.Subsystem
}

func (o *VirtualMachines) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type NodeFull struct {
	// Unique id of the node
	ID string `json:"id"`
	// Fully qualified name of the node
	Hostname string `json:"hostname"`
	// Status of the node
	Status NodeFullStatus `json:"status"`
	// Information about CPU architecture (32/64 bits)
	ArchitectureDescription *string `json:"architectureDescription,omitempty"`
	// A human intended description of the node (not used)
	Description *string `json:"description,omitempty"`
	// IP addresses of the node (IPv4 and IPv6)
	IPAddresses []string `json:"ipAddresses"`
	// Date and time of the latest run, if one is available (node time). Up to API v11, format was: "YYYY-MM-dd HH:mm"
	LastRunDate *types.Date `json:"lastRunDate,omitempty"`
	// Date and time of the latest generated inventory, if one is available (node time). Up to API v11, format was: "YYYY-MM-dd HH:mm"
	LastInventoryDate *types.Date `json:"lastInventoryDate,omitempty"`
	// Information about the underlying machine
	Machine *Machine `json:"machine,omitempty"`
	// Information about the operating system
	Os *NodeFullOs `json:"os,omitempty"`
	// Management agents running on the node
	ManagementTechnology []ManagementTechnology `json:"managementTechnology"`
	// Rudder policy server managing the node
	PolicyServerID string `json:"policyServerId"`
	// Node properties (either set by user or filled by third party sources)
	Properties []NodeFullProperties `json:"properties"`
	// Rudder policy mode for the node (`default` follows the global configuration)
	PolicyMode *NodeFullPolicyMode `json:"policyMode,omitempty"`
	// Size of RAM in MB
	RAM      *int64            `json:"ram,omitempty"`
	Timezone *NodeFullTimezone `json:"timezone,omitempty"`
	// User accounts declared in the node
	Accounts []string `json:"accounts,omitempty"`
	// BIOS information
	Bios *Bios `json:"bios,omitempty"`
	// Physical controller information
	Controllers []Controllers `json:"controllers,omitempty"`
	// Environment variables defined on the node in the context of the agent
	EnvironmentVariables []EnvironmentVariables `json:"environmentVariables,omitempty"`
	// File system declared on the node
	FileSystems []FileSystems `json:"fileSystems,omitempty"`
	// Additional information about management technology
	ManagementTechnologyDetails *ManagementTechnologyDetails `json:"managementTechnologyDetails,omitempty"`
	// Memory slots
	Memories []Memories `json:"memories,omitempty"`
	// Detailed information about registered network interfaces on the node
	NetworkInterfaces []NetworkInterfaces `json:"networkInterfaces,omitempty"`
	// Physical port information objects
	Ports []Ports `json:"ports,omitempty"`
	// Process running (at inventory time)
	Processes []Processes `json:"processes,omitempty"`
	// CPU information
	Processors []Processors `json:"processors,omitempty"`
	// Physical extension slot information
	Slots []Slots `json:"slots,omitempty"`
	// Software installed on the node (can have thousands items)
	Software []Software `json:"software,omitempty"`
	// Software that can be updated on that machine
	SoftwareUpdate []SoftwareUpdate `json:"softwareUpdate,omitempty"`
	// Sound card information
	Sound []Sound `json:"sound,omitempty"`
	// Storage (disks) information objects
	Storage []Storage `json:"storage,omitempty"`
	// Video card information
	Videos []Videos `json:"videos,omitempty"`
	// Hosted virtual machine information
	VirtualMachines []VirtualMachines `json:"virtualMachines,omitempty"`
}

func (n NodeFull) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NodeFull) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NodeFull) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *NodeFull) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *NodeFull) GetStatus() NodeFullStatus {
	if o == nil {
		return NodeFullStatus("")
	}
	return o.Status
}

func (o *NodeFull) GetArchitectureDescription() *string {
	if o == nil {
		return nil
	}
	return o.ArchitectureDescription
}

func (o *NodeFull) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *NodeFull) GetIPAddresses() []string {
	if o == nil {
		return []string{}
	}
	return o.IPAddresses
}

func (o *NodeFull) GetLastRunDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.LastRunDate
}

func (o *NodeFull) GetLastInventoryDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.LastInventoryDate
}

func (o *NodeFull) GetMachine() *Machine {
	if o == nil {
		return nil
	}
	return o.Machine
}

func (o *NodeFull) GetOs() *NodeFullOs {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *NodeFull) GetManagementTechnology() []ManagementTechnology {
	if o == nil {
		return []ManagementTechnology{}
	}
	return o.ManagementTechnology
}

func (o *NodeFull) GetPolicyServerID() string {
	if o == nil {
		return ""
	}
	return o.PolicyServerID
}

func (o *NodeFull) GetProperties() []NodeFullProperties {
	if o == nil {
		return []NodeFullProperties{}
	}
	return o.Properties
}

func (o *NodeFull) GetPolicyMode() *NodeFullPolicyMode {
	if o == nil {
		return nil
	}
	return o.PolicyMode
}

func (o *NodeFull) GetRAM() *int64 {
	if o == nil {
		return nil
	}
	return o.RAM
}

func (o *NodeFull) GetTimezone() *NodeFullTimezone {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *NodeFull) GetAccounts() []string {
	if o == nil {
		return nil
	}
	return o.Accounts
}

func (o *NodeFull) GetBios() *Bios {
	if o == nil {
		return nil
	}
	return o.Bios
}

func (o *NodeFull) GetControllers() []Controllers {
	if o == nil {
		return nil
	}
	return o.Controllers
}

func (o *NodeFull) GetEnvironmentVariables() []EnvironmentVariables {
	if o == nil {
		return nil
	}
	return o.EnvironmentVariables
}

func (o *NodeFull) GetFileSystems() []FileSystems {
	if o == nil {
		return nil
	}
	return o.FileSystems
}

func (o *NodeFull) GetManagementTechnologyDetails() *ManagementTechnologyDetails {
	if o == nil {
		return nil
	}
	return o.ManagementTechnologyDetails
}

func (o *NodeFull) GetMemories() []Memories {
	if o == nil {
		return nil
	}
	return o.Memories
}

func (o *NodeFull) GetNetworkInterfaces() []NetworkInterfaces {
	if o == nil {
		return nil
	}
	return o.NetworkInterfaces
}

func (o *NodeFull) GetPorts() []Ports {
	if o == nil {
		return nil
	}
	return o.Ports
}

func (o *NodeFull) GetProcesses() []Processes {
	if o == nil {
		return nil
	}
	return o.Processes
}

func (o *NodeFull) GetProcessors() []Processors {
	if o == nil {
		return nil
	}
	return o.Processors
}

func (o *NodeFull) GetSlots() []Slots {
	if o == nil {
		return nil
	}
	return o.Slots
}

func (o *NodeFull) GetSoftware() []Software {
	if o == nil {
		return nil
	}
	return o.Software
}

func (o *NodeFull) GetSoftwareUpdate() []SoftwareUpdate {
	if o == nil {
		return nil
	}
	return o.SoftwareUpdate
}

func (o *NodeFull) GetSound() []Sound {
	if o == nil {
		return nil
	}
	return o.Sound
}

func (o *NodeFull) GetStorage() []Storage {
	if o == nil {
		return nil
	}
	return o.Storage
}

func (o *NodeFull) GetVideos() []Videos {
	if o == nil {
		return nil
	}
	return o.Videos
}

func (o *NodeFull) GetVirtualMachines() []VirtualMachines {
	if o == nil {
		return nil
	}
	return o.VirtualMachines
}
