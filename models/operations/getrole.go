package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetRoleResult - Result of the request
type GetRoleResult string

const (
	GetRoleResultSuccess GetRoleResult = "success"
	GetRoleResultError   GetRoleResult = "error"
)

func (e GetRoleResult) ToPointer() *GetRoleResult {
	return &e
}
func (e *GetRoleResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetRoleResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRoleResult: %v", v)
	}
}

// GetRoleAction - The id of the action
type GetRoleAction string

const (
	GetRoleActionGetRole GetRoleAction = "getRole"
)

func (e GetRoleAction) ToPointer() *GetRoleAction {
	return &e
}
func (e *GetRoleAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getRole":
		*e = GetRoleAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRoleAction: %v", v)
	}
}

// ID - Name of the role
type ID string

const (
	IDInventory     ID = "inventory"
	IDCompliance    ID = "compliance"
	IDAdministrator ID = "administrator"
	IDEtc           ID = "etc"
)

func (e ID) ToPointer() *ID {
	return &e
}
func (e *ID) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "inventory":
		fallthrough
	case "compliance":
		fallthrough
	case "administrator":
		fallthrough
	case "etc":
		*e = ID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ID: %v", v)
	}
}

type Rights string

const (
	RightsNodeRead       Rights = "node_read"
	RightsUserAccountAll Rights = "userAccount_all"
)

func (e Rights) ToPointer() *Rights {
	return &e
}
func (e *Rights) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "node_read":
		fallthrough
	case "userAccount_all":
		*e = Rights(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Rights: %v", v)
	}
}

type GetRoleData struct {
	// Name of the role
	ID ID `json:"id"`
	// Role's rights
	Rights []Rights `json:"rights"`
}

func (o *GetRoleData) GetID() ID {
	if o == nil {
		return ID("")
	}
	return o.ID
}

func (o *GetRoleData) GetRights() []Rights {
	if o == nil {
		return []Rights{}
	}
	return o.Rights
}

// GetRoleResponseBody - Roles information
type GetRoleResponseBody struct {
	// Result of the request
	Result GetRoleResult `json:"result"`
	// The id of the action
	Action GetRoleAction `json:"action"`
	Data   []GetRoleData `json:"data"`
}

func (o *GetRoleResponseBody) GetResult() GetRoleResult {
	if o == nil {
		return GetRoleResult("")
	}
	return o.Result
}

func (o *GetRoleResponseBody) GetAction() GetRoleAction {
	if o == nil {
		return GetRoleAction("")
	}
	return o.Action
}

func (o *GetRoleResponseBody) GetData() []GetRoleData {
	if o == nil {
		return []GetRoleData{}
	}
	return o.Data
}

type GetRoleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Roles information
	Object *GetRoleResponseBody
}

func (o *GetRoleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRoleResponse) GetObject() *GetRoleResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
