package shared

type TeamAttributesBilling struct {
	CustomerBillingID *string `json:"customer_billing_id,omitempty"`
	ID                *int64  `json:"id,omitempty"`
}

type TeamAttributes struct {
	Address     *string                `json:"address,omitempty"`
	Billing     *TeamAttributesBilling `json:"billing,omitempty"`
	CreatedAt   *string                `json:"created_at,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	Description *string                `json:"description,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Owner       *UserInclude           `json:"owner,omitempty"`
	Projects    []ProjectInclude       `json:"projects,omitempty"`
	Slug        *string                `json:"slug,omitempty"`
	UpdatedAt   *string                `json:"updated_at,omitempty"`
	Users       []UserInclude          `json:"users,omitempty"`
}

type TeamRelationships struct {
	Billing  *LazySideload `json:"billing,omitempty"`
	Owner    *LazySideload `json:"owner,omitempty"`
	Projects *LazySideload `json:"projects,omitempty"`
	Users    *LazySideload `json:"users,omitempty"`
}

type Team struct {
	Attributes    *TeamAttributes    `json:"attributes,omitempty"`
	ID            *string            `json:"id,omitempty"`
	Relationships *TeamRelationships `json:"relationships,omitempty"`
}
