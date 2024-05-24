package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/types"
)

type GetCampaignResultsRequest struct {
	// Id of the campaign
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Max number of elements in response
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Offset of data in response (skip X elements)
	Offset *int64          `queryParam:"style=form,explode=true,name=offset"`
	Before *types.Date     `queryParam:"style=form,explode=true,name=before"`
	After  *types.Date     `queryParam:"style=form,explode=true,name=after"`
	Order  *string         `queryParam:"style=form,explode=true,name=order"`
	Asc    *components.Asc `queryParam:"style=form,explode=true,name=asc"`
}

func (g GetCampaignResultsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetCampaignResultsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetCampaignResultsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetCampaignResultsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetCampaignResultsRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetCampaignResultsRequest) GetBefore() *types.Date {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *GetCampaignResultsRequest) GetAfter() *types.Date {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *GetCampaignResultsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetCampaignResultsRequest) GetAsc() *components.Asc {
	if o == nil {
		return nil
	}
	return o.Asc
}

// GetCampaignResultsResult - Result of the request
type GetCampaignResultsResult string

const (
	GetCampaignResultsResultSuccess GetCampaignResultsResult = "success"
	GetCampaignResultsResultError   GetCampaignResultsResult = "error"
)

func (e GetCampaignResultsResult) ToPointer() *GetCampaignResultsResult {
	return &e
}
func (e *GetCampaignResultsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetCampaignResultsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCampaignResultsResult: %v", v)
	}
}

// GetCampaignResultsAction - The id of the action
type GetCampaignResultsAction string

const (
	GetCampaignResultsActionGetCampaignResults GetCampaignResultsAction = "getCampaignResults"
)

func (e GetCampaignResultsAction) ToPointer() *GetCampaignResultsAction {
	return &e
}
func (e *GetCampaignResultsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getCampaignResults":
		*e = GetCampaignResultsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCampaignResultsAction: %v", v)
	}
}

type GetCampaignResultsData struct {
	CampaignHistory []components.CampaignEventResult `json:"campaignHistory"`
}

func (o *GetCampaignResultsData) GetCampaignHistory() []components.CampaignEventResult {
	if o == nil {
		return []components.CampaignEventResult{}
	}
	return o.CampaignHistory
}

// GetCampaignResultsResponseBody - Campaign result history result
type GetCampaignResultsResponseBody struct {
	// Result of the request
	Result GetCampaignResultsResult `json:"result"`
	// The id of the action
	Action GetCampaignResultsAction `json:"action"`
	Data   GetCampaignResultsData   `json:"data"`
}

func (o *GetCampaignResultsResponseBody) GetResult() GetCampaignResultsResult {
	if o == nil {
		return GetCampaignResultsResult("")
	}
	return o.Result
}

func (o *GetCampaignResultsResponseBody) GetAction() GetCampaignResultsAction {
	if o == nil {
		return GetCampaignResultsAction("")
	}
	return o.Action
}

func (o *GetCampaignResultsResponseBody) GetData() GetCampaignResultsData {
	if o == nil {
		return GetCampaignResultsData{}
	}
	return o.Data
}

type GetCampaignResultsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Campaign result history result
	Object *GetCampaignResultsResponseBody
}

func (o *GetCampaignResultsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCampaignResultsResponse) GetObject() *GetCampaignResultsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
