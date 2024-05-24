package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// PurgeSoftwareResult - Result of the request
type PurgeSoftwareResult string

const (
	PurgeSoftwareResultSuccess PurgeSoftwareResult = "success"
	PurgeSoftwareResultError   PurgeSoftwareResult = "error"
)

func (e PurgeSoftwareResult) ToPointer() *PurgeSoftwareResult {
	return &e
}
func (e *PurgeSoftwareResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = PurgeSoftwareResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PurgeSoftwareResult: %v", v)
	}
}

// PurgeSoftwareAction - The id of the action
type PurgeSoftwareAction string

const (
	PurgeSoftwareActionPurgeSoftware PurgeSoftwareAction = "purgeSoftware"
)

func (e PurgeSoftwareAction) ToPointer() *PurgeSoftwareAction {
	return &e
}
func (e *PurgeSoftwareAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "purgeSoftware":
		*e = PurgeSoftwareAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PurgeSoftwareAction: %v", v)
	}
}

// PurgeSoftwareResponseBody - Purge Software
type PurgeSoftwareResponseBody struct {
	// Result of the request
	Result PurgeSoftwareResult `json:"result"`
	// The id of the action
	Action PurgeSoftwareAction `json:"action"`
	Data   []string            `json:"data"`
}

func (o *PurgeSoftwareResponseBody) GetResult() PurgeSoftwareResult {
	if o == nil {
		return PurgeSoftwareResult("")
	}
	return o.Result
}

func (o *PurgeSoftwareResponseBody) GetAction() PurgeSoftwareAction {
	if o == nil {
		return PurgeSoftwareAction("")
	}
	return o.Action
}

func (o *PurgeSoftwareResponseBody) GetData() []string {
	if o == nil {
		return []string{}
	}
	return o.Data
}

type PurgeSoftwareResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Purge Software
	Object *PurgeSoftwareResponseBody
}

func (o *PurgeSoftwareResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PurgeSoftwareResponse) GetObject() *PurgeSoftwareResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
