package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type File struct {
	FileName string `multipartForm:"name=file"`
	Content  []byte `multipartForm:"content"`
}

func (o *File) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *File) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

type Signature struct {
	FileName string `multipartForm:"name=signature"`
	Content  []byte `multipartForm:"content"`
}

func (o *Signature) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *Signature) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

type UploadInventoryRequestBody struct {
	// The inventory file. The original file name is used to check extension, that should be .xml[.gz] or .ocs[.gz]
	File *File `multipartForm:"file"`
	// The signature file. The original file name is used to check extension, that should be ${originalInventoryFileName}.sign[.gz]
	Signature *Signature `multipartForm:"file"`
}

func (o *UploadInventoryRequestBody) GetFile() *File {
	if o == nil {
		return nil
	}
	return o.File
}

func (o *UploadInventoryRequestBody) GetSignature() *Signature {
	if o == nil {
		return nil
	}
	return o.Signature
}

// UploadInventoryResult - Result of the request
type UploadInventoryResult string

const (
	UploadInventoryResultSuccess UploadInventoryResult = "success"
	UploadInventoryResultError   UploadInventoryResult = "error"
)

func (e UploadInventoryResult) ToPointer() *UploadInventoryResult {
	return &e
}
func (e *UploadInventoryResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UploadInventoryResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UploadInventoryResult: %v", v)
	}
}

// UploadInventoryAction - The id of the action
type UploadInventoryAction string

const (
	UploadInventoryActionUploadInventory UploadInventoryAction = "uploadInventory"
)

func (e UploadInventoryAction) ToPointer() *UploadInventoryAction {
	return &e
}
func (e *UploadInventoryAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "uploadInventory":
		*e = UploadInventoryAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UploadInventoryAction: %v", v)
	}
}

// UploadInventoryResponseBody - Inventory uploaded
type UploadInventoryResponseBody struct {
	// Result of the request
	Result UploadInventoryResult `json:"result"`
	// The id of the action
	Action UploadInventoryAction `json:"action"`
	Data   string                `json:"data"`
}

func (o *UploadInventoryResponseBody) GetResult() UploadInventoryResult {
	if o == nil {
		return UploadInventoryResult("")
	}
	return o.Result
}

func (o *UploadInventoryResponseBody) GetAction() UploadInventoryAction {
	if o == nil {
		return UploadInventoryAction("")
	}
	return o.Action
}

func (o *UploadInventoryResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type UploadInventoryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Inventory uploaded
	Object *UploadInventoryResponseBody
}

func (o *UploadInventoryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UploadInventoryResponse) GetObject() *UploadInventoryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
