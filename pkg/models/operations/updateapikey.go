package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type UpdateAPIKeyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateAPIKeySecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateAPIKeyRequest struct {
	PathParams UpdateAPIKeyPathParams
	Request    *shared.UpdateAPIKey `request:"mediaType=application/json"`
	Security   UpdateAPIKeySecurity
}

type UpdateAPIKey200ApplicationJSON struct {
	Data *shared.APIKey `json:"data,omitempty"`
}

type UpdateAPIKeyResponse struct {
	ContentType                          string
	StatusCode                           int
	ErrorObject                          *shared.ErrorObject
	UpdateAPIKey200ApplicationJSONObject *UpdateAPIKey200ApplicationJSON
}
