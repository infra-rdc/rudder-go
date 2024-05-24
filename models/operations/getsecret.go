package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetSecretRequest struct {
	// Unique name of the secret
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetSecretRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// GetSecretResult - Result of the request
type GetSecretResult string

const (
	GetSecretResultSuccess GetSecretResult = "success"
	GetSecretResultError   GetSecretResult = "error"
)

func (e GetSecretResult) ToPointer() *GetSecretResult {
	return &e
}
func (e *GetSecretResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetSecretResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSecretResult: %v", v)
	}
}

// GetSecretAction - The id of the action
type GetSecretAction string

const (
	GetSecretActionGetSecret GetSecretAction = "getSecret"
)

func (e GetSecretAction) ToPointer() *GetSecretAction {
	return &e
}
func (e *GetSecretAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getSecret":
		*e = GetSecretAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSecretAction: %v", v)
	}
}

type GetSecretSecrets struct {
	// The name of the secret used as a reference on the value
	Name string `json:"name"`
	// The description of the secret to identify it more easily
	Description string `json:"description"`
}

func (o *GetSecretSecrets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetSecretSecrets) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

type GetSecretData struct {
	// List of API version and status
	Secrets []GetSecretSecrets `json:"secrets"`
}

func (o *GetSecretData) GetSecrets() []GetSecretSecrets {
	if o == nil {
		return []GetSecretSecrets{}
	}
	return o.Secrets
}

// GetSecretResponseBody - Secret information
type GetSecretResponseBody struct {
	// Result of the request
	Result GetSecretResult `json:"result"`
	// The id of the action
	Action GetSecretAction `json:"action"`
	Data   GetSecretData   `json:"data"`
}

func (o *GetSecretResponseBody) GetResult() GetSecretResult {
	if o == nil {
		return GetSecretResult("")
	}
	return o.Result
}

func (o *GetSecretResponseBody) GetAction() GetSecretAction {
	if o == nil {
		return GetSecretAction("")
	}
	return o.Action
}

func (o *GetSecretResponseBody) GetData() GetSecretData {
	if o == nil {
		return GetSecretData{}
	}
	return o.Data
}

type GetSecretResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Secret information
	Object *GetSecretResponseBody
}

func (o *GetSecretResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSecretResponse) GetObject() *GetSecretResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
