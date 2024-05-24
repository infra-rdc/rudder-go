package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/types"
)

type GetEventsCampaignRequest struct {
	// Id of the campaign
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Type of the campaigns we want
	CampaignType *components.CampaignType `queryParam:"style=form,explode=true,name=campaignType"`
	// Status of the campaign events we want
	State *components.CampaignEventStatus `queryParam:"style=form,explode=true,name=state"`
	// Max number of elements in response
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Offset of data in response (skip X elements)
	Offset *int64          `queryParam:"style=form,explode=true,name=offset"`
	Before *types.Date     `queryParam:"style=form,explode=true,name=before"`
	After  *types.Date     `queryParam:"style=form,explode=true,name=after"`
	Order  *string         `queryParam:"style=form,explode=true,name=order"`
	Asc    *components.Asc `queryParam:"style=form,explode=true,name=asc"`
}

func (g GetEventsCampaignRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetEventsCampaignRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetEventsCampaignRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetEventsCampaignRequest) GetCampaignType() *components.CampaignType {
	if o == nil {
		return nil
	}
	return o.CampaignType
}

func (o *GetEventsCampaignRequest) GetState() *components.CampaignEventStatus {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *GetEventsCampaignRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetEventsCampaignRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetEventsCampaignRequest) GetBefore() *types.Date {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *GetEventsCampaignRequest) GetAfter() *types.Date {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *GetEventsCampaignRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetEventsCampaignRequest) GetAsc() *components.Asc {
	if o == nil {
		return nil
	}
	return o.Asc
}

// GetEventsCampaignResult - Result of the request
type GetEventsCampaignResult string

const (
	GetEventsCampaignResultSuccess GetEventsCampaignResult = "success"
	GetEventsCampaignResultError   GetEventsCampaignResult = "error"
)

func (e GetEventsCampaignResult) ToPointer() *GetEventsCampaignResult {
	return &e
}
func (e *GetEventsCampaignResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetEventsCampaignResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEventsCampaignResult: %v", v)
	}
}

// GetEventsCampaignAction - The id of the action
type GetEventsCampaignAction string

const (
	GetEventsCampaignActionGetEventsCampaign GetEventsCampaignAction = "getEventsCampaign"
)

func (e GetEventsCampaignAction) ToPointer() *GetEventsCampaignAction {
	return &e
}
func (e *GetEventsCampaignAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getEventsCampaign":
		*e = GetEventsCampaignAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEventsCampaignAction: %v", v)
	}
}

type GetEventsCampaignData struct {
	CampaignEvents []components.CampaignEventDetails `json:"campaignEvents"`
}

func (o *GetEventsCampaignData) GetCampaignEvents() []components.CampaignEventDetails {
	if o == nil {
		return []components.CampaignEventDetails{}
	}
	return o.CampaignEvents
}

// GetEventsCampaignResponseBody - Campaign details result
type GetEventsCampaignResponseBody struct {
	// Result of the request
	Result GetEventsCampaignResult `json:"result"`
	// The id of the action
	Action GetEventsCampaignAction `json:"action"`
	Data   GetEventsCampaignData   `json:"data"`
}

func (o *GetEventsCampaignResponseBody) GetResult() GetEventsCampaignResult {
	if o == nil {
		return GetEventsCampaignResult("")
	}
	return o.Result
}

func (o *GetEventsCampaignResponseBody) GetAction() GetEventsCampaignAction {
	if o == nil {
		return GetEventsCampaignAction("")
	}
	return o.Action
}

func (o *GetEventsCampaignResponseBody) GetData() GetEventsCampaignData {
	if o == nil {
		return GetEventsCampaignData{}
	}
	return o.Data
}

type GetEventsCampaignResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Campaign details result
	Object *GetEventsCampaignResponseBody
}

func (o *GetEventsCampaignResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetEventsCampaignResponse) GetObject() *GetEventsCampaignResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
