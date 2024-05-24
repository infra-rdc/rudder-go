package components

import (
	"encoding/json"
	"fmt"
)

type Status string

const (
	StatusDeployed          Status = "Deployed"
	StatusPendingDeployment Status = "Pending deployment"
	StatusCancelled         Status = "Cancelled"
	StatusPendingValidation Status = "Pending validation"
	StatusOpen              Status = "Open"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Deployed":
		fallthrough
	case "Pending deployment":
		fallthrough
	case "Cancelled":
		fallthrough
	case "Pending validation":
		fallthrough
	case "Open":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type Rules struct {
	Action *string `json:"action,omitempty"`
}

func (o *Rules) GetAction() *string {
	if o == nil {
		return nil
	}
	return o.Action
}

type Changes struct {
	Rules []Rules `json:"rules,omitempty"`
}

func (o *Changes) GetRules() []Rules {
	if o == nil {
		return nil
	}
	return o.Rules
}

// ChangeRequest - Content of the change request
type ChangeRequest struct {
	ID          *int64   `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Status      *Status  `json:"status,omitempty"`
	Acceptable  *bool    `json:"acceptable,omitempty"`
	CreatedBy   *string  `json:"created by,omitempty"`
	Changes     *Changes `json:"changes,omitempty"`
}

func (o *ChangeRequest) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ChangeRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ChangeRequest) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ChangeRequest) GetAcceptable() *bool {
	if o == nil {
		return nil
	}
	return o.Acceptable
}

func (o *ChangeRequest) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *ChangeRequest) GetChanges() *Changes {
	if o == nil {
		return nil
	}
	return o.Changes
}
