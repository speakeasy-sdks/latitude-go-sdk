package shared

type TeamInclude struct {
	Address     *string                `json:"address,omitempty"`
	Currency    map[string]interface{} `json:"currency,omitempty"`
	Description *string                `json:"description,omitempty"`
	ID          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Slug        *string                `json:"slug,omitempty"`
	Status      *string                `json:"status,omitempty"`
}
