package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteCampaignEventRequest struct {
	// Id of the campaign event
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteCampaignEventRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteCampaignEventResult - Result of the request
type DeleteCampaignEventResult string

const (
	DeleteCampaignEventResultSuccess DeleteCampaignEventResult = "success"
	DeleteCampaignEventResultError   DeleteCampaignEventResult = "error"
)

func (e DeleteCampaignEventResult) ToPointer() *DeleteCampaignEventResult {
	return &e
}
func (e *DeleteCampaignEventResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteCampaignEventResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteCampaignEventResult: %v", v)
	}
}

// DeleteCampaignEventAction - The id of the action
type DeleteCampaignEventAction string

const (
	DeleteCampaignEventActionDeleteCampaignEvent DeleteCampaignEventAction = "deleteCampaignEvent"
)

func (e DeleteCampaignEventAction) ToPointer() *DeleteCampaignEventAction {
	return &e
}
func (e *DeleteCampaignEventAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleteCampaignEvent":
		*e = DeleteCampaignEventAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteCampaignEventAction: %v", v)
	}
}

type DeleteCampaignEventData struct {
	CampaignEvents []components.CampaignEventDetails `json:"campaignEvents"`
}

func (o *DeleteCampaignEventData) GetCampaignEvents() []components.CampaignEventDetails {
	if o == nil {
		return []components.CampaignEventDetails{}
	}
	return o.CampaignEvents
}

// DeleteCampaignEventResponseBody - Campaign event details result
type DeleteCampaignEventResponseBody struct {
	// Result of the request
	Result DeleteCampaignEventResult `json:"result"`
	// The id of the action
	Action DeleteCampaignEventAction `json:"action"`
	Data   DeleteCampaignEventData   `json:"data"`
}

func (o *DeleteCampaignEventResponseBody) GetResult() DeleteCampaignEventResult {
	if o == nil {
		return DeleteCampaignEventResult("")
	}
	return o.Result
}

func (o *DeleteCampaignEventResponseBody) GetAction() DeleteCampaignEventAction {
	if o == nil {
		return DeleteCampaignEventAction("")
	}
	return o.Action
}

func (o *DeleteCampaignEventResponseBody) GetData() DeleteCampaignEventData {
	if o == nil {
		return DeleteCampaignEventData{}
	}
	return o.Data
}

type DeleteCampaignEventResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Campaign event details result
	Object *DeleteCampaignEventResponseBody
}

func (o *DeleteCampaignEventResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteCampaignEventResponse) GetObject() *DeleteCampaignEventResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
