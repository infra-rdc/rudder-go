package components

import (
	"encoding/json"
	"fmt"
)

type CampaignType string

const (
	CampaignTypeSystemUpdate   CampaignType = "system-update"
	CampaignTypeSoftwareUpdate CampaignType = "software-update"
)

func (e CampaignType) ToPointer() *CampaignType {
	return &e
}
func (e *CampaignType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system-update":
		fallthrough
	case "software-update":
		*e = CampaignType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CampaignType: %v", v)
	}
}
