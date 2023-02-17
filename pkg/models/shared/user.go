package shared

type UserAttributesRole struct {
	CreatedAt *string `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type UserAttributes struct {
	AuthenticationFactorID *string             `json:"authentication_factor_id,omitempty"`
	Email                  *string             `json:"email,omitempty"`
	FirstName              *string             `json:"first_name,omitempty"`
	LastName               *string             `json:"last_name,omitempty"`
	Role                   *UserAttributesRole `json:"role,omitempty"`
	Teams                  []TeamInclude       `json:"teams,omitempty"`
}

type User struct {
	Attributes *UserAttributes `json:"attributes,omitempty"`
	ID         *string         `json:"id,omitempty"`
}
