package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetGroupCategoryDetailsRequest struct {
	GroupCategoryID string `pathParam:"style=simple,explode=false,name=groupCategoryId"`
}

func (o *GetGroupCategoryDetailsRequest) GetGroupCategoryID() string {
	if o == nil {
		return ""
	}
	return o.GroupCategoryID
}

// GetGroupCategoryDetailsResult - Result of the request
type GetGroupCategoryDetailsResult string

const (
	GetGroupCategoryDetailsResultSuccess GetGroupCategoryDetailsResult = "success"
	GetGroupCategoryDetailsResultError   GetGroupCategoryDetailsResult = "error"
)

func (e GetGroupCategoryDetailsResult) ToPointer() *GetGroupCategoryDetailsResult {
	return &e
}
func (e *GetGroupCategoryDetailsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetGroupCategoryDetailsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetGroupCategoryDetailsResult: %v", v)
	}
}

// GetGroupCategoryDetailsAction - The id of the action
type GetGroupCategoryDetailsAction string

const (
	GetGroupCategoryDetailsActionGetGroupCategoryDetails GetGroupCategoryDetailsAction = "GetGroupCategoryDetails"
)

func (e GetGroupCategoryDetailsAction) ToPointer() *GetGroupCategoryDetailsAction {
	return &e
}
func (e *GetGroupCategoryDetailsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GetGroupCategoryDetails":
		*e = GetGroupCategoryDetailsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetGroupCategoryDetailsAction: %v", v)
	}
}

type GetGroupCategoryDetailsData struct {
	GroupCategories []components.GroupCategory `json:"groupCategories"`
}

func (o *GetGroupCategoryDetailsData) GetGroupCategories() []components.GroupCategory {
	if o == nil {
		return []components.GroupCategory{}
	}
	return o.GroupCategories
}

// GetGroupCategoryDetailsResponseBody - Groups category information
type GetGroupCategoryDetailsResponseBody struct {
	// Result of the request
	Result GetGroupCategoryDetailsResult `json:"result"`
	// The id of the action
	Action GetGroupCategoryDetailsAction `json:"action"`
	Data   GetGroupCategoryDetailsData   `json:"data"`
}

func (o *GetGroupCategoryDetailsResponseBody) GetResult() GetGroupCategoryDetailsResult {
	if o == nil {
		return GetGroupCategoryDetailsResult("")
	}
	return o.Result
}

func (o *GetGroupCategoryDetailsResponseBody) GetAction() GetGroupCategoryDetailsAction {
	if o == nil {
		return GetGroupCategoryDetailsAction("")
	}
	return o.Action
}

func (o *GetGroupCategoryDetailsResponseBody) GetData() GetGroupCategoryDetailsData {
	if o == nil {
		return GetGroupCategoryDetailsData{}
	}
	return o.Data
}

type GetGroupCategoryDetailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups category information
	Object *GetGroupCategoryDetailsResponseBody
}

func (o *GetGroupCategoryDetailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetGroupCategoryDetailsResponse) GetObject() *GetGroupCategoryDetailsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
