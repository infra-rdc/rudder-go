package components

type GroupCategoryUpdate struct {
	// The parent category of the groups
	Parent string `json:"parent"`
	// Name of the group category
	Name string `json:"name"`
	// Group category description
	Description *string `json:"description,omitempty"`
}

func (o *GroupCategoryUpdate) GetParent() string {
	if o == nil {
		return ""
	}
	return o.Parent
}

func (o *GroupCategoryUpdate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GroupCategoryUpdate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
