package components

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
)

type CallsType string

const (
	CallsTypeTechniqueMethodCall CallsType = "technique-method-call"
	CallsTypeTechniqueBlock      CallsType = "technique-block"
)

type Calls struct {
	TechniqueMethodCall *TechniqueMethodCall
	TechniqueBlock      *TechniqueBlock

	Type CallsType
}

func CreateCallsTechniqueMethodCall(techniqueMethodCall TechniqueMethodCall) Calls {
	typ := CallsTypeTechniqueMethodCall

	return Calls{
		TechniqueMethodCall: &techniqueMethodCall,
		Type:                typ,
	}
}

func CreateCallsTechniqueBlock(techniqueBlock TechniqueBlock) Calls {
	typ := CallsTypeTechniqueBlock

	return Calls{
		TechniqueBlock: &techniqueBlock,
		Type:           typ,
	}
}

func (u *Calls) UnmarshalJSON(data []byte) error {

	var techniqueBlock TechniqueBlock = TechniqueBlock{}
	if err := utils.UnmarshalJSON(data, &techniqueBlock, "", true, true); err == nil {
		u.TechniqueBlock = &techniqueBlock
		u.Type = CallsTypeTechniqueBlock
		return nil
	}

	var techniqueMethodCall TechniqueMethodCall = TechniqueMethodCall{}
	if err := utils.UnmarshalJSON(data, &techniqueMethodCall, "", true, true); err == nil {
		u.TechniqueMethodCall = &techniqueMethodCall
		u.Type = CallsTypeTechniqueMethodCall
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Calls", string(data))
}

func (u Calls) MarshalJSON() ([]byte, error) {
	if u.TechniqueMethodCall != nil {
		return utils.MarshalJSON(u.TechniqueMethodCall, "", true)
	}

	if u.TechniqueBlock != nil {
		return utils.MarshalJSON(u.TechniqueBlock, "", true)
	}

	return nil, errors.New("could not marshal union type Calls: all fields are null")
}

// TechniqueBlockType - Kind of reporting logic
type TechniqueBlockType string

const (
	TechniqueBlockTypeWorst                TechniqueBlockType = "worst"
	TechniqueBlockTypeFocus                TechniqueBlockType = "focus"
	TechniqueBlockTypeWorstCaseWeightedOne TechniqueBlockType = "worst-case-weighted-one"
	TechniqueBlockTypeWorstCaseWeightedSum TechniqueBlockType = "worst-case-weighted-sum"
)

func (e TechniqueBlockType) ToPointer() *TechniqueBlockType {
	return &e
}
func (e *TechniqueBlockType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "worst":
		fallthrough
	case "focus":
		fallthrough
	case "worst-case-weighted-one":
		fallthrough
	case "worst-case-weighted-sum":
		*e = TechniqueBlockType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TechniqueBlockType: %v", v)
	}
}

type ReportingLogic struct {
	// Kind of reporting logic
	Type *TechniqueBlockType `json:"type,omitempty"`
	// reporting value used for some reporting logic (i.e. focus)
	Value *string `json:"value,omitempty"`
}

func (o *ReportingLogic) GetType() *TechniqueBlockType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ReportingLogic) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type TechniqueBlock struct {
	// Method call id
	ID *string `json:"id,omitempty"`
	// Component is used in reporting to identify this method call. You can see it as a name
	Component *string `json:"component,omitempty"`
	// Condition to run this method.
	Condition *string `json:"condition,omitempty"`
	// Method and blocks contained by this block
	Calls          []Calls         `json:"calls,omitempty"`
	ReportingLogic *ReportingLogic `json:"reportingLogic,omitempty"`
}

func (o *TechniqueBlock) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TechniqueBlock) GetComponent() *string {
	if o == nil {
		return nil
	}
	return o.Component
}

func (o *TechniqueBlock) GetCondition() *string {
	if o == nil {
		return nil
	}
	return o.Condition
}

func (o *TechniqueBlock) GetCalls() []Calls {
	if o == nil {
		return nil
	}
	return o.Calls
}

func (o *TechniqueBlock) GetReportingLogic() *ReportingLogic {
	if o == nil {
		return nil
	}
	return o.ReportingLogic
}
