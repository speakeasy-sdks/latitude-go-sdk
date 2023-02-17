package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type DestroyVirtualNetworkPathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type DestroyVirtualNetworkSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type DestroyVirtualNetworkRequest struct {
	PathParams DestroyVirtualNetworkPathParams
	Security   DestroyVirtualNetworkSecurity
}

type DestroyVirtualNetworkResponse struct {
	ContentType    string
	StatusCode     int
	ErrorObject    *shared.ErrorObject
	VirtualNetwork *shared.VirtualNetwork
}
