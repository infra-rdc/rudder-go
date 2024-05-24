package components

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/types"
)

// CampaignEventResultStatus - Campaign status
type CampaignEventResultStatus string

const (
	CampaignEventResultStatusSuccess   CampaignEventResultStatus = "success"
	CampaignEventResultStatusError     CampaignEventResultStatus = "error"
	CampaignEventResultStatusScheduled CampaignEventResultStatus = "scheduled"
	CampaignEventResultStatusMissing   CampaignEventResultStatus = "missing"
)

func (e CampaignEventResultStatus) ToPointer() *CampaignEventResultStatus {
	return &e
}
func (e *CampaignEventResultStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		fallthrough
	case "scheduled":
		fallthrough
	case "missing":
		*e = CampaignEventResultStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CampaignEventResultStatus: %v", v)
	}
}

// Nodes - Campaign result for a Node
type Nodes struct {
	// Node id
	NodeID *string `json:"nodeId,omitempty"`
	// Campaign status
	Status *CampaignEventResultStatus `json:"status,omitempty"`
	// Number of software updated
	NbPackages *int64      `json:"nbPackages,omitempty"`
	Date       *types.Date `json:"date,omitempty"`
}

func (n Nodes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *Nodes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Nodes) GetNodeID() *string {
	if o == nil {
		return nil
	}
	return o.NodeID
}

func (o *Nodes) GetStatus() *CampaignEventResultStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Nodes) GetNbPackages() *int64 {
	if o == nil {
		return nil
	}
	return o.NbPackages
}

func (o *Nodes) GetDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.Date
}

type CampaignEventResult struct {
	// Campaign event id
	ID *string `json:"id,omitempty"`
	// Campaign result for all Nodes
	Nodes []Nodes `json:"nodes,omitempty"`
}

func (o *CampaignEventResult) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CampaignEventResult) GetNodes() []Nodes {
	if o == nil {
		return nil
	}
	return o.Nodes
}
