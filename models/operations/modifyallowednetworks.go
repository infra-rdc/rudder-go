package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type AllowedNetworks struct {
	// List of networks to add to existing allowed networks for that server
	Add []string `json:"add,omitempty"`
	// List of networks to remove from existing allowed networks for that server
	Delete []string `json:"delete,omitempty"`
}

func (o *AllowedNetworks) GetAdd() []string {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *AllowedNetworks) GetDelete() []string {
	if o == nil {
		return nil
	}
	return o.Delete
}

type ModifyAllowedNetworksRequestBody struct {
	AllowedNetworks *AllowedNetworks `json:"allowed_networks,omitempty"`
}

func (o *ModifyAllowedNetworksRequestBody) GetAllowedNetworks() *AllowedNetworks {
	if o == nil {
		return nil
	}
	return o.AllowedNetworks
}

type ModifyAllowedNetworksRequest struct {
	// Policy server ID for which you want to manage allowed networks.
	NodeID      string                           `pathParam:"style=simple,explode=false,name=nodeId"`
	RequestBody ModifyAllowedNetworksRequestBody `request:"mediaType=application/json"`
}

func (o *ModifyAllowedNetworksRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

func (o *ModifyAllowedNetworksRequest) GetRequestBody() ModifyAllowedNetworksRequestBody {
	if o == nil {
		return ModifyAllowedNetworksRequestBody{}
	}
	return o.RequestBody
}

// ModifyAllowedNetworksResult - Result of the request
type ModifyAllowedNetworksResult string

const (
	ModifyAllowedNetworksResultSuccess ModifyAllowedNetworksResult = "success"
	ModifyAllowedNetworksResultError   ModifyAllowedNetworksResult = "error"
)

func (e ModifyAllowedNetworksResult) ToPointer() *ModifyAllowedNetworksResult {
	return &e
}
func (e *ModifyAllowedNetworksResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ModifyAllowedNetworksResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ModifyAllowedNetworksResult: %v", v)
	}
}

// ModifyAllowedNetworksAction - The id of the action
type ModifyAllowedNetworksAction string

const (
	ModifyAllowedNetworksActionModifyDiffAllowedNetworks ModifyAllowedNetworksAction = "modifyDiffAllowedNetworks"
)

func (e ModifyAllowedNetworksAction) ToPointer() *ModifyAllowedNetworksAction {
	return &e
}
func (e *ModifyAllowedNetworksAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "modifyDiffAllowedNetworks":
		*e = ModifyAllowedNetworksAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ModifyAllowedNetworksAction: %v", v)
	}
}

// ModifyAllowedNetworksData - Information about the allowed_networks settings
type ModifyAllowedNetworksData struct {
	// Array of allowed networks
	Settings []string `json:"settings"`
}

func (o *ModifyAllowedNetworksData) GetSettings() []string {
	if o == nil {
		return []string{}
	}
	return o.Settings
}

// ModifyAllowedNetworksResponseBody - Allowed Networks
type ModifyAllowedNetworksResponseBody struct {
	// Result of the request
	Result ModifyAllowedNetworksResult `json:"result"`
	// The id of the action
	Action ModifyAllowedNetworksAction `json:"action"`
	// Information about the allowed_networks settings
	Data ModifyAllowedNetworksData `json:"data"`
}

func (o *ModifyAllowedNetworksResponseBody) GetResult() ModifyAllowedNetworksResult {
	if o == nil {
		return ModifyAllowedNetworksResult("")
	}
	return o.Result
}

func (o *ModifyAllowedNetworksResponseBody) GetAction() ModifyAllowedNetworksAction {
	if o == nil {
		return ModifyAllowedNetworksAction("")
	}
	return o.Action
}

func (o *ModifyAllowedNetworksResponseBody) GetData() ModifyAllowedNetworksData {
	if o == nil {
		return ModifyAllowedNetworksData{}
	}
	return o.Data
}

type ModifyAllowedNetworksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Allowed Networks
	Object *ModifyAllowedNetworksResponseBody
}

func (o *ModifyAllowedNetworksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ModifyAllowedNetworksResponse) GetObject() *ModifyAllowedNetworksResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
