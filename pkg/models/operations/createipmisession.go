package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type CreateIpmiSessionPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type CreateIpmiSessionSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type CreateIpmiSessionRequest struct {
	PathParams CreateIpmiSessionPathParams
	Security   CreateIpmiSessionSecurity
}

type CreateIpmiSessionResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
	IpmiSession *shared.IpmiSession
}
