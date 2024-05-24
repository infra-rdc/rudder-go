package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// AddSecretResult - Result of the request
type AddSecretResult string

const (
	AddSecretResultSuccess AddSecretResult = "success"
	AddSecretResultError   AddSecretResult = "error"
)

func (e AddSecretResult) ToPointer() *AddSecretResult {
	return &e
}
func (e *AddSecretResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = AddSecretResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddSecretResult: %v", v)
	}
}

// AddSecretAction - The id of the action
type AddSecretAction string

const (
	AddSecretActionAddSecret AddSecretAction = "addSecret"
)

func (e AddSecretAction) ToPointer() *AddSecretAction {
	return &e
}
func (e *AddSecretAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "addSecret":
		*e = AddSecretAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddSecretAction: %v", v)
	}
}

type AddSecretSecrets struct {
	// The name of the secret used as a reference on the value
	Name string `json:"name"`
	// The description of the secret to identify it more easily
	Description string `json:"description"`
}

func (o *AddSecretSecrets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AddSecretSecrets) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

type AddSecretData struct {
	// List of API version and status
	Secrets []AddSecretSecrets `json:"secrets"`
}

func (o *AddSecretData) GetSecrets() []AddSecretSecrets {
	if o == nil {
		return []AddSecretSecrets{}
	}
	return o.Secrets
}

// AddSecretResponseBody - Secrets information
type AddSecretResponseBody struct {
	// Result of the request
	Result AddSecretResult `json:"result"`
	// The id of the action
	Action AddSecretAction `json:"action"`
	Data   AddSecretData   `json:"data"`
}

func (o *AddSecretResponseBody) GetResult() AddSecretResult {
	if o == nil {
		return AddSecretResult("")
	}
	return o.Result
}

func (o *AddSecretResponseBody) GetAction() AddSecretAction {
	if o == nil {
		return AddSecretAction("")
	}
	return o.Action
}

func (o *AddSecretResponseBody) GetData() AddSecretData {
	if o == nil {
		return AddSecretData{}
	}
	return o.Data
}

type AddSecretResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Secrets information
	Object *AddSecretResponseBody
}

func (o *AddSecretResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AddSecretResponse) GetObject() *AddSecretResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
