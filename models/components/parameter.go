package components

type Parameter struct {
	// Name of the parameter
	ID string `json:"id"`
	// Value of the parameter
	Value any `json:"value,omitempty"`
	// Description of the parameter
	Description *string `json:"description,omitempty"`
	// Is the global parameter overridable by node
	Overridable *bool `json:"overridable,omitempty"`
}

func (o *Parameter) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Parameter) GetValue() any {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *Parameter) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Parameter) GetOverridable() *bool {
	if o == nil {
		return nil
	}
	return o.Overridable
}
