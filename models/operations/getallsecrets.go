package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetAllSecretsResult - Result of the request
type GetAllSecretsResult string

const (
	GetAllSecretsResultSuccess GetAllSecretsResult = "success"
	GetAllSecretsResultError   GetAllSecretsResult = "error"
)

func (e GetAllSecretsResult) ToPointer() *GetAllSecretsResult {
	return &e
}
func (e *GetAllSecretsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetAllSecretsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllSecretsResult: %v", v)
	}
}

// GetAllSecretsAction - The id of the action
type GetAllSecretsAction string

const (
	GetAllSecretsActionGetAllSecrets GetAllSecretsAction = "getAllSecrets"
)

func (e GetAllSecretsAction) ToPointer() *GetAllSecretsAction {
	return &e
}
func (e *GetAllSecretsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getAllSecrets":
		*e = GetAllSecretsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllSecretsAction: %v", v)
	}
}

type Secrets struct {
	// The name of the secret used as a reference on the value
	Name string `json:"name"`
	// The description of the secret to identify it more easily
	Description string `json:"description"`
}

func (o *Secrets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Secrets) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

type GetAllSecretsData struct {
	// List of API version and status
	Secrets []Secrets `json:"secrets"`
}

func (o *GetAllSecretsData) GetSecrets() []Secrets {
	if o == nil {
		return []Secrets{}
	}
	return o.Secrets
}

// GetAllSecretsResponseBody - Secrets information
type GetAllSecretsResponseBody struct {
	// Result of the request
	Result GetAllSecretsResult `json:"result"`
	// The id of the action
	Action GetAllSecretsAction `json:"action"`
	Data   GetAllSecretsData   `json:"data"`
}

func (o *GetAllSecretsResponseBody) GetResult() GetAllSecretsResult {
	if o == nil {
		return GetAllSecretsResult("")
	}
	return o.Result
}

func (o *GetAllSecretsResponseBody) GetAction() GetAllSecretsAction {
	if o == nil {
		return GetAllSecretsAction("")
	}
	return o.Action
}

func (o *GetAllSecretsResponseBody) GetData() GetAllSecretsData {
	if o == nil {
		return GetAllSecretsData{}
	}
	return o.Data
}

type GetAllSecretsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Secrets information
	Object *GetAllSecretsResponseBody
}

func (o *GetAllSecretsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAllSecretsResponse) GetObject() *GetAllSecretsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
