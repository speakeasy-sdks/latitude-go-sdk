package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetIPPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetIPQueryParams struct {
	ExtraFieldsIPAddresses *string `queryParam:"style=form,explode=true,name=extra_fields[ip_addresses]"`
	Include                *string `queryParam:"style=form,explode=true,name=include"`
}

type GetIPSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetIPRequest struct {
	PathParams  GetIPPathParams
	QueryParams GetIPQueryParams
	Security    GetIPSecurity
}

type GetIPResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
	IPAddress   *shared.IPAddress
}
