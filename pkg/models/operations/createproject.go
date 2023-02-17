package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type CreateProjectRequestBodyDataAttributesEnvironmentEnum string

const (
	CreateProjectRequestBodyDataAttributesEnvironmentEnumDevelopment CreateProjectRequestBodyDataAttributesEnvironmentEnum = "Development"
	CreateProjectRequestBodyDataAttributesEnvironmentEnumStaging     CreateProjectRequestBodyDataAttributesEnvironmentEnum = "Staging"
	CreateProjectRequestBodyDataAttributesEnvironmentEnumProduction  CreateProjectRequestBodyDataAttributesEnvironmentEnum = "Production"
)

type CreateProjectRequestBodyDataAttributes struct {
	Description *string                                                `json:"description,omitempty"`
	Environment *CreateProjectRequestBodyDataAttributesEnvironmentEnum `json:"environment,omitempty"`
	Name        string                                                 `json:"name"`
}

type CreateProjectRequestBodyDataTypeEnum string

const (
	CreateProjectRequestBodyDataTypeEnumProjects CreateProjectRequestBodyDataTypeEnum = "projects"
)

type CreateProjectRequestBodyData struct {
	Attributes *CreateProjectRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       CreateProjectRequestBodyDataTypeEnum    `json:"type"`
}

type CreateProjectRequestBody struct {
	Data *CreateProjectRequestBodyData `json:"data,omitempty"`
}

type CreateProjectSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type CreateProjectRequest struct {
	Request  *CreateProjectRequestBody `request:"mediaType=application/json"`
	Security CreateProjectSecurity
}

type CreateProject201ApplicationJSON struct {
	Data *shared.Project `json:"data,omitempty"`
}

type CreateProjectResponse struct {
	ContentType                           string
	StatusCode                            int
	CreateProject201ApplicationJSONObject *CreateProject201ApplicationJSON
	ErrorObject                           *shared.ErrorObject
}
