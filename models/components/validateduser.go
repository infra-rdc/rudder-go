package components

// ValidatedUser - list of users with their workflow settings
type ValidatedUser struct {
	Username string `json:"username"`
	// whether the user's actions generate chanque-request or not
	IsValidated bool `json:"isValidated"`
	// whether the user exists in file or not
	UserExists bool `json:"userExists"`
}

func (o *ValidatedUser) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ValidatedUser) GetIsValidated() bool {
	if o == nil {
		return false
	}
	return o.IsValidated
}

func (o *ValidatedUser) GetUserExists() bool {
	if o == nil {
		return false
	}
	return o.UserExists
}
