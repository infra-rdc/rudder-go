package components

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/types"
)

// CampaignEventNodeResultStatus - Campaign result
type CampaignEventNodeResultStatus string

const (
	CampaignEventNodeResultStatusSuccess   CampaignEventNodeResultStatus = "success"
	CampaignEventNodeResultStatusError     CampaignEventNodeResultStatus = "error"
	CampaignEventNodeResultStatusScheduled CampaignEventNodeResultStatus = "scheduled"
	CampaignEventNodeResultStatusMissing   CampaignEventNodeResultStatus = "missing"
)

func (e CampaignEventNodeResultStatus) ToPointer() *CampaignEventNodeResultStatus {
	return &e
}
func (e *CampaignEventNodeResultStatus) UnmarshalJSON(data []byte) error {
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
		*e = CampaignEventNodeResultStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CampaignEventNodeResultStatus: %v", v)
	}
}

type SoftwareUpdated struct {
	// Name of the software updated
	Name *string `json:"name,omitempty"`
	// Version of the software before update
	Before *string `json:"before,omitempty"`
	// Version of the software after update
	After *string `json:"after,omitempty"`
}

func (o *SoftwareUpdated) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SoftwareUpdated) GetBefore() *string {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *SoftwareUpdated) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

type CampaignEventNodeResultResult struct {
	// Campaign result
	Status *CampaignEventNodeResultStatus `json:"status,omitempty"`
	// List of updated software
	SoftwareUpdated []SoftwareUpdated `json:"software-updated,omitempty"`
	// campaign standard output
	Output *string `json:"output,omitempty"`
	// campaign standard errors
	Errors *string `json:"errors,omitempty"`
}

func (o *CampaignEventNodeResultResult) GetStatus() *CampaignEventNodeResultStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *CampaignEventNodeResultResult) GetSoftwareUpdated() []SoftwareUpdated {
	if o == nil {
		return nil
	}
	return o.SoftwareUpdated
}

func (o *CampaignEventNodeResultResult) GetOutput() *string {
	if o == nil {
		return nil
	}
	return o.Output
}

func (o *CampaignEventNodeResultResult) GetErrors() *string {
	if o == nil {
		return nil
	}
	return o.Errors
}

type CampaignEventNodeResultNodes struct {
	// Node id
	NodeID *string                        `json:"nodeId,omitempty"`
	Result *CampaignEventNodeResultResult `json:"result,omitempty"`
	Date   *types.Date                    `json:"date,omitempty"`
}

func (c CampaignEventNodeResultNodes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CampaignEventNodeResultNodes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CampaignEventNodeResultNodes) GetNodeID() *string {
	if o == nil {
		return nil
	}
	return o.NodeID
}

func (o *CampaignEventNodeResultNodes) GetResult() *CampaignEventNodeResultResult {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *CampaignEventNodeResultNodes) GetDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.Date
}

type CampaignEventNodeResult struct {
	// Campaign event id
	ID    *string                        `json:"id,omitempty"`
	Nodes []CampaignEventNodeResultNodes `json:"nodes,omitempty"`
}

func (o *CampaignEventNodeResult) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CampaignEventNodeResult) GetNodes() []CampaignEventNodeResultNodes {
	if o == nil {
		return nil
	}
	return o.Nodes
}
