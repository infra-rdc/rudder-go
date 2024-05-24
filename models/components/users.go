package components

type Users struct {
	// If you want to provide hashed password set this property to `true` otherwise we will hash the plain password and store the hash
	IsPreHahed *bool   `json:"isPreHahed,omitempty"`
	Username   *string `json:"username,omitempty"`
	// this password will be hashed for you if the `isPreHashed` is set on false
	Password *string `json:"password,omitempty"`
	// Defined user's permissions
	Role []string `json:"role,omitempty"`
}

func (o *Users) GetIsPreHahed() *bool {
	if o == nil {
		return nil
	}
	return o.IsPreHahed
}

func (o *Users) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *Users) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *Users) GetRole() []string {
	if o == nil {
		return nil
	}
	return o.Role
}
