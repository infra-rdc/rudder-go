package components

// Timezone information of the node.
type Timezone struct {
	// The standard name of the timezone
	Name string `json:"name"`
	// Timezone offset compared to UTC, in +/-HHMM format
	Offset string `json:"offset"`
}

func (o *Timezone) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Timezone) GetOffset() string {
	if o == nil {
		return ""
	}
	return o.Offset
}
