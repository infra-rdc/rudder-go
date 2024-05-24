package components

type Secrets struct {
	// The name of the secret used as a reference on the value
	Name *string `json:"name,omitempty"`
	// The description of the secret to identify it more easily
	Description *string `json:"description,omitempty"`
	// The value of the secret it will not be exposed in the interface
	Value *string `json:"value,omitempty"`
}

func (o *Secrets) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Secrets) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Secrets) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
