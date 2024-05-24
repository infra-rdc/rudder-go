package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type GetGlobalComplianceRequest struct {
	// Number of digits after comma in compliance percent figures
	Precision *int64 `default:"2" queryParam:"style=form,explode=true,name=precision"`
}

func (g GetGlobalComplianceRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetGlobalComplianceRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetGlobalComplianceRequest) GetPrecision() *int64 {
	if o == nil {
		return nil
	}
	return o.Precision
}

// GetGlobalComplianceResult - Result of the request
type GetGlobalComplianceResult string

const (
	GetGlobalComplianceResultSuccess GetGlobalComplianceResult = "success"
	GetGlobalComplianceResultError   GetGlobalComplianceResult = "error"
)

func (e GetGlobalComplianceResult) ToPointer() *GetGlobalComplianceResult {
	return &e
}
func (e *GetGlobalComplianceResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetGlobalComplianceResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetGlobalComplianceResult: %v", v)
	}
}

// GetGlobalComplianceAction - The id of the action
type GetGlobalComplianceAction string

const (
	GetGlobalComplianceActionGetGlobalCompliance GetGlobalComplianceAction = "getGlobalCompliance"
)

func (e GetGlobalComplianceAction) ToPointer() *GetGlobalComplianceAction {
	return &e
}
func (e *GetGlobalComplianceAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getGlobalCompliance":
		*e = GetGlobalComplianceAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetGlobalComplianceAction: %v", v)
	}
}

type ComplianceDetails struct {
	SuccessAlreadyOK           *float32 `json:"successAlreadyOK,omitempty"`
	NoReport                   *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable       *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error                      *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired            *float32 `json:"successRepaired,omitempty"`
}

func (o *ComplianceDetails) GetSuccessAlreadyOK() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessAlreadyOK
}

func (o *ComplianceDetails) GetNoReport() *float32 {
	if o == nil {
		return nil
	}
	return o.NoReport
}

func (o *ComplianceDetails) GetSuccessNotApplicable() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessNotApplicable
}

func (o *ComplianceDetails) GetUnexpectedMissingComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedMissingComponent
}

func (o *ComplianceDetails) GetError() *float32 {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ComplianceDetails) GetUnexpectedUnknownComponent() *float32 {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnknownComponent
}

func (o *ComplianceDetails) GetSuccessRepaired() *float32 {
	if o == nil {
		return nil
	}
	return o.SuccessRepaired
}

type GlobalCompliance struct {
	// Global compliance level (`-1` when no policies are defined)
	Compliance        float64            `json:"compliance"`
	ComplianceDetails *ComplianceDetails `json:"complianceDetails,omitempty"`
}

func (o *GlobalCompliance) GetCompliance() float64 {
	if o == nil {
		return 0.0
	}
	return o.Compliance
}

func (o *GlobalCompliance) GetComplianceDetails() *ComplianceDetails {
	if o == nil {
		return nil
	}
	return o.ComplianceDetails
}

type GetGlobalComplianceData struct {
	GlobalCompliance GlobalCompliance `json:"globalCompliance"`
}

func (o *GetGlobalComplianceData) GetGlobalCompliance() GlobalCompliance {
	if o == nil {
		return GlobalCompliance{}
	}
	return o.GlobalCompliance
}

// GetGlobalComplianceResponseBody - Success
type GetGlobalComplianceResponseBody struct {
	// Result of the request
	Result GetGlobalComplianceResult `json:"result"`
	// The id of the action
	Action GetGlobalComplianceAction `json:"action"`
	Data   GetGlobalComplianceData   `json:"data"`
}

func (o *GetGlobalComplianceResponseBody) GetResult() GetGlobalComplianceResult {
	if o == nil {
		return GetGlobalComplianceResult("")
	}
	return o.Result
}

func (o *GetGlobalComplianceResponseBody) GetAction() GetGlobalComplianceAction {
	if o == nil {
		return GetGlobalComplianceAction("")
	}
	return o.Action
}

func (o *GetGlobalComplianceResponseBody) GetData() GetGlobalComplianceData {
	if o == nil {
		return GetGlobalComplianceData{}
	}
	return o.Data
}

type GetGlobalComplianceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetGlobalComplianceResponseBody
}

func (o *GetGlobalComplianceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetGlobalComplianceResponse) GetObject() *GetGlobalComplianceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
