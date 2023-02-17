package shared

type VirtualNetworkAttributesRegionSite struct {
	Facility *string `json:"facility,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Slug     *string `json:"slug,omitempty"`
}

type VirtualNetworkAttributesRegion struct {
	City    *string                             `json:"city,omitempty"`
	Country *string                             `json:"country,omitempty"`
	Site    *VirtualNetworkAttributesRegionSite `json:"site,omitempty"`
}

type VirtualNetworkAttributes struct {
	AssignmentsCount *int64                          `json:"assignments_count,omitempty"`
	Description      *string                         `json:"description,omitempty"`
	Name             *string                         `json:"name,omitempty"`
	Region           *VirtualNetworkAttributesRegion `json:"region,omitempty"`
	Vid              *int64                          `json:"vid,omitempty"`
}

type VirtualNetworkTypeEnum string

const (
	VirtualNetworkTypeEnumVirtualNetworks VirtualNetworkTypeEnum = "virtual_networks"
)

type VirtualNetwork struct {
	Attributes *VirtualNetworkAttributes `json:"attributes,omitempty"`
	ID         *string                   `json:"id,omitempty"`
	Type       *VirtualNetworkTypeEnum   `json:"type,omitempty"`
}
