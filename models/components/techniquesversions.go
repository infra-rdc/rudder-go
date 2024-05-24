package components

type TechniquesVersions struct {
	// Technique name
	Name *string `json:"name,omitempty"`
	// Available versions for this technique
	Versions []string `json:"versions,omitempty"`
}

func (o *TechniquesVersions) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TechniquesVersions) GetVersions() []string {
	if o == nil {
		return nil
	}
	return o.Versions
}
