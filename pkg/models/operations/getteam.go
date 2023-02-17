package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetTeamSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetTeamRequest struct {
	Security GetTeamSecurity
}

type GetTeamResponse struct {
	ContentType string
	StatusCode  int
	Teams       *shared.Teams
}
