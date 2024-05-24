package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteSecretRequest struct {
	// Unique name of the secret
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *DeleteSecretRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// DeleteSecretResult - Result of the request
type DeleteSecretResult string

const (
	DeleteSecretResultSuccess DeleteSecretResult = "success"
	DeleteSecretResultError   DeleteSecretResult = "error"
)

func (e DeleteSecretResult) ToPointer() *DeleteSecretResult {
	return &e
}
func (e *DeleteSecretResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteSecretResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteSecretResult: %v", v)
	}
}

// DeleteSecretAction - The id of the action
type DeleteSecretAction string

const (
	DeleteSecretActionDeleteSecret DeleteSecretAction = "deleteSecret"
)

func (e DeleteSecretAction) ToPointer() *DeleteSecretAction {
	return &e
}
func (e *DeleteSecretAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleteSecret":
		*e = DeleteSecretAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteSecretAction: %v", v)
	}
}

type DeleteSecretSecrets struct {
	// The name of the secret used as a reference on the value
	Name string `json:"name"`
	// The description of the secret to identify it more easily
	Description string `json:"description"`
}

func (o *DeleteSecretSecrets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DeleteSecretSecrets) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

type DeleteSecretData struct {
	// List of API version and status
	Secrets []DeleteSecretSecrets `json:"secrets"`
}

func (o *DeleteSecretData) GetSecrets() []DeleteSecretSecrets {
	if o == nil {
		return []DeleteSecretSecrets{}
	}
	return o.Secrets
}

// DeleteSecretResponseBody - Secret information
type DeleteSecretResponseBody struct {
	// Result of the request
	Result DeleteSecretResult `json:"result"`
	// The id of the action
	Action DeleteSecretAction `json:"action"`
	Data   DeleteSecretData   `json:"data"`
}

func (o *DeleteSecretResponseBody) GetResult() DeleteSecretResult {
	if o == nil {
		return DeleteSecretResult("")
	}
	return o.Result
}

func (o *DeleteSecretResponseBody) GetAction() DeleteSecretAction {
	if o == nil {
		return DeleteSecretAction("")
	}
	return o.Action
}

func (o *DeleteSecretResponseBody) GetData() DeleteSecretData {
	if o == nil {
		return DeleteSecretData{}
	}
	return o.Data
}

type DeleteSecretResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Secret information
	Object *DeleteSecretResponseBody
}

func (o *DeleteSecretResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteSecretResponse) GetObject() *DeleteSecretResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
