package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetPlanPathParams struct {
	PlanID string `pathParam:"style=simple,explode=false,name=plan_id"`
}

type GetPlanSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetPlanRequest struct {
	PathParams GetPlanPathParams
	Security   GetPlanSecurity
}

type GetPlanResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
	Plan        *shared.Plan
}
