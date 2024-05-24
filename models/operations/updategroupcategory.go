package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type UpdateGroupCategoryRequest struct {
	GroupCategoryID     string                         `pathParam:"style=simple,explode=false,name=groupCategoryId"`
	GroupCategoryUpdate components.GroupCategoryUpdate `request:"mediaType=application/json"`
}

func (o *UpdateGroupCategoryRequest) GetGroupCategoryID() string {
	if o == nil {
		return ""
	}
	return o.GroupCategoryID
}

func (o *UpdateGroupCategoryRequest) GetGroupCategoryUpdate() components.GroupCategoryUpdate {
	if o == nil {
		return components.GroupCategoryUpdate{}
	}
	return o.GroupCategoryUpdate
}

// UpdateGroupCategoryResult - Result of the request
type UpdateGroupCategoryResult string

const (
	UpdateGroupCategoryResultSuccess UpdateGroupCategoryResult = "success"
	UpdateGroupCategoryResultError   UpdateGroupCategoryResult = "error"
)

func (e UpdateGroupCategoryResult) ToPointer() *UpdateGroupCategoryResult {
	return &e
}
func (e *UpdateGroupCategoryResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateGroupCategoryResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateGroupCategoryResult: %v", v)
	}
}

// UpdateGroupCategoryAction - The id of the action
type UpdateGroupCategoryAction string

const (
	UpdateGroupCategoryActionUpdateGroupCategory UpdateGroupCategoryAction = "UpdateGroupCategory"
)

func (e UpdateGroupCategoryAction) ToPointer() *UpdateGroupCategoryAction {
	return &e
}
func (e *UpdateGroupCategoryAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UpdateGroupCategory":
		*e = UpdateGroupCategoryAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateGroupCategoryAction: %v", v)
	}
}

type UpdateGroupCategoryData struct {
	GroupCategories []components.GroupCategory `json:"groupCategories"`
}

func (o *UpdateGroupCategoryData) GetGroupCategories() []components.GroupCategory {
	if o == nil {
		return []components.GroupCategory{}
	}
	return o.GroupCategories
}

// UpdateGroupCategoryResponseBody - Groups category information
type UpdateGroupCategoryResponseBody struct {
	// Result of the request
	Result UpdateGroupCategoryResult `json:"result"`
	// The id of the action
	Action UpdateGroupCategoryAction `json:"action"`
	Data   UpdateGroupCategoryData   `json:"data"`
}

func (o *UpdateGroupCategoryResponseBody) GetResult() UpdateGroupCategoryResult {
	if o == nil {
		return UpdateGroupCategoryResult("")
	}
	return o.Result
}

func (o *UpdateGroupCategoryResponseBody) GetAction() UpdateGroupCategoryAction {
	if o == nil {
		return UpdateGroupCategoryAction("")
	}
	return o.Action
}

func (o *UpdateGroupCategoryResponseBody) GetData() UpdateGroupCategoryData {
	if o == nil {
		return UpdateGroupCategoryData{}
	}
	return o.Data
}

type UpdateGroupCategoryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups category information
	Object *UpdateGroupCategoryResponseBody
}

func (o *UpdateGroupCategoryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateGroupCategoryResponse) GetObject() *UpdateGroupCategoryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
