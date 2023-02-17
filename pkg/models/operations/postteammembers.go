package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type PostTeamMembersRequestBodyDataAttributesRoleEnum string

const (
	PostTeamMembersRequestBodyDataAttributesRoleEnumOwner         PostTeamMembersRequestBodyDataAttributesRoleEnum = "owner"
	PostTeamMembersRequestBodyDataAttributesRoleEnumAdministrator PostTeamMembersRequestBodyDataAttributesRoleEnum = "administrator"
	PostTeamMembersRequestBodyDataAttributesRoleEnumCollaborator  PostTeamMembersRequestBodyDataAttributesRoleEnum = "collaborator"
	PostTeamMembersRequestBodyDataAttributesRoleEnumBilling       PostTeamMembersRequestBodyDataAttributesRoleEnum = "billing"
)

type PostTeamMembersRequestBodyDataAttributes struct {
	Email     string                                           `json:"email"`
	FirstName *string                                          `json:"first_name,omitempty"`
	LastName  *string                                          `json:"last_name,omitempty"`
	Role      PostTeamMembersRequestBodyDataAttributesRoleEnum `json:"role"`
}

type PostTeamMembersRequestBodyDataTypeEnum string

const (
	PostTeamMembersRequestBodyDataTypeEnumMemberships PostTeamMembersRequestBodyDataTypeEnum = "memberships"
)

type PostTeamMembersRequestBodyData struct {
	Attributes *PostTeamMembersRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       PostTeamMembersRequestBodyDataTypeEnum    `json:"type"`
}

type PostTeamMembersRequestBody struct {
	Data PostTeamMembersRequestBodyData `json:"data"`
}

type PostTeamMembersSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type PostTeamMembersRequest struct {
	Request  *PostTeamMembersRequestBody `request:"mediaType=application/json"`
	Security PostTeamMembersSecurity
}

type PostTeamMembersResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
	Membership  *shared.Membership
}
