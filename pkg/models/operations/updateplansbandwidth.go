package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type UpdatePlansBandwidthRequestBodyDataAttributes struct {
	Project    *string `json:"project,omitempty"`
	Quantity   *int64  `json:"quantity,omitempty"`
	RegionSlug *string `json:"region_slug,omitempty"`
}

type UpdatePlansBandwidthRequestBodyDataTypeEnum string

const (
	UpdatePlansBandwidthRequestBodyDataTypeEnumBandwidthPackages UpdatePlansBandwidthRequestBodyDataTypeEnum = "bandwidth_packages"
)

type UpdatePlansBandwidthRequestBodyData struct {
	Attributes *UpdatePlansBandwidthRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       *UpdatePlansBandwidthRequestBodyDataTypeEnum   `json:"type,omitempty"`
}

type UpdatePlansBandwidthRequestBody struct {
	Data *UpdatePlansBandwidthRequestBodyData `json:"data,omitempty"`
}

type UpdatePlansBandwidthSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type UpdatePlansBandwidthRequest struct {
	Request  *UpdatePlansBandwidthRequestBody `request:"mediaType=application/json"`
	Security UpdatePlansBandwidthSecurity
}

type UpdatePlansBandwidthResponse struct {
	ContentType    string
	StatusCode     int
	ErrorObject    *shared.ErrorObject
	PlansBandwidth *shared.PlansBandwidth
}
