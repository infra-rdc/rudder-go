package components

type RuleCategoryUpdate struct {
	// The parent category of the rules
	Parent string `json:"parent"`
	// Name of the rule category
	Name string `json:"name"`
	// Rules category description
	Description *string `json:"description,omitempty"`
}

func (o *RuleCategoryUpdate) GetParent() string {
	if o == nil {
		return ""
	}
	return o.Parent
}

func (o *RuleCategoryUpdate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RuleCategoryUpdate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
