package shared

type MembershipDataAttributesRoleEnum string

const (
	MembershipDataAttributesRoleEnumOwner         MembershipDataAttributesRoleEnum = "owner"
	MembershipDataAttributesRoleEnumAdministrator MembershipDataAttributesRoleEnum = "administrator"
	MembershipDataAttributesRoleEnumCollaborator  MembershipDataAttributesRoleEnum = "collaborator"
	MembershipDataAttributesRoleEnumBilling       MembershipDataAttributesRoleEnum = "billing"
)

type MembershipDataAttributes struct {
	Email     *string                           `json:"email,omitempty"`
	FirstName *string                           `json:"first_name,omitempty"`
	LastName  *string                           `json:"last_name,omitempty"`
	Role      *MembershipDataAttributesRoleEnum `json:"role,omitempty"`
}

type MembershipData struct {
	Attributes *MembershipDataAttributes `json:"attributes,omitempty"`
	ID         *string                   `json:"id,omitempty"`
}

type Membership struct {
	Data *MembershipData `json:"data,omitempty"`
}
