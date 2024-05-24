package components

import (
	"encoding/json"
	"fmt"
)

type CampaignID string

const (
	CampaignIDSystemUpdate   CampaignID = "system-update"
	CampaignIDSoftwareUpdate CampaignID = "software-update"
)

func (e CampaignID) ToPointer() *CampaignID {
	return &e
}
func (e *CampaignID) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system-update":
		fallthrough
	case "software-update":
		*e = CampaignID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CampaignID: %v", v)
	}
}
