package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type DeleteProjectUserDataPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
	UserDataID      string `pathParam:"style=simple,explode=false,name=user_data_id"`
}

type DeleteProjectUserDataSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteProjectUserDataRequest struct {
	PathParams DeleteProjectUserDataPathParams
	Security   DeleteProjectUserDataSecurity
}

type DeleteProjectUserDataResponse struct {
	ContentType string
	StatusCode  int
}
