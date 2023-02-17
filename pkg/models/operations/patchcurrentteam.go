package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PatchCurrentTeamPathParams struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
}

type PatchCurrentTeamRequestBodyDataAttributes struct {
	Address     *string `json:"address,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type PatchCurrentTeamRequestBodyDataTypeEnum string

const (
	PatchCurrentTeamRequestBodyDataTypeEnumTeams PatchCurrentTeamRequestBodyDataTypeEnum = "teams"
)

type PatchCurrentTeamRequestBodyData struct {
	Attributes *PatchCurrentTeamRequestBodyDataAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       PatchCurrentTeamRequestBodyDataTypeEnum    `json:"type"`
}

type PatchCurrentTeamRequestBody struct {
	Data PatchCurrentTeamRequestBodyData `json:"data"`
}

type PatchCurrentTeamSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PatchCurrentTeamRequest struct {
	PathParams PatchCurrentTeamPathParams
	Request    *PatchCurrentTeamRequestBody `request:"mediaType=application/json"`
	Security   PatchCurrentTeamSecurity
}

type PatchCurrentTeam200ApplicationJSON struct {
	Data *shared.Team `json:"data,omitempty"`
}

type PatchCurrentTeamResponse struct {
	ContentType                              string
	StatusCode                               int
	ErrorObject                              *shared.ErrorObject
	PatchCurrentTeam200ApplicationJSONObject *PatchCurrentTeam200ApplicationJSON
}
