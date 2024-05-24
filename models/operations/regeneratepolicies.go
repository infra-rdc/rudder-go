package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// RegeneratePoliciesResult - Result of the request
type RegeneratePoliciesResult string

const (
	RegeneratePoliciesResultSuccess RegeneratePoliciesResult = "success"
	RegeneratePoliciesResultError   RegeneratePoliciesResult = "error"
)

func (e RegeneratePoliciesResult) ToPointer() *RegeneratePoliciesResult {
	return &e
}
func (e *RegeneratePoliciesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = RegeneratePoliciesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RegeneratePoliciesResult: %v", v)
	}
}

// RegeneratePoliciesAction - The id of the action
type RegeneratePoliciesAction string

const (
	RegeneratePoliciesActionRegeneratePolicies RegeneratePoliciesAction = "regeneratePolicies"
)

func (e RegeneratePoliciesAction) ToPointer() *RegeneratePoliciesAction {
	return &e
}
func (e *RegeneratePoliciesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "regeneratePolicies":
		*e = RegeneratePoliciesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RegeneratePoliciesAction: %v", v)
	}
}

type Policies string

const (
	PoliciesStarted Policies = "Started"
)

func (e Policies) ToPointer() *Policies {
	return &e
}
func (e *Policies) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Started":
		*e = Policies(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Policies: %v", v)
	}
}

type RegeneratePoliciesData struct {
	Policies Policies `json:"policies"`
}

func (o *RegeneratePoliciesData) GetPolicies() Policies {
	if o == nil {
		return Policies("")
	}
	return o.Policies
}

// RegeneratePoliciesResponseBody - Success
type RegeneratePoliciesResponseBody struct {
	// Result of the request
	Result RegeneratePoliciesResult `json:"result"`
	// The id of the action
	Action RegeneratePoliciesAction `json:"action"`
	Data   RegeneratePoliciesData   `json:"data"`
}

func (o *RegeneratePoliciesResponseBody) GetResult() RegeneratePoliciesResult {
	if o == nil {
		return RegeneratePoliciesResult("")
	}
	return o.Result
}

func (o *RegeneratePoliciesResponseBody) GetAction() RegeneratePoliciesAction {
	if o == nil {
		return RegeneratePoliciesAction("")
	}
	return o.Action
}

func (o *RegeneratePoliciesResponseBody) GetData() RegeneratePoliciesData {
	if o == nil {
		return RegeneratePoliciesData{}
	}
	return o.Data
}

type RegeneratePoliciesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *RegeneratePoliciesResponseBody
}

func (o *RegeneratePoliciesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RegeneratePoliciesResponse) GetObject() *RegeneratePoliciesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
