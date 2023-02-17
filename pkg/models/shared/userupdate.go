package shared

type UserUpdateAttributes struct {
	AuthenticationFactorID *string `json:"authentication_factor_id,omitempty"`
	Email                  *string `json:"email,omitempty"`
	FirstName              *string `json:"first_name,omitempty"`
	LastName               *string `json:"last_name,omitempty"`
	Role                   *string `json:"role,omitempty"`
}

type UserUpdate struct {
	Attributes *UserUpdateAttributes `json:"attributes,omitempty"`
	ID         *string               `json:"id,omitempty"`
}
