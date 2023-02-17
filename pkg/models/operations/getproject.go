package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetProjectPathParams struct {
	IDOrSlug string `pathParam:"style=simple,explode=false,name=id_or_slug"`
}

type GetProjectQueryParams struct {
	ExtraFieldsProjects *string `queryParam:"style=form,explode=true,name=extra_fields[projects]"`
}

type GetProjectSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetProjectRequest struct {
	PathParams  GetProjectPathParams
	QueryParams GetProjectQueryParams
	Security    GetProjectSecurity
}

type GetProject200ApplicationJSON struct {
	Data *shared.Project `json:"data,omitempty"`
}

type GetProjectResponse struct {
	ContentType                        string
	StatusCode                         int
	GetProject200ApplicationJSONObject *GetProject200ApplicationJSON
}
