package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// UpdatePoliciesResult - Result of the request
type UpdatePoliciesResult string

const (
	UpdatePoliciesResultSuccess UpdatePoliciesResult = "success"
	UpdatePoliciesResultError   UpdatePoliciesResult = "error"
)

func (e UpdatePoliciesResult) ToPointer() *UpdatePoliciesResult {
	return &e
}
func (e *UpdatePoliciesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdatePoliciesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdatePoliciesResult: %v", v)
	}
}

// UpdatePoliciesAction - The id of the action
type UpdatePoliciesAction string

const (
	UpdatePoliciesActionUpdatePolicies UpdatePoliciesAction = "updatePolicies"
)

func (e UpdatePoliciesAction) ToPointer() *UpdatePoliciesAction {
	return &e
}
func (e *UpdatePoliciesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updatePolicies":
		*e = UpdatePoliciesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdatePoliciesAction: %v", v)
	}
}

type UpdatePoliciesPolicies string

const (
	UpdatePoliciesPoliciesStarted UpdatePoliciesPolicies = "Started"
)

func (e UpdatePoliciesPolicies) ToPointer() *UpdatePoliciesPolicies {
	return &e
}
func (e *UpdatePoliciesPolicies) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Started":
		*e = UpdatePoliciesPolicies(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdatePoliciesPolicies: %v", v)
	}
}

type UpdatePoliciesData struct {
	Policies UpdatePoliciesPolicies `json:"policies"`
}

func (o *UpdatePoliciesData) GetPolicies() UpdatePoliciesPolicies {
	if o == nil {
		return UpdatePoliciesPolicies("")
	}
	return o.Policies
}

// UpdatePoliciesResponseBody - Success
type UpdatePoliciesResponseBody struct {
	// Result of the request
	Result UpdatePoliciesResult `json:"result"`
	// The id of the action
	Action UpdatePoliciesAction `json:"action"`
	Data   UpdatePoliciesData   `json:"data"`
}

func (o *UpdatePoliciesResponseBody) GetResult() UpdatePoliciesResult {
	if o == nil {
		return UpdatePoliciesResult("")
	}
	return o.Result
}

func (o *UpdatePoliciesResponseBody) GetAction() UpdatePoliciesAction {
	if o == nil {
		return UpdatePoliciesAction("")
	}
	return o.Action
}

func (o *UpdatePoliciesResponseBody) GetData() UpdatePoliciesData {
	if o == nil {
		return UpdatePoliciesData{}
	}
	return o.Data
}

type UpdatePoliciesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *UpdatePoliciesResponseBody
}

func (o *UpdatePoliciesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdatePoliciesResponse) GetObject() *UpdatePoliciesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
