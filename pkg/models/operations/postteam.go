package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PostTeamRequestBodyDataAttributesCurrencyEnum string

const (
	PostTeamRequestBodyDataAttributesCurrencyEnumUsd PostTeamRequestBodyDataAttributesCurrencyEnum = "USD"
	PostTeamRequestBodyDataAttributesCurrencyEnumBrl PostTeamRequestBodyDataAttributesCurrencyEnum = "BRL"
)

type PostTeamRequestBodyDataAttributes struct {
	Address     *string                                       `json:"address,omitempty"`
	Currency    PostTeamRequestBodyDataAttributesCurrencyEnum `json:"currency"`
	Description *string                                       `json:"description,omitempty"`
	Name        string                                        `json:"name"`
}

type PostTeamRequestBodyDataTypeEnum string

const (
	PostTeamRequestBodyDataTypeEnumTeams PostTeamRequestBodyDataTypeEnum = "teams"
)

type PostTeamRequestBodyData struct {
	Attributes *PostTeamRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       PostTeamRequestBodyDataTypeEnum    `json:"type"`
}

type PostTeamRequestBody struct {
	Data PostTeamRequestBodyData `json:"data"`
}

type PostTeamSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PostTeamRequest struct {
	Request  *PostTeamRequestBody `request:"mediaType=application/json"`
	Security PostTeamSecurity
}

type PostTeam201ApplicationJSON struct {
	Data *shared.Team `json:"data,omitempty"`
}

type PostTeamResponse struct {
	ContentType                      string
	StatusCode                       int
	ErrorObject                      *shared.ErrorObject
	PostTeam201ApplicationJSONObject *PostTeam201ApplicationJSON
}
