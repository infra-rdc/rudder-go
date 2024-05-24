package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type Archive struct {
	FileName string `multipartForm:"name=archive"`
	Content  []byte `multipartForm:"content"`
}

func (o *Archive) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *Archive) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

// Merge - Optional merge algo of the import. Default `override-all` means what is in the archive is the new reality. `keep-rule-groups` will keep existing target definition for existing rules (ignore archive value).
type Merge string

const (
	MergeOverrideAll    Merge = "override-all"
	MergeKeepRuleGroups Merge = "keep-rule-groups"
)

func (e Merge) ToPointer() *Merge {
	return &e
}
func (e *Merge) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "override-all":
		fallthrough
	case "keep-rule-groups":
		*e = Merge(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Merge: %v", v)
	}
}

type ImportRequestBody struct {
	// The ZIP archive file containing policies in a conventional layout and serialization format
	Archive *Archive `multipartForm:"file"`
	// Optional merge algo of the import. Default `override-all` means what is in the archive is the new reality. `keep-rule-groups` will keep existing target definition for existing rules (ignore archive value).
	Merge *Merge `multipartForm:"name=merge"`
}

func (o *ImportRequestBody) GetArchive() *Archive {
	if o == nil {
		return nil
	}
	return o.Archive
}

func (o *ImportRequestBody) GetMerge() *Merge {
	if o == nil {
		return nil
	}
	return o.Merge
}

// ImportResult - Result of the request
type ImportResult string

const (
	ImportResultSuccess ImportResult = "success"
	ImportResultError   ImportResult = "error"
)

func (e ImportResult) ToPointer() *ImportResult {
	return &e
}
func (e *ImportResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ImportResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ImportResult: %v", v)
	}
}

// ImportAction - The id of the action
type ImportAction string

const (
	ImportActionImport ImportAction = "import"
)

func (e ImportAction) ToPointer() *ImportAction {
	return &e
}
func (e *ImportAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "import":
		*e = ImportAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ImportAction: %v", v)
	}
}

// ImportData - Details about archive import process
type ImportData struct {
	ArchiveImported *bool `json:"archiveImported,omitempty"`
}

func (o *ImportData) GetArchiveImported() *bool {
	if o == nil {
		return nil
	}
	return o.ArchiveImported
}

// ImportResponseBody - Archive imported
type ImportResponseBody struct {
	// Result of the request
	Result ImportResult `json:"result"`
	// The id of the action
	Action ImportAction `json:"action"`
	// Details about archive import process
	Data ImportData `json:"data"`
}

func (o *ImportResponseBody) GetResult() ImportResult {
	if o == nil {
		return ImportResult("")
	}
	return o.Result
}

func (o *ImportResponseBody) GetAction() ImportAction {
	if o == nil {
		return ImportAction("")
	}
	return o.Action
}

func (o *ImportResponseBody) GetData() ImportData {
	if o == nil {
		return ImportData{}
	}
	return o.Data
}

type ImportResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Archive imported
	Object *ImportResponseBody
}

func (o *ImportResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ImportResponse) GetObject() *ImportResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
