package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetProjectSSHKeyPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
	SSHKeyID        string `pathParam:"style=simple,explode=false,name=ssh_key_id"`
}

type GetProjectSSHKeyQueryParams struct {
	Include *string `queryParam:"style=form,explode=true,name=include"`
}

type GetProjectSSHKeySecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetProjectSSHKeyRequest struct {
	PathParams  GetProjectSSHKeyPathParams
	QueryParams GetProjectSSHKeyQueryParams
	Security    GetProjectSSHKeySecurity
}

type GetProjectSSHKey200ApplicationJSON struct {
	Data *shared.SSHKeyData `json:"data,omitempty"`
}

type GetProjectSSHKeyResponse struct {
	ContentType                              string
	StatusCode                               int
	GetProjectSSHKey200ApplicationJSONObject *GetProjectSSHKey200ApplicationJSON
}
