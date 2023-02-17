package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetVirtualNetworksAssignmentsQueryParams struct {
	FilterServer           *string `queryParam:"style=form,explode=true,name=filter[server]"`
	FilterVid              *string `queryParam:"style=form,explode=true,name=filter[vid]"`
	FilterVirtualNetworkID *string `queryParam:"style=form,explode=true,name=filter[virtual_network_id]"`
	Include                *string `queryParam:"style=form,explode=true,name=include"`
}

type GetVirtualNetworksAssignmentsSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetVirtualNetworksAssignmentsRequest struct {
	QueryParams GetVirtualNetworksAssignmentsQueryParams
	Security    GetVirtualNetworksAssignmentsSecurity
}

type GetVirtualNetworksAssignmentsResponse struct {
	ContentType               string
	StatusCode                int
	VirtualNetworkAssignments *shared.VirtualNetworkAssignments
}
