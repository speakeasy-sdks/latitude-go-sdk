package shared

type TeamMembersDataRole struct {
	CreatedAt *string `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type TeamMembersData struct {
	Email      *string              `json:"email,omitempty"`
	FirstName  *string              `json:"first_name,omitempty"`
	LastName   *string              `json:"last_name,omitempty"`
	MfaEnabled *bool                `json:"mfa_enabled,omitempty"`
	Role       *TeamMembersDataRole `json:"role,omitempty"`
}

type TeamMembers struct {
	Data []TeamMembersData `json:"data,omitempty"`
}
