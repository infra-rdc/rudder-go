package components

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
)

// GroupUpdateComposition - Boolean operator to use between each  `where` criteria.
type GroupUpdateComposition string

const (
	GroupUpdateCompositionAnd GroupUpdateComposition = "and"
	GroupUpdateCompositionOr  GroupUpdateComposition = "or"
)

func (e GroupUpdateComposition) ToPointer() *GroupUpdateComposition {
	return &e
}
func (e *GroupUpdateComposition) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = GroupUpdateComposition(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupUpdateComposition: %v", v)
	}
}

type GroupUpdateWhere struct {
	// Type of the object
	ObjectType *string `json:"objectType,omitempty"`
	// Attribute to compare
	Attribute *string `json:"attribute,omitempty"`
	// Comparator to use
	Comparator *string `json:"comparator,omitempty"`
	// Value to compare against
	Value *string `json:"value,omitempty"`
}

func (o *GroupUpdateWhere) GetObjectType() *string {
	if o == nil {
		return nil
	}
	return o.ObjectType
}

func (o *GroupUpdateWhere) GetAttribute() *string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *GroupUpdateWhere) GetComparator() *string {
	if o == nil {
		return nil
	}
	return o.Comparator
}

func (o *GroupUpdateWhere) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// GroupUpdateQuery - The criteria defining the group. If not provided, the group will be empty.
type GroupUpdateQuery struct {
	// What kind of data we want to include. Here we can get policy servers/relay by setting `nodeAndPolicyServer`. Only used if `where` is defined.
	Select *string `default:"node" json:"select"`
	// Boolean operator to use between each  `where` criteria.
	Composition *GroupUpdateComposition `default:"and" json:"composition"`
	// List of criteria
	Where []GroupUpdateWhere `json:"where,omitempty"`
}

func (g GroupUpdateQuery) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupUpdateQuery) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GroupUpdateQuery) GetSelect() *string {
	if o == nil {
		return nil
	}
	return o.Select
}

func (o *GroupUpdateQuery) GetComposition() *GroupUpdateComposition {
	if o == nil {
		return nil
	}
	return o.Composition
}

func (o *GroupUpdateQuery) GetWhere() []GroupUpdateWhere {
	if o == nil {
		return nil
	}
	return o.Where
}

type GroupUpdate struct {
	// Id of the new group's category
	Category *string `json:"category,omitempty"`
	// Name of the group
	DisplayName *string `json:"displayName,omitempty"`
	// Group description
	Description *string `json:"description,omitempty"`
	// The criteria defining the group. If not provided, the group will be empty.
	Query *GroupUpdateQuery `json:"query,omitempty"`
	// Should the group be dynamically refreshed (if not, it is a static group)
	Dynamic *bool `default:"true" json:"dynamic"`
	// Enable or disable the group
	Enabled *bool `default:"true" json:"enabled"`
}

func (g GroupUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GroupUpdate) GetCategory() *string {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *GroupUpdate) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *GroupUpdate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GroupUpdate) GetQuery() *GroupUpdateQuery {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GroupUpdate) GetDynamic() *bool {
	if o == nil {
		return nil
	}
	return o.Dynamic
}

func (o *GroupUpdate) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}
