package shared

type PlansBandwidthAttributesProjectsPackages struct {
	Contracted *int64   `json:"contracted,omitempty"`
	Currency   *string  `json:"currency,omitempty"`
	RegionSlug *string  `json:"region_slug,omitempty"`
	TotalPrice *float64 `json:"total_price,omitempty"`
	UnitPrice  *float64 `json:"unit_price,omitempty"`
}

type PlansBandwidthAttributesProjectsProject struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

type PlansBandwidthAttributesProjects struct {
	Packages []PlansBandwidthAttributesProjectsPackages `json:"packages,omitempty"`
	Project  *PlansBandwidthAttributesProjectsProject   `json:"project,omitempty"`
}

type PlansBandwidthAttributes struct {
	Projects []PlansBandwidthAttributesProjects `json:"projects,omitempty"`
}

type PlansBandwidthTypeEnum string

const (
	PlansBandwidthTypeEnumBandwidthPackages PlansBandwidthTypeEnum = "bandwidth_packages"
)

type PlansBandwidth struct {
	Attributes *PlansBandwidthAttributes `json:"attributes,omitempty"`
	ID         *string                   `json:"id,omitempty"`
	Type       *PlansBandwidthTypeEnum   `json:"type,omitempty"`
}
