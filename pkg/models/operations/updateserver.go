package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type UpdateServerPathParams struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

type UpdateServerRequestBodyAttributes struct {
	Hostname *string `json:"hostname,omitempty"`
}

type UpdateServerRequestBodyTypeEnum string

const (
	UpdateServerRequestBodyTypeEnumServers UpdateServerRequestBodyTypeEnum = "servers"
)

type UpdateServerRequestBody struct {
	Attributes *UpdateServerRequestBodyAttributes `json:"attributes,omitempty"`
	ID         *string                            `json:"id,omitempty"`
	Type       *UpdateServerRequestBodyTypeEnum   `json:"type,omitempty"`
}

type UpdateServerSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateServerRequest struct {
	PathParams UpdateServerPathParams
	Request    *UpdateServerRequestBody `request:"mediaType=application/json"`
	Security   UpdateServerSecurity
}

type UpdateServerResponse struct {
	ContentType string
	StatusCode  int
	Server      *shared.Server
}
