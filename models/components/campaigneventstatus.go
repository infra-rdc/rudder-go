package components

import (
	"encoding/json"
	"fmt"
)

type CampaignEventStatus string

const (
	CampaignEventStatusScheduled CampaignEventStatus = "scheduled"
	CampaignEventStatusRunning   CampaignEventStatus = "running"
	CampaignEventStatusFinished  CampaignEventStatus = "finished"
	CampaignEventStatusSkipped   CampaignEventStatus = "skipped"
)

func (e CampaignEventStatus) ToPointer() *CampaignEventStatus {
	return &e
}
func (e *CampaignEventStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "scheduled":
		fallthrough
	case "running":
		fallthrough
	case "finished":
		fallthrough
	case "skipped":
		*e = CampaignEventStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CampaignEventStatus: %v", v)
	}
}
