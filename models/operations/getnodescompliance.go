package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type GetNodesComplianceRequest struct {
	// Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports)
	Level *int64 `default:"10" queryParam:"style=form,explode=true,name=level"`
	// Number of digits after comma in compliance percent figures
	Precision *int64 `default:"2" queryParam:"style=form,explode=true,name=precision"`
}

func (g GetNodesComplianceRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetNodesComplianceRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetNodesComplianceRequest) GetLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *GetNodesComplianceRequest) GetPrecision() *int64 {
	if o == nil {
		return nil
	}
	return o.Precision
}

// GetNodesComplianceResult - Result of the request
type GetNodesComplianceResult string

const (
	GetNodesComplianceResultSuccess GetNodesComplianceResult = "success"
	GetNodesComplianceResultError   GetNodesComplianceResult = "error"
)

func (e GetNodesComplianceResult) ToPointer() *GetNodesComplianceResult {
	return &e
}
func (e *GetNodesComplianceResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetNodesComplianceResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNodesComplianceResult: %v", v)
	}
}

// GetNodesComplianceAction - The id of the action
type GetNodesComplianceAction string

const (
	GetNodesComplianceActionGetNodesCompliance GetNodesComplianceAction = "getNodesCompliance"
)

func (e GetNodesComplianceAction) ToPointer() *GetNodesComplianceAction {
	return &e
}
func (e *GetNodesComplianceAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getNodesCompliance":
		*e = GetNodesComplianceAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNodesComplianceAction: %v", v)
	}
}

type GetNodesComplianceMode string

const (
	GetNodesComplianceModeFullCompliance  GetNodesComplianceMode = "full-compliance"
	GetNodesComplianceModeChangesOnly     GetNodesComplianceMode = "changes-only"
	GetNodesComplianceModeReportsDisabled GetNodesComplianceMode = "reports-disabled"
)

func (e GetNodesComplianceMode) ToPointer() *GetNodesComplianceMode {
	return &e
}
func (e *GetNodesComplianceMode) UnmarshalJSON(data []byte) error {
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
		*e = GetNodesComplianceMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNodesComplianceMode: %v", v)
	}
}

type GetNodesComplianceComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *GetNodesComplianceComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *GetNodesComplianceComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *GetNodesComplianceComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *GetNodesComplianceComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *GetNodesComplianceComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetNodesComplianceComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *GetNodesComplianceComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type Nodes struct {
	// id of the node
	ID   string                 `json:"id"`
	Mode GetNodesComplianceMode `json:"mode"`
	// Rule compliance level
	Compliance        float32                             `json:"compliance"`
	ComplianceDetails GetNodesComplianceComplianceDetails `json:"complianceDetails"`
}

func (o *Nodes) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Nodes) GetMode() GetNodesComplianceMode {
	if o == nil {
		return GetNodesComplianceMode("")
	}
	return o.Mode
}

func (o *Nodes) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *Nodes) GetComplianceDetails() GetNodesComplianceComplianceDetails {
	if o == nil {
		return GetNodesComplianceComplianceDetails{}
	}
	return o.ComplianceDetails
}

type GetNodesComplianceData struct {
	Nodes []Nodes `json:"nodes"`
}

func (o *GetNodesComplianceData) GetNodes() []Nodes {
	if o == nil {
		return []Nodes{}
	}
	return o.Nodes
}

// GetNodesComplianceResponseBody - Success
type GetNodesComplianceResponseBody struct {
	// Result of the request
	Result GetNodesComplianceResult `json:"result"`
	// The id of the action
	Action GetNodesComplianceAction `json:"action"`
	Data   GetNodesComplianceData   `json:"data"`
}

func (o *GetNodesComplianceResponseBody) GetResult() GetNodesComplianceResult {
	if o == nil {
		return GetNodesComplianceResult("")
	}
	return o.Result
}

func (o *GetNodesComplianceResponseBody) GetAction() GetNodesComplianceAction {
	if o == nil {
		return GetNodesComplianceAction("")
	}
	return o.Action
}

func (o *GetNodesComplianceResponseBody) GetData() GetNodesComplianceData {
	if o == nil {
		return GetNodesComplianceData{}
	}
	return o.Data
}

type GetNodesComplianceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetNodesComplianceResponseBody
}

func (o *GetNodesComplianceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetNodesComplianceResponse) GetObject() *GetNodesComplianceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
