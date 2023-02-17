package shared

type SSHKeyDataAttributes struct {
	CreatedAt   *string      `json:"created_at,omitempty"`
	Fingerprint *string      `json:"fingerprint,omitempty"`
	Name        *string      `json:"name,omitempty"`
	PublicKey   *string      `json:"public_key,omitempty"`
	UpdatedAt   *string      `json:"updated_at,omitempty"`
	User        *UserInclude `json:"user,omitempty"`
}

type SSHKeyDataRelationships struct {
	Projects map[string]interface{} `json:"projects,omitempty"`
	User     map[string]interface{} `json:"user,omitempty"`
}

type SSHKeyDataTypeEnum string

const (
	SSHKeyDataTypeEnumSSHKeys SSHKeyDataTypeEnum = "ssh_keys"
)

type SSHKeyData struct {
	Attributes    *SSHKeyDataAttributes    `json:"attributes,omitempty"`
	ID            *string                  `json:"id,omitempty"`
	Relationships *SSHKeyDataRelationships `json:"relationships,omitempty"`
	Type          SSHKeyDataTypeEnum       `json:"type"`
}
