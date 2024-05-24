package components

type TechniqueParameter struct {
	// parameter id
	ID *string `json:"id,omitempty"`
	// Parameter name
	Name *string `json:"name,omitempty"`
	// description of this parameter
	Description *string `json:"description,omitempty"`
	// May the value given when creating a directive be empty
	MayBeEmpty *bool `json:"mayBeEmpty,omitempty"`
}

func (o *TechniqueParameter) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TechniqueParameter) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TechniqueParameter) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TechniqueParameter) GetMayBeEmpty() *bool {
	if o == nil {
		return nil
	}
	return o.MayBeEmpty
}
