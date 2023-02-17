package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PutVpnSessionPathParams struct {
	VpnSessionID string `pathParam:"style=simple,explode=false,name=vpn_session_id"`
}

type PutVpnSessionSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PutVpnSessionRequest struct {
	PathParams PutVpnSessionPathParams
	Security   PutVpnSessionSecurity
}

type PutVpnSessionResponse struct {
	ContentType            string
	StatusCode             int
	ErrorObject            *shared.ErrorObject
	VpnSessionWithPassword *shared.VpnSessionWithPassword
}
