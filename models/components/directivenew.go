package components

type DirectiveNewTags struct {
	// Value of the `name` key
	Name *string `json:"name,omitempty"`
}

func (o *DirectiveNewTags) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DirectiveNewParameters - Directive parameters (depends on the source technique)
type DirectiveNewParameters struct {
}

type DirectiveNew struct {
	// The id of the directive the clone will be based onto. If this parameter if provided,  the new directive will be a clone of this source. Other value will override values from the source.
	Source *string `json:"source,omitempty"`
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
	System *bool              `json:"system,omitempty"`
	Tags   []DirectiveNewTags `json:"tags,omitempty"`
	// Directive parameters (depends on the source technique)
	Parameters *DirectiveNewParameters `json:"parameters,omitempty"`
}

func (o *DirectiveNew) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *DirectiveNew) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DirectiveNew) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *DirectiveNew) GetShortDescription() *string {
	if o == nil {
		return nil
	}
	return o.ShortDescription
}

func (o *DirectiveNew) GetLongDescription() *string {
	if o == nil {
		return nil
	}
	return o.LongDescription
}

func (o *DirectiveNew) GetTechniqueName() *string {
	if o == nil {
		return nil
	}
	return o.TechniqueName
}

func (o *DirectiveNew) GetTechniqueVersion() *string {
	if o == nil {
		return nil
	}
	return o.TechniqueVersion
}

func (o *DirectiveNew) GetPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *DirectiveNew) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *DirectiveNew) GetSystem() *bool {
	if o == nil {
		return nil
	}
	return o.System
}

func (o *DirectiveNew) GetTags() []DirectiveNewTags {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *DirectiveNew) GetParameters() *DirectiveNewParameters {
	if o == nil {
		return nil
	}
	return o.Parameters
}
