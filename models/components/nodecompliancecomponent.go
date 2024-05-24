package components

import (
	"encoding/json"
	"fmt"
)

type NodeComplianceComponentMode string

const (
	NodeComplianceComponentModeFullCompliance  NodeComplianceComponentMode = "full-compliance"
	NodeComplianceComponentModeChangesOnly     NodeComplianceComponentMode = "changes-only"
	NodeComplianceComponentModeReportsDisabled NodeComplianceComponentMode = "reports-disabled"
)

func (e NodeComplianceComponentMode) ToPointer() *NodeComplianceComponentMode {
	return &e
}
func (e *NodeComplianceComponentMode) UnmarshalJSON(data []byte) error {
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
		*e = NodeComplianceComponentMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeComplianceComponentMode: %v", v)
	}
}

type NodeComplianceComponentComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *NodeComplianceComponentComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *NodeComplianceComponentComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *NodeComplianceComponentComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *NodeComplianceComponentComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *NodeComplianceComponentComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *NodeComplianceComponentComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *NodeComplianceComponentComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type Reports struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (o *Reports) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *Reports) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type Values struct {
	Value   string    `json:"value"`
	Reports []Reports `json:"reports"`
}

func (o *Values) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *Values) GetReports() []Reports {
	if o == nil {
		return []Reports{}
	}
	return o.Reports
}

type NodeComplianceComponent struct {
	// id of the node
	ID string `json:"id"`
	// Name of the node
	Name string                       `json:"name"`
	Mode *NodeComplianceComponentMode `json:"mode,omitempty"`
	// Directive compliance level
	Compliance        float32                                  `json:"compliance"`
	ComplianceDetails NodeComplianceComponentComplianceDetails `json:"complianceDetails"`
	Values            []Values                                 `json:"values"`
}

func (o *NodeComplianceComponent) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *NodeComplianceComponent) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *NodeComplianceComponent) GetMode() *NodeComplianceComponentMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *NodeComplianceComponent) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *NodeComplianceComponent) GetComplianceDetails() NodeComplianceComponentComplianceDetails {
	if o == nil {
		return NodeComplianceComponentComplianceDetails{}
	}
	return o.ComplianceDetails
}

func (o *NodeComplianceComponent) GetValues() []Values {
	if o == nil {
		return []Values{}
	}
	return o.Values
}
