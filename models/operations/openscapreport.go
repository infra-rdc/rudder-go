package operations

import (
	"github.com/infra-rdc/rudder-go/models/components"
)

type OpenscapReportRequest struct {
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
}

func (o *OpenscapReportRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

type OpenscapReportResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Promote response
	Res *string
}

func (o *OpenscapReportResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OpenscapReportResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
