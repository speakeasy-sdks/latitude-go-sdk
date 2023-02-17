package shared

type RegionsDataAttributesCountry struct {
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

type RegionsDataAttributes struct {
	Country *RegionsDataAttributesCountry `json:"country,omitempty"`
	Name    *string                       `json:"name,omitempty"`
	Slug    *string                       `json:"slug,omitempty"`
}

type RegionsData struct {
	Attributes *RegionsDataAttributes `json:"attributes,omitempty"`
	ID         *string                `json:"id,omitempty"`
}

type Regions struct {
	Data []RegionsData `json:"data,omitempty"`
}
