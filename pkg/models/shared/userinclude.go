package shared

type UserIncludeRole struct {
	CreatedAt *string `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type UserInclude struct {
	AuthenticationFactorID *string          `json:"authentication_factor_id,omitempty"`
	CreatedAt              *string          `json:"created_at,omitempty"`
	Email                  *string          `json:"email,omitempty"`
	FirstName              *string          `json:"first_name,omitempty"`
	ID                     *string          `json:"id,omitempty"`
	LastName               *string          `json:"last_name,omitempty"`
	Role                   *UserIncludeRole `json:"role,omitempty"`
	UpdatedAt              *string          `json:"updated_at,omitempty"`
}
