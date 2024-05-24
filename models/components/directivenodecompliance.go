package components

type DirectiveNodeComplianceComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *DirectiveNodeComplianceComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *DirectiveNodeComplianceComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *DirectiveNodeComplianceComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *DirectiveNodeComplianceComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *DirectiveNodeComplianceComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DirectiveNodeComplianceComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *DirectiveNodeComplianceComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type DirectiveNodeCompliance struct {
	// id of the node
	ID string `json:"id"`
	// Name of the node
	Name string `json:"name"`
	// Directive compliance level
	Compliance        float32                                  `json:"compliance"`
	ComplianceDetails DirectiveNodeComplianceComplianceDetails `json:"complianceDetails"`
	Rules             RuleComplianceComponent                  `json:"rules"`
}

func (o *DirectiveNodeCompliance) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DirectiveNodeCompliance) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DirectiveNodeCompliance) GetCompliance() float32 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *DirectiveNodeCompliance) GetComplianceDetails() DirectiveNodeComplianceComplianceDetails {
	if o == nil {
		return DirectiveNodeComplianceComplianceDetails{}
	}
	return o.ComplianceDetails
}

func (o *DirectiveNodeCompliance) GetRules() RuleComplianceComponent {
	if o == nil {
		return RuleComplianceComponent{}
	}
	return o.Rules
}
