package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetAllowedNetworksRequest struct {
	// Policy server ID for which you want to manage allowed networks.
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
}

func (o *GetAllowedNetworksRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

// GetAllowedNetworksResult - Result of the request
type GetAllowedNetworksResult string

const (
	GetAllowedNetworksResultSuccess GetAllowedNetworksResult = "success"
	GetAllowedNetworksResultError   GetAllowedNetworksResult = "error"
)

func (e GetAllowedNetworksResult) ToPointer() *GetAllowedNetworksResult {
	return &e
}
func (e *GetAllowedNetworksResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetAllowedNetworksResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllowedNetworksResult: %v", v)
	}
}

// GetAllowedNetworksAction - The id of the action
type GetAllowedNetworksAction string

const (
	GetAllowedNetworksActionGetAllowedNetworks GetAllowedNetworksAction = "getAllowedNetworks"
)

func (e GetAllowedNetworksAction) ToPointer() *GetAllowedNetworksAction {
	return &e
}
func (e *GetAllowedNetworksAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getAllowedNetworks":
		*e = GetAllowedNetworksAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllowedNetworksAction: %v", v)
	}
}

// GetAllowedNetworksData - Information about the allowed_networks settings
type GetAllowedNetworksData struct {
	// Array of allowed networks
	AllowedNetwork []string `json:"allowed_network"`
}

func (o *GetAllowedNetworksData) GetAllowedNetwork() []string {
	if o == nil {
		return []string{}
	}
	return o.AllowedNetwork
}

// GetAllowedNetworksResponseBody - Allowed Networks
type GetAllowedNetworksResponseBody struct {
	// Target policy server ID
	ID string `json:"id"`
	// Result of the request
	Result GetAllowedNetworksResult `json:"result"`
	// The id of the action
	Action GetAllowedNetworksAction `json:"action"`
	// Information about the allowed_networks settings
	Data GetAllowedNetworksData `json:"data"`
}

func (o *GetAllowedNetworksResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetAllowedNetworksResponseBody) GetResult() GetAllowedNetworksResult {
	if o == nil {
		return GetAllowedNetworksResult("")
	}
	return o.Result
}

func (o *GetAllowedNetworksResponseBody) GetAction() GetAllowedNetworksAction {
	if o == nil {
		return GetAllowedNetworksAction("")
	}
	return o.Action
}

func (o *GetAllowedNetworksResponseBody) GetData() GetAllowedNetworksData {
	if o == nil {
		return GetAllowedNetworksData{}
	}
	return o.Data
}

type GetAllowedNetworksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Allowed Networks
	Object *GetAllowedNetworksResponseBody
}

func (o *GetAllowedNetworksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAllowedNetworksResponse) GetObject() *GetAllowedNetworksResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
