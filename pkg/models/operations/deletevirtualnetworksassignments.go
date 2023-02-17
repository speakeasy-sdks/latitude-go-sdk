package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type DeleteVirtualNetworksAssignmentsPathParams struct {
	AssignmentID string `pathParam:"style=simple,explode=false,name=assignment_id"`
}

type DeleteVirtualNetworksAssignmentsSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteVirtualNetworksAssignmentsRequest struct {
	PathParams DeleteVirtualNetworksAssignmentsPathParams
	Security   DeleteVirtualNetworksAssignmentsSecurity
}

type DeleteVirtualNetworksAssignmentsResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
}
