package components

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/types"
)

// BaseSeverity - CVSS V3 Severity
type BaseSeverity string

const (
	BaseSeverityCritical BaseSeverity = "critical"
	BaseSeverityHigh     BaseSeverity = "high"
	BaseSeverityMedium   BaseSeverity = "medium"
	BaseSeverityLow      BaseSeverity = "low"
	BaseSeverityNone     BaseSeverity = "none"
)

func (e BaseSeverity) ToPointer() *BaseSeverity {
	return &e
}
func (e *BaseSeverity) UnmarshalJSON(data []byte) error {
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
		*e = BaseSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BaseSeverity: %v", v)
	}
}

// Cvssv3 - CVSS V3 of the CVE
type Cvssv3 struct {
	// CVSS V3 base score
	BaseScore *float64 `json:"baseScore,omitempty"`
	// CVSS V3 vector
	Vector *string `json:"vector,omitempty"`
	// CVSS V3 Severity
	BaseSeverity *BaseSeverity `json:"baseSeverity,omitempty"`
}

func (o *Cvssv3) GetBaseScore() *float64 {
	if o == nil {
		return nil
	}
	return o.BaseScore
}

func (o *Cvssv3) GetVector() *string {
	if o == nil {
		return nil
	}
	return o.Vector
}

func (o *Cvssv3) GetBaseSeverity() *BaseSeverity {
	if o == nil {
		return nil
	}
	return o.BaseSeverity
}

// Cvssv2 - CVSS V2 of the CVE
type Cvssv2 struct {
	// CVSS V2 base score
	BaseScore *int64 `json:"baseScore,omitempty"`
	// CVSS V2 vector
	Vector *string `json:"vector,omitempty"`
}

func (o *Cvssv2) GetBaseScore() *int64 {
	if o == nil {
		return nil
	}
	return o.BaseScore
}

func (o *Cvssv2) GetVector() *string {
	if o == nil {
		return nil
	}
	return o.Vector
}

type CveDetails struct {
	// CVE id
	ID *string `json:"id,omitempty"`
	// Date the CVE was published
	PublishedDate *types.Date `json:"publishedDate,omitempty"`
	// last Date the CVE was modified
	LastModifiedDate *types.Date `json:"lastModifiedDate,omitempty"`
	// CVE Description
	Description *string `json:"description,omitempty"`
	// List of CWE (Common Weakness Enumeration) of the CVE
	CweIds []string `json:"cweIds,omitempty"`
	// CVSS V3 of the CVE
	Cvssv3 *Cvssv3 `json:"cvssv3,omitempty"`
	// CVSS V2 of the CVE
	Cvssv2 *Cvssv2 `json:"cvssv2,omitempty"`
}

func (c CveDetails) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CveDetails) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CveDetails) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CveDetails) GetPublishedDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.PublishedDate
}

func (o *CveDetails) GetLastModifiedDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.LastModifiedDate
}

func (o *CveDetails) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CveDetails) GetCweIds() []string {
	if o == nil {
		return nil
	}
	return o.CweIds
}

func (o *CveDetails) GetCvssv3() *Cvssv3 {
	if o == nil {
		return nil
	}
	return o.Cvssv3
}

func (o *CveDetails) GetCvssv2() *Cvssv2 {
	if o == nil {
		return nil
	}
	return o.Cvssv2
}
