package operations

import (
	"github.com/infra-rdc/rudder-go/models/components"
)

type ApplyPolicyRequest struct {
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
}

func (o *ApplyPolicyRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

type ApplyPolicyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Agent output
	AgentOutput *string
}

func (o *ApplyPolicyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ApplyPolicyResponse) GetAgentOutput() *string {
	if o == nil {
		return nil
	}
	return o.AgentOutput
}
