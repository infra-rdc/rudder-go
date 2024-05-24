package components

import (
	"encoding/json"
	"fmt"
)

// AgentKeyStatus - Certification status of the security token (reset to `undefined` to trust a new certificate). If `certified`, inventory signature check will be enforced
type AgentKeyStatus string

const (
	AgentKeyStatusCertified AgentKeyStatus = "certified"
	AgentKeyStatusUndefined AgentKeyStatus = "undefined"
)

func (e AgentKeyStatus) ToPointer() *AgentKeyStatus {
	return &e
}
func (e *AgentKeyStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "certified":
		fallthrough
	case "undefined":
		*e = AgentKeyStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AgentKeyStatus: %v", v)
	}
}

// AgentKey - Information about agent key or certificate
type AgentKey struct {
	// Certificate (or public key for <6.0 agents) used by the agent. Be careful write a "\n" after header line and before footer line, JSON does not keep formatting in string.
	Value string `json:"value"`
	// Certification status of the security token (reset to `undefined` to trust a new certificate). If `certified`, inventory signature check will be enforced
	Status *AgentKeyStatus `json:"status,omitempty"`
}

func (o *AgentKey) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *AgentKey) GetStatus() *AgentKeyStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
