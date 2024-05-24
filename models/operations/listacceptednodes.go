package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type ListAcceptedNodesRequest struct {
	// Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don't provide a base level, they will be added to `default` level, so if you only want os details, use `minimal,os` as the value for this parameter.
	// * **minimal** includes: `id`, `hostname` and `status`
	// * **default** includes **minimal** plus `architectureDescription`, `description`, `ipAddresses`, `lastRunDate`, `lastInventoryDate`, `machine`, `os`, `managementTechnology`, `policyServerId`, `properties` (be careful! Only node own properties, if you also need inherited properties, look at the dedicated `/nodes/{id}/inheritedProperties` endpoint), `policyMode `, `ram` and `timezone`
	// * **full** includes: **default** plus `accounts`, `bios`, `controllers`, `environmentVariables`, `fileSystems`, `managementTechnologyDetails`, `memories`, `networkInterfaces`, `ports`, `processes`, `processors`, `slots`, `software`, `softwareUpdate`, `sound`, `storage`, `videos` and `virtualMachines`
	Include *string `default:"default" queryParam:"style=form,explode=true,name=include"`
	// The criterion you want to find for your nodes. Replaces the `where`, `composition` and `select` parameters in a single parameter.
	Query *components.NodeQuery `queryParam:"serialization=json,name=query"`
	// The criterion you want to find for your nodes
	Where []components.NodeWhere `queryParam:"serialization=json,name=where"`
	// Boolean operator to use between each  `where` criteria.
	Composition *components.NodeComposition `default:"and" queryParam:"style=form,explode=true,name=composition"`
	// What kind of data we want to include. Here we can get policy servers/relay by setting `nodeAndPolicyServer`. Only used if `where` is defined.
	Select *string `default:"node" queryParam:"style=form,explode=true,name=select"`
}

func (l ListAcceptedNodesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAcceptedNodesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAcceptedNodesRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *ListAcceptedNodesRequest) GetQuery() *components.NodeQuery {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListAcceptedNodesRequest) GetWhere() []components.NodeWhere {
	if o == nil {
		return nil
	}
	return o.Where
}

func (o *ListAcceptedNodesRequest) GetComposition() *components.NodeComposition {
	if o == nil {
		return nil
	}
	return o.Composition
}

func (o *ListAcceptedNodesRequest) GetSelect() *string {
	if o == nil {
		return nil
	}
	return o.Select
}

// ListAcceptedNodesResult - Result of the request
type ListAcceptedNodesResult string

const (
	ListAcceptedNodesResultSuccess ListAcceptedNodesResult = "success"
	ListAcceptedNodesResultError   ListAcceptedNodesResult = "error"
)

func (e ListAcceptedNodesResult) ToPointer() *ListAcceptedNodesResult {
	return &e
}
func (e *ListAcceptedNodesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListAcceptedNodesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAcceptedNodesResult: %v", v)
	}
}

// ListAcceptedNodesAction - The id of the action
type ListAcceptedNodesAction string

const (
	ListAcceptedNodesActionListAcceptedNodes ListAcceptedNodesAction = "listAcceptedNodes"
)

func (e ListAcceptedNodesAction) ToPointer() *ListAcceptedNodesAction {
	return &e
}
func (e *ListAcceptedNodesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listAcceptedNodes":
		*e = ListAcceptedNodesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAcceptedNodesAction: %v", v)
	}
}

// ListAcceptedNodesData - Information about the nodes
type ListAcceptedNodesData struct {
	Nodes []components.NodeFull `json:"nodes"`
}

func (o *ListAcceptedNodesData) GetNodes() []components.NodeFull {
	if o == nil {
		return []components.NodeFull{}
	}
	return o.Nodes
}

// ListAcceptedNodesResponseBody - Nodes
type ListAcceptedNodesResponseBody struct {
	// Result of the request
	Result ListAcceptedNodesResult `json:"result"`
	// The id of the action
	Action ListAcceptedNodesAction `json:"action"`
	// Information about the nodes
	Data ListAcceptedNodesData `json:"data"`
}

func (o *ListAcceptedNodesResponseBody) GetResult() ListAcceptedNodesResult {
	if o == nil {
		return ListAcceptedNodesResult("")
	}
	return o.Result
}

func (o *ListAcceptedNodesResponseBody) GetAction() ListAcceptedNodesAction {
	if o == nil {
		return ListAcceptedNodesAction("")
	}
	return o.Action
}

func (o *ListAcceptedNodesResponseBody) GetData() ListAcceptedNodesData {
	if o == nil {
		return ListAcceptedNodesData{}
	}
	return o.Data
}

type ListAcceptedNodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Nodes
	Object *ListAcceptedNodesResponseBody
}

func (o *ListAcceptedNodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAcceptedNodesResponse) GetObject() *ListAcceptedNodesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
