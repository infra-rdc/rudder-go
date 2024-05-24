package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type GetRulesComplianceRequest struct {
	// Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports)
	Level *int64 `default:"10" queryParam:"style=form,explode=true,name=level"`
	// Number of digits after comma in compliance percent figures
	Precision *int64 `default:"2" queryParam:"style=form,explode=true,name=precision"`
}

func (g GetRulesComplianceRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRulesComplianceRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRulesComplianceRequest) GetLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *GetRulesComplianceRequest) GetPrecision() *int64 {
	if o == nil {
		return nil
	}
	return o.Precision
}

// GetRulesComplianceResult - Result of the request
type GetRulesComplianceResult string

const (
	GetRulesComplianceResultSuccess GetRulesComplianceResult = "success"
	GetRulesComplianceResultError   GetRulesComplianceResult = "error"
)

func (e GetRulesComplianceResult) ToPointer() *GetRulesComplianceResult {
	return &e
}
func (e *GetRulesComplianceResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetRulesComplianceResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesComplianceResult: %v", v)
	}
}

// GetRulesComplianceAction - The id of the action
type GetRulesComplianceAction string

const (
	GetRulesComplianceActionGetRulesCompliance GetRulesComplianceAction = "getRulesCompliance"
)

func (e GetRulesComplianceAction) ToPointer() *GetRulesComplianceAction {
	return &e
}
func (e *GetRulesComplianceAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getRulesCompliance":
		*e = GetRulesComplianceAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesComplianceAction: %v", v)
	}
}

type GetRulesComplianceMode string

const (
	GetRulesComplianceModeFullCompliance  GetRulesComplianceMode = "full-compliance"
	GetRulesComplianceModeChangesOnly     GetRulesComplianceMode = "changes-only"
	GetRulesComplianceModeReportsDisabled GetRulesComplianceMode = "reports-disabled"
)

func (e GetRulesComplianceMode) ToPointer() *GetRulesComplianceMode {
	return &e
}
func (e *GetRulesComplianceMode) UnmarshalJSON(data []byte) error {
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
		*e = GetRulesComplianceMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesComplianceMode: %v", v)
	}
}

type GetRulesComplianceComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *GetRulesComplianceComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *GetRulesComplianceComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *GetRulesComplianceComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *GetRulesComplianceComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *GetRulesComplianceComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetRulesComplianceComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *GetRulesComplianceComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type Rules struct {
	// id of the rule
	ID   string                 `json:"id"`
	Mode GetRulesComplianceMode `json:"mode"`
	// Rule compliance level
	Compliance        float32                             `json:"compliance"`
	ComplianceDetails GetRulesComplianceComplianceDetails `json:"complianceDetails"`
}

func (o *Rules) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Rules) GetMode() GetRulesComplianceMode {
	if o == nil {
		return GetRulesComplianceMode("")
	}
	return o.Mode
}

func (o *Rules) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *Rules) GetComplianceDetails() GetRulesComplianceComplianceDetails {
	if o == nil {
		return GetRulesComplianceComplianceDetails{}
	}
	return o.ComplianceDetails
}

type GetRulesComplianceData struct {
	Rules []Rules `json:"rules"`
}

func (o *GetRulesComplianceData) GetRules() []Rules {
	if o == nil {
		return []Rules{}
	}
	return o.Rules
}

// GetRulesComplianceResponseBody - Success
type GetRulesComplianceResponseBody struct {
	// Result of the request
	Result GetRulesComplianceResult `json:"result"`
	// The id of the action
	Action GetRulesComplianceAction `json:"action"`
	Data   GetRulesComplianceData   `json:"data"`
}

func (o *GetRulesComplianceResponseBody) GetResult() GetRulesComplianceResult {
	if o == nil {
		return GetRulesComplianceResult("")
	}
	return o.Result
}

func (o *GetRulesComplianceResponseBody) GetAction() GetRulesComplianceAction {
	if o == nil {
		return GetRulesComplianceAction("")
	}
	return o.Action
}

func (o *GetRulesComplianceResponseBody) GetData() GetRulesComplianceData {
	if o == nil {
		return GetRulesComplianceData{}
	}
	return o.Data
}

type GetRulesComplianceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetRulesComplianceResponseBody
}

func (o *GetRulesComplianceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRulesComplianceResponse) GetObject() *GetRulesComplianceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
