package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type ServerStartRescueModePathParams struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

type ServerStartRescueModeSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type ServerStartRescueModeRequest struct {
	PathParams ServerStartRescueModePathParams
	Security   ServerStartRescueModeSecurity
}

type ServerStartRescueModeResponse struct {
	ContentType  string
	StatusCode   int
	ErrorObject  *shared.ErrorObject
	ServerRescue *shared.ServerRescue
}
