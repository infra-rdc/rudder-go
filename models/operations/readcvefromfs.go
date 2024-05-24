package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ReadCVEfromFSResult - Result of the request
type ReadCVEfromFSResult string

const (
	ReadCVEfromFSResultSuccess ReadCVEfromFSResult = "success"
	ReadCVEfromFSResultError   ReadCVEfromFSResult = "error"
)

func (e ReadCVEfromFSResult) ToPointer() *ReadCVEfromFSResult {
	return &e
}
func (e *ReadCVEfromFSResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReadCVEfromFSResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReadCVEfromFSResult: %v", v)
	}
}

// ReadCVEfromFSAction - The id of the action
type ReadCVEfromFSAction string

const (
	ReadCVEfromFSActionReadCvEfromFs ReadCVEfromFSAction = "readCVEfromFS"
)

func (e ReadCVEfromFSAction) ToPointer() *ReadCVEfromFSAction {
	return &e
}
func (e *ReadCVEfromFSAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "readCVEfromFS":
		*e = ReadCVEfromFSAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReadCVEfromFSAction: %v", v)
	}
}

type ReadCVEfromFSData struct {
	CVEs int64 `json:"CVEs"`
}

func (o *ReadCVEfromFSData) GetCVEs() int64 {
	if o == nil {
		return 0
	}
	return o.CVEs
}

// ReadCVEfromFSResponseBody - updated CVE count
type ReadCVEfromFSResponseBody struct {
	// Result of the request
	Result ReadCVEfromFSResult `json:"result"`
	// The id of the action
	Action ReadCVEfromFSAction `json:"action"`
	Data   ReadCVEfromFSData   `json:"data"`
}

func (o *ReadCVEfromFSResponseBody) GetResult() ReadCVEfromFSResult {
	if o == nil {
		return ReadCVEfromFSResult("")
	}
	return o.Result
}

func (o *ReadCVEfromFSResponseBody) GetAction() ReadCVEfromFSAction {
	if o == nil {
		return ReadCVEfromFSAction("")
	}
	return o.Action
}

func (o *ReadCVEfromFSResponseBody) GetData() ReadCVEfromFSData {
	if o == nil {
		return ReadCVEfromFSData{}
	}
	return o.Data
}

type ReadCVEfromFSResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// updated CVE count
	Object *ReadCVEfromFSResponseBody
}

func (o *ReadCVEfromFSResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReadCVEfromFSResponse) GetObject() *ReadCVEfromFSResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
