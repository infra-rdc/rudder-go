package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// AddUserResult - Result of the request
type AddUserResult string

const (
	AddUserResultSuccess AddUserResult = "success"
	AddUserResultError   AddUserResult = "error"
)

func (e AddUserResult) ToPointer() *AddUserResult {
	return &e
}
func (e *AddUserResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = AddUserResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddUserResult: %v", v)
	}
}

// AddUserAction - The id of the action
type AddUserAction string

const (
	AddUserActionAddUser AddUserAction = "addUser"
)

func (e AddUserAction) ToPointer() *AddUserAction {
	return &e
}
func (e *AddUserAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "addUser":
		*e = AddUserAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddUserAction: %v", v)
	}
}

type AddedUser struct {
	Username *string `json:"username,omitempty"`
	// this password will be hashed for you if the `isPreHashed` is set on false
	Password *string `json:"password,omitempty"`
	// defined user's permissions
	Role []string `json:"role,omitempty"`
}

func (o *AddedUser) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *AddedUser) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *AddedUser) GetRole() []string {
	if o == nil {
		return nil
	}
	return o.Role
}

type AddUserData struct {
	AddedUser AddedUser `json:"addedUser"`
}

func (o *AddUserData) GetAddedUser() AddedUser {
	if o == nil {
		return AddedUser{}
	}
	return o.AddedUser
}

// AddUserResponseBody - Updated
type AddUserResponseBody struct {
	// Result of the request
	Result AddUserResult `json:"result"`
	// The id of the action
	Action AddUserAction `json:"action"`
	Data   AddUserData   `json:"data"`
}

func (o *AddUserResponseBody) GetResult() AddUserResult {
	if o == nil {
		return AddUserResult("")
	}
	return o.Result
}

func (o *AddUserResponseBody) GetAction() AddUserAction {
	if o == nil {
		return AddUserAction("")
	}
	return o.Action
}

func (o *AddUserResponseBody) GetData() AddUserData {
	if o == nil {
		return AddUserData{}
	}
	return o.Data
}

type AddUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Updated
	Object *AddUserResponseBody
}

func (o *AddUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AddUserResponse) GetObject() *AddUserResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
