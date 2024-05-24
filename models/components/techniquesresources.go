package components

type TechniquesResources struct {
	// Resource's name
	Name *string `json:"name,omitempty"`
	// State of the resource
	State *string `json:"state,omitempty"`
}

func (o *TechniquesResources) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TechniquesResources) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}
