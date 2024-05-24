package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// FileWatcherStartResult - Result of the request
type FileWatcherStartResult string

const (
	FileWatcherStartResultSuccess FileWatcherStartResult = "success"
	FileWatcherStartResultError   FileWatcherStartResult = "error"
)

func (e FileWatcherStartResult) ToPointer() *FileWatcherStartResult {
	return &e
}
func (e *FileWatcherStartResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = FileWatcherStartResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileWatcherStartResult: %v", v)
	}
}

// FileWatcherStartAction - The id of the action
type FileWatcherStartAction string

const (
	FileWatcherStartActionFileWatcherStart FileWatcherStartAction = "fileWatcherStart"
)

func (e FileWatcherStartAction) ToPointer() *FileWatcherStartAction {
	return &e
}
func (e *FileWatcherStartAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fileWatcherStart":
		*e = FileWatcherStartAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileWatcherStartAction: %v", v)
	}
}

// FileWatcherStartResponseBody - Started
type FileWatcherStartResponseBody struct {
	// Result of the request
	Result FileWatcherStartResult `json:"result"`
	// The id of the action
	Action FileWatcherStartAction `json:"action"`
	Data   string                 `json:"data"`
}

func (o *FileWatcherStartResponseBody) GetResult() FileWatcherStartResult {
	if o == nil {
		return FileWatcherStartResult("")
	}
	return o.Result
}

func (o *FileWatcherStartResponseBody) GetAction() FileWatcherStartAction {
	if o == nil {
		return FileWatcherStartAction("")
	}
	return o.Action
}

func (o *FileWatcherStartResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type FileWatcherStartResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Started
	Object *FileWatcherStartResponseBody
}

func (o *FileWatcherStartResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FileWatcherStartResponse) GetObject() *FileWatcherStartResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
