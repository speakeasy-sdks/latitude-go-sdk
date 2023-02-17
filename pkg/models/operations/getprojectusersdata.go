package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetProjectUsersDataPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
}

type GetProjectUsersDataQueryParams struct {
	ExtraFieldsUserData *string `queryParam:"style=form,explode=true,name=extra_fields[user_data]"`
	Include             *string `queryParam:"style=form,explode=true,name=include"`
}

type GetProjectUsersDataSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetProjectUsersDataRequest struct {
	PathParams  GetProjectUsersDataPathParams
	QueryParams GetProjectUsersDataQueryParams
	Security    GetProjectUsersDataSecurity
}

type GetProjectUsersData200ApplicationJSON struct {
	Data []shared.UserData `json:"data,omitempty"`
}

type GetProjectUsersDataResponse struct {
	ContentType                                 string
	StatusCode                                  int
	GetProjectUsersData200ApplicationJSONObject *GetProjectUsersData200ApplicationJSON
}
