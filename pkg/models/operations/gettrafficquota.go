package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetTrafficQuotaQueryParams struct {
	FilterProject *string `queryParam:"style=form,explode=true,name=filter[project]"`
}

type GetTrafficQuotaSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetTrafficQuotaRequest struct {
	QueryParams GetTrafficQuotaQueryParams
	Security    GetTrafficQuotaSecurity
}

type GetTrafficQuotaResponse struct {
	ContentType  string
	StatusCode   int
	ErrorObject  *shared.ErrorObject
	TrafficQuota *shared.TrafficQuota
}
