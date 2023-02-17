package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type UpdateCurrentVersionRequestBodyDataAttributes struct {
	NewVersion *string `json:"new_version,omitempty"`
}

type UpdateCurrentVersionRequestBodyDataTypeEnum string

const (
	UpdateCurrentVersionRequestBodyDataTypeEnumAPIVersion UpdateCurrentVersionRequestBodyDataTypeEnum = "api_version"
)

type UpdateCurrentVersionRequestBodyData struct {
	Attributes *UpdateCurrentVersionRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       *UpdateCurrentVersionRequestBodyDataTypeEnum   `json:"type,omitempty"`
}

type UpdateCurrentVersionRequestBody struct {
	Data *UpdateCurrentVersionRequestBodyData `json:"data,omitempty"`
}

type UpdateCurrentVersionSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateCurrentVersionRequest struct {
	Request  *UpdateCurrentVersionRequestBody `request:"mediaType=application/json"`
	Security UpdateCurrentVersionSecurity
}

type UpdateCurrentVersionResponse struct {
	ContentType string
	StatusCode  int
}
