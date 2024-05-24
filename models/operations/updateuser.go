package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type UpdateUserRequest struct {
	// Username of an user (unique)
	Username string           `pathParam:"style=simple,explode=false,name=username"`
	Users    components.Users `request:"mediaType=application/json"`
}

func (o *UpdateUserRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *UpdateUserRequest) GetUsers() components.Users {
	if o == nil {
		return components.Users{}
	}
	return o.Users
}

// UpdateUserResult - Result of the request
type UpdateUserResult string

const (
	UpdateUserResultSuccess UpdateUserResult = "success"
	UpdateUserResultError   UpdateUserResult = "error"
)

func (e UpdateUserResult) ToPointer() *UpdateUserResult {
	return &e
}
func (e *UpdateUserResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateUserResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateUserResult: %v", v)
	}
}

// UpdateUserAction - The id of the action
type UpdateUserAction string

const (
	UpdateUserActionUpdateUser UpdateUserAction = "updateUser"
)

func (e UpdateUserAction) ToPointer() *UpdateUserAction {
	return &e
}
func (e *UpdateUserAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updateUser":
		*e = UpdateUserAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateUserAction: %v", v)
	}
}

type UpdatedUser struct {
	// New Username
	Username string `json:"username"`
	// New password given
	Password string `json:"password"`
	// defined user's permissions
	Role []string `json:"role"`
}

func (o *UpdatedUser) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *UpdatedUser) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *UpdatedUser) GetRole() []string {
	if o == nil {
		return []string{}
	}
	return o.Role
}

type UpdateUserData struct {
	UpdatedUser UpdatedUser `json:"updatedUser"`
}

func (o *UpdateUserData) GetUpdatedUser() UpdatedUser {
	if o == nil {
		return UpdatedUser{}
	}
	return o.UpdatedUser
}

// UpdateUserResponseBody - Updated
type UpdateUserResponseBody struct {
	// Result of the request
	Result UpdateUserResult `json:"result"`
	// The id of the action
	Action UpdateUserAction `json:"action"`
	Data   UpdateUserData   `json:"data"`
}

func (o *UpdateUserResponseBody) GetResult() UpdateUserResult {
	if o == nil {
		return UpdateUserResult("")
	}
	return o.Result
}

func (o *UpdateUserResponseBody) GetAction() UpdateUserAction {
	if o == nil {
		return UpdateUserAction("")
	}
	return o.Action
}

func (o *UpdateUserResponseBody) GetData() UpdateUserData {
	if o == nil {
		return UpdateUserData{}
	}
	return o.Data
}

type UpdateUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Updated
	Object *UpdateUserResponseBody
}

func (o *UpdateUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateUserResponse) GetObject() *UpdateUserResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
