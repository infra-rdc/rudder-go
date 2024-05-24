package components

import (
	"encoding/json"
	"fmt"
)

// CheckStatus - The severity status of the check
type CheckStatus string

const (
	CheckStatusCritical CheckStatus = "Critical"
	CheckStatusWarning  CheckStatus = "Warning"
	CheckStatusOk       CheckStatus = "Ok"
)

func (e CheckStatus) ToPointer() *CheckStatus {
	return &e
}
func (e *CheckStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Critical":
		fallthrough
	case "Warning":
		fallthrough
	case "Ok":
		*e = CheckStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CheckStatus: %v", v)
	}
}

type Check struct {
	// Name of the check
	Name string `json:"name"`
	// Message about the check
	Msg string `json:"msg"`
	// The severity status of the check
	Status CheckStatus `json:"status"`
}

func (o *Check) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Check) GetMsg() string {
	if o == nil {
		return ""
	}
	return o.Msg
}

func (o *Check) GetStatus() CheckStatus {
	if o == nil {
		return CheckStatus("")
	}
	return o.Status
}
