package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type DeleteProjectSSHKeyPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
	SSHKeyID        string `pathParam:"style=simple,explode=false,name=ssh_key_id"`
}

type DeleteProjectSSHKeySecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteProjectSSHKeyRequest struct {
	PathParams DeleteProjectSSHKeyPathParams
	Security   DeleteProjectSSHKeySecurity
}

type DeleteProjectSSHKeyResponse struct {
	ContentType string
	StatusCode  int
}
