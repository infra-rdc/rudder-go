package operations

import (
	"errors"
	"fmt"
	"io"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type GetDirectiveComplianceIDRequest struct {
	// format of export
	Format *string `queryParam:"style=form,explode=false,name=format"`
	// Id of the directive
	DirectiveID string `pathParam:"style=simple,explode=false,name=directiveId"`
}

func (o *GetDirectiveComplianceIDRequest) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *GetDirectiveComplianceIDRequest) GetDirectiveID() string {
	if o == nil {
		return ""
	}
	return o.DirectiveID
}

type GetDirectiveComplianceIDResponseBodyType string

const (
	GetDirectiveComplianceIDResponseBodyTypeComplianceDirectiveID GetDirectiveComplianceIDResponseBodyType = "compliance-directive-id"
	GetDirectiveComplianceIDResponseBodyTypeStream                GetDirectiveComplianceIDResponseBodyType = "stream"
)

// GetDirectiveComplianceIDResponseBody - Success
type GetDirectiveComplianceIDResponseBody struct {
	ComplianceDirectiveID *components.ComplianceDirectiveID
	Stream                io.ReadCloser

	Type GetDirectiveComplianceIDResponseBodyType
}

func CreateGetDirectiveComplianceIDResponseBodyComplianceDirectiveID(complianceDirectiveID components.ComplianceDirectiveID) GetDirectiveComplianceIDResponseBody {
	typ := GetDirectiveComplianceIDResponseBodyTypeComplianceDirectiveID

	return GetDirectiveComplianceIDResponseBody{
		ComplianceDirectiveID: &complianceDirectiveID,
		Type:                  typ,
	}
}

func CreateGetDirectiveComplianceIDResponseBodyStream(stream io.ReadCloser) GetDirectiveComplianceIDResponseBody {
	typ := GetDirectiveComplianceIDResponseBodyTypeStream

	return GetDirectiveComplianceIDResponseBody{
		Stream: &stream,
		Type:   typ,
	}
}

func (u *GetDirectiveComplianceIDResponseBody) UnmarshalJSON(data []byte) error {

	var complianceDirectiveID components.ComplianceDirectiveID = components.ComplianceDirectiveID{}
	if err := utils.UnmarshalJSON(data, &complianceDirectiveID, "", true, true); err == nil {
		u.ComplianceDirectiveID = &complianceDirectiveID
		u.Type = GetDirectiveComplianceIDResponseBodyTypeComplianceDirectiveID
		return nil
	}

	var stream io.ReadCloser = io.ReadCloser{}
	if err := utils.UnmarshalJSON(data, &stream, "", true, true); err == nil {
		u.Stream = &stream
		u.Type = GetDirectiveComplianceIDResponseBodyTypeStream
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetDirectiveComplianceIDResponseBody", string(data))
}

func (u GetDirectiveComplianceIDResponseBody) MarshalJSON() ([]byte, error) {
	if u.ComplianceDirectiveID != nil {
		return utils.MarshalJSON(u.ComplianceDirectiveID, "", true)
	}

	if u.Stream != nil {
		return utils.MarshalJSON(u.Stream, "", true)
	}

	return nil, errors.New("could not marshal union type GetDirectiveComplianceIDResponseBody: all fields are null")
}

type GetDirectiveComplianceIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	OneOf *GetDirectiveComplianceIDResponseBody
}

func (o *GetDirectiveComplianceIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetDirectiveComplianceIDResponse) GetOneOf() *GetDirectiveComplianceIDResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
