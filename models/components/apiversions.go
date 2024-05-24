package components

type APIVersions struct {
	// Latest API version available
	Latest *int64 `json:"latest,omitempty"`
	// List of API version and status
	All []APIVersion `json:"all,omitempty"`
}

func (o *APIVersions) GetLatest() *int64 {
	if o == nil {
		return nil
	}
	return o.Latest
}

func (o *APIVersions) GetAll() []APIVersion {
	if o == nil {
		return nil
	}
	return o.All
}
