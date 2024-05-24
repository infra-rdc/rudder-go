package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type ListPendingNodesRequest struct {
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

func (l ListPendingNodesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPendingNodesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPendingNodesRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *ListPendingNodesRequest) GetQuery() *components.NodeQuery {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListPendingNodesRequest) GetWhere() []components.NodeWhere {
	if o == nil {
		return nil
	}
	return o.Where
}

func (o *ListPendingNodesRequest) GetComposition() *components.NodeComposition {
	if o == nil {
		return nil
	}
	return o.Composition
}

func (o *ListPendingNodesRequest) GetSelect() *string {
	if o == nil {
		return nil
	}
	return o.Select
}

// ListPendingNodesResult - Result of the request
type ListPendingNodesResult string

const (
	ListPendingNodesResultSuccess ListPendingNodesResult = "success"
	ListPendingNodesResultError   ListPendingNodesResult = "error"
)

func (e ListPendingNodesResult) ToPointer() *ListPendingNodesResult {
	return &e
}
func (e *ListPendingNodesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListPendingNodesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPendingNodesResult: %v", v)
	}
}

// ListPendingNodesAction - The id of the action
type ListPendingNodesAction string

const (
	ListPendingNodesActionListPendingNodes ListPendingNodesAction = "listPendingNodes"
)

func (e ListPendingNodesAction) ToPointer() *ListPendingNodesAction {
	return &e
}
func (e *ListPendingNodesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listPendingNodes":
		*e = ListPendingNodesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPendingNodesAction: %v", v)
	}
}

// ListPendingNodesData - Information about the nodes
type ListPendingNodesData struct {
	Nodes []components.NodeFull `json:"nodes"`
}

func (o *ListPendingNodesData) GetNodes() []components.NodeFull {
	if o == nil {
		return []components.NodeFull{}
	}
	return o.Nodes
}

// ListPendingNodesResponseBody - Nodes
type ListPendingNodesResponseBody struct {
	// Result of the request
	Result ListPendingNodesResult `json:"result"`
	// The id of the action
	Action ListPendingNodesAction `json:"action"`
	// Information about the nodes
	Data ListPendingNodesData `json:"data"`
}

func (o *ListPendingNodesResponseBody) GetResult() ListPendingNodesResult {
	if o == nil {
		return ListPendingNodesResult("")
	}
	return o.Result
}

func (o *ListPendingNodesResponseBody) GetAction() ListPendingNodesAction {
	if o == nil {
		return ListPendingNodesAction("")
	}
	return o.Action
}

func (o *ListPendingNodesResponseBody) GetData() ListPendingNodesData {
	if o == nil {
		return ListPendingNodesData{}
	}
	return o.Data
}

type ListPendingNodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Nodes
	Object *ListPendingNodesResponseBody
}

func (o *ListPendingNodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListPendingNodesResponse) GetObject() *ListPendingNodesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
