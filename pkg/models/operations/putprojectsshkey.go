package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PutProjectSSHKeyPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
	SSHKeyID        string `pathParam:"style=simple,explode=false,name=ssh_key_id"`
}

type PutProjectSSHKeyRequestBodyDataAttributes struct {
	Name *string `json:"name,omitempty"`
}

type PutProjectSSHKeyRequestBodyDataTypeEnum string

const (
	PutProjectSSHKeyRequestBodyDataTypeEnumSSHKeys PutProjectSSHKeyRequestBodyDataTypeEnum = "ssh_keys"
)

type PutProjectSSHKeyRequestBodyData struct {
	Attributes *PutProjectSSHKeyRequestBodyDataAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       PutProjectSSHKeyRequestBodyDataTypeEnum    `json:"type"`
}

type PutProjectSSHKeyRequestBody struct {
	Data PutProjectSSHKeyRequestBodyData `json:"data"`
}

type PutProjectSSHKeySecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PutProjectSSHKeyRequest struct {
	PathParams PutProjectSSHKeyPathParams
	Request    *PutProjectSSHKeyRequestBody `request:"mediaType=application/json"`
	Security   PutProjectSSHKeySecurity
}

type PutProjectSSHKey200ApplicationJSON struct {
	Data *shared.SSHKeyData `json:"data,omitempty"`
}

type PutProjectSSHKeyResponse struct {
	ContentType                              string
	StatusCode                               int
	PutProjectSSHKey200ApplicationJSONObject *PutProjectSSHKey200ApplicationJSON
}
