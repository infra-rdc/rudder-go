package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type GetNodeComplianceRequest struct {
	// Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports)
	Level *int64 `default:"10" queryParam:"style=form,explode=true,name=level"`
	// Number of digits after comma in compliance percent figures
	Precision *int64 `default:"2" queryParam:"style=form,explode=true,name=precision"`
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
}

func (g GetNodeComplianceRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetNodeComplianceRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetNodeComplianceRequest) GetLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *GetNodeComplianceRequest) GetPrecision() *int64 {
	if o == nil {
		return nil
	}
	return o.Precision
}

func (o *GetNodeComplianceRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

// GetNodeComplianceResult - Result of the request
type GetNodeComplianceResult string

const (
	GetNodeComplianceResultSuccess GetNodeComplianceResult = "success"
	GetNodeComplianceResultError   GetNodeComplianceResult = "error"
)

func (e GetNodeComplianceResult) ToPointer() *GetNodeComplianceResult {
	return &e
}
func (e *GetNodeComplianceResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetNodeComplianceResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNodeComplianceResult: %v", v)
	}
}

// GetNodeComplianceAction - The id of the action
type GetNodeComplianceAction string

const (
	GetNodeComplianceActionGetNodeCompliance GetNodeComplianceAction = "getNodeCompliance"
)

func (e GetNodeComplianceAction) ToPointer() *GetNodeComplianceAction {
	return &e
}
func (e *GetNodeComplianceAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getNodeCompliance":
		*e = GetNodeComplianceAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNodeComplianceAction: %v", v)
	}
}

type GetNodeComplianceMode string

const (
	GetNodeComplianceModeFullCompliance  GetNodeComplianceMode = "full-compliance"
	GetNodeComplianceModeChangesOnly     GetNodeComplianceMode = "changes-only"
	GetNodeComplianceModeReportsDisabled GetNodeComplianceMode = "reports-disabled"
)

func (e GetNodeComplianceMode) ToPointer() *GetNodeComplianceMode {
	return &e
}
func (e *GetNodeComplianceMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "full-compliance":
		fallthrough
	case "changes-only":
		fallthrough
	case "reports-disabled":
		*e = GetNodeComplianceMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNodeComplianceMode: %v", v)
	}
}

type GetNodeComplianceComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *GetNodeComplianceComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *GetNodeComplianceComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *GetNodeComplianceComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *GetNodeComplianceComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *GetNodeComplianceComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetNodeComplianceComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *GetNodeComplianceComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type GetNodeComplianceNodes struct {
	// id of the node
	ID   string                `json:"id"`
	Mode GetNodeComplianceMode `json:"mode"`
	// Rule compliance level
	Compliance        float32                            `json:"compliance"`
	ComplianceDetails GetNodeComplianceComplianceDetails `json:"complianceDetails"`
}

func (o *GetNodeComplianceNodes) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetNodeComplianceNodes) GetMode() GetNodeComplianceMode {
	if o == nil {
		return GetNodeComplianceMode("")
	}
	return o.Mode
}

func (o *GetNodeComplianceNodes) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *GetNodeComplianceNodes) GetComplianceDetails() GetNodeComplianceComplianceDetails {
	if o == nil {
		return GetNodeComplianceComplianceDetails{}
	}
	return o.ComplianceDetails
}

type GetNodeComplianceData struct {
	Nodes []GetNodeComplianceNodes `json:"nodes"`
}

func (o *GetNodeComplianceData) GetNodes() []GetNodeComplianceNodes {
	if o == nil {
		return []GetNodeComplianceNodes{}
	}
	return o.Nodes
}

// GetNodeComplianceResponseBody - Success
type GetNodeComplianceResponseBody struct {
	// Result of the request
	Result GetNodeComplianceResult `json:"result"`
	// The id of the action
	Action GetNodeComplianceAction `json:"action"`
	Data   GetNodeComplianceData   `json:"data"`
}

func (o *GetNodeComplianceResponseBody) GetResult() GetNodeComplianceResult {
	if o == nil {
		return GetNodeComplianceResult("")
	}
	return o.Result
}

func (o *GetNodeComplianceResponseBody) GetAction() GetNodeComplianceAction {
	if o == nil {
		return GetNodeComplianceAction("")
	}
	return o.Action
}

func (o *GetNodeComplianceResponseBody) GetData() GetNodeComplianceData {
	if o == nil {
		return GetNodeComplianceData{}
	}
	return o.Data
}

type GetNodeComplianceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetNodeComplianceResponseBody
}

func (o *GetNodeComplianceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetNodeComplianceResponse) GetObject() *GetNodeComplianceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
