package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type NodeDetailsRequest struct {
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
	// Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don't provide a base level, they will be added to `default` level, so if you only want os details, use `minimal,os` as the value for this parameter.
	// * **minimal** includes: `id`, `hostname` and `status`
	// * **default** includes **minimal** plus `architectureDescription`, `description`, `ipAddresses`, `lastRunDate`, `lastInventoryDate`, `machine`, `os`, `managementTechnology`, `policyServerId`, `properties` (be careful! Only node own properties, if you also need inherited properties, look at the dedicated `/nodes/{id}/inheritedProperties` endpoint), `policyMode `, `ram` and `timezone`
	// * **full** includes: **default** plus `accounts`, `bios`, `controllers`, `environmentVariables`, `fileSystems`, `managementTechnologyDetails`, `memories`, `networkInterfaces`, `ports`, `processes`, `processors`, `slots`, `software`, `softwareUpdate`, `sound`, `storage`, `videos` and `virtualMachines`
	Include *string `default:"default" queryParam:"style=form,explode=true,name=include"`
}

func (n NodeDetailsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NodeDetailsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NodeDetailsRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

func (o *NodeDetailsRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

// NodeDetailsResult - Result of the request
type NodeDetailsResult string

const (
	NodeDetailsResultSuccess NodeDetailsResult = "success"
	NodeDetailsResultError   NodeDetailsResult = "error"
)

func (e NodeDetailsResult) ToPointer() *NodeDetailsResult {
	return &e
}
func (e *NodeDetailsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = NodeDetailsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeDetailsResult: %v", v)
	}
}

// NodeDetailsAction - The id of the action
type NodeDetailsAction string

const (
	NodeDetailsActionNodeDetails NodeDetailsAction = "nodeDetails"
)

func (e NodeDetailsAction) ToPointer() *NodeDetailsAction {
	return &e
}
func (e *NodeDetailsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "nodeDetails":
		*e = NodeDetailsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeDetailsAction: %v", v)
	}
}

// NodeDetailsData - Information about the node
type NodeDetailsData struct {
	Nodes []components.NodeFull `json:"nodes"`
}

func (o *NodeDetailsData) GetNodes() []components.NodeFull {
	if o == nil {
		return []components.NodeFull{}
	}
	return o.Nodes
}

// NodeDetailsResponseBody - Nodes
type NodeDetailsResponseBody struct {
	// Result of the request
	Result NodeDetailsResult `json:"result"`
	// The id of the action
	Action NodeDetailsAction `json:"action"`
	// Information about the node
	Data NodeDetailsData `json:"data"`
}

func (o *NodeDetailsResponseBody) GetResult() NodeDetailsResult {
	if o == nil {
		return NodeDetailsResult("")
	}
	return o.Result
}

func (o *NodeDetailsResponseBody) GetAction() NodeDetailsAction {
	if o == nil {
		return NodeDetailsAction("")
	}
	return o.Action
}

func (o *NodeDetailsResponseBody) GetData() NodeDetailsData {
	if o == nil {
		return NodeDetailsData{}
	}
	return o.Data
}

type NodeDetailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Nodes
	Object *NodeDetailsResponseBody
}

func (o *NodeDetailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *NodeDetailsResponse) GetObject() *NodeDetailsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
