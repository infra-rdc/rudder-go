package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type CreateArchiveRequest struct {
	// Type of archive to make
	ArchiveKind components.ArchiveKind `pathParam:"style=simple,explode=false,name=archiveKind"`
}

func (o *CreateArchiveRequest) GetArchiveKind() components.ArchiveKind {
	if o == nil {
		return components.ArchiveKind("")
	}
	return o.ArchiveKind
}

// CreateArchiveResult - Result of the request
type CreateArchiveResult string

const (
	CreateArchiveResultSuccess CreateArchiveResult = "success"
	CreateArchiveResultError   CreateArchiveResult = "error"
)

func (e CreateArchiveResult) ToPointer() *CreateArchiveResult {
	return &e
}
func (e *CreateArchiveResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = CreateArchiveResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateArchiveResult: %v", v)
	}
}

// CreateArchiveAction - The kind of the archive
type CreateArchiveAction string

const (
	CreateArchiveActionArchiveFull       CreateArchiveAction = "archiveFull"
	CreateArchiveActionArchiveGroups     CreateArchiveAction = "archiveGroups"
	CreateArchiveActionArchiveRules      CreateArchiveAction = "archiveRules"
	CreateArchiveActionArchiveDirectives CreateArchiveAction = "archiveDirectives"
	CreateArchiveActionArchiveParameters CreateArchiveAction = "archiveParameters"
)

func (e CreateArchiveAction) ToPointer() *CreateArchiveAction {
	return &e
}
func (e *CreateArchiveAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "archiveFull":
		fallthrough
	case "archiveGroups":
		fallthrough
	case "archiveRules":
		fallthrough
	case "archiveDirectives":
		fallthrough
	case "archiveParameters":
		*e = CreateArchiveAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateArchiveAction: %v", v)
	}
}

type CreateArchiveFull struct {
	Commiter  string `json:"commiter"`
	GitCommit string `json:"gitCommit"`
	ID        string `json:"id"`
}

func (o *CreateArchiveFull) GetCommiter() string {
	if o == nil {
		return ""
	}
	return o.Commiter
}

func (o *CreateArchiveFull) GetGitCommit() string {
	if o == nil {
		return ""
	}
	return o.GitCommit
}

func (o *CreateArchiveFull) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type CreateArchiveData struct {
	Full CreateArchiveFull `json:"full"`
}

func (o *CreateArchiveData) GetFull() CreateArchiveFull {
	if o == nil {
		return CreateArchiveFull{}
	}
	return o.Full
}

// CreateArchiveResponseBody - Success
type CreateArchiveResponseBody struct {
	// Result of the request
	Result CreateArchiveResult `json:"result"`
	// The kind of the archive
	Action CreateArchiveAction `json:"action"`
	Data   CreateArchiveData   `json:"data"`
}

func (o *CreateArchiveResponseBody) GetResult() CreateArchiveResult {
	if o == nil {
		return CreateArchiveResult("")
	}
	return o.Result
}

func (o *CreateArchiveResponseBody) GetAction() CreateArchiveAction {
	if o == nil {
		return CreateArchiveAction("")
	}
	return o.Action
}

func (o *CreateArchiveResponseBody) GetData() CreateArchiveData {
	if o == nil {
		return CreateArchiveData{}
	}
	return o.Data
}

type CreateArchiveResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *CreateArchiveResponseBody
}

func (o *CreateArchiveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateArchiveResponse) GetObject() *CreateArchiveResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
