package shared

type APIKeyAttributes struct {
	CreatedAt      *string `json:"created_at,omitempty"`
	LastUsedAt     *string `json:"last_used_at,omitempty"`
	Name           *string `json:"name,omitempty"`
	TokenLastSlice *string `json:"token_last_slice,omitempty"`
	UpdatedAt      *string `json:"updated_at,omitempty"`
}

type APIKey struct {
	Attributes *APIKeyAttributes `json:"attributes,omitempty"`
	ID         *string           `json:"id,omitempty"`
}
