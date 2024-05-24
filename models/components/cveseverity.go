package components

import (
	"encoding/json"
	"fmt"
)

type CveSeverity string

const (
	CveSeverityCritical CveSeverity = "critical"
	CveSeverityHigh     CveSeverity = "high"
	CveSeverityMedium   CveSeverity = "medium"
	CveSeverityLow      CveSeverity = "low"
	CveSeverityNone     CveSeverity = "none"
	CveSeverityUnknown  CveSeverity = "unknown"
)

func (e CveSeverity) ToPointer() *CveSeverity {
	return &e
}
func (e *CveSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "critical":
		fallthrough
	case "high":
		fallthrough
	case "medium":
		fallthrough
	case "low":
		fallthrough
	case "none":
		fallthrough
	case "unknown":
		*e = CveSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CveSeverity: %v", v)
	}
}
