package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type DeleteProjectPathParams struct {
	IDOrSlug string `pathParam:"style=simple,explode=false,name=id_or_slug"`
}

type DeleteProjectSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteProjectRequest struct {
	PathParams DeleteProjectPathParams
	Security   DeleteProjectSecurity
}

type DeleteProject200ApplicationJSON struct {
	Data *shared.Project `json:"data,omitempty"`
}

type DeleteProjectResponse struct {
	ContentType                           string
	StatusCode                            int
	DeleteProject200ApplicationJSONObject *DeleteProject200ApplicationJSON
	ErrorObject                           *shared.ErrorObject
}
