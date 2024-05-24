package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/types"
)

type GetAllCampaignEventsRequest struct {
	// Type of the campaigns we want
	CampaignType *components.CampaignType `queryParam:"style=form,explode=true,name=campaignType"`
	// Status of the campaign events we want
	State *components.CampaignEventStatus `queryParam:"style=form,explode=true,name=state"`
	// id of the campaigns we want
	CampaignID *components.CampaignID `queryParam:"style=form,explode=true,name=campaignId"`
	// Max number of elements in response
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Offset of data in response (skip X elements)
	Offset *int64          `queryParam:"style=form,explode=true,name=offset"`
	Before *types.Date     `queryParam:"style=form,explode=true,name=before"`
	After  *types.Date     `queryParam:"style=form,explode=true,name=after"`
	Order  *string         `queryParam:"style=form,explode=true,name=order"`
	Asc    *components.Asc `queryParam:"style=form,explode=true,name=asc"`
}

func (g GetAllCampaignEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAllCampaignEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAllCampaignEventsRequest) GetCampaignType() *components.CampaignType {
	if o == nil {
		return nil
	}
	return o.CampaignType
}

func (o *GetAllCampaignEventsRequest) GetState() *components.CampaignEventStatus {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *GetAllCampaignEventsRequest) GetCampaignID() *components.CampaignID {
	if o == nil {
		return nil
	}
	return o.CampaignID
}

func (o *GetAllCampaignEventsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetAllCampaignEventsRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetAllCampaignEventsRequest) GetBefore() *types.Date {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *GetAllCampaignEventsRequest) GetAfter() *types.Date {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *GetAllCampaignEventsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetAllCampaignEventsRequest) GetAsc() *components.Asc {
	if o == nil {
		return nil
	}
	return o.Asc
}

// GetAllCampaignEventsResult - Result of the request
type GetAllCampaignEventsResult string

const (
	GetAllCampaignEventsResultSuccess GetAllCampaignEventsResult = "success"
	GetAllCampaignEventsResultError   GetAllCampaignEventsResult = "error"
)

func (e GetAllCampaignEventsResult) ToPointer() *GetAllCampaignEventsResult {
	return &e
}
func (e *GetAllCampaignEventsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetAllCampaignEventsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllCampaignEventsResult: %v", v)
	}
}

// GetAllCampaignEventsAction - The id of the action
type GetAllCampaignEventsAction string

const (
	GetAllCampaignEventsActionGetCampaignEvent GetAllCampaignEventsAction = "getCampaignEvent"
)

func (e GetAllCampaignEventsAction) ToPointer() *GetAllCampaignEventsAction {
	return &e
}
func (e *GetAllCampaignEventsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getCampaignEvent":
		*e = GetAllCampaignEventsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllCampaignEventsAction: %v", v)
	}
}

type GetAllCampaignEventsData struct {
	CampaignEvents []components.CampaignEventDetails `json:"campaignEvents"`
}

func (o *GetAllCampaignEventsData) GetCampaignEvents() []components.CampaignEventDetails {
	if o == nil {
		return []components.CampaignEventDetails{}
	}
	return o.CampaignEvents
}

// GetAllCampaignEventsResponseBody - Campaign event details result
type GetAllCampaignEventsResponseBody struct {
	// Result of the request
	Result GetAllCampaignEventsResult `json:"result"`
	// The id of the action
	Action GetAllCampaignEventsAction `json:"action"`
	Data   GetAllCampaignEventsData   `json:"data"`
}

func (o *GetAllCampaignEventsResponseBody) GetResult() GetAllCampaignEventsResult {
	if o == nil {
		return GetAllCampaignEventsResult("")
	}
	return o.Result
}

func (o *GetAllCampaignEventsResponseBody) GetAction() GetAllCampaignEventsAction {
	if o == nil {
		return GetAllCampaignEventsAction("")
	}
	return o.Action
}

func (o *GetAllCampaignEventsResponseBody) GetData() GetAllCampaignEventsData {
	if o == nil {
		return GetAllCampaignEventsData{}
	}
	return o.Data
}

type GetAllCampaignEventsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Campaign event details result
	Object *GetAllCampaignEventsResponseBody
}

func (o *GetAllCampaignEventsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAllCampaignEventsResponse) GetObject() *GetAllCampaignEventsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
