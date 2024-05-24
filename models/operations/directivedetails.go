package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DirectiveDetailsRequest struct {
	// Id of the directive
	DirectiveID string `pathParam:"style=simple,explode=false,name=directiveId"`
}

func (o *DirectiveDetailsRequest) GetDirectiveID() string {
	if o == nil {
		return ""
	}
	return o.DirectiveID
}

// DirectiveDetailsResult - Result of the request
type DirectiveDetailsResult string

const (
	DirectiveDetailsResultSuccess DirectiveDetailsResult = "success"
	DirectiveDetailsResultError   DirectiveDetailsResult = "error"
)

func (e DirectiveDetailsResult) ToPointer() *DirectiveDetailsResult {
	return &e
}
func (e *DirectiveDetailsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DirectiveDetailsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DirectiveDetailsResult: %v", v)
	}
}

// DirectiveDetailsAction - The id of the action
type DirectiveDetailsAction string

const (
	DirectiveDetailsActionDirectiveDetails DirectiveDetailsAction = "directiveDetails"
)

func (e DirectiveDetailsAction) ToPointer() *DirectiveDetailsAction {
	return &e
}
func (e *DirectiveDetailsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "directiveDetails":
		*e = DirectiveDetailsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DirectiveDetailsAction: %v", v)
	}
}

type DirectiveDetailsData struct {
	Directives []components.Directive `json:"directives"`
}

func (o *DirectiveDetailsData) GetDirectives() []components.Directive {
	if o == nil {
		return []components.Directive{}
	}
	return o.Directives
}

// DirectiveDetailsResponseBody - Directives information
type DirectiveDetailsResponseBody struct {
	// Result of the request
	Result DirectiveDetailsResult `json:"result"`
	// The id of the action
	Action DirectiveDetailsAction `json:"action"`
	Data   DirectiveDetailsData   `json:"data"`
}

func (o *DirectiveDetailsResponseBody) GetResult() DirectiveDetailsResult {
	if o == nil {
		return DirectiveDetailsResult("")
	}
	return o.Result
}

func (o *DirectiveDetailsResponseBody) GetAction() DirectiveDetailsAction {
	if o == nil {
		return DirectiveDetailsAction("")
	}
	return o.Action
}

func (o *DirectiveDetailsResponseBody) GetData() DirectiveDetailsData {
	if o == nil {
		return DirectiveDetailsData{}
	}
	return o.Data
}

type DirectiveDetailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Directives information
	Object *DirectiveDetailsResponseBody
}

func (o *DirectiveDetailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DirectiveDetailsResponse) GetObject() *DirectiveDetailsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
