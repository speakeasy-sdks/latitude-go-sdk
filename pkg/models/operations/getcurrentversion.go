package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetCurrentVersionSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetCurrentVersionRequest struct {
	Security GetCurrentVersionSecurity
}

type GetCurrentVersionResponse struct {
	ContentType string
	StatusCode  int
}
