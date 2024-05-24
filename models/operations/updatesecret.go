package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// UpdateSecretResult - Result of the request
type UpdateSecretResult string

const (
	UpdateSecretResultSuccess UpdateSecretResult = "success"
	UpdateSecretResultError   UpdateSecretResult = "error"
)

func (e UpdateSecretResult) ToPointer() *UpdateSecretResult {
	return &e
}
func (e *UpdateSecretResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateSecretResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateSecretResult: %v", v)
	}
}

// UpdateSecretAction - The id of the action
type UpdateSecretAction string

const (
	UpdateSecretActionUpdateSecret UpdateSecretAction = "updateSecret"
)

func (e UpdateSecretAction) ToPointer() *UpdateSecretAction {
	return &e
}
func (e *UpdateSecretAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updateSecret":
		*e = UpdateSecretAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateSecretAction: %v", v)
	}
}

type UpdateSecretSecrets struct {
	// The name of the secret used as a reference on the value
	Name string `json:"name"`
	// The description of the secret to identify it more easily
	Description string `json:"description"`
}

func (o *UpdateSecretSecrets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateSecretSecrets) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

type UpdateSecretData struct {
	// List of API version and status
	Secrets []UpdateSecretSecrets `json:"secrets"`
}

func (o *UpdateSecretData) GetSecrets() []UpdateSecretSecrets {
	if o == nil {
		return []UpdateSecretSecrets{}
	}
	return o.Secrets
}

// UpdateSecretResponseBody - Secrets information
type UpdateSecretResponseBody struct {
	// Result of the request
	Result UpdateSecretResult `json:"result"`
	// The id of the action
	Action UpdateSecretAction `json:"action"`
	Data   UpdateSecretData   `json:"data"`
}

func (o *UpdateSecretResponseBody) GetResult() UpdateSecretResult {
	if o == nil {
		return UpdateSecretResult("")
	}
	return o.Result
}

func (o *UpdateSecretResponseBody) GetAction() UpdateSecretAction {
	if o == nil {
		return UpdateSecretAction("")
	}
	return o.Action
}

func (o *UpdateSecretResponseBody) GetData() UpdateSecretData {
	if o == nil {
		return UpdateSecretData{}
	}
	return o.Data
}

type UpdateSecretResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Secrets information
	Object *UpdateSecretResponseBody
}

func (o *UpdateSecretResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateSecretResponse) GetObject() *UpdateSecretResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
