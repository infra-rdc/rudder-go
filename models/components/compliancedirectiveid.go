package components

import (
	"encoding/json"
	"fmt"
)

// Result of the request
type Result string

const (
	ResultSuccess Result = "success"
	ResultError   Result = "error"
)

func (e Result) ToPointer() *Result {
	return &e
}
func (e *Result) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = Result(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Result: %v", v)
	}
}

// Action - The id of the action
type Action string

const (
	ActionGetDirectiveComplianceID Action = "getDirectiveComplianceId"
)

func (e Action) ToPointer() *Action {
	return &e
}
func (e *Action) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getDirectiveComplianceId":
		*e = Action(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Action: %v", v)
	}
}

type ComplianceDirectiveIDMode string

const (
	ComplianceDirectiveIDModeFullCompliance  ComplianceDirectiveIDMode = "full-compliance"
	ComplianceDirectiveIDModeChangesOnly     ComplianceDirectiveIDMode = "changes-only"
	ComplianceDirectiveIDModeReportsDisabled ComplianceDirectiveIDMode = "reports-disabled"
)

func (e ComplianceDirectiveIDMode) ToPointer() *ComplianceDirectiveIDMode {
	return &e
}
func (e *ComplianceDirectiveIDMode) UnmarshalJSON(data []byte) error {
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
		*e = ComplianceDirectiveIDMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ComplianceDirectiveIDMode: %v", v)
	}
}

type ComplianceDirectiveIDComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *ComplianceDirectiveIDComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *ComplianceDirectiveIDComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *ComplianceDirectiveIDComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *ComplianceDirectiveIDComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *ComplianceDirectiveIDComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ComplianceDirectiveIDComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *ComplianceDirectiveIDComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type DirectiveCompliance struct {
	// id of the directive
	ID string `json:"id"`
	// Name of the directive
	Name string                    `json:"name"`
	Mode ComplianceDirectiveIDMode `json:"mode"`
	// Directive compliance level
	Compliance        float32                                `json:"compliance"`
	ComplianceDetails ComplianceDirectiveIDComplianceDetails `json:"complianceDetails"`
	Rules             DirectiveRuleCompliance                `json:"rules"`
	Nodes             DirectiveNodeCompliance                `json:"nodes"`
}

func (o *DirectiveCompliance) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DirectiveCompliance) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DirectiveCompliance) GetMode() ComplianceDirectiveIDMode {
	if o == nil {
		return ComplianceDirectiveIDMode("")
	}
	return o.Mode
}

func (o *DirectiveCompliance) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *DirectiveCompliance) GetComplianceDetails() ComplianceDirectiveIDComplianceDetails {
	if o == nil {
		return ComplianceDirectiveIDComplianceDetails{}
	}
	return o.ComplianceDetails
}

func (o *DirectiveCompliance) GetRules() DirectiveRuleCompliance {
	if o == nil {
		return DirectiveRuleCompliance{}
	}
	return o.Rules
}

func (o *DirectiveCompliance) GetNodes() DirectiveNodeCompliance {
	if o == nil {
		return DirectiveNodeCompliance{}
	}
	return o.Nodes
}

type Data struct {
	DirectiveCompliance DirectiveCompliance `json:"directiveCompliance"`
}

func (o *Data) GetDirectiveCompliance() DirectiveCompliance {
	if o == nil {
		return DirectiveCompliance{}
	}
	return o.DirectiveCompliance
}

type ComplianceDirectiveID struct {
	// Result of the request
	Result Result `json:"result"`
	// The id of the action
	Action Action `json:"action"`
	Data   Data   `json:"data"`
}

func (o *ComplianceDirectiveID) GetResult() Result {
	if o == nil {
		return Result("")
	}
	return o.Result
}

func (o *ComplianceDirectiveID) GetAction() Action {
	if o == nil {
		return Action("")
	}
	return o.Action
}

func (o *ComplianceDirectiveID) GetData() Data {
	if o == nil {
		return Data{}
	}
	return o.Data
}
