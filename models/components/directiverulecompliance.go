package components

import (
	"encoding/json"
	"fmt"
)

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

type ComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *ComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *ComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *ComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *ComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *ComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *ComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type DirectiveRuleComplianceComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *DirectiveRuleComplianceComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *DirectiveRuleComplianceComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *DirectiveRuleComplianceComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *DirectiveRuleComplianceComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *DirectiveRuleComplianceComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DirectiveRuleComplianceComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *DirectiveRuleComplianceComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type Components struct {
	// Name of the component
	Name string `json:"name"`
	// Directive compliance level
	Compliance        float32                                  `json:"compliance"`
	ComplianceDetails DirectiveRuleComplianceComplianceDetails `json:"complianceDetails"`
	Nodes             NodeComplianceComponent                  `json:"nodes"`
}

func (o *Components) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Components) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *Components) GetComplianceDetails() DirectiveRuleComplianceComplianceDetails {
	if o == nil {
		return DirectiveRuleComplianceComplianceDetails{}
	}
	return o.ComplianceDetails
}

func (o *Components) GetNodes() NodeComplianceComponent {
	if o == nil {
		return NodeComplianceComponent{}
	}
	return o.Nodes
}

type DirectiveRuleCompliance struct {
	// id of the rule
	ID string `json:"id"`
	// Name of the rule
	Name string `json:"name"`
	Mode *Mode  `json:"mode,omitempty"`
	// Directive compliance level
	Compliance        float32           `json:"compliance"`
	ComplianceDetails ComplianceDetails `json:"complianceDetails"`
	Components        []Components      `json:"components,omitempty"`
}

func (o *DirectiveRuleCompliance) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DirectiveRuleCompliance) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DirectiveRuleCompliance) GetMode() *Mode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *DirectiveRuleCompliance) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *DirectiveRuleCompliance) GetComplianceDetails() ComplianceDetails {
	if o == nil {
		return ComplianceDetails{}
	}
	return o.ComplianceDetails
}

func (o *DirectiveRuleCompliance) GetComponents() []Components {
	if o == nil {
		return nil
	}
	return o.Components
}
