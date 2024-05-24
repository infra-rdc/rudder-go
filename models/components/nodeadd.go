package components

import (
	"encoding/json"
	"fmt"
)

// NodeAddStatus - Target status of the node
type NodeAddStatus string

const (
	NodeAddStatusAccepted NodeAddStatus = "accepted"
	NodeAddStatusPending  NodeAddStatus = "pending"
)

func (e NodeAddStatus) ToPointer() *NodeAddStatus {
	return &e
}
func (e *NodeAddStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "accepted":
		fallthrough
	case "pending":
		*e = NodeAddStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeAddStatus: %v", v)
	}
}

// MachineType - The kind of machine for the node (use vm for a generic VM)
type MachineType string

const (
	MachineTypePhysical    MachineType = "physical"
	MachineTypeVM          MachineType = "vm"
	MachineTypeAixlpar     MachineType = "aixlpar"
	MachineTypeBsdjail     MachineType = "bsdjail"
	MachineTypeHyperv      MachineType = "hyperv"
	MachineTypeQemu        MachineType = "qemu"
	MachineTypeSolariszone MachineType = "solariszone"
	MachineTypeVbox        MachineType = "vbox"
	MachineTypeVmware      MachineType = "vmware"
	MachineTypeXen         MachineType = "xen"
)

func (e MachineType) ToPointer() *MachineType {
	return &e
}
func (e *MachineType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "physical":
		fallthrough
	case "vm":
		fallthrough
	case "aixlpar":
		fallthrough
	case "bsdjail":
		fallthrough
	case "hyperv":
		fallthrough
	case "qemu":
		fallthrough
	case "solariszone":
		fallthrough
	case "vbox":
		fallthrough
	case "vmware":
		fallthrough
	case "xen":
		*e = MachineType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MachineType: %v", v)
	}
}

// NodeAddType - list of groups to include in rule application
type NodeAddType string

const (
	NodeAddTypePhysical NodeAddType = "Physical"
	NodeAddTypeVirtual  NodeAddType = "Virtual"
)

func (e NodeAddType) ToPointer() *NodeAddType {
	return &e
}
func (e *NodeAddType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Physical":
		fallthrough
	case "Virtual":
		*e = NodeAddType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeAddType: %v", v)
	}
}

// Provider - The kind of virtual machine for the node
type Provider string

const (
	ProviderAixlpar     Provider = "aixlpar"
	ProviderBsdjail     Provider = "bsdjail"
	ProviderHyperv      Provider = "hyperv"
	ProviderQemu        Provider = "qemu"
	ProviderSolariszone Provider = "solariszone"
	ProviderVbox        Provider = "vbox"
	ProviderVmware      Provider = "vmware"
	ProviderXen         Provider = "xen"
)

func (e Provider) ToPointer() *Provider {
	return &e
}
func (e *Provider) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aixlpar":
		fallthrough
	case "bsdjail":
		fallthrough
	case "hyperv":
		fallthrough
	case "qemu":
		fallthrough
	case "solariszone":
		fallthrough
	case "vbox":
		fallthrough
	case "vmware":
		fallthrough
	case "xen":
		*e = Provider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Provider: %v", v)
	}
}

// NodeAddMachine - The kind of machine for the node (use vm for a generic VM)
type NodeAddMachine struct {
	// list of groups to include in rule application
	Type NodeAddType `json:"type"`
	// The kind of virtual machine for the node
	Provider *Provider `json:"provider,omitempty"`
	// Manufacturer of the machine
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Serial number of the machine
	SerialNumber *string `json:"serialNumber,omitempty"`
}

func (o *NodeAddMachine) GetType() NodeAddType {
	if o == nil {
		return NodeAddType("")
	}
	return o.Type
}

func (o *NodeAddMachine) GetProvider() *Provider {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *NodeAddMachine) GetManufacturer() *string {
	if o == nil {
		return nil
	}
	return o.Manufacturer
}

func (o *NodeAddMachine) GetSerialNumber() *string {
	if o == nil {
		return nil
	}
	return o.SerialNumber
}

// State - Node lifecycle state. Can only be specified when status=accepted. If not specified, enable is used
type State string

const (
	StateEnabled       State = "enabled"
	StateIgnored       State = "ignored"
	StateEmptyPolicies State = "empty-policies"
	StateInitializing  State = "initializing"
	StatePreparingEol  State = "preparing-eol"
)

func (e State) ToPointer() *State {
	return &e
}
func (e *State) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enabled":
		fallthrough
	case "ignored":
		fallthrough
	case "empty-policies":
		fallthrough
	case "initializing":
		fallthrough
	case "preparing-eol":
		*e = State(v)
		return nil
	default:
		return fmt.Errorf("invalid value for State: %v", v)
	}
}

// NodeAddPolicyMode - The policy mode for the node. Can only be specified when status=accepted. If not specified, the default (global) mode will be used
type NodeAddPolicyMode string

const (
	NodeAddPolicyModeEnforce NodeAddPolicyMode = "enforce"
	NodeAddPolicyModeAudit   NodeAddPolicyMode = "audit"
)

func (e NodeAddPolicyMode) ToPointer() *NodeAddPolicyMode {
	return &e
}
func (e *NodeAddPolicyMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enforce":
		fallthrough
	case "audit":
		*e = NodeAddPolicyMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeAddPolicyMode: %v", v)
	}
}

type NodeAddProperties struct {
	// Property name
	Name string `json:"name"`
	// Property value (can be a string or JSON object)
	Value any `json:"value"`
}

func (o *NodeAddProperties) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *NodeAddProperties) GetValue() any {
	if o == nil {
		return nil
	}
	return o.Value
}

type NodeAdd struct {
	// The Rudder node unique identifier in /opt/rudder/etc/uuid.hive
	ID string `json:"id"`
	// The fully qualified name of the node
	Hostname string `json:"hostname"`
	// Target status of the node
	Status NodeAddStatus `json:"status"`
	Os     Os            `json:"os"`
	// The policy server ID for that node. By default, "root"
	PolicyServerID *string `json:"policyServerId,omitempty"`
	// The kind of machine for the node (use vm for a generic VM)
	MachineType *MachineType `json:"machineType,omitempty"`
	// The kind of machine for the node (use vm for a generic VM)
	Machine *NodeAddMachine `json:"machine,omitempty"`
	// Node lifecycle state. Can only be specified when status=accepted. If not specified, enable is used
	State *State `json:"state,omitempty"`
	// The policy mode for the node. Can only be specified when status=accepted. If not specified, the default (global) mode will be used
	PolicyMode *NodeAddPolicyMode `json:"policyMode,omitempty"`
	// Information about agent key or certificate
	AgentKey *AgentKey `json:"agentKey,omitempty"`
	// Node properties (either set by user or filled by third party sources)
	Properties []NodeAddProperties `json:"properties"`
	// an array of IPs.
	IPAddresses []string `json:"ipAddresses"`
	// Timezone information of the node.
	Timezone *Timezone `json:"timezone,omitempty"`
}

func (o *NodeAdd) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *NodeAdd) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *NodeAdd) GetStatus() NodeAddStatus {
	if o == nil {
		return NodeAddStatus("")
	}
	return o.Status
}

func (o *NodeAdd) GetOs() Os {
	if o == nil {
		return Os{}
	}
	return o.Os
}

func (o *NodeAdd) GetPolicyServerID() *string {
	if o == nil {
		return nil
	}
	return o.PolicyServerID
}

func (o *NodeAdd) GetMachineType() *MachineType {
	if o == nil {
		return nil
	}
	return o.MachineType
}

func (o *NodeAdd) GetMachine() *NodeAddMachine {
	if o == nil {
		return nil
	}
	return o.Machine
}

func (o *NodeAdd) GetState() *State {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *NodeAdd) GetPolicyMode() *NodeAddPolicyMode {
	if o == nil {
		return nil
	}
	return o.PolicyMode
}

func (o *NodeAdd) GetAgentKey() *AgentKey {
	if o == nil {
		return nil
	}
	return o.AgentKey
}

func (o *NodeAdd) GetProperties() []NodeAddProperties {
	if o == nil {
		return []NodeAddProperties{}
	}
	return o.Properties
}

func (o *NodeAdd) GetIPAddresses() []string {
	if o == nil {
		return []string{}
	}
	return o.IPAddresses
}

func (o *NodeAdd) GetTimezone() *Timezone {
	if o == nil {
		return nil
	}
	return o.Timezone
}
