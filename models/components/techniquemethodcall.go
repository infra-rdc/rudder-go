package components

type TechniqueMethodCallParameters struct {
	// Parameter name
	Name *string `json:"name,omitempty"`
	// Parameter value
	Value *string `json:"value,omitempty"`
}

func (o *TechniqueMethodCallParameters) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TechniqueMethodCallParameters) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type TechniqueMethodCall struct {
	// Method call id
	ID *string `json:"id,omitempty"`
	// Component is used in reporting to identify this method call. You can see it as a name
	Component *string `json:"component,omitempty"`
	// Id of the method called
	Method *string `json:"method,omitempty"`
	// Condition to run this method.
	Condition *string `json:"condition,omitempty"`
	// Should the reporting of this method be disabled
	DisableReporting *bool `json:"disableReporting,omitempty"`
	// Parameters for this method call
	Parameters []TechniqueMethodCallParameters `json:"parameters,omitempty"`
}

func (o *TechniqueMethodCall) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TechniqueMethodCall) GetComponent() *string {
	if o == nil {
		return nil
	}
	return o.Component
}

func (o *TechniqueMethodCall) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *TechniqueMethodCall) GetCondition() *string {
	if o == nil {
		return nil
	}
	return o.Condition
}

func (o *TechniqueMethodCall) GetDisableReporting() *bool {
	if o == nil {
		return nil
	}
	return o.DisableReporting
}

func (o *TechniqueMethodCall) GetParameters() []TechniqueMethodCallParameters {
	if o == nil {
		return nil
	}
	return o.Parameters
}
