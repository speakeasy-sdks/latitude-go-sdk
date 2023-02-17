package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PatchUserProfilePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type PatchUserProfileRequestBodyDataAttributesRoleEnum string

const (
	PatchUserProfileRequestBodyDataAttributesRoleEnumAdministrator PatchUserProfileRequestBodyDataAttributesRoleEnum = "administrator"
	PatchUserProfileRequestBodyDataAttributesRoleEnumBilling       PatchUserProfileRequestBodyDataAttributesRoleEnum = "billing"
	PatchUserProfileRequestBodyDataAttributesRoleEnumCollaborator  PatchUserProfileRequestBodyDataAttributesRoleEnum = "collaborator"
	PatchUserProfileRequestBodyDataAttributesRoleEnumOwner         PatchUserProfileRequestBodyDataAttributesRoleEnum = "owner"
)

type PatchUserProfileRequestBodyDataAttributes struct {
	AuthenticationFactorID *string                                            `json:"authentication_factor_id,omitempty"`
	FirstName              *string                                            `json:"first_name,omitempty"`
	LastName               *string                                            `json:"last_name,omitempty"`
	Role                   *PatchUserProfileRequestBodyDataAttributesRoleEnum `json:"role,omitempty"`
}

type PatchUserProfileRequestBodyDataTypeEnum string

const (
	PatchUserProfileRequestBodyDataTypeEnumUsers PatchUserProfileRequestBodyDataTypeEnum = "users"
)

type PatchUserProfileRequestBodyData struct {
	Attributes *PatchUserProfileRequestBodyDataAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       PatchUserProfileRequestBodyDataTypeEnum    `json:"type"`
}

type PatchUserProfileRequestBody struct {
	Data PatchUserProfileRequestBodyData `json:"data"`
}

type PatchUserProfileSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PatchUserProfileRequest struct {
	PathParams PatchUserProfilePathParams
	Request    *PatchUserProfileRequestBody `request:"mediaType=application/json"`
	Security   PatchUserProfileSecurity
}

type PatchUserProfile200ApplicationJSON struct {
	Data *shared.UserUpdate `json:"data,omitempty"`
}

type PatchUserProfileResponse struct {
	ContentType                              string
	StatusCode                               int
	ErrorObject                              *shared.ErrorObject
	PatchUserProfile200ApplicationJSONObject *PatchUserProfile200ApplicationJSON
}
