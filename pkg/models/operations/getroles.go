package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetRolesSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetRolesRequest struct {
	Security GetRolesSecurity
}

type GetRoles200ApplicationJSON struct {
	Data []shared.RoleData `json:"data,omitempty"`
}

type GetRolesResponse struct {
	ContentType                      string
	StatusCode                       int
	GetRoles200ApplicationJSONObject *GetRoles200ApplicationJSON
}
