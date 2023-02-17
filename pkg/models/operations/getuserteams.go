package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetUserTeamsSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetUserTeamsRequest struct {
	Security GetUserTeamsSecurity
}

type GetUserTeamsResponse struct {
	ContentType string
	StatusCode  int
	Teams       *shared.Teams
}
