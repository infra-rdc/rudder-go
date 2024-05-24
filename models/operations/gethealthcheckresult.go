package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetHealthcheckResultResult - Result of the request
type GetHealthcheckResultResult string

const (
	GetHealthcheckResultResultSuccess GetHealthcheckResultResult = "success"
	GetHealthcheckResultResultError   GetHealthcheckResultResult = "error"
)

func (e GetHealthcheckResultResult) ToPointer() *GetHealthcheckResultResult {
	return &e
}
func (e *GetHealthcheckResultResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetHealthcheckResultResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetHealthcheckResultResult: %v", v)
	}
}

// GetHealthcheckResultAction - The id of the action
type GetHealthcheckResultAction string

const (
	GetHealthcheckResultActionGetHealthcheckResult GetHealthcheckResultAction = "getHealthcheckResult"
)

func (e GetHealthcheckResultAction) ToPointer() *GetHealthcheckResultAction {
	return &e
}
func (e *GetHealthcheckResultAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getHealthcheckResult":
		*e = GetHealthcheckResultAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetHealthcheckResultAction: %v", v)
	}
}

// GetHealthcheckResultResponseBody - Healthcheck information
type GetHealthcheckResultResponseBody struct {
	// Result of the request
	Result GetHealthcheckResultResult `json:"result"`
	// The id of the action
	Action GetHealthcheckResultAction `json:"action"`
	Data   []components.Check         `json:"data"`
}

func (o *GetHealthcheckResultResponseBody) GetResult() GetHealthcheckResultResult {
	if o == nil {
		return GetHealthcheckResultResult("")
	}
	return o.Result
}

func (o *GetHealthcheckResultResponseBody) GetAction() GetHealthcheckResultAction {
	if o == nil {
		return GetHealthcheckResultAction("")
	}
	return o.Action
}

func (o *GetHealthcheckResultResponseBody) GetData() []components.Check {
	if o == nil {
		return []components.Check{}
	}
	return o.Data
}

type GetHealthcheckResultResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Healthcheck information
	Object *GetHealthcheckResultResponseBody
}

func (o *GetHealthcheckResultResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetHealthcheckResultResponse) GetObject() *GetHealthcheckResultResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
