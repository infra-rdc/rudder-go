package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// Value - New value of the allowed networks
type Value struct {
}

type SetAllowedNetworksRequestBody struct {
	// New value of the allowed networks
	Value *Value `json:"value,omitempty"`
}

func (o *SetAllowedNetworksRequestBody) GetValue() *Value {
	if o == nil {
		return nil
	}
	return o.Value
}

type SetAllowedNetworksRequest struct {
	// Policy server ID for which you want to manage allowed networks.
	NodeID      string                        `pathParam:"style=simple,explode=false,name=nodeId"`
	RequestBody SetAllowedNetworksRequestBody `request:"mediaType=application/json"`
}

func (o *SetAllowedNetworksRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

func (o *SetAllowedNetworksRequest) GetRequestBody() SetAllowedNetworksRequestBody {
	if o == nil {
		return SetAllowedNetworksRequestBody{}
	}
	return o.RequestBody
}

// SetAllowedNetworksResult - Result of the request
type SetAllowedNetworksResult string

const (
	SetAllowedNetworksResultSuccess SetAllowedNetworksResult = "success"
	SetAllowedNetworksResultError   SetAllowedNetworksResult = "error"
)

func (e SetAllowedNetworksResult) ToPointer() *SetAllowedNetworksResult {
	return &e
}
func (e *SetAllowedNetworksResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = SetAllowedNetworksResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SetAllowedNetworksResult: %v", v)
	}
}

// SetAllowedNetworksAction - The id of the action
type SetAllowedNetworksAction string

const (
	SetAllowedNetworksActionModifyAllowedNetworks SetAllowedNetworksAction = "modifyAllowedNetworks"
)

func (e SetAllowedNetworksAction) ToPointer() *SetAllowedNetworksAction {
	return &e
}
func (e *SetAllowedNetworksAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "modifyAllowedNetworks":
		*e = SetAllowedNetworksAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SetAllowedNetworksAction: %v", v)
	}
}

// SetAllowedNetworksData - Information about the allowed_networks settings
type SetAllowedNetworksData struct {
	// Array of allowed networks
	Settings []string `json:"settings"`
}

func (o *SetAllowedNetworksData) GetSettings() []string {
	if o == nil {
		return []string{}
	}
	return o.Settings
}

// SetAllowedNetworksResponseBody - Allowed Networks
type SetAllowedNetworksResponseBody struct {
	// Result of the request
	Result SetAllowedNetworksResult `json:"result"`
	// The id of the action
	Action SetAllowedNetworksAction `json:"action"`
	// The id of the modified node
	ID *string `json:"id,omitempty"`
	// Information about the allowed_networks settings
	Data SetAllowedNetworksData `json:"data"`
}

func (o *SetAllowedNetworksResponseBody) GetResult() SetAllowedNetworksResult {
	if o == nil {
		return SetAllowedNetworksResult("")
	}
	return o.Result
}

func (o *SetAllowedNetworksResponseBody) GetAction() SetAllowedNetworksAction {
	if o == nil {
		return SetAllowedNetworksAction("")
	}
	return o.Action
}

func (o *SetAllowedNetworksResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetAllowedNetworksResponseBody) GetData() SetAllowedNetworksData {
	if o == nil {
		return SetAllowedNetworksData{}
	}
	return o.Data
}

type SetAllowedNetworksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Allowed Networks
	Object *SetAllowedNetworksResponseBody
}

func (o *SetAllowedNetworksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SetAllowedNetworksResponse) GetObject() *SetAllowedNetworksResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
