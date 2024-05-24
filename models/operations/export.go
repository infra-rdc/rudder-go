package operations

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/infra-rdc/rudder-go/models/components"
)

type Include string

const (
	IncludeAllDefault Include = "all (default)"
	IncludeNone       Include = "none"
	IncludeDirectives Include = "directives"
	IncludeTechniques Include = "techniques"
	IncludeGroups     Include = "groups"
)

func (e Include) ToPointer() *Include {
	return &e
}
func (e *Include) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "all (default)":
		fallthrough
	case "none":
		fallthrough
	case "directives":
		fallthrough
	case "techniques":
		fallthrough
	case "groups":
		*e = Include(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Include: %v", v)
	}
}

type ExportRequest struct {
	// IDs (optionally with revision, '+' need to be escaped as '%2B') of rules to include
	Rules []string `queryParam:"style=form,explode=false,name=rules"`
	// IDs (optionally with revision, '+' need to be escaped as '%2B') of directives to include
	Directives []string `queryParam:"style=form,explode=false,name=directives"`
	// IDs, ie technique name/technique version (optionally with revision, '+' need to be escaped as '%2B') of techniques to include
	Techniques []string `queryParam:"style=form,explode=false,name=techniques"`
	// IDs (optionally with revision, '+' need to be escaped as '%2B') of groups to include
	Groups []string `queryParam:"style=form,explode=false,name=groups"`
	// Scope of dependencies to include in archive, where rule as directives and groups dependencies, directives have techniques dependencies, and techniques and groups don't have dependencies. 'none' means no dependencies will be include, 'all' means that the whole tree will,  'directives' and 'groups' means to include them specifically, 'techniques' means to include both directives and techniques.
	Include []Include `queryParam:"style=form,explode=false,name=include"`
}

func (o *ExportRequest) GetRules() []string {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *ExportRequest) GetDirectives() []string {
	if o == nil {
		return nil
	}
	return o.Directives
}

func (o *ExportRequest) GetTechniques() []string {
	if o == nil {
		return nil
	}
	return o.Techniques
}

func (o *ExportRequest) GetGroups() []string {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *ExportRequest) GetInclude() []Include {
	if o == nil {
		return nil
	}
	return o.Include
}

type ExportResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A zip archive with the queried content.
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Stream io.ReadCloser
}

func (o *ExportResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ExportResponse) GetStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Stream
}
