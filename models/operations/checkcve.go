package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// CheckCVEResult - Result of the request
type CheckCVEResult string

const (
	CheckCVEResultSuccess CheckCVEResult = "success"
	CheckCVEResultError   CheckCVEResult = "error"
)

func (e CheckCVEResult) ToPointer() *CheckCVEResult {
	return &e
}
func (e *CheckCVEResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = CheckCVEResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CheckCVEResult: %v", v)
	}
}

// CheckCVEAction - The id of the action
type CheckCVEAction string

const (
	CheckCVEActionCheckCve CheckCVEAction = "checkCVE"
)

func (e CheckCVEAction) ToPointer() *CheckCVEAction {
	return &e
}
func (e *CheckCVEAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "checkCVE":
		*e = CheckCVEAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CheckCVEAction: %v", v)
	}
}

type CheckCVEData struct {
	CveChecks []components.CveCheck `json:"cveChecks"`
}

func (o *CheckCVEData) GetCveChecks() []components.CveCheck {
	if o == nil {
		return []components.CveCheck{}
	}
	return o.CveChecks
}

// CheckCVEResponseBody - CVE check result
type CheckCVEResponseBody struct {
	// Result of the request
	Result CheckCVEResult `json:"result"`
	// The id of the action
	Action CheckCVEAction `json:"action"`
	Data   CheckCVEData   `json:"data"`
}

func (o *CheckCVEResponseBody) GetResult() CheckCVEResult {
	if o == nil {
		return CheckCVEResult("")
	}
	return o.Result
}

func (o *CheckCVEResponseBody) GetAction() CheckCVEAction {
	if o == nil {
		return CheckCVEAction("")
	}
	return o.Action
}

func (o *CheckCVEResponseBody) GetData() CheckCVEData {
	if o == nil {
		return CheckCVEData{}
	}
	return o.Data
}

type CheckCVEResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// CVE check result
	Object *CheckCVEResponseBody
}

func (o *CheckCVEResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CheckCVEResponse) GetObject() *CheckCVEResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
