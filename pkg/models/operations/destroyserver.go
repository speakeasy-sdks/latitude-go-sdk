package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type DestroyServerPathParams struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

type DestroyServerSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type DestroyServerRequest struct {
	PathParams DestroyServerPathParams
	Security   DestroyServerSecurity
}

type DestroyServerResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
}
