package components

type RuleComplianceComponentComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *RuleComplianceComponentComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *RuleComplianceComponentComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *RuleComplianceComponentComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *RuleComplianceComponentComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *RuleComplianceComponentComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *RuleComplianceComponentComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *RuleComplianceComponentComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type RuleComplianceComponentComponentsComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *RuleComplianceComponentComponentsComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *RuleComplianceComponentComponentsComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *RuleComplianceComponentComponentsComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *RuleComplianceComponentComponentsComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *RuleComplianceComponentComponentsComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *RuleComplianceComponentComponentsComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *RuleComplianceComponentComponentsComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type RuleComplianceComponentReports struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (o *RuleComplianceComponentReports) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *RuleComplianceComponentReports) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type RuleComplianceComponentValues struct {
	Value   string                           `json:"value"`
	Reports []RuleComplianceComponentReports `json:"reports"`
}

func (o *RuleComplianceComponentValues) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *RuleComplianceComponentValues) GetReports() []RuleComplianceComponentReports {
	if o == nil {
		return []RuleComplianceComponentReports{}
	}
	return o.Reports
}

type RuleComplianceComponentComponents struct {
	// Name of the component
	Name string `json:"name"`
	// directive compliance level
	Compliance        float32                                            `json:"compliance"`
	ComplianceDetails RuleComplianceComponentComponentsComplianceDetails `json:"complianceDetails"`
	Values            []RuleComplianceComponentValues                    `json:"values"`
}

func (o *RuleComplianceComponentComponents) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RuleComplianceComponentComponents) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *RuleComplianceComponentComponents) GetComplianceDetails() RuleComplianceComponentComponentsComplianceDetails {
	if o == nil {
		return RuleComplianceComponentComponentsComplianceDetails{}
	}
	return o.ComplianceDetails
}

func (o *RuleComplianceComponentComponents) GetValues() []RuleComplianceComponentValues {
	if o == nil {
		return []RuleComplianceComponentValues{}
	}
	return o.Values
}

type RuleComplianceComponent struct {
	// id of the rule
	ID string `json:"id"`
	// Name of the rule
	Name string `json:"name"`
	// Rule compliance level
	Compliance        float32                                  `json:"compliance"`
	ComplianceDetails RuleComplianceComponentComplianceDetails `json:"complianceDetails"`
	Components        []RuleComplianceComponentComponents      `json:"components"`
}

func (o *RuleComplianceComponent) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RuleComplianceComponent) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RuleComplianceComponent) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *RuleComplianceComponent) GetComplianceDetails() RuleComplianceComponentComplianceDetails {
	if o == nil {
		return RuleComplianceComponentComplianceDetails{}
	}
	return o.ComplianceDetails
}

func (o *RuleComplianceComponent) GetComponents() []RuleComplianceComponentComponents {
	if o == nil {
		return []RuleComplianceComponentComponents{}
	}
	return o.Components
}
