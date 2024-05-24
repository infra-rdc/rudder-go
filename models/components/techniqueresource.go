package components

type TechniqueResource struct {
	// resource name. this is the relative path to the resource
	Name *string `json:"name,omitempty"`
	// State of the resource file. it can be a value between new, modified, deleted, untouched
	State *string `json:"state,omitempty"`
}

func (o *TechniqueResource) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TechniqueResource) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}
