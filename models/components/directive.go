package components

import (
	"encoding/json"
	"fmt"
)

// PolicyMode - Policy mode of the directive
type PolicyMode string

const (
	PolicyModeEnforce PolicyMode = "enforce"
	PolicyModeAudit   PolicyMode = "audit"
)

func (e PolicyMode) ToPointer() *PolicyMode {
	return &e
}
func (e *PolicyMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enforce":
		fallthrough
	case "audit":
		*e = PolicyMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PolicyMode: %v", v)
	}
}

type Tags struct {
	// Value of the `name` key
	Name *string `json:"name,omitempty"`
}

func (o *Tags) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// Parameters - Directive parameters (depends on the source technique)
type Parameters struct {
}

type Directive struct {
	// Directive id
	ID *string `json:"id,omitempty"`
	// Human readable name of the directive
	DisplayName *string `json:"displayName,omitempty"`
	// One line directive description
	ShortDescription *string `json:"shortDescription,omitempty"`
	// Description of the technique (rendered as markdown)
	LongDescription *string `json:"longDescription,omitempty"`
	// Directive id
	TechniqueName *string `json:"techniqueName,omitempty"`
	// Directive id
	TechniqueVersion *string `json:"techniqueVersion,omitempty"`
	// Directive priority. `0` has highest priority.
	Priority *int64 `json:"priority,omitempty"`
	// Is the directive enabled
	Enabled *bool `json:"enabled,omitempty"`
	// If true it is an internal Rudder directive
	System *bool `json:"system,omitempty"`
	// Policy mode of the directive
	PolicyMode *PolicyMode `json:"policyMode,omitempty"`
	Tags       []Tags      `json:"tags,omitempty"`
	// Directive parameters (depends on the source technique)
	Parameters *Parameters `json:"parameters,omitempty"`
}

func (o *Directive) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Directive) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Directive) GetShortDescription() *string {
	if o == nil {
		return nil
	}
	return o.ShortDescription
}

func (o *Directive) GetLongDescription() *string {
	if o == nil {
		return nil
	}
	return o.LongDescription
}

func (o *Directive) GetTechniqueName() *string {
	if o == nil {
		return nil
	}
	return o.TechniqueName
}

func (o *Directive) GetTechniqueVersion() *string {
	if o == nil {
		return nil
	}
	return o.TechniqueVersion
}

func (o *Directive) GetPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *Directive) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Directive) GetSystem() *bool {
	if o == nil {
		return nil
	}
	return o.System
}

func (o *Directive) GetPolicyMode() *PolicyMode {
	if o == nil {
		return nil
	}
	return o.PolicyMode
}

func (o *Directive) GetTags() []Tags {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Directive) GetParameters() *Parameters {
	if o == nil {
		return nil
	}
	return o.Parameters
}
