package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type SaveCampaignEventRequest struct {
	// Id of the campaign event
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *SaveCampaignEventRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// SaveCampaignEventResult - Result of the request
type SaveCampaignEventResult string

const (
	SaveCampaignEventResultSuccess SaveCampaignEventResult = "success"
	SaveCampaignEventResultError   SaveCampaignEventResult = "error"
)

func (e SaveCampaignEventResult) ToPointer() *SaveCampaignEventResult {
	return &e
}
func (e *SaveCampaignEventResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = SaveCampaignEventResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SaveCampaignEventResult: %v", v)
	}
}

// SaveCampaignEventAction - The id of the action
type SaveCampaignEventAction string

const (
	SaveCampaignEventActionSaveCampaignEvent SaveCampaignEventAction = "saveCampaignEvent"
)

func (e SaveCampaignEventAction) ToPointer() *SaveCampaignEventAction {
	return &e
}
func (e *SaveCampaignEventAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "saveCampaignEvent":
		*e = SaveCampaignEventAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SaveCampaignEventAction: %v", v)
	}
}

type SaveCampaignEventData struct {
	CampaignEvents []components.CampaignEventDetails `json:"campaignEvents"`
}

func (o *SaveCampaignEventData) GetCampaignEvents() []components.CampaignEventDetails {
	if o == nil {
		return []components.CampaignEventDetails{}
	}
	return o.CampaignEvents
}

// SaveCampaignEventResponseBody - Campaign event details result
type SaveCampaignEventResponseBody struct {
	// Result of the request
	Result SaveCampaignEventResult `json:"result"`
	// The id of the action
	Action SaveCampaignEventAction `json:"action"`
	Data   SaveCampaignEventData   `json:"data"`
}

func (o *SaveCampaignEventResponseBody) GetResult() SaveCampaignEventResult {
	if o == nil {
		return SaveCampaignEventResult("")
	}
	return o.Result
}

func (o *SaveCampaignEventResponseBody) GetAction() SaveCampaignEventAction {
	if o == nil {
		return SaveCampaignEventAction("")
	}
	return o.Action
}

func (o *SaveCampaignEventResponseBody) GetData() SaveCampaignEventData {
	if o == nil {
		return SaveCampaignEventData{}
	}
	return o.Data
}

type SaveCampaignEventResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Campaign event details result
	Object *SaveCampaignEventResponseBody
}

func (o *SaveCampaignEventResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SaveCampaignEventResponse) GetObject() *SaveCampaignEventResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
