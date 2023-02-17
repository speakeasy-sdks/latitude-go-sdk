package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetPlansBandwidthSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetPlansBandwidthRequest struct {
	Security GetPlansBandwidthSecurity
}

type GetPlansBandwidthResponse struct {
	ContentType    string
	StatusCode     int
	PlansBandwidth *shared.PlansBandwidth
}
