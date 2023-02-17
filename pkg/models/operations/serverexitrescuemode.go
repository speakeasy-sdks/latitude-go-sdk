package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type ServerExitRescueModePathParams struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

type ServerExitRescueModeSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type ServerExitRescueModeRequest struct {
	PathParams ServerExitRescueModePathParams
	Security   ServerExitRescueModeSecurity
}

type ServerExitRescueModeResponse struct {
	ContentType  string
	StatusCode   int
	ErrorObject  *shared.ErrorObject
	ServerRescue *shared.ServerRescue
}
