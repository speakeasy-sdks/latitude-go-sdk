package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetServerPathParams struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

type GetServerQueryParams struct {
	ExtraFieldsServers *string `queryParam:"style=form,explode=true,name=extra_fields[servers]"`
}

type GetServerSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetServerRequest struct {
	PathParams  GetServerPathParams
	QueryParams GetServerQueryParams
	Security    GetServerSecurity
}

type GetServerResponse struct {
	ContentType string
	StatusCode  int
	Server      *shared.Server
}
