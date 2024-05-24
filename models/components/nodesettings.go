package components

import (
	"encoding/json"
	"fmt"
)

type NodeSettingsProperties struct {
	// Property name
	Name string `json:"name"`
	// Property value (can be a string or JSON object)
	Value any `json:"value"`
}

func (o *NodeSettingsProperties) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *NodeSettingsProperties) GetValue() any {
	if o == nil {
		return nil
	}
	return o.Value
}

// NodeSettingsPolicyMode - In which mode the node will apply its configuration policy. Use `default` to use the global mode.
type NodeSettingsPolicyMode string

const (
	NodeSettingsPolicyModeAudit   NodeSettingsPolicyMode = "audit"
	NodeSettingsPolicyModeEnforce NodeSettingsPolicyMode = "enforce"
	NodeSettingsPolicyModeDefault NodeSettingsPolicyMode = "default"
)

func (e NodeSettingsPolicyMode) ToPointer() *NodeSettingsPolicyMode {
	return &e
}
func (e *NodeSettingsPolicyMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "audit":
		fallthrough
	case "enforce":
		fallthrough
	case "default":
		*e = NodeSettingsPolicyMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeSettingsPolicyMode: %v", v)
	}
}

// NodeSettingsState - The node life cycle state. See [dedicated doc](https://docs.rudder.io/reference/current/usage/advanced_node_management.html#node-lifecycle) for more information.
type NodeSettingsState string

const (
	NodeSettingsStateEnabled       NodeSettingsState = "enabled"
	NodeSettingsStateIgnored       NodeSettingsState = "ignored"
	NodeSettingsStateEmptyPolicies NodeSettingsState = "empty-policies"
	NodeSettingsStateInitializing  NodeSettingsState = "initializing"
	NodeSettingsStatePreparingEol  NodeSettingsState = "preparing-eol"
)

func (e NodeSettingsState) ToPointer() *NodeSettingsState {
	return &e
}
func (e *NodeSettingsState) UnmarshalJSON(data []byte) error {
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
		*e = NodeSettingsState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeSettingsState: %v", v)
	}
}

type NodeSettings struct {
	Properties []NodeSettingsProperties `json:"properties,omitempty"`
	// In which mode the node will apply its configuration policy. Use `default` to use the global mode.
	PolicyMode *NodeSettingsPolicyMode `json:"policyMode,omitempty"`
	// The node life cycle state. See [dedicated doc](https://docs.rudder.io/reference/current/usage/advanced_node_management.html#node-lifecycle) for more information.
	State *NodeSettingsState `json:"state,omitempty"`
	// Information about agent key or certificate
	AgentKey *AgentKey `json:"agentKey,omitempty"`
}

func (o *NodeSettings) GetProperties() []NodeSettingsProperties {
	if o == nil {
		return nil
	}
	return o.Properties
}

func (o *NodeSettings) GetPolicyMode() *NodeSettingsPolicyMode {
	if o == nil {
		return nil
	}
	return o.PolicyMode
}

func (o *NodeSettings) GetState() *NodeSettingsState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *NodeSettings) GetAgentKey() *AgentKey {
	if o == nil {
		return nil
	}
	return o.AgentKey
}
