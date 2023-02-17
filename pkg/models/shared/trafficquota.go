package shared

type TrafficQuotaDataAttributesQuotaPerProjectQuotaPerRegionQuotaInMbps struct {
	Additional *int64 `json:"additional,omitempty"`
	Granted    *int64 `json:"granted,omitempty"`
	Total      *int64 `json:"total,omitempty"`
}

type TrafficQuotaDataAttributesQuotaPerProjectQuotaPerRegionQuotaInTb struct {
	Additional *int64 `json:"additional,omitempty"`
	Granted    *int64 `json:"granted,omitempty"`
	Total      *int64 `json:"total,omitempty"`
}

type TrafficQuotaDataAttributesQuotaPerProjectQuotaPerRegion struct {
	QuotaInMbps *TrafficQuotaDataAttributesQuotaPerProjectQuotaPerRegionQuotaInMbps `json:"quota_in_mbps,omitempty"`
	QuotaInTb   *TrafficQuotaDataAttributesQuotaPerProjectQuotaPerRegionQuotaInTb   `json:"quota_in_tb,omitempty"`
	RegionID    *int64                                                              `json:"region_id,omitempty"`
	RegionSlug  *string                                                             `json:"region_slug,omitempty"`
}

type TrafficQuotaDataAttributesQuotaPerProject struct {
	BillingMethod  *string                                                   `json:"billing_method,omitempty"`
	ProjectID      *int64                                                    `json:"project_id,omitempty"`
	ProjectSlug    *string                                                   `json:"project_slug,omitempty"`
	QuotaPerRegion []TrafficQuotaDataAttributesQuotaPerProjectQuotaPerRegion `json:"quota_per_region,omitempty"`
}

type TrafficQuotaDataAttributes struct {
	QuotaPerProject []TrafficQuotaDataAttributesQuotaPerProject `json:"quota_per_project,omitempty"`
}

type TrafficQuotaDataTypeEnum string

const (
	TrafficQuotaDataTypeEnumTrafficQuota TrafficQuotaDataTypeEnum = "traffic_quota"
)

type TrafficQuotaData struct {
	Attributes *TrafficQuotaDataAttributes `json:"attributes,omitempty"`
	ID         *string                     `json:"id,omitempty"`
	Type       *TrafficQuotaDataTypeEnum   `json:"type,omitempty"`
}

type TrafficQuota struct {
	Data *TrafficQuotaData `json:"data,omitempty"`
}
