package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// QueueInformationResult - Result of the request
type QueueInformationResult string

const (
	QueueInformationResultSuccess QueueInformationResult = "success"
	QueueInformationResultError   QueueInformationResult = "error"
)

func (e QueueInformationResult) ToPointer() *QueueInformationResult {
	return &e
}
func (e *QueueInformationResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = QueueInformationResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueueInformationResult: %v", v)
	}
}

// QueueInformationAction - The id of the action
type QueueInformationAction string

const (
	QueueInformationActionQueueInformation QueueInformationAction = "queueInformation"
)

func (e QueueInformationAction) ToPointer() *QueueInformationAction {
	return &e
}
func (e *QueueInformationAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "queueInformation":
		*e = QueueInformationAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueueInformationAction: %v", v)
	}
}

// QueueInformationData - Information about the service
type QueueInformationData struct {
	QueueMaxSize int64 `json:"queueMaxSize"`
	// Is the inventory queue full
	QueueSaturated bool `json:"queueSaturated"`
}

func (o *QueueInformationData) GetQueueMaxSize() int64 {
	if o == nil {
		return 0
	}
	return o.QueueMaxSize
}

func (o *QueueInformationData) GetQueueSaturated() bool {
	if o == nil {
		return false
	}
	return o.QueueSaturated
}

// QueueInformationResponseBody - Inventories information
type QueueInformationResponseBody struct {
	// Result of the request
	Result QueueInformationResult `json:"result"`
	// The id of the action
	Action QueueInformationAction `json:"action"`
	// Information about the service
	Data QueueInformationData `json:"data"`
}

func (o *QueueInformationResponseBody) GetResult() QueueInformationResult {
	if o == nil {
		return QueueInformationResult("")
	}
	return o.Result
}

func (o *QueueInformationResponseBody) GetAction() QueueInformationAction {
	if o == nil {
		return QueueInformationAction("")
	}
	return o.Action
}

func (o *QueueInformationResponseBody) GetData() QueueInformationData {
	if o == nil {
		return QueueInformationData{}
	}
	return o.Data
}

type QueueInformationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Inventories information
	Object *QueueInformationResponseBody
}

func (o *QueueInformationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *QueueInformationResponse) GetObject() *QueueInformationResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
