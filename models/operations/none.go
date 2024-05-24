package operations

import (
	"github.com/infra-rdc/rudder-go/models/components"
)

type NoneResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK status message
	Res *string
}

func (o *NoneResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *NoneResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
