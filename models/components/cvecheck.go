package components

// Score - CVE score
type Score struct {
	Value    *string `json:"value,omitempty"`
	Severity *string `json:"severity,omitempty"`
}

func (o *Score) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *Score) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

// Packages affected by this CVE
type Packages struct {
	// Name of the package affected
	Name *string `json:"name,omitempty"`
	// Version of the package affected
	Version *string `json:"version,omitempty"`
}

func (o *Packages) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Packages) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

type CveCheck struct {
	// CVE id
	CveID *string `json:"cveId,omitempty"`
	// CVE score
	Score *Score `json:"score,omitempty"`
	// Id of Nodes affected by this CVE
	Nodes []string `json:"nodes,omitempty"`
	// Packages affected by this CVE
	Packages []Packages `json:"packages,omitempty"`
}

func (o *CveCheck) GetCveID() *string {
	if o == nil {
		return nil
	}
	return o.CveID
}

func (o *CveCheck) GetScore() *Score {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *CveCheck) GetNodes() []string {
	if o == nil {
		return nil
	}
	return o.Nodes
}

func (o *CveCheck) GetPackages() []Packages {
	if o == nil {
		return nil
	}
	return o.Packages
}
