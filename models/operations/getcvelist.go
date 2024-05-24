package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/types"
)

// GetCVEListRequestBody - List of CVE ids you want
type GetCVEListRequestBody struct {
	CveIds []string `json:"cveIds,omitempty"`
	// Only send score of the cve, and not the whole detailed list
	OnlyScore *bool `default:"false" json:"onlyScore"`
	// Only send CVE with a score higher than the value
	MinScore *string `json:"minScore,omitempty"`
	// Only send CVE with a score lower than the value
	MaxScore *string `json:"maxScore,omitempty"`
	// Only send CVE with a publication date more recent than the value
	PublishedDate *types.Date `json:"publishedDate,omitempty"`
}

func (g GetCVEListRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetCVEListRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetCVEListRequestBody) GetCveIds() []string {
	if o == nil {
		return nil
	}
	return o.CveIds
}

func (o *GetCVEListRequestBody) GetOnlyScore() *bool {
	if o == nil {
		return nil
	}
	return o.OnlyScore
}

func (o *GetCVEListRequestBody) GetMinScore() *string {
	if o == nil {
		return nil
	}
	return o.MinScore
}

func (o *GetCVEListRequestBody) GetMaxScore() *string {
	if o == nil {
		return nil
	}
	return o.MaxScore
}

func (o *GetCVEListRequestBody) GetPublishedDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.PublishedDate
}

// GetCVEListResult - Result of the request
type GetCVEListResult string

const (
	GetCVEListResultSuccess GetCVEListResult = "success"
	GetCVEListResultError   GetCVEListResult = "error"
)

func (e GetCVEListResult) ToPointer() *GetCVEListResult {
	return &e
}
func (e *GetCVEListResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetCVEListResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCVEListResult: %v", v)
	}
}

// GetCVEListAction - The id of the action
type GetCVEListAction string

const (
	GetCVEListActionGetCveList GetCVEListAction = "getCVEList"
)

func (e GetCVEListAction) ToPointer() *GetCVEListAction {
	return &e
}
func (e *GetCVEListAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getCVEList":
		*e = GetCVEListAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCVEListAction: %v", v)
	}
}

type GetCVEListData struct {
	CVEs []components.CveDetails `json:"CVEs"`
}

func (o *GetCVEListData) GetCVEs() []components.CveDetails {
	if o == nil {
		return []components.CveDetails{}
	}
	return o.CVEs
}

// GetCVEListResponseBody - CVE list
type GetCVEListResponseBody struct {
	// Result of the request
	Result GetCVEListResult `json:"result"`
	// The id of the action
	Action GetCVEListAction `json:"action"`
	Data   GetCVEListData   `json:"data"`
}

func (o *GetCVEListResponseBody) GetResult() GetCVEListResult {
	if o == nil {
		return GetCVEListResult("")
	}
	return o.Result
}

func (o *GetCVEListResponseBody) GetAction() GetCVEListAction {
	if o == nil {
		return GetCVEListAction("")
	}
	return o.Action
}

func (o *GetCVEListResponseBody) GetData() GetCVEListData {
	if o == nil {
		return GetCVEListData{}
	}
	return o.Data
}

type GetCVEListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// CVE list
	Object *GetCVEListResponseBody
}

func (o *GetCVEListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCVEListResponse) GetObject() *GetCVEListResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
