package shared

type ProjectAttributesBilling struct {
	Method         *string `json:"method,omitempty"`
	SubscriptionID *string `json:"subscription_id,omitempty"`
	Type           *string `json:"type,omitempty"`
}

type ProjectAttributesBillingMethodEnum string

const (
	ProjectAttributesBillingMethodEnumNormal                 ProjectAttributesBillingMethodEnum = "Normal"
	ProjectAttributesBillingMethodEnumNinetyFivethPercentile ProjectAttributesBillingMethodEnum = "95th percentile"
	ProjectAttributesBillingMethodEnumNull                   ProjectAttributesBillingMethodEnum = "null"
)

type ProjectAttributesBillingTypeEnum string

const (
	ProjectAttributesBillingTypeEnumYearly  ProjectAttributesBillingTypeEnum = "Yearly"
	ProjectAttributesBillingTypeEnumMonthly ProjectAttributesBillingTypeEnum = "Monthly"
	ProjectAttributesBillingTypeEnumHourly  ProjectAttributesBillingTypeEnum = "Hourly"
	ProjectAttributesBillingTypeEnumNull    ProjectAttributesBillingTypeEnum = "null"
)

type ProjectAttributesEnvironmentEnum string

const (
	ProjectAttributesEnvironmentEnumDevelopment ProjectAttributesEnvironmentEnum = "Development"
	ProjectAttributesEnvironmentEnumStaging     ProjectAttributesEnvironmentEnum = "Staging"
	ProjectAttributesEnvironmentEnumProduction  ProjectAttributesEnvironmentEnum = "Production"
	ProjectAttributesEnvironmentEnumNull        ProjectAttributesEnvironmentEnum = "null"
)

type ProjectAttributesStats struct {
	IPAddresses *float64 `json:"ip_addresses,omitempty"`
	Prefixes    *float64 `json:"prefixes,omitempty"`
	Servers     *float64 `json:"servers,omitempty"`
	Vlans       *float64 `json:"vlans,omitempty"`
}

type ProjectAttributes struct {
	Billing       *ProjectAttributesBilling           `json:"billing,omitempty"`
	BillingMethod *ProjectAttributesBillingMethodEnum `json:"billing_method,omitempty"`
	BillingType   *ProjectAttributesBillingTypeEnum   `json:"billing_type,omitempty"`
	Cost          *string                             `json:"cost,omitempty"`
	CreatedAt     *string                             `json:"created_at,omitempty"`
	Description   *string                             `json:"description,omitempty"`
	Environment   *ProjectAttributesEnvironmentEnum   `json:"environment,omitempty"`
	Name          *string                             `json:"name,omitempty"`
	Slug          *string                             `json:"slug,omitempty"`
	Stats         *ProjectAttributesStats             `json:"stats,omitempty"`
	Team          *TeamInclude                        `json:"team,omitempty"`
	UpdatedAt     *string                             `json:"updated_at,omitempty"`
}

type ProjectRelationships struct {
	Billing *LazySideload `json:"billing,omitempty"`
	Team    *LazySideload `json:"team,omitempty"`
}

type Project struct {
	Attributes    *ProjectAttributes    `json:"attributes,omitempty"`
	ID            *string               `json:"id,omitempty"`
	Relationships *ProjectRelationships `json:"relationships,omitempty"`
}
