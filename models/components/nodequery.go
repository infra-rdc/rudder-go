package components

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
)

// Composition - Boolean operator to use between each  `where` criteria.
type Composition string

const (
	CompositionAnd Composition = "and"
	CompositionOr  Composition = "or"
)

func (e Composition) ToPointer() *Composition {
	return &e
}
func (e *Composition) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = Composition(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Composition: %v", v)
	}
}

type Where struct {
	// Type of the object
	ObjectType *string `json:"objectType,omitempty"`
	// Attribute to compare
	Attribute *string `json:"attribute,omitempty"`
	// Comparator to use
	Comparator *string `json:"comparator,omitempty"`
	// Value to compare against
	Value *string `json:"value,omitempty"`
}

func (o *Where) GetObjectType() *string {
	if o == nil {
		return nil
	}
	return o.ObjectType
}

func (o *Where) GetAttribute() *string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *Where) GetComparator() *string {
	if o == nil {
		return nil
	}
	return o.Comparator
}

func (o *Where) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type NodeQuery struct {
	// What kind of data we want to include. Here we can get policy servers/relay by setting `nodeAndPolicyServer`. Only used if `where` is defined.
	Select *string `default:"node" json:"select"`
	// Boolean operator to use between each  `where` criteria.
	Composition *Composition `default:"and" json:"composition"`
	// List of criteria
	Where []Where `json:"where,omitempty"`
}

func (n NodeQuery) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NodeQuery) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NodeQuery) GetSelect() *string {
	if o == nil {
		return nil
	}
	return o.Select
}

func (o *NodeQuery) GetComposition() *Composition {
	if o == nil {
		return nil
	}
	return o.Composition
}

func (o *NodeQuery) GetWhere() []Where {
	if o == nil {
		return nil
	}
	return o.Where
}
