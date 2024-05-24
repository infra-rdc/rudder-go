package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetUserInfoResult - Result of the request
type GetUserInfoResult string

const (
	GetUserInfoResultSuccess GetUserInfoResult = "success"
	GetUserInfoResultError   GetUserInfoResult = "error"
)

func (e GetUserInfoResult) ToPointer() *GetUserInfoResult {
	return &e
}
func (e *GetUserInfoResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetUserInfoResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetUserInfoResult: %v", v)
	}
}

// GetUserInfoAction - The id of the action
type GetUserInfoAction string

const (
	GetUserInfoActionGetUserInfo GetUserInfoAction = "getUserInfo"
)

func (e GetUserInfoAction) ToPointer() *GetUserInfoAction {
	return &e
}
func (e *GetUserInfoAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getUserInfo":
		*e = GetUserInfoAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetUserInfoAction: %v", v)
	}
}

type Digest string

const (
	DigestBcrypt Digest = "BCRYPT"
	DigestSha512 Digest = "SHA512"
	DigestSha256 Digest = "SHA256"
	DigestSha1   Digest = "SHA1"
	DigestMd5    Digest = "MD5"
)

func (e Digest) ToPointer() *Digest {
	return &e
}
func (e *Digest) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BCRYPT":
		fallthrough
	case "SHA512":
		fallthrough
	case "SHA256":
		fallthrough
	case "SHA1":
		fallthrough
	case "MD5":
		*e = Digest(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Digest: %v", v)
	}
}

type GetUserInfoData struct {
	Digest Digest             `json:"digest"`
	Users  []components.Users `json:"users"`
}

func (o *GetUserInfoData) GetDigest() Digest {
	if o == nil {
		return Digest("")
	}
	return o.Digest
}

func (o *GetUserInfoData) GetUsers() []components.Users {
	if o == nil {
		return []components.Users{}
	}
	return o.Users
}

// GetUserInfoResponseBody - Users information
type GetUserInfoResponseBody struct {
	// Result of the request
	Result GetUserInfoResult `json:"result"`
	// The id of the action
	Action GetUserInfoAction `json:"action"`
	Data   GetUserInfoData   `json:"data"`
}

func (o *GetUserInfoResponseBody) GetResult() GetUserInfoResult {
	if o == nil {
		return GetUserInfoResult("")
	}
	return o.Result
}

func (o *GetUserInfoResponseBody) GetAction() GetUserInfoAction {
	if o == nil {
		return GetUserInfoAction("")
	}
	return o.Action
}

func (o *GetUserInfoResponseBody) GetData() GetUserInfoData {
	if o == nil {
		return GetUserInfoData{}
	}
	return o.Data
}

type GetUserInfoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Users information
	Object *GetUserInfoResponseBody
}

func (o *GetUserInfoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserInfoResponse) GetObject() *GetUserInfoResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
