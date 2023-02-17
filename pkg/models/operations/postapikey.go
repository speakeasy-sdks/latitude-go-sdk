package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PostAPIKeySecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PostAPIKeyRequest struct {
	Request  *shared.CreateAPIKey `request:"mediaType=application/json"`
	Security PostAPIKeySecurity
}

type PostAPIKey201ApplicationJSON struct {
	Data *shared.APIKey `json:"data,omitempty"`
}

type PostAPIKeyResponse struct {
	ContentType                        string
	StatusCode                         int
	ErrorObject                        *shared.ErrorObject
	PostAPIKey201ApplicationJSONObject *PostAPIKey201ApplicationJSON
}
