package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetProjectSSHKeysPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
}

type GetProjectSSHKeysQueryParams struct {
	Include *string `queryParam:"style=form,explode=true,name=include"`
}

type GetProjectSSHKeysSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetProjectSSHKeysRequest struct {
	PathParams  GetProjectSSHKeysPathParams
	QueryParams GetProjectSSHKeysQueryParams
	Security    GetProjectSSHKeysSecurity
}

type GetProjectSSHKeysResponse struct {
	ContentType string
	StatusCode  int
	SSHKey      *shared.SSHKey
}
