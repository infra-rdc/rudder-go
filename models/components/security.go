package components

type Security struct {
	APITokens string `security:"scheme,type=apiKey,subtype=header,name=X-API-Token"`
}

func (o *Security) GetAPITokens() string {
	if o == nil {
		return ""
	}
	return o.APITokens
}
