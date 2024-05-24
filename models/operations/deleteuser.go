package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteUserRequest struct {
	// Username of an user (unique)
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *DeleteUserRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

// DeleteUserResult - Result of the request
type DeleteUserResult string

const (
	DeleteUserResultSuccess DeleteUserResult = "success"
	DeleteUserResultError   DeleteUserResult = "error"
)

func (e DeleteUserResult) ToPointer() *DeleteUserResult {
	return &e
}
func (e *DeleteUserResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteUserResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteUserResult: %v", v)
	}
}

// DeleteUserAction - The id of the action
type DeleteUserAction string

const (
	DeleteUserActionDeleteUser DeleteUserAction = "deleteUser"
)

func (e DeleteUserAction) ToPointer() *DeleteUserAction {
	return &e
}
func (e *DeleteUserAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleteUser":
		*e = DeleteUserAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteUserAction: %v", v)
	}
}

type DeletedUser struct {
	// Username of the deleted user
	Username string `json:"username"`
}

func (o *DeletedUser) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type DeleteUserData struct {
	DeletedUser DeletedUser `json:"deletedUser"`
}

func (o *DeleteUserData) GetDeletedUser() DeletedUser {
	if o == nil {
		return DeletedUser{}
	}
	return o.DeletedUser
}

// DeleteUserResponseBody - Deleted user
type DeleteUserResponseBody struct {
	// Result of the request
	Result DeleteUserResult `json:"result"`
	// The id of the action
	Action DeleteUserAction `json:"action"`
	Data   DeleteUserData   `json:"data"`
}

func (o *DeleteUserResponseBody) GetResult() DeleteUserResult {
	if o == nil {
		return DeleteUserResult("")
	}
	return o.Result
}

func (o *DeleteUserResponseBody) GetAction() DeleteUserAction {
	if o == nil {
		return DeleteUserAction("")
	}
	return o.Action
}

func (o *DeleteUserResponseBody) GetData() DeleteUserData {
	if o == nil {
		return DeleteUserData{}
	}
	return o.Data
}

type DeleteUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Deleted user
	Object *DeleteUserResponseBody
}

func (o *DeleteUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteUserResponse) GetObject() *DeleteUserResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
