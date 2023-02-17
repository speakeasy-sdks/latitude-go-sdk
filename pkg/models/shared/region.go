package shared

type RegionDataAttributesCountry struct {
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

type RegionDataAttributes struct {
	Country *RegionDataAttributesCountry `json:"country,omitempty"`
	Name    *string                      `json:"name,omitempty"`
	Slug    *string                      `json:"slug,omitempty"`
}

type RegionData struct {
	Attributes *RegionDataAttributes `json:"attributes,omitempty"`
	ID         *string               `json:"id,omitempty"`
}

type Region struct {
	Data *RegionData `json:"data,omitempty"`
}
