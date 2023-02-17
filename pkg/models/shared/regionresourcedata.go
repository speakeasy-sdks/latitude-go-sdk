package shared

type RegionResourceDataSite struct {
	Facility *string `json:"facility,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Slug     *string `json:"slug,omitempty"`
}

type RegionResourceData struct {
	City    *string                 `json:"city,omitempty"`
	Country *string                 `json:"country,omitempty"`
	Site    *RegionResourceDataSite `json:"site,omitempty"`
}
