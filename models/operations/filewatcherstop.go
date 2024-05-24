package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// FileWatcherStopResult - Result of the request
type FileWatcherStopResult string

const (
	FileWatcherStopResultSuccess FileWatcherStopResult = "success"
	FileWatcherStopResultError   FileWatcherStopResult = "error"
)

func (e FileWatcherStopResult) ToPointer() *FileWatcherStopResult {
	return &e
}
func (e *FileWatcherStopResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = FileWatcherStopResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileWatcherStopResult: %v", v)
	}
}

// FileWatcherStopAction - The id of the action
type FileWatcherStopAction string

const (
	FileWatcherStopActionFileWatcherStop FileWatcherStopAction = "fileWatcherStop"
)

func (e FileWatcherStopAction) ToPointer() *FileWatcherStopAction {
	return &e
}
func (e *FileWatcherStopAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fileWatcherStop":
		*e = FileWatcherStopAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileWatcherStopAction: %v", v)
	}
}

// FileWatcherStopResponseBody - Stopped
type FileWatcherStopResponseBody struct {
	// Result of the request
	Result FileWatcherStopResult `json:"result"`
	// The id of the action
	Action FileWatcherStopAction `json:"action"`
	Data   string                `json:"data"`
}

func (o *FileWatcherStopResponseBody) GetResult() FileWatcherStopResult {
	if o == nil {
		return FileWatcherStopResult("")
	}
	return o.Result
}

func (o *FileWatcherStopResponseBody) GetAction() FileWatcherStopAction {
	if o == nil {
		return FileWatcherStopAction("")
	}
	return o.Action
}

func (o *FileWatcherStopResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type FileWatcherStopResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stopped
	Object *FileWatcherStopResponseBody
}

func (o *FileWatcherStopResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FileWatcherStopResponse) GetObject() *FileWatcherStopResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
