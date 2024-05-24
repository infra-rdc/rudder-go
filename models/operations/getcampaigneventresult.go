package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetCampaignEventResultRequest struct {
	// Id of the campaign event
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCampaignEventResultRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetCampaignEventResultResult - Result of the request
type GetCampaignEventResultResult string

const (
	GetCampaignEventResultResultSuccess GetCampaignEventResultResult = "success"
	GetCampaignEventResultResultError   GetCampaignEventResultResult = "error"
)

func (e GetCampaignEventResultResult) ToPointer() *GetCampaignEventResultResult {
	return &e
}
func (e *GetCampaignEventResultResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetCampaignEventResultResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCampaignEventResultResult: %v", v)
	}
}

// GetCampaignEventResultAction - The id of the action
type GetCampaignEventResultAction string

const (
	GetCampaignEventResultActionGetCampaignEventResult GetCampaignEventResultAction = "getCampaignEventResult"
)

func (e GetCampaignEventResultAction) ToPointer() *GetCampaignEventResultAction {
	return &e
}
func (e *GetCampaignEventResultAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getCampaignEventResult":
		*e = GetCampaignEventResultAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCampaignEventResultAction: %v", v)
	}
}

type GetCampaignEventResultData struct {
	EventResult []components.CampaignEventResult `json:"eventResult"`
}

func (o *GetCampaignEventResultData) GetEventResult() []components.CampaignEventResult {
	if o == nil {
		return []components.CampaignEventResult{}
	}
	return o.EventResult
}

// GetCampaignEventResultResponseBody - Campaign result history result
type GetCampaignEventResultResponseBody struct {
	// Result of the request
	Result GetCampaignEventResultResult `json:"result"`
	// The id of the action
	Action GetCampaignEventResultAction `json:"action"`
	// Campaign event id
	ID   *string                    `json:"id,omitempty"`
	Data GetCampaignEventResultData `json:"data"`
}

func (o *GetCampaignEventResultResponseBody) GetResult() GetCampaignEventResultResult {
	if o == nil {
		return GetCampaignEventResultResult("")
	}
	return o.Result
}

func (o *GetCampaignEventResultResponseBody) GetAction() GetCampaignEventResultAction {
	if o == nil {
		return GetCampaignEventResultAction("")
	}
	return o.Action
}

func (o *GetCampaignEventResultResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetCampaignEventResultResponseBody) GetData() GetCampaignEventResultData {
	if o == nil {
		return GetCampaignEventResultData{}
	}
	return o.Data
}

type GetCampaignEventResultResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Campaign result history result
	Object *GetCampaignEventResultResponseBody
}

func (o *GetCampaignEventResultResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCampaignEventResultResponse) GetObject() *GetCampaignEventResultResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
