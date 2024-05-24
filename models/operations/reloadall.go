package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ReloadAllResult - Result of the request
type ReloadAllResult string

const (
	ReloadAllResultSuccess ReloadAllResult = "success"
	ReloadAllResultError   ReloadAllResult = "error"
)

func (e ReloadAllResult) ToPointer() *ReloadAllResult {
	return &e
}
func (e *ReloadAllResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadAllResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadAllResult: %v", v)
	}
}

// ReloadAllAction - The id of the action
type ReloadAllAction string

const (
	ReloadAllActionReloadAll ReloadAllAction = "reloadAll"
)

func (e ReloadAllAction) ToPointer() *ReloadAllAction {
	return &e
}
func (e *ReloadAllAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "reloadAll":
		*e = ReloadAllAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadAllAction: %v", v)
	}
}

type Groups string

const (
	GroupsStarted Groups = "Started"
)

func (e Groups) ToPointer() *Groups {
	return &e
}
func (e *Groups) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Started":
		*e = Groups(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Groups: %v", v)
	}
}

type Techniques string

const (
	TechniquesStarted Techniques = "Started"
)

func (e Techniques) ToPointer() *Techniques {
	return &e
}
func (e *Techniques) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Started":
		*e = Techniques(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Techniques: %v", v)
	}
}

type ReloadAllData struct {
	Groups     Groups     `json:"groups"`
	Techniques Techniques `json:"techniques"`
}

func (o *ReloadAllData) GetGroups() Groups {
	if o == nil {
		return Groups("")
	}
	return o.Groups
}

func (o *ReloadAllData) GetTechniques() Techniques {
	if o == nil {
		return Techniques("")
	}
	return o.Techniques
}

// ReloadAllResponseBody - Service reload
type ReloadAllResponseBody struct {
	// Result of the request
	Result ReloadAllResult `json:"result"`
	// The id of the action
	Action ReloadAllAction `json:"action"`
	Data   ReloadAllData   `json:"data"`
}

func (o *ReloadAllResponseBody) GetResult() ReloadAllResult {
	if o == nil {
		return ReloadAllResult("")
	}
	return o.Result
}

func (o *ReloadAllResponseBody) GetAction() ReloadAllAction {
	if o == nil {
		return ReloadAllAction("")
	}
	return o.Action
}

func (o *ReloadAllResponseBody) GetData() ReloadAllData {
	if o == nil {
		return ReloadAllData{}
	}
	return o.Data
}

type ReloadAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Service reload
	Object *ReloadAllResponseBody
}

func (o *ReloadAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadAllResponse) GetObject() *ReloadAllResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
