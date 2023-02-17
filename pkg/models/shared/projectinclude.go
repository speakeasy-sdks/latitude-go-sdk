package shared

type ProjectIncludeBilling struct {
	Method         *string `json:"method,omitempty"`
	SubscriptionID *string `json:"subscription_id,omitempty"`
	Type           *string `json:"type,omitempty"`
}

type ProjectIncludeStats struct {
	IPAddresses *int64 `json:"ip_addresses,omitempty"`
	Prefixes    *int64 `json:"prefixes,omitempty"`
	Servers     *int64 `json:"servers,omitempty"`
	Vlans       *int64 `json:"vlans,omitempty"`
}

type ProjectInclude struct {
	BandwidthAlert *bool                  `json:"bandwidth_alert,omitempty"`
	Billing        *ProjectIncludeBilling `json:"billing,omitempty"`
	BillingMethod  *string                `json:"billing_method,omitempty"`
	BillingType    *string                `json:"billing_type,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Environment    *string                `json:"environment,omitempty"`
	ID             *int64                 `json:"id,omitempty"`
	Name           *string                `json:"name,omitempty"`
	Slug           *string                `json:"slug,omitempty"`
	Stats          *ProjectIncludeStats   `json:"stats,omitempty"`
}
