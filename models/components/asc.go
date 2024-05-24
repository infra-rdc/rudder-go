package components

import (
	"encoding/json"
	"fmt"
)

type Asc string

const (
	AscAsc  Asc = "asc"
	AscDesc Asc = "desc"
)

func (e Asc) ToPointer() *Asc {
	return &e
}
func (e *Asc) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = Asc(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Asc: %v", v)
	}
}
