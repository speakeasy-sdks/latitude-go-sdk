package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type DestroyTeamMemberPathParams struct {
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type DestroyTeamMemberSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type DestroyTeamMemberRequest struct {
	PathParams DestroyTeamMemberPathParams
	Security   DestroyTeamMemberSecurity
}

type DestroyTeamMemberResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
}
