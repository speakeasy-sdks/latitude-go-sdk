package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PostVpnSessionRequestBody struct {
	Data *interface{} `json:"data,omitempty"`
}

type PostVpnSessionSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PostVpnSessionRequest struct {
	Request  *PostVpnSessionRequestBody `request:"mediaType=application/json"`
	Security PostVpnSessionSecurity
}

type PostVpnSessionResponse struct {
	ContentType            string
	StatusCode             int
	VpnSessionWithPassword *shared.VpnSessionWithPassword
}
