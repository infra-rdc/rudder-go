package components

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
)

// GroupComposition - Boolean operator to use between each  `where` criteria.
type GroupComposition string

const (
	GroupCompositionAnd GroupComposition = "and"
	GroupCompositionOr  GroupComposition = "or"
)

func (e GroupComposition) ToPointer() *GroupComposition {
	return &e
}
func (e *GroupComposition) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = GroupComposition(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupComposition: %v", v)
	}
}

type GroupWhere struct {
	// Type of the object
	ObjectType *string `json:"objectType,omitempty"`
	// Attribute to compare
	Attribute *string `json:"attribute,omitempty"`
	// Comparator to use
	Comparator *string `json:"comparator,omitempty"`
	// Value to compare against
	Value *string `json:"value,omitempty"`
}

func (o *GroupWhere) GetObjectType() *string {
	if o == nil {
		return nil
	}
	return o.ObjectType
}

func (o *GroupWhere) GetAttribute() *string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *GroupWhere) GetComparator() *string {
	if o == nil {
		return nil
	}
	return o.Comparator
}

func (o *GroupWhere) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// Query - The criteria defining the group
type Query struct {
	// What kind of data we want to include. Here we can get policy servers/relay by setting `nodeAndPolicyServer`. Only used if `where` is defined.
	Select *string `default:"node" json:"select"`
	// Boolean operator to use between each  `where` criteria.
	Composition *GroupComposition `default:"and" json:"composition"`
	// List of criteria
	Where []GroupWhere `json:"where,omitempty"`
}

func (q Query) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(q, "", false)
}

func (q *Query) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &q, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Query) GetSelect() *string {
	if o == nil {
		return nil
	}
	return o.Select
}

func (o *Query) GetComposition() *GroupComposition {
	if o == nil {
		return nil
	}
	return o.Composition
}

func (o *Query) GetWhere() []GroupWhere {
	if o == nil {
		return nil
	}
	return o.Where
}

type Properties struct {
	// Property name
	Name string `json:"name"`
	// Property value (can be a string or JSON object)
	Value any `json:"value"`
}

func (o *Properties) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Properties) GetValue() any {
	if o == nil {
		return nil
	}
	return o.Value
}

type Group struct {
	// Group id
	ID *string `default:"{autogenerated}" json:"id"`
	// Name of the group
	DisplayName *string `json:"displayName,omitempty"`
	// Group description
	Description *string `json:"description,omitempty"`
	// The criteria defining the group
	Query *Query `json:"query,omitempty"`
	// List of nodes in the group
	NodeIds []string `json:"nodeIds,omitempty"`
	// Should the group be dynamically refreshed (if not, it is a static group)
	Dynamic *bool `default:"true" json:"dynamic"`
	// Enable or disable the group
	Enabled    *bool    `default:"true" json:"enabled"`
	GroupClass []string `json:"groupClass,omitempty"`
	// Group properties
	Properties []Properties `json:"properties,omitempty"`
}

func (g Group) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *Group) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Group) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Group) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Group) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Group) GetQuery() *Query {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *Group) GetNodeIds() []string {
	if o == nil {
		return nil
	}
	return o.NodeIds
}

func (o *Group) GetDynamic() *bool {
	if o == nil {
		return nil
	}
	return o.Dynamic
}

func (o *Group) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Group) GetGroupClass() []string {
	if o == nil {
		return nil
	}
	return o.GroupClass
}

func (o *Group) GetProperties() []Properties {
	if o == nil {
		return nil
	}
	return o.Properties
}
