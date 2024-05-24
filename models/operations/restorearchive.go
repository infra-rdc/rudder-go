package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ArchiveRestoreKind - What archive to restore (latest archive, latest commit in configuration repository, or archive with ID as given by listArchive)
type ArchiveRestoreKind string

const (
	ArchiveRestoreKindLatestArchive ArchiveRestoreKind = "latestArchive"
	ArchiveRestoreKindLatestCommit  ArchiveRestoreKind = "latestCommit"
	ArchiveRestoreKindArchiveID     ArchiveRestoreKind = "archive ID"
)

func (e ArchiveRestoreKind) ToPointer() *ArchiveRestoreKind {
	return &e
}
func (e *ArchiveRestoreKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "latestArchive":
		fallthrough
	case "latestCommit":
		fallthrough
	case "archive ID":
		*e = ArchiveRestoreKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ArchiveRestoreKind: %v", v)
	}
}

type RestoreArchiveRequest struct {
	// Type of archive to make
	ArchiveKind components.ArchiveKind `pathParam:"style=simple,explode=false,name=archiveKind"`
	// What archive to restore (latest archive, latest commit in configuration repository, or archive with ID as given by listArchive)
	ArchiveRestoreKind ArchiveRestoreKind `pathParam:"style=simple,explode=false,name=archiveRestoreKind"`
}

func (o *RestoreArchiveRequest) GetArchiveKind() components.ArchiveKind {
	if o == nil {
		return components.ArchiveKind("")
	}
	return o.ArchiveKind
}

func (o *RestoreArchiveRequest) GetArchiveRestoreKind() ArchiveRestoreKind {
	if o == nil {
		return ArchiveRestoreKind("")
	}
	return o.ArchiveRestoreKind
}

// RestoreArchiveResult - Result of the request
type RestoreArchiveResult string

const (
	RestoreArchiveResultSuccess RestoreArchiveResult = "success"
	RestoreArchiveResultError   RestoreArchiveResult = "error"
)

func (e RestoreArchiveResult) ToPointer() *RestoreArchiveResult {
	return &e
}
func (e *RestoreArchiveResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = RestoreArchiveResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RestoreArchiveResult: %v", v)
	}
}

// RestoreArchiveAction - The kind of the archive
type RestoreArchiveAction string

const (
	RestoreArchiveActionRestoreFullLatestArchive       RestoreArchiveAction = "restoreFullLatestArchive"
	RestoreArchiveActionRestoreGroupLatestArchive      RestoreArchiveAction = "restoreGroupLatestArchive"
	RestoreArchiveActionRestoreRulesLatestArchive      RestoreArchiveAction = "restoreRulesLatestArchive"
	RestoreArchiveActionRestoreDirectivesLatestArchive RestoreArchiveAction = "restoreDirectivesLatestArchive"
	RestoreArchiveActionRestoreParametersLatestArchive RestoreArchiveAction = "restoreParametersLatestArchive"
	RestoreArchiveActionRestoreFullLatestCommit        RestoreArchiveAction = "restoreFullLatestCommit"
	RestoreArchiveActionRestoreGroupLatestCommit       RestoreArchiveAction = "restoreGroupLatestCommit"
	RestoreArchiveActionRestoreRulesLatestCommit       RestoreArchiveAction = "restoreRulesLatestCommit"
	RestoreArchiveActionRestoreDirectivesLatestCommit  RestoreArchiveAction = "restoreDirectivesLatestCommit"
	RestoreArchiveActionRestoreParametersLatestCommit  RestoreArchiveAction = "restoreParametersLatestCommit"
	RestoreArchiveActionArchiveFullDateRestore         RestoreArchiveAction = "archiveFullDateRestore"
	RestoreArchiveActionArchiveGroupDateRestore        RestoreArchiveAction = "archiveGroupDateRestore"
	RestoreArchiveActionArchiveRulesDateRestore        RestoreArchiveAction = "archiveRulesDateRestore"
	RestoreArchiveActionArchiveDirectivesDateRestore   RestoreArchiveAction = "archiveDirectivesDateRestore"
	RestoreArchiveActionArchiveParametersDateRestore   RestoreArchiveAction = "archiveParametersDateRestore"
)

func (e RestoreArchiveAction) ToPointer() *RestoreArchiveAction {
	return &e
}
func (e *RestoreArchiveAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "restoreFullLatestArchive":
		fallthrough
	case "restoreGroupLatestArchive":
		fallthrough
	case "restoreRulesLatestArchive":
		fallthrough
	case "restoreDirectivesLatestArchive":
		fallthrough
	case "restoreParametersLatestArchive":
		fallthrough
	case "restoreFullLatestCommit":
		fallthrough
	case "restoreGroupLatestCommit":
		fallthrough
	case "restoreRulesLatestCommit":
		fallthrough
	case "restoreDirectivesLatestCommit":
		fallthrough
	case "restoreParametersLatestCommit":
		fallthrough
	case "archiveFullDateRestore":
		fallthrough
	case "archiveGroupDateRestore":
		fallthrough
	case "archiveRulesDateRestore":
		fallthrough
	case "archiveDirectivesDateRestore":
		fallthrough
	case "archiveParametersDateRestore":
		*e = RestoreArchiveAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RestoreArchiveAction: %v", v)
	}
}

type RestoreArchiveData struct {
	Full       *string `json:"full,omitempty"`
	Groups     *string `json:"groups,omitempty"`
	Rules      *string `json:"rules,omitempty"`
	Directive  *string `json:"directive,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
}

func (o *RestoreArchiveData) GetFull() *string {
	if o == nil {
		return nil
	}
	return o.Full
}

func (o *RestoreArchiveData) GetGroups() *string {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *RestoreArchiveData) GetRules() *string {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *RestoreArchiveData) GetDirective() *string {
	if o == nil {
		return nil
	}
	return o.Directive
}

func (o *RestoreArchiveData) GetParameters() *string {
	if o == nil {
		return nil
	}
	return o.Parameters
}

// RestoreArchiveResponseBody - Success
type RestoreArchiveResponseBody struct {
	// Result of the request
	Result RestoreArchiveResult `json:"result"`
	// The kind of the archive
	Action RestoreArchiveAction `json:"action"`
	Data   RestoreArchiveData   `json:"data"`
}

func (o *RestoreArchiveResponseBody) GetResult() RestoreArchiveResult {
	if o == nil {
		return RestoreArchiveResult("")
	}
	return o.Result
}

func (o *RestoreArchiveResponseBody) GetAction() RestoreArchiveAction {
	if o == nil {
		return RestoreArchiveAction("")
	}
	return o.Action
}

func (o *RestoreArchiveResponseBody) GetData() RestoreArchiveData {
	if o == nil {
		return RestoreArchiveData{}
	}
	return o.Data
}

type RestoreArchiveResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *RestoreArchiveResponseBody
}

func (o *RestoreArchiveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RestoreArchiveResponse) GetObject() *RestoreArchiveResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
