package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PutProjectUserDataPathParams struct {
	ProjectIDOrSlug string `pathParam:"style=simple,explode=false,name=project_id_or_slug"`
	UserDataID      string `pathParam:"style=simple,explode=false,name=user_data_id"`
}

type PutProjectUserDataRequestBodyDataAttributes struct {
	Content     *string `json:"content,omitempty"`
	Description *string `json:"description,omitempty"`
}

type PutProjectUserDataRequestBodyDataTypeEnum string

const (
	PutProjectUserDataRequestBodyDataTypeEnumUserData PutProjectUserDataRequestBodyDataTypeEnum = "user_data"
)

type PutProjectUserDataRequestBodyData struct {
	Attributes *PutProjectUserDataRequestBodyDataAttributes `json:"attributes,omitempty"`
	ID         string                                       `json:"id"`
	Type       PutProjectUserDataRequestBodyDataTypeEnum    `json:"type"`
}

type PutProjectUserDataRequestBody struct {
	Data PutProjectUserDataRequestBodyData `json:"data"`
}

type PutProjectUserDataSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PutProjectUserDataRequest struct {
	PathParams PutProjectUserDataPathParams
	Request    *PutProjectUserDataRequestBody `request:"mediaType=application/json"`
	Security   PutProjectUserDataSecurity
}

type PutProjectUserDataResponse struct {
	ContentType string
	StatusCode  int
	UserData    *shared.UserData
}
