package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type DeleteAPIKeyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteAPIKeySecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteAPIKeyRequest struct {
	PathParams DeleteAPIKeyPathParams
	Security   DeleteAPIKeySecurity
}

type DeleteAPIKeyResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
}
