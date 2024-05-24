package components

import (
	"encoding/json"
	"fmt"
)

type NodeComposition string

const (
	NodeCompositionAnd NodeComposition = "and"
	NodeCompositionOr  NodeComposition = "or"
)

func (e NodeComposition) ToPointer() *NodeComposition {
	return &e
}
func (e *NodeComposition) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = NodeComposition(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeComposition: %v", v)
	}
}
