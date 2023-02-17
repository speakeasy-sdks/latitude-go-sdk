package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type CreateServerActionPathParams struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

type CreateServerActionRequestBodyDataAttributesActionEnum string

const (
	CreateServerActionRequestBodyDataAttributesActionEnumPowerOn  CreateServerActionRequestBodyDataAttributesActionEnum = "power_on"
	CreateServerActionRequestBodyDataAttributesActionEnumPowerOff CreateServerActionRequestBodyDataAttributesActionEnum = "power_off"
	CreateServerActionRequestBodyDataAttributesActionEnumReboot   CreateServerActionRequestBodyDataAttributesActionEnum = "reboot"
)

type CreateServerActionRequestBodyDataAttributes struct {
	Action *CreateServerActionRequestBodyDataAttributesActionEnum `json:"action,omitempty"`
}

type CreateServerActionRequestBodyDataTypeEnum string

const (
	CreateServerActionRequestBodyDataTypeEnumActions CreateServerActionRequestBodyDataTypeEnum = "actions"
)

type CreateServerActionRequestBodyData struct {
	Attributes *CreateServerActionRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       CreateServerActionRequestBodyDataTypeEnum    `json:"type"`
}

type CreateServerActionRequestBody struct {
	Data CreateServerActionRequestBodyData `json:"data"`
}

type CreateServerActionSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type CreateServerActionRequest struct {
	PathParams CreateServerActionPathParams
	Request    *CreateServerActionRequestBody `request:"mediaType=application/json"`
	Security   CreateServerActionSecurity
}

type CreateServerActionResponse struct {
	ContentType  string
	StatusCode   int
	ErrorObject  *shared.ErrorObject
	ServerAction *shared.ServerAction
}
