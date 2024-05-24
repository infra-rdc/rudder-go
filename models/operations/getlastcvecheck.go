package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetLastCVECheckRequest struct {
	// Id of node groups you want to get from last CVE check
	GroupID *string `queryParam:"style=form,explode=true,name=groupId"`
	// Id of nodes you want to get from last CVE check
	NodeID *string `queryParam:"style=form,explode=true,name=nodeId"`
	// Id of CVE you want to get from last CVE check
	CveID *string `queryParam:"style=form,explode=true,name=cveId"`
	// Name of packages you want to get from last CVE check
	Package *string `queryParam:"style=form,explode=true,name=package"`
	// Severity of the CVE you want to get from last CVE check
	Severity *components.CveSeverity `queryParam:"style=form,explode=true,name=severity"`
}

func (o *GetLastCVECheckRequest) GetGroupID() *string {
	if o == nil {
		return nil
	}
	return o.GroupID
}

func (o *GetLastCVECheckRequest) GetNodeID() *string {
	if o == nil {
		return nil
	}
	return o.NodeID
}

func (o *GetLastCVECheckRequest) GetCveID() *string {
	if o == nil {
		return nil
	}
	return o.CveID
}

func (o *GetLastCVECheckRequest) GetPackage() *string {
	if o == nil {
		return nil
	}
	return o.Package
}

func (o *GetLastCVECheckRequest) GetSeverity() *components.CveSeverity {
	if o == nil {
		return nil
	}
	return o.Severity
}

// GetLastCVECheckResult - Result of the request
type GetLastCVECheckResult string

const (
	GetLastCVECheckResultSuccess GetLastCVECheckResult = "success"
	GetLastCVECheckResultError   GetLastCVECheckResult = "error"
)

func (e GetLastCVECheckResult) ToPointer() *GetLastCVECheckResult {
	return &e
}
func (e *GetLastCVECheckResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetLastCVECheckResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetLastCVECheckResult: %v", v)
	}
}

// GetLastCVECheckAction - The id of the action
type GetLastCVECheckAction string

const (
	GetLastCVECheckActionGetLastCveCheck GetLastCVECheckAction = "getLastCVECheck"
)

func (e GetLastCVECheckAction) ToPointer() *GetLastCVECheckAction {
	return &e
}
func (e *GetLastCVECheckAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getLastCVECheck":
		*e = GetLastCVECheckAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetLastCVECheckAction: %v", v)
	}
}

type GetLastCVECheckData struct {
	CVEChecks []components.CveCheck `json:"CVEChecks"`
}

func (o *GetLastCVECheckData) GetCVEChecks() []components.CveCheck {
	if o == nil {
		return []components.CveCheck{}
	}
	return o.CVEChecks
}

// GetLastCVECheckResponseBody - Last CVE check
type GetLastCVECheckResponseBody struct {
	// Result of the request
	Result GetLastCVECheckResult `json:"result"`
	// The id of the action
	Action GetLastCVECheckAction `json:"action"`
	Data   GetLastCVECheckData   `json:"data"`
}

func (o *GetLastCVECheckResponseBody) GetResult() GetLastCVECheckResult {
	if o == nil {
		return GetLastCVECheckResult("")
	}
	return o.Result
}

func (o *GetLastCVECheckResponseBody) GetAction() GetLastCVECheckAction {
	if o == nil {
		return GetLastCVECheckAction("")
	}
	return o.Action
}

func (o *GetLastCVECheckResponseBody) GetData() GetLastCVECheckData {
	if o == nil {
		return GetLastCVECheckData{}
	}
	return o.Data
}

type GetLastCVECheckResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Last CVE check
	Object *GetLastCVECheckResponseBody
}

func (o *GetLastCVECheckResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetLastCVECheckResponse) GetObject() *GetLastCVECheckResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
