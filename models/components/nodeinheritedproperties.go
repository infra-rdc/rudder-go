package components

import (
	"encoding/json"
	"fmt"
)

// NodeInheritedPropertiesProvider - Property provider (if the property is not a simple node property)
type NodeInheritedPropertiesProvider string

const (
	NodeInheritedPropertiesProviderInherited                 NodeInheritedPropertiesProvider = "inherited"
	NodeInheritedPropertiesProviderOverridden                NodeInheritedPropertiesProvider = "overridden"
	NodeInheritedPropertiesProviderPluginNameLikeDatasources NodeInheritedPropertiesProvider = "plugin name like datasources"
)

func (e NodeInheritedPropertiesProvider) ToPointer() *NodeInheritedPropertiesProvider {
	return &e
}
func (e *NodeInheritedPropertiesProvider) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "inherited":
		fallthrough
	case "overridden":
		fallthrough
	case "plugin name like datasources":
		*e = NodeInheritedPropertiesProvider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeInheritedPropertiesProvider: %v", v)
	}
}

// NodeInheritedPropertiesKind - The kind of object from which the property's value is inherited
type NodeInheritedPropertiesKind string

const (
	NodeInheritedPropertiesKindGlobal NodeInheritedPropertiesKind = "global"
	NodeInheritedPropertiesKindGroup  NodeInheritedPropertiesKind = "group"
)

func (e NodeInheritedPropertiesKind) ToPointer() *NodeInheritedPropertiesKind {
	return &e
}
func (e *NodeInheritedPropertiesKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "global":
		fallthrough
	case "group":
		*e = NodeInheritedPropertiesKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeInheritedPropertiesKind: %v", v)
	}
}

type Hierarchy struct {
	// The kind of object from which the property's value is inherited
	Kind NodeInheritedPropertiesKind `json:"kind"`
	// a string representation of the value. If it was a json value, it will be escaped.
	Value string `json:"value"`
	// in the case of a group, its identifier
	ID *string `json:"id,omitempty"`
	// in the case of a group, its name
	Name *string `json:"name,omitempty"`
}

func (o *Hierarchy) GetKind() NodeInheritedPropertiesKind {
	if o == nil {
		return NodeInheritedPropertiesKind("")
	}
	return o.Kind
}

func (o *Hierarchy) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *Hierarchy) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Hierarchy) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type NodeInheritedPropertiesProperties struct {
	// Property name
	Name string `json:"name"`
	// Resolved (ie, with inheritance and overriding done) property value (can be a string or JSON object)
	Value any `json:"value"`
	// Property provider (if the property is not a simple node property)
	Provider *NodeInheritedPropertiesProvider `json:"provider,omitempty"`
	// A description of the inheritance hierarchy for that property, with most direct parent at head and oldest one at tail
	Hierarchy []Hierarchy `json:"hierarchy,omitempty"`
	// The original value (ie, before overriding and inheritance resolution) for that node
	Origval any `json:"origval,omitempty"`
}

func (o *NodeInheritedPropertiesProperties) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *NodeInheritedPropertiesProperties) GetValue() any {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *NodeInheritedPropertiesProperties) GetProvider() *NodeInheritedPropertiesProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *NodeInheritedPropertiesProperties) GetHierarchy() []Hierarchy {
	if o == nil {
		return nil
	}
	return o.Hierarchy
}

func (o *NodeInheritedPropertiesProperties) GetOrigval() any {
	if o == nil {
		return nil
	}
	return o.Origval
}

type NodeInheritedProperties struct {
	// Unique id of the node
	ID string `json:"id"`
	// Node properties (either set by user or filled by third party sources)
	Properties []NodeInheritedPropertiesProperties `json:"properties"`
}

func (o *NodeInheritedProperties) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *NodeInheritedProperties) GetProperties() []NodeInheritedPropertiesProperties {
	if o == nil {
		return []NodeInheritedPropertiesProperties{}
	}
	return o.Properties
}
