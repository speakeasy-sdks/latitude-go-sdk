package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetAPIKeysSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetAPIKeysRequest struct {
	Security GetAPIKeysSecurity
}

type GetAPIKeysResponse struct {
	ContentType string
	StatusCode  int
	APIKey      *shared.APIKey
}
