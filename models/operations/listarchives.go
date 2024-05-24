package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ListArchivesRequest struct {
	// Type of archive to make
	ArchiveKind components.ArchiveKind `pathParam:"style=simple,explode=false,name=archiveKind"`
}

func (o *ListArchivesRequest) GetArchiveKind() components.ArchiveKind {
	if o == nil {
		return components.ArchiveKind("")
	}
	return o.ArchiveKind
}

// ListArchivesResult - Result of the request
type ListArchivesResult string

const (
	ListArchivesResultSuccess ListArchivesResult = "success"
	ListArchivesResultError   ListArchivesResult = "error"
)

func (e ListArchivesResult) ToPointer() *ListArchivesResult {
	return &e
}
func (e *ListArchivesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListArchivesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListArchivesResult: %v", v)
	}
}

// ListArchivesAction - The kind of the archive
type ListArchivesAction string

const (
	ListArchivesActionArchiveFull       ListArchivesAction = "archiveFull"
	ListArchivesActionArchiveGroups     ListArchivesAction = "archiveGroups"
	ListArchivesActionArchiveRules      ListArchivesAction = "archiveRules"
	ListArchivesActionArchiveDirectives ListArchivesAction = "archiveDirectives"
	ListArchivesActionArchiveParameters ListArchivesAction = "archiveParameters"
)

func (e ListArchivesAction) ToPointer() *ListArchivesAction {
	return &e
}
func (e *ListArchivesAction) UnmarshalJSON(data []byte) error {
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
		*e = ListArchivesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListArchivesAction: %v", v)
	}
}

type Full struct {
	Commiter  string `json:"commiter"`
	GitCommit string `json:"gitCommit"`
	ID        string `json:"id"`
}

func (o *Full) GetCommiter() string {
	if o == nil {
		return ""
	}
	return o.Commiter
}

func (o *Full) GetGitCommit() string {
	if o == nil {
		return ""
	}
	return o.GitCommit
}

func (o *Full) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListArchivesData struct {
	Full []Full `json:"full"`
}

func (o *ListArchivesData) GetFull() []Full {
	if o == nil {
		return []Full{}
	}
	return o.Full
}

// ListArchivesResponseBody - Success
type ListArchivesResponseBody struct {
	// Result of the request
	Result ListArchivesResult `json:"result"`
	// The kind of the archive
	Action ListArchivesAction `json:"action"`
	Data   ListArchivesData   `json:"data"`
}

func (o *ListArchivesResponseBody) GetResult() ListArchivesResult {
	if o == nil {
		return ListArchivesResult("")
	}
	return o.Result
}

func (o *ListArchivesResponseBody) GetAction() ListArchivesAction {
	if o == nil {
		return ListArchivesAction("")
	}
	return o.Action
}

func (o *ListArchivesResponseBody) GetData() ListArchivesData {
	if o == nil {
		return ListArchivesData{}
	}
	return o.Data
}

type ListArchivesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *ListArchivesResponseBody
}

func (o *ListArchivesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListArchivesResponse) GetObject() *ListArchivesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
