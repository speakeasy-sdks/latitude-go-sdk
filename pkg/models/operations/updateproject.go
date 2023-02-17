package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type UpdateProjectPathParams struct {
	IDOrSlug string `pathParam:"style=simple,explode=false,name=id_or_slug"`
}

type UpdateProjectRequestBodyDataAttributesEnvironmentEnum string

const (
	UpdateProjectRequestBodyDataAttributesEnvironmentEnumDevelopment UpdateProjectRequestBodyDataAttributesEnvironmentEnum = "Development"
	UpdateProjectRequestBodyDataAttributesEnvironmentEnumStaging     UpdateProjectRequestBodyDataAttributesEnvironmentEnum = "Staging"
	UpdateProjectRequestBodyDataAttributesEnvironmentEnumProduction  UpdateProjectRequestBodyDataAttributesEnvironmentEnum = "Production"
)

type UpdateProjectRequestBodyDataAttributes struct {
	BandwidthAlert *bool                                                  `json:"bandwidth_alert,omitempty"`
	Description    *string                                                `json:"description,omitempty"`
	Environment    *UpdateProjectRequestBodyDataAttributesEnvironmentEnum `json:"environment,omitempty"`
	Name           *string                                                `json:"name,omitempty"`
}

type UpdateProjectRequestBodyDataTypeEnum string

const (
	UpdateProjectRequestBodyDataTypeEnumProjects UpdateProjectRequestBodyDataTypeEnum = "projects"
)

type UpdateProjectRequestBodyData struct {
	Attributes *UpdateProjectRequestBodyDataAttributes `json:"attributes,omitempty"`
	ID         string                                  `json:"id"`
	Type       UpdateProjectRequestBodyDataTypeEnum    `json:"type"`
}

type UpdateProjectRequestBody struct {
	Data UpdateProjectRequestBodyData `json:"data"`
}

type UpdateProjectSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateProjectRequest struct {
	PathParams UpdateProjectPathParams
	Request    *UpdateProjectRequestBody `request:"mediaType=application/json"`
	Security   UpdateProjectSecurity
}

type UpdateProject200ApplicationJSON struct {
	Data *shared.Project `json:"data,omitempty"`
}

type UpdateProjectResponse struct {
	ContentType                           string
	StatusCode                            int
	ErrorObject                           *shared.ErrorObject
	UpdateProject200ApplicationJSONObject *UpdateProject200ApplicationJSON
}
