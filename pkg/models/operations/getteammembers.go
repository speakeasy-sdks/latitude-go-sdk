package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetTeamMembersSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetTeamMembersRequest struct {
	Security GetTeamMembersSecurity
}

type GetTeamMembersResponse struct {
	ContentType string
	StatusCode  int
	TeamMembers *shared.TeamMembers
}
