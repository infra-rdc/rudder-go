package components

import (
	"encoding/json"
	"fmt"
)

type NodeDeleteMode string

const (
	NodeDeleteModeMove  NodeDeleteMode = "move"
	NodeDeleteModeErase NodeDeleteMode = "erase"
)

func (e NodeDeleteMode) ToPointer() *NodeDeleteMode {
	return &e
}
func (e *NodeDeleteMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "move":
		fallthrough
	case "erase":
		*e = NodeDeleteMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeDeleteMode: %v", v)
	}
}
