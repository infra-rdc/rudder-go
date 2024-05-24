package components

import (
	"encoding/json"
	"fmt"
)

type ArchiveKind string

const (
	ArchiveKindFull       ArchiveKind = "full"
	ArchiveKindGroups     ArchiveKind = "groups"
	ArchiveKindRules      ArchiveKind = "rules"
	ArchiveKindDirectives ArchiveKind = "directives"
	ArchiveKindParameters ArchiveKind = "parameters"
)

func (e ArchiveKind) ToPointer() *ArchiveKind {
	return &e
}
func (e *ArchiveKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "full":
		fallthrough
	case "groups":
		fallthrough
	case "rules":
		fallthrough
	case "directives":
		fallthrough
	case "parameters":
		*e = ArchiveKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ArchiveKind: %v", v)
	}
}
