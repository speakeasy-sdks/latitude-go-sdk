package shared

type PlanDataAttributesAvailableInPricingBRL struct {
	Month *float64 `json:"month,omitempty"`
}

type PlanDataAttributesAvailableInPricingUSD struct {
	Month *float64 `json:"month,omitempty"`
}

type PlanDataAttributesAvailableInPricing struct {
	Brl *PlanDataAttributesAvailableInPricingBRL `json:"BRL,omitempty"`
	Usd *PlanDataAttributesAvailableInPricingUSD `json:"USD,omitempty"`
}

type PlanDataAttributesAvailableInRegion struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

type PlanDataAttributesAvailableInSites struct {
	ID         *string  `json:"id,omitempty"`
	InStock    *bool    `json:"in_stock,omitempty"`
	Instant    []string `json:"instant,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Slug       *string  `json:"slug,omitempty"`
	StockLevel *string  `json:"stock_level,omitempty"`
}

type PlanDataAttributesAvailableIn struct {
	Pricing *PlanDataAttributesAvailableInPricing `json:"pricing,omitempty"`
	Region  *PlanDataAttributesAvailableInRegion  `json:"region,omitempty"`
	Sites   []PlanDataAttributesAvailableInSites  `json:"sites,omitempty"`
}

type PlanDataAttributesSpecsCpus struct {
	Clock *float64 `json:"clock,omitempty"`
	Cores *float64 `json:"cores,omitempty"`
	Count *float64 `json:"count,omitempty"`
	Type  *string  `json:"type,omitempty"`
}

type PlanDataAttributesSpecsDrivesTypeEnum string

const (
	PlanDataAttributesSpecsDrivesTypeEnumSsd PlanDataAttributesSpecsDrivesTypeEnum = "SSD"
	PlanDataAttributesSpecsDrivesTypeEnumHdd PlanDataAttributesSpecsDrivesTypeEnum = "HDD"
)

type PlanDataAttributesSpecsDrives struct {
	Count *float64                               `json:"count,omitempty"`
	Size  *string                                `json:"size,omitempty"`
	Type  *PlanDataAttributesSpecsDrivesTypeEnum `json:"type,omitempty"`
}

type PlanDataAttributesSpecsFeatures struct {
	CustomScripts *bool `json:"custom_scripts,omitempty"`
	Raid          *bool `json:"raid,omitempty"`
	SSH           *bool `json:"ssh,omitempty"`
}

type PlanDataAttributesSpecsMemory struct {
	Total *string `json:"total,omitempty"`
}

type PlanDataAttributesSpecsNics struct {
	Count *float64 `json:"count,omitempty"`
	Type  *string  `json:"type,omitempty"`
}

type PlanDataAttributesSpecs struct {
	Cpus     []PlanDataAttributesSpecsCpus    `json:"cpus,omitempty"`
	Drives   []PlanDataAttributesSpecsDrives  `json:"drives,omitempty"`
	Features *PlanDataAttributesSpecsFeatures `json:"features,omitempty"`
	Memory   *PlanDataAttributesSpecsMemory   `json:"memory,omitempty"`
	Nics     []PlanDataAttributesSpecsNics    `json:"nics,omitempty"`
}

type PlanDataAttributes struct {
	AvailableIn []PlanDataAttributesAvailableIn `json:"available_in,omitempty"`
	Line        *string                         `json:"line,omitempty"`
	Name        *string                         `json:"name,omitempty"`
	Slug        *string                         `json:"slug,omitempty"`
	Specs       *PlanDataAttributesSpecs        `json:"specs,omitempty"`
}

type PlanDataTypeEnum string

const (
	PlanDataTypeEnumPlans PlanDataTypeEnum = "plans"
)

type PlanData struct {
	Attributes *PlanDataAttributes `json:"attributes,omitempty"`
	ID         *string             `json:"id,omitempty"`
	Type       *PlanDataTypeEnum   `json:"type,omitempty"`
}
