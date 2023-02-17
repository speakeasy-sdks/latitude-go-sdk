package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetRegionPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetRegionSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetRegionRequest struct {
	PathParams GetRegionPathParams
	Security   GetRegionSecurity
}

type GetRegionResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
	Region      *shared.Region
}
