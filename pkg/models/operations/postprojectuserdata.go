package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PostProjectUserDataPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
}

type PostProjectUserDataRequestBodyDataAttributes struct {
	Content     string `json:"content"`
	Description string `json:"description"`
}

type PostProjectUserDataRequestBodyDataTypeEnum string

const (
	PostProjectUserDataRequestBodyDataTypeEnumUserData PostProjectUserDataRequestBodyDataTypeEnum = "user_data"
)

type PostProjectUserDataRequestBodyData struct {
	Attributes *PostProjectUserDataRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       PostProjectUserDataRequestBodyDataTypeEnum    `json:"type"`
}

type PostProjectUserDataRequestBody struct {
	Data PostProjectUserDataRequestBodyData `json:"data"`
}

type PostProjectUserDataSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PostProjectUserDataRequest struct {
	PathParams PostProjectUserDataPathParams
	Request    *PostProjectUserDataRequestBody `request:"mediaType=application/json"`
	Security   PostProjectUserDataSecurity
}

type PostProjectUserDataResponse struct {
	ContentType string
	StatusCode  int
	UserData    *shared.UserData
}
