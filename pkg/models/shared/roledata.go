package shared

type RoleDataAttributes struct {
	Name *string `json:"name,omitempty"`
}

type RoleDataTypeEnum string

const (
	RoleDataTypeEnumRoles RoleDataTypeEnum = "roles"
)

type RoleData struct {
	Attributes *RoleDataAttributes `json:"attributes,omitempty"`
	ID         *string             `json:"id,omitempty"`
	Type       *RoleDataTypeEnum   `json:"type,omitempty"`
}
