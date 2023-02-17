package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetProjectUserDataPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
	UserDataID      string `pathParam:"style=simple,explode=false,name=user_data_id"`
}

type GetProjectUserDataQueryParams struct {
	ExtraFieldsUserData *string `queryParam:"style=form,explode=true,name=extra_fields[user_data]"`
	Include             *string `queryParam:"style=form,explode=true,name=include"`
}

type GetProjectUserDataSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetProjectUserDataRequest struct {
	PathParams  GetProjectUserDataPathParams
	QueryParams GetProjectUserDataQueryParams
	Security    GetProjectUserDataSecurity
}

type GetProjectUserDataResponse struct {
	ContentType string
	StatusCode  int
	UserData    *shared.UserData
}
