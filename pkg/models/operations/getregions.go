package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetRegionsSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetRegionsRequest struct {
	Security GetRegionsSecurity
}

type GetRegionsResponse struct {
	ContentType string
	StatusCode  int
	Regions     *shared.Regions
}
