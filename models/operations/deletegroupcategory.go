package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteGroupCategoryRequest struct {
	GroupCategoryID string `pathParam:"style=simple,explode=false,name=groupCategoryId"`
}

func (o *DeleteGroupCategoryRequest) GetGroupCategoryID() string {
	if o == nil {
		return ""
	}
	return o.GroupCategoryID
}

// DeleteGroupCategoryResult - Result of the request
type DeleteGroupCategoryResult string

const (
	DeleteGroupCategoryResultSuccess DeleteGroupCategoryResult = "success"
	DeleteGroupCategoryResultError   DeleteGroupCategoryResult = "error"
)

func (e DeleteGroupCategoryResult) ToPointer() *DeleteGroupCategoryResult {
	return &e
}
func (e *DeleteGroupCategoryResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteGroupCategoryResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteGroupCategoryResult: %v", v)
	}
}

// DeleteGroupCategoryAction - The id of the action
type DeleteGroupCategoryAction string

const (
	DeleteGroupCategoryActionDeleteGroupCategory DeleteGroupCategoryAction = "DeleteGroupCategory"
)

func (e DeleteGroupCategoryAction) ToPointer() *DeleteGroupCategoryAction {
	return &e
}
func (e *DeleteGroupCategoryAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DeleteGroupCategory":
		*e = DeleteGroupCategoryAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteGroupCategoryAction: %v", v)
	}
}

type DeleteGroupCategoryData struct {
	GroupCategories []components.GroupCategory `json:"groupCategories"`
}

func (o *DeleteGroupCategoryData) GetGroupCategories() []components.GroupCategory {
	if o == nil {
		return []components.GroupCategory{}
	}
	return o.GroupCategories
}

// DeleteGroupCategoryResponseBody - Groups category information
type DeleteGroupCategoryResponseBody struct {
	// Result of the request
	Result DeleteGroupCategoryResult `json:"result"`
	// The id of the action
	Action DeleteGroupCategoryAction `json:"action"`
	Data   DeleteGroupCategoryData   `json:"data"`
}

func (o *DeleteGroupCategoryResponseBody) GetResult() DeleteGroupCategoryResult {
	if o == nil {
		return DeleteGroupCategoryResult("")
	}
	return o.Result
}

func (o *DeleteGroupCategoryResponseBody) GetAction() DeleteGroupCategoryAction {
	if o == nil {
		return DeleteGroupCategoryAction("")
	}
	return o.Action
}

func (o *DeleteGroupCategoryResponseBody) GetData() DeleteGroupCategoryData {
	if o == nil {
		return DeleteGroupCategoryData{}
	}
	return o.Data
}

type DeleteGroupCategoryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups category information
	Object *DeleteGroupCategoryResponseBody
}

func (o *DeleteGroupCategoryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteGroupCategoryResponse) GetObject() *DeleteGroupCategoryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
