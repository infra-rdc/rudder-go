package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ScheduleCampaignRequest struct {
	// Id of the campaign
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ScheduleCampaignRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// ScheduleCampaignResult - Result of the request
type ScheduleCampaignResult string

const (
	ScheduleCampaignResultSuccess ScheduleCampaignResult = "success"
	ScheduleCampaignResultError   ScheduleCampaignResult = "error"
)

func (e ScheduleCampaignResult) ToPointer() *ScheduleCampaignResult {
	return &e
}
func (e *ScheduleCampaignResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ScheduleCampaignResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScheduleCampaignResult: %v", v)
	}
}

// ScheduleCampaignAction - The id of the action
type ScheduleCampaignAction string

const (
	ScheduleCampaignActionScheduleCampaign ScheduleCampaignAction = "scheduleCampaign"
)

func (e ScheduleCampaignAction) ToPointer() *ScheduleCampaignAction {
	return &e
}
func (e *ScheduleCampaignAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "scheduleCampaign":
		*e = ScheduleCampaignAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScheduleCampaignAction: %v", v)
	}
}

type ScheduleCampaignData struct {
	CampaignEvents []components.CampaignEventDetails `json:"campaignEvents"`
}

func (o *ScheduleCampaignData) GetCampaignEvents() []components.CampaignEventDetails {
	if o == nil {
		return []components.CampaignEventDetails{}
	}
	return o.CampaignEvents
}

// ScheduleCampaignResponseBody - Campaign events details result
type ScheduleCampaignResponseBody struct {
	// Result of the request
	Result ScheduleCampaignResult `json:"result"`
	// The id of the action
	Action ScheduleCampaignAction `json:"action"`
	Data   ScheduleCampaignData   `json:"data"`
}

func (o *ScheduleCampaignResponseBody) GetResult() ScheduleCampaignResult {
	if o == nil {
		return ScheduleCampaignResult("")
	}
	return o.Result
}

func (o *ScheduleCampaignResponseBody) GetAction() ScheduleCampaignAction {
	if o == nil {
		return ScheduleCampaignAction("")
	}
	return o.Action
}

func (o *ScheduleCampaignResponseBody) GetData() ScheduleCampaignData {
	if o == nil {
		return ScheduleCampaignData{}
	}
	return o.Data
}

type ScheduleCampaignResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Campaign events details result
	Object *ScheduleCampaignResponseBody
}

func (o *ScheduleCampaignResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ScheduleCampaignResponse) GetObject() *ScheduleCampaignResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
