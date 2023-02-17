package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type DeleteVpnSessionPathParams struct {
	VpnSessionID string `pathParam:"style=simple,explode=false,name=vpn_session_id"`
}

type DeleteVpnSessionSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteVpnSessionRequest struct {
	PathParams DeleteVpnSessionPathParams
	Security   DeleteVpnSessionSecurity
}

type DeleteVpnSessionResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
}
