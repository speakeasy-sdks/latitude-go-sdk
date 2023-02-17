package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetVirtualNetworksQueryParams struct {
	FilterProject *string `queryParam:"style=form,explode=true,name=filter[project]"`
	FilterSite    *string `queryParam:"style=form,explode=true,name=filter[site]"`
}

type GetVirtualNetworksSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetVirtualNetworksRequest struct {
	QueryParams GetVirtualNetworksQueryParams
	Security    GetVirtualNetworksSecurity
}

type GetVirtualNetworksResponse struct {
	ContentType     string
	StatusCode      int
	VirtualNetworks *shared.VirtualNetworks
}
