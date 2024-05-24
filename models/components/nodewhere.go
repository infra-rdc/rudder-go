package components

type NodeWhere struct {
	// Type of the object
	ObjectType *string `json:"objectType,omitempty"`
	// Attribute to compare
	Attribute *string `json:"attribute,omitempty"`
	// Comparator to use
	Comparator *string `json:"comparator,omitempty"`
	// Value to compare against
	Value *string `json:"value,omitempty"`
}

func (o *NodeWhere) GetObjectType() *string {
	if o == nil {
		return nil
	}
	return o.ObjectType
}

func (o *NodeWhere) GetAttribute() *string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *NodeWhere) GetComparator() *string {
	if o == nil {
		return nil
	}
	return o.Comparator
}

func (o *NodeWhere) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
