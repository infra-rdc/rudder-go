package components

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
)

type Constraints struct {
	// Can this parameter be empty?
	AllowEmptyString bool `json:"allow_empty_string"`
	// Can this parameter allow trailing/ending spaces, or even a full whitespace string ?
	AllowWhitespaceString bool `json:"allow_whitespace_string"`
	// Maximum size of a parameter
	MaxLength int64 `json:"max_length"`
	// Minimal size of a parameter
	MinLength int64 `json:"min_length"`
	// A regex to validate this parameter
	Regex string `json:"regex"`
	// A regexp to invalidate this parameter
	NotRegex string `json:"not_regex"`
	// List of items authorized for this parameter
	Select []string `json:"select"`
}

func (o *Constraints) GetAllowEmptyString() bool {
	if o == nil {
		return false
	}
	return o.AllowEmptyString
}

func (o *Constraints) GetAllowWhitespaceString() bool {
	if o == nil {
		return false
	}
	return o.AllowWhitespaceString
}

func (o *Constraints) GetMaxLength() int64 {
	if o == nil {
		return 0
	}
	return o.MaxLength
}

func (o *Constraints) GetMinLength() int64 {
	if o == nil {
		return 0
	}
	return o.MinLength
}

func (o *Constraints) GetRegex() string {
	if o == nil {
		return ""
	}
	return o.Regex
}

func (o *Constraints) GetNotRegex() string {
	if o == nil {
		return ""
	}
	return o.NotRegex
}

func (o *Constraints) GetSelect() []string {
	if o == nil {
		return []string{}
	}
	return o.Select
}

// MethodParameterType - Type of the parameter
type MethodParameterType string

const (
	MethodParameterTypeString     MethodParameterType = "String"
	MethodParameterTypeHereString MethodParameterType = "HereString"
)

func (e MethodParameterType) ToPointer() *MethodParameterType {
	return &e
}
func (e *MethodParameterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "String":
		fallthrough
	case "HereString":
		*e = MethodParameterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MethodParameterType: %v", v)
	}
}

type MethodParameter struct {
	// Parameter name
	Name string `json:"name"`
	// Description of this parameter
	Description string      `json:"description"`
	Constraints Constraints `json:"constraints"`
	// Type of the parameter
	Type *MethodParameterType `default:"String" json:"_type"`
}

func (m MethodParameter) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MethodParameter) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MethodParameter) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MethodParameter) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *MethodParameter) GetConstraints() Constraints {
	if o == nil {
		return Constraints{}
	}
	return o.Constraints
}

func (o *MethodParameter) GetType() *MethodParameterType {
	if o == nil {
		return nil
	}
	return o.Type
}
