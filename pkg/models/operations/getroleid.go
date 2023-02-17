package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetRoleIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetRoleIDSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetRoleIDRequest struct {
	PathParams GetRoleIDPathParams
	Security   GetRoleIDSecurity
}

type GetRoleIDResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
	Role        *shared.Role
}
