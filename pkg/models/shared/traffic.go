package shared

type TrafficDataAttributesRegionsData struct {
	AvgInboundSpeedMbps  *float64 `json:"avg_inbound_speed_mbps,omitempty"`
	AvgOutboundSpeedMbps *float64 `json:"avg_outbound_speed_mbps,omitempty"`
	Date                 *string  `json:"date,omitempty"`
	InboundGb            *int64   `json:"inbound_gb,omitempty"`
	OutboundGb           *int64   `json:"outbound_gb,omitempty"`
}

type TrafficDataAttributesRegions struct {
	Data                            []TrafficDataAttributesRegionsData `json:"data,omitempty"`
	RegionSlug                      *string                            `json:"region_slug,omitempty"`
	TotalInbound95thPercentileMbps  *float64                           `json:"total_inbound_95th_percentile_mbps,omitempty"`
	TotalInboundGb                  *int64                             `json:"total_inbound_gb,omitempty"`
	TotalOutbound95thPercentileMbps *float64                           `json:"total_outbound_95th_percentile_mbps,omitempty"`
	TotalOutboundGb                 *int64                             `json:"total_outbound_gb,omitempty"`
}

type TrafficDataAttributes struct {
	FromDate                        *int64                         `json:"from_date,omitempty"`
	Regions                         []TrafficDataAttributesRegions `json:"regions,omitempty"`
	ToDate                          *int64                         `json:"to_date,omitempty"`
	TotalInbound95thPercentileMbps  *float64                       `json:"total_inbound_95th_percentile_mbps,omitempty"`
	TotalInboundGb                  *int64                         `json:"total_inbound_gb,omitempty"`
	TotalOutbound95thPercentileMbps *float64                       `json:"total_outbound_95th_percentile_mbps,omitempty"`
	TotalOutboundGb                 *int64                         `json:"total_outbound_gb,omitempty"`
}

type TrafficDataTypeEnum string

const (
	TrafficDataTypeEnumTraffic TrafficDataTypeEnum = "traffic"
)

type TrafficData struct {
	Attributes *TrafficDataAttributes `json:"attributes,omitempty"`
	ID         *string                `json:"id,omitempty"`
	Type       *TrafficDataTypeEnum   `json:"type,omitempty"`
}

type Traffic struct {
	Data *TrafficData `json:"data,omitempty"`
}
