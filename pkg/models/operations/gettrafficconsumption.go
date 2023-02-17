package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetTrafficConsumptionQueryParams struct {
	FilterFromDate *int64 `queryParam:"style=form,explode=true,name=filter[from_date]"`
	FilterProject  *int64 `queryParam:"style=form,explode=true,name=filter[project]"`
	FilterServer   *int64 `queryParam:"style=form,explode=true,name=filter[server]"`
	FilterToDate   *int64 `queryParam:"style=form,explode=true,name=filter[to_date]"`
}

type GetTrafficConsumptionSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetTrafficConsumptionRequest struct {
	QueryParams GetTrafficConsumptionQueryParams
	Security    GetTrafficConsumptionSecurity
}

type GetTrafficConsumptionResponse struct {
	ContentType string
	StatusCode  int
	Traffic     *shared.Traffic
}
