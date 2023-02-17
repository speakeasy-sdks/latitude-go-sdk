package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetUserProfileSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetUserProfileRequest struct {
	Security GetUserProfileSecurity
}

type GetUserProfile200ApplicationJSON struct {
	Data *shared.User `json:"data,omitempty"`
}

type GetUserProfileResponse struct {
	ContentType                            string
	StatusCode                             int
	GetUserProfile200ApplicationJSONObject *GetUserProfile200ApplicationJSON
}
