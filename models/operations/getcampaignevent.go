package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetCampaignEventRequest struct {
	// Id of the campaign event
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCampaignEventRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetCampaignEventResult - Result of the request
type GetCampaignEventResult string

const (
	GetCampaignEventResultSuccess GetCampaignEventResult = "success"
	GetCampaignEventResultError   GetCampaignEventResult = "error"
)

func (e GetCampaignEventResult) ToPointer() *GetCampaignEventResult {
	return &e
}
func (e *GetCampaignEventResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetCampaignEventResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCampaignEventResult: %v", v)
	}
}

// GetCampaignEventAction - The id of the action
type GetCampaignEventAction string

const (
	GetCampaignEventActionGetCampaignEvent GetCampaignEventAction = "getCampaignEvent"
)

func (e GetCampaignEventAction) ToPointer() *GetCampaignEventAction {
	return &e
}
func (e *GetCampaignEventAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getCampaignEvent":
		*e = GetCampaignEventAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCampaignEventAction: %v", v)
	}
}

type GetCampaignEventData struct {
	CampaignEvents []components.CampaignEventDetails `json:"campaignEvents"`
}

func (o *GetCampaignEventData) GetCampaignEvents() []components.CampaignEventDetails {
	if o == nil {
		return []components.CampaignEventDetails{}
	}
	return o.CampaignEvents
}

// GetCampaignEventResponseBody - Campaign details result
type GetCampaignEventResponseBody struct {
	// Result of the request
	Result GetCampaignEventResult `json:"result"`
	// The id of the action
	Action GetCampaignEventAction `json:"action"`
	Data   GetCampaignEventData   `json:"data"`
}

func (o *GetCampaignEventResponseBody) GetResult() GetCampaignEventResult {
	if o == nil {
		return GetCampaignEventResult("")
	}
	return o.Result
}

func (o *GetCampaignEventResponseBody) GetAction() GetCampaignEventAction {
	if o == nil {
		return GetCampaignEventAction("")
	}
	return o.Action
}

func (o *GetCampaignEventResponseBody) GetData() GetCampaignEventData {
	if o == nil {
		return GetCampaignEventData{}
	}
	return o.Data
}

type GetCampaignEventResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Campaign details result
	Object *GetCampaignEventResponseBody
}

func (o *GetCampaignEventResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCampaignEventResponse) GetObject() *GetCampaignEventResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
