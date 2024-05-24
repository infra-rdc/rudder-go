package components

// APIEndpoints - objects with two fields, the first one has the endpoint name as key and its description as value, the second one has HTTP verb to use (GET, POST PUT, DELETE) as key and the supported version an API path for value.
type APIEndpoints struct {
	// The endpoint name for key and its description for value
	EndpointName *string `json:"endpointName,omitempty"`
	HTTPVerb     any     `json:"httpVerb,omitempty"`
}

func (o *APIEndpoints) GetEndpointName() *string {
	if o == nil {
		return nil
	}
	return o.EndpointName
}

func (o *APIEndpoints) GetHTTPVerb() any {
	if o == nil {
		return nil
	}
	return o.HTTPVerb
}
