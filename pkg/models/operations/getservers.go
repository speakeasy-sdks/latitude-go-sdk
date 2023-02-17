package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetServersQueryParams struct {
	ExtraFieldsServers    *string `queryParam:"style=form,explode=true,name=extra_fields[servers]"`
	FilterCreatedAtGte    *string `queryParam:"style=form,explode=true,name=filter[created_at_gte]"`
	FilterCreatedAtLte    *string `queryParam:"style=form,explode=true,name=filter[created_at_lte]"`
	FilterHostname        *string `queryParam:"style=form,explode=true,name=filter[hostname]"`
	FilterLabel           *string `queryParam:"style=form,explode=true,name=filter[label]"`
	FilterOperatingSystem *string `queryParam:"style=form,explode=true,name=filter[operating_system]"`
	FilterPlan            *string `queryParam:"style=form,explode=true,name=filter[plan]"`
	FilterProject         *string `queryParam:"style=form,explode=true,name=filter[project]"`
	FilterRegion          *string `queryParam:"style=form,explode=true,name=filter[region]"`
	FilterStatus          *string `queryParam:"style=form,explode=true,name=filter[status]"`
}

type GetServersSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetServersRequest struct {
	QueryParams GetServersQueryParams
	Security    GetServersSecurity
}

type GetServersResponse struct {
	ContentType string
	StatusCode  int
	Servers     *shared.Servers
}
