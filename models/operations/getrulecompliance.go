package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type GetRuleComplianceRequest struct {
	// Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports)
	Level *int64 `default:"10" queryParam:"style=form,explode=true,name=level"`
	// Number of digits after comma in compliance percent figures
	Precision *int64 `default:"2" queryParam:"style=form,explode=true,name=precision"`
	// Id of the target rule
	RuleID string `pathParam:"style=simple,explode=false,name=ruleId"`
}

func (g GetRuleComplianceRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRuleComplianceRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRuleComplianceRequest) GetLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *GetRuleComplianceRequest) GetPrecision() *int64 {
	if o == nil {
		return nil
	}
	return o.Precision
}

func (o *GetRuleComplianceRequest) GetRuleID() string {
	if o == nil {
		return ""
	}
	return o.RuleID
}

// GetRuleComplianceResult - Result of the request
type GetRuleComplianceResult string

const (
	GetRuleComplianceResultSuccess GetRuleComplianceResult = "success"
	GetRuleComplianceResultError   GetRuleComplianceResult = "error"
)

func (e GetRuleComplianceResult) ToPointer() *GetRuleComplianceResult {
	return &e
}
func (e *GetRuleComplianceResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetRuleComplianceResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRuleComplianceResult: %v", v)
	}
}

// GetRuleComplianceAction - The id of the action
type GetRuleComplianceAction string

const (
	GetRuleComplianceActionGetRuleCompliance GetRuleComplianceAction = "getRuleCompliance"
)

func (e GetRuleComplianceAction) ToPointer() *GetRuleComplianceAction {
	return &e
}
func (e *GetRuleComplianceAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getRuleCompliance":
		*e = GetRuleComplianceAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRuleComplianceAction: %v", v)
	}
}

type GetRuleComplianceMode string

const (
	GetRuleComplianceModeFullCompliance  GetRuleComplianceMode = "full-compliance"
	GetRuleComplianceModeChangesOnly     GetRuleComplianceMode = "changes-only"
	GetRuleComplianceModeReportsDisabled GetRuleComplianceMode = "reports-disabled"
)

func (e GetRuleComplianceMode) ToPointer() *GetRuleComplianceMode {
	return &e
}
func (e *GetRuleComplianceMode) UnmarshalJSON(data []byte) error {
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
		*e = GetRuleComplianceMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRuleComplianceMode: %v", v)
	}
}

type GetRuleComplianceComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *GetRuleComplianceComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *GetRuleComplianceComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *GetRuleComplianceComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *GetRuleComplianceComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *GetRuleComplianceComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetRuleComplianceComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *GetRuleComplianceComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type GetRuleComplianceRules struct {
	// id of the rule
	ID   string                `json:"id"`
	Mode GetRuleComplianceMode `json:"mode"`
	// Rule compliance level
	Compliance        float32                            `json:"compliance"`
	ComplianceDetails GetRuleComplianceComplianceDetails `json:"complianceDetails"`
}

func (o *GetRuleComplianceRules) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetRuleComplianceRules) GetMode() GetRuleComplianceMode {
	if o == nil {
		return GetRuleComplianceMode("")
	}
	return o.Mode
}

func (o *GetRuleComplianceRules) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *GetRuleComplianceRules) GetComplianceDetails() GetRuleComplianceComplianceDetails {
	if o == nil {
		return GetRuleComplianceComplianceDetails{}
	}
	return o.ComplianceDetails
}

type GetRuleComplianceData struct {
	Rules []GetRuleComplianceRules `json:"rules"`
}

func (o *GetRuleComplianceData) GetRules() []GetRuleComplianceRules {
	if o == nil {
		return []GetRuleComplianceRules{}
	}
	return o.Rules
}

// GetRuleComplianceResponseBody - Success
type GetRuleComplianceResponseBody struct {
	// Result of the request
	Result GetRuleComplianceResult `json:"result"`
	// The id of the action
	Action GetRuleComplianceAction `json:"action"`
	Data   GetRuleComplianceData   `json:"data"`
}

func (o *GetRuleComplianceResponseBody) GetResult() GetRuleComplianceResult {
	if o == nil {
		return GetRuleComplianceResult("")
	}
	return o.Result
}

func (o *GetRuleComplianceResponseBody) GetAction() GetRuleComplianceAction {
	if o == nil {
		return GetRuleComplianceAction("")
	}
	return o.Action
}

func (o *GetRuleComplianceResponseBody) GetData() GetRuleComplianceData {
	if o == nil {
		return GetRuleComplianceData{}
	}
	return o.Data
}

type GetRuleComplianceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetRuleComplianceResponseBody
}

func (o *GetRuleComplianceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRuleComplianceResponse) GetObject() *GetRuleComplianceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
