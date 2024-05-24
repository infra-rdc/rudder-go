package components

type All struct {
	// API Version
	Version int64 `json:"version"`
	// Status of the API, either maintained or deprecated
	Status any `json:"status"`
}

func (o *All) GetVersion() int64 {
	if o == nil {
		return 0
	}
	return o.Version
}

func (o *All) GetStatus() any {
	if o == nil {
		return nil
	}
	return o.Status
}

type APIVersion struct {
	// Latest API version available
	Latest *int64 `json:"latest,omitempty"`
	// List of API version and status
	All []All `json:"all,omitempty"`
}

func (o *APIVersion) GetLatest() *int64 {
	if o == nil {
		return nil
	}
	return o.Latest
}

func (o *APIVersion) GetAll() []All {
	if o == nil {
		return nil
	}
	return o.All
}
