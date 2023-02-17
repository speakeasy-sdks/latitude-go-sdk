package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetVirtualNetworkPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetVirtualNetworkSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetVirtualNetworkRequest struct {
	PathParams GetVirtualNetworkPathParams
	Security   GetVirtualNetworkSecurity
}

type GetVirtualNetwork200ApplicationJSON struct {
	Data *shared.VirtualNetwork `json:"data,omitempty"`
}

type GetVirtualNetworkResponse struct {
	ContentType                               string
	StatusCode                                int
	GetVirtualNetwork200ApplicationJSONObject *GetVirtualNetwork200ApplicationJSON
}
