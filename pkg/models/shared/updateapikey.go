package shared

type UpdateAPIKeyDataAttributes struct {
	Name *string `json:"name,omitempty"`
}

type UpdateAPIKeyDataTypeEnum string

const (
	UpdateAPIKeyDataTypeEnumAPIKeys UpdateAPIKeyDataTypeEnum = "api_keys"
)

type UpdateAPIKeyData struct {
	Attributes *UpdateAPIKeyDataAttributes `json:"attributes,omitempty"`
	ID         *string                     `json:"id,omitempty"`
	Type       UpdateAPIKeyDataTypeEnum    `json:"type"`
}

type UpdateAPIKey struct {
	Data *UpdateAPIKeyData `json:"data,omitempty"`
}
