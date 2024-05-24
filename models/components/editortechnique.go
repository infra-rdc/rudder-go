package components

import (
	"errors"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
)

type EditorTechniqueCallsType string

const (
	EditorTechniqueCallsTypeTechniqueMethodCall EditorTechniqueCallsType = "technique-method-call"
	EditorTechniqueCallsTypeTechniqueBlock      EditorTechniqueCallsType = "technique-block"
)

type EditorTechniqueCalls struct {
	TechniqueMethodCall *TechniqueMethodCall
	TechniqueBlock      *TechniqueBlock

	Type EditorTechniqueCallsType
}

func CreateEditorTechniqueCallsTechniqueMethodCall(techniqueMethodCall TechniqueMethodCall) EditorTechniqueCalls {
	typ := EditorTechniqueCallsTypeTechniqueMethodCall

	return EditorTechniqueCalls{
		TechniqueMethodCall: &techniqueMethodCall,
		Type:                typ,
	}
}

func CreateEditorTechniqueCallsTechniqueBlock(techniqueBlock TechniqueBlock) EditorTechniqueCalls {
	typ := EditorTechniqueCallsTypeTechniqueBlock

	return EditorTechniqueCalls{
		TechniqueBlock: &techniqueBlock,
		Type:           typ,
	}
}

func (u *EditorTechniqueCalls) UnmarshalJSON(data []byte) error {

	var techniqueBlock TechniqueBlock = TechniqueBlock{}
	if err := utils.UnmarshalJSON(data, &techniqueBlock, "", true, true); err == nil {
		u.TechniqueBlock = &techniqueBlock
		u.Type = EditorTechniqueCallsTypeTechniqueBlock
		return nil
	}

	var techniqueMethodCall TechniqueMethodCall = TechniqueMethodCall{}
	if err := utils.UnmarshalJSON(data, &techniqueMethodCall, "", true, true); err == nil {
		u.TechniqueMethodCall = &techniqueMethodCall
		u.Type = EditorTechniqueCallsTypeTechniqueMethodCall
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for EditorTechniqueCalls", string(data))
}

func (u EditorTechniqueCalls) MarshalJSON() ([]byte, error) {
	if u.TechniqueMethodCall != nil {
		return utils.MarshalJSON(u.TechniqueMethodCall, "", true)
	}

	if u.TechniqueBlock != nil {
		return utils.MarshalJSON(u.TechniqueBlock, "", true)
	}

	return nil, errors.New("could not marshal union type EditorTechniqueCalls: all fields are null")
}

type EditorTechnique struct {
	// Technique id
	ID *string `json:"id,omitempty"`
	// Technique name
	Name *string `json:"name,omitempty"`
	// version of this technique
	Version *string `json:"version,omitempty"`
	// category of this technique
	Category *string `json:"category,omitempty"`
	// description of this technique
	Description *string `json:"description,omitempty"`
	// Source of the technique, always editor here
	Source *string `json:"source,omitempty"`
	// Parameters for this technique
	Parameters []TechniqueParameter `json:"parameters,omitempty"`
	// Resources for this technique
	Resources []TechniqueResource `json:"resources,omitempty"`
	// Method and blocks contained by this technique
	Calls []EditorTechniqueCalls `json:"calls,omitempty"`
}

func (o *EditorTechnique) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EditorTechnique) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *EditorTechnique) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *EditorTechnique) GetCategory() *string {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *EditorTechnique) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *EditorTechnique) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *EditorTechnique) GetParameters() []TechniqueParameter {
	if o == nil {
		return nil
	}
	return o.Parameters
}

func (o *EditorTechnique) GetResources() []TechniqueResource {
	if o == nil {
		return nil
	}
	return o.Resources
}

func (o *EditorTechnique) GetCalls() []EditorTechniqueCalls {
	if o == nil {
		return nil
	}
	return o.Calls
}
