package sdkerrors

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// Result of the request
type Result string

const (
	ResultError Result = "error"
)

func (e Result) ToPointer() *Result {
	return &e
}
func (e *Result) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		*e = Result(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Result: %v", v)
	}
}

// Action - The id of the action
type Action string

const (
	ActionDeleteParameter Action = "deleteParameter"
)

func (e Action) ToPointer() *Action {
	return &e
}
func (e *Action) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleteParameter":
		*e = Action(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Action: %v", v)
	}
}

// DeleteParameterResponseBody - Non existing global property
type DeleteParameterResponseBody struct {
	// Id of the global property
	ID string `json:"id"`
	// Result of the request
	Result Result `json:"result"`
	// The id of the action
	Action       Action                  `json:"action"`
	ErrorDetails *string                 `json:"errorDetails,omitempty"`
	HTTPMeta     components.HTTPMetadata `json:"-"`
}

var _ error = &DeleteParameterResponseBody{}

func (e *DeleteParameterResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
