package components

type PluginsSettings struct {
	// username to use for Rudder account
	Username string `json:"username"`
	// password to access Rudder account
	Password string `json:"password"`
	// URL for getting plugins
	URL *string `json:"url,omitempty"`
	// if an authenticated proxy is necessary, username of proxy
	ProxyUser *string `json:"proxyUser,omitempty"`
	// if an authenticated proxy is necessary, password of the proxy
	ProxyPassword *string `json:"proxyPassword,omitempty"`
	// proxy URL to use
	ProxyURL *string `json:"proxyUrl,omitempty"`
}

func (o *PluginsSettings) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *PluginsSettings) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *PluginsSettings) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *PluginsSettings) GetProxyUser() *string {
	if o == nil {
		return nil
	}
	return o.ProxyUser
}

func (o *PluginsSettings) GetProxyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ProxyPassword
}

func (o *PluginsSettings) GetProxyURL() *string {
	if o == nil {
		return nil
	}
	return o.ProxyURL
}
