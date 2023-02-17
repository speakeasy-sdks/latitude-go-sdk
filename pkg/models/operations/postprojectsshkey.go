package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PostProjectSSHKeyPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
}

type PostProjectSSHKeyRequestBodyDataAttributes struct {
	Name      *string `json:"name,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
}

type PostProjectSSHKeyRequestBodyDataTypeEnum string

const (
	PostProjectSSHKeyRequestBodyDataTypeEnumSSHKeys PostProjectSSHKeyRequestBodyDataTypeEnum = "ssh_keys"
)

type PostProjectSSHKeyRequestBodyData struct {
	Attributes *PostProjectSSHKeyRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       PostProjectSSHKeyRequestBodyDataTypeEnum    `json:"type"`
}

type PostProjectSSHKeyRequestBody struct {
	Data PostProjectSSHKeyRequestBodyData `json:"data"`
}

type PostProjectSSHKeySecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PostProjectSSHKeyRequest struct {
	PathParams PostProjectSSHKeyPathParams
	Request    *PostProjectSSHKeyRequestBody `request:"mediaType=application/json"`
	Security   PostProjectSSHKeySecurity
}

type PostProjectSSHKey201ApplicationJSON struct {
	Data *shared.SSHKeyData `json:"data,omitempty"`
}

type PostProjectSSHKeyResponse struct {
	ContentType                               string
	StatusCode                                int
	PostProjectSSHKey201ApplicationJSONObject *PostProjectSSHKey201ApplicationJSON
}
