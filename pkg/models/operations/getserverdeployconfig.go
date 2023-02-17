package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetServerDeployConfigPathParams struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

type GetServerDeployConfigSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetServerDeployConfigRequest struct {
	PathParams GetServerDeployConfigPathParams
	Security   GetServerDeployConfigSecurity
}

type GetServerDeployConfigResponse struct {
	ContentType  string
	StatusCode   int
	DeployConfig *shared.DeployConfig
}
