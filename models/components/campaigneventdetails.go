package components

import (
	"encoding/json"
	"fmt"
)

type Value string

const (
	ValueScheduled Value = "scheduled"
	ValueRunning   Value = "running"
	ValueFinished  Value = "finished"
	ValueSkipped   Value = "skipped"
)

func (e Value) ToPointer() *Value {
	return &e
}
func (e *Value) UnmarshalJSON(data []byte) error {
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
		*e = Value(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Value: %v", v)
	}
}

type CampaignEventDetailsStatus struct {
	Value *Value `json:"value,omitempty"`
	// used only for skipped status
	Reason *string `json:"reason,omitempty"`
}

func (o *CampaignEventDetailsStatus) GetValue() *Value {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *CampaignEventDetailsStatus) GetReason() *string {
	if o == nil {
		return nil
	}
	return o.Reason
}

// CampaignEventDetailsCampaignType - Type of campaign
type CampaignEventDetailsCampaignType string

const (
	CampaignEventDetailsCampaignTypeSystemUpdate   CampaignEventDetailsCampaignType = "system-update"
	CampaignEventDetailsCampaignTypeSoftwareUpdate CampaignEventDetailsCampaignType = "software-update"
)

func (e CampaignEventDetailsCampaignType) ToPointer() *CampaignEventDetailsCampaignType {
	return &e
}
func (e *CampaignEventDetailsCampaignType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system-update":
		fallthrough
	case "software-update":
		*e = CampaignEventDetailsCampaignType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CampaignEventDetailsCampaignType: %v", v)
	}
}

type CampaignEventDetails struct {
	// Campaign event id
	ID *string `json:"id,omitempty"`
	// Campaign event name
	Name *string `json:"name,omitempty"`
	// Campaign event start date
	Start *string `json:"start,omitempty"`
	// Campaign event end date
	End    *string                     `json:"end,omitempty"`
	Status *CampaignEventDetailsStatus `json:"status,omitempty"`
	// Type of campaign
	CampaignType *CampaignEventDetailsCampaignType `json:"campaignType,omitempty"`
	// Id of the campaign for this event
	CampaignID *string `json:"campaignId,omitempty"`
}

func (o *CampaignEventDetails) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CampaignEventDetails) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CampaignEventDetails) GetStart() *string {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *CampaignEventDetails) GetEnd() *string {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *CampaignEventDetails) GetStatus() *CampaignEventDetailsStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *CampaignEventDetails) GetCampaignType() *CampaignEventDetailsCampaignType {
	if o == nil {
		return nil
	}
	return o.CampaignType
}

func (o *CampaignEventDetails) GetCampaignID() *string {
	if o == nil {
		return nil
	}
	return o.CampaignID
}
