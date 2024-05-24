package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// FileWatcherRestartResult - Result of the request
type FileWatcherRestartResult string

const (
	FileWatcherRestartResultSuccess FileWatcherRestartResult = "success"
	FileWatcherRestartResultError   FileWatcherRestartResult = "error"
)

func (e FileWatcherRestartResult) ToPointer() *FileWatcherRestartResult {
	return &e
}
func (e *FileWatcherRestartResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = FileWatcherRestartResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileWatcherRestartResult: %v", v)
	}
}

// FileWatcherRestartAction - The id of the action
type FileWatcherRestartAction string

const (
	FileWatcherRestartActionFileWatcherRestart FileWatcherRestartAction = "fileWatcherRestart"
)

func (e FileWatcherRestartAction) ToPointer() *FileWatcherRestartAction {
	return &e
}
func (e *FileWatcherRestartAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fileWatcherRestart":
		*e = FileWatcherRestartAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileWatcherRestartAction: %v", v)
	}
}

// FileWatcherRestartResponseBody - Started
type FileWatcherRestartResponseBody struct {
	// Result of the request
	Result FileWatcherRestartResult `json:"result"`
	// The id of the action
	Action FileWatcherRestartAction `json:"action"`
	Data   string                   `json:"data"`
}

func (o *FileWatcherRestartResponseBody) GetResult() FileWatcherRestartResult {
	if o == nil {
		return FileWatcherRestartResult("")
	}
	return o.Result
}

func (o *FileWatcherRestartResponseBody) GetAction() FileWatcherRestartAction {
	if o == nil {
		return FileWatcherRestartAction("")
	}
	return o.Action
}

func (o *FileWatcherRestartResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type FileWatcherRestartResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Started
	Object *FileWatcherRestartResponseBody
}

func (o *FileWatcherRestartResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FileWatcherRestartResponse) GetObject() *FileWatcherRestartResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
