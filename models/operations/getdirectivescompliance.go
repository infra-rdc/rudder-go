package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetDirectivesComplianceResult - Result of the request
type GetDirectivesComplianceResult string

const (
	GetDirectivesComplianceResultSuccess GetDirectivesComplianceResult = "success"
	GetDirectivesComplianceResultError   GetDirectivesComplianceResult = "error"
)

func (e GetDirectivesComplianceResult) ToPointer() *GetDirectivesComplianceResult {
	return &e
}
func (e *GetDirectivesComplianceResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetDirectivesComplianceResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDirectivesComplianceResult: %v", v)
	}
}

// GetDirectivesComplianceAction - The id of the action
type GetDirectivesComplianceAction string

const (
	GetDirectivesComplianceActionGetDirectiveComplianceID GetDirectivesComplianceAction = "getDirectiveComplianceId"
)

func (e GetDirectivesComplianceAction) ToPointer() *GetDirectivesComplianceAction {
	return &e
}
func (e *GetDirectivesComplianceAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getDirectiveComplianceId":
		*e = GetDirectivesComplianceAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDirectivesComplianceAction: %v", v)
	}
}

type Mode string

const (
	ModeFullCompliance  Mode = "full-compliance"
	ModeChangesOnly     Mode = "changes-only"
	ModeReportsDisabled Mode = "reports-disabled"
)

func (e Mode) ToPointer() *Mode {
	return &e
}
func (e *Mode) UnmarshalJSON(data []byte) error {
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
		*e = Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mode: %v", v)
	}
}

type GetDirectivesComplianceComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *GetDirectivesComplianceComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *GetDirectivesComplianceComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *GetDirectivesComplianceComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *GetDirectivesComplianceComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *GetDirectivesComplianceComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetDirectivesComplianceComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *GetDirectivesComplianceComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type DirectivesCompliance struct {
	// id of the directive
	ID string `json:"id"`
	// Name of the directive
	Name string `json:"name"`
	Mode Mode   `json:"mode"`
	// Directive compliance level
	Compliance        float32                                  `json:"compliance"`
	ComplianceDetails GetDirectivesComplianceComplianceDetails `json:"complianceDetails"`
	Rules             components.DirectiveRuleCompliance       `json:"rules"`
	Nodes             components.DirectiveNodeCompliance       `json:"nodes"`
}

func (o *DirectivesCompliance) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DirectivesCompliance) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DirectivesCompliance) GetMode() Mode {
	if o == nil {
		return Mode("")
	}
	return o.Mode
}

func (o *DirectivesCompliance) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *DirectivesCompliance) GetComplianceDetails() GetDirectivesComplianceComplianceDetails {
	if o == nil {
		return GetDirectivesComplianceComplianceDetails{}
	}
	return o.ComplianceDetails
}

func (o *DirectivesCompliance) GetRules() components.DirectiveRuleCompliance {
	if o == nil {
		return components.DirectiveRuleCompliance{}
	}
	return o.Rules
}

func (o *DirectivesCompliance) GetNodes() components.DirectiveNodeCompliance {
	if o == nil {
		return components.DirectiveNodeCompliance{}
	}
	return o.Nodes
}

type GetDirectivesComplianceData struct {
	DirectivesCompliance DirectivesCompliance `json:"directivesCompliance"`
}

func (o *GetDirectivesComplianceData) GetDirectivesCompliance() DirectivesCompliance {
	if o == nil {
		return DirectivesCompliance{}
	}
	return o.DirectivesCompliance
}

// GetDirectivesComplianceResponseBody - Success
type GetDirectivesComplianceResponseBody struct {
	// Result of the request
	Result GetDirectivesComplianceResult `json:"result"`
	// The id of the action
	Action GetDirectivesComplianceAction `json:"action"`
	Data   GetDirectivesComplianceData   `json:"data"`
}

func (o *GetDirectivesComplianceResponseBody) GetResult() GetDirectivesComplianceResult {
	if o == nil {
		return GetDirectivesComplianceResult("")
	}
	return o.Result
}

func (o *GetDirectivesComplianceResponseBody) GetAction() GetDirectivesComplianceAction {
	if o == nil {
		return GetDirectivesComplianceAction("")
	}
	return o.Action
}

func (o *GetDirectivesComplianceResponseBody) GetData() GetDirectivesComplianceData {
	if o == nil {
		return GetDirectivesComplianceData{}
	}
	return o.Data
}

type GetDirectivesComplianceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetDirectivesComplianceResponseBody
}

func (o *GetDirectivesComplianceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetDirectivesComplianceResponse) GetObject() *GetDirectivesComplianceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
