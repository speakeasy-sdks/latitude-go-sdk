package shared

type CreateAPIKeyDataAttributes struct {
	Name string `json:"name"`
}

type CreateAPIKeyDataTypeEnum string

const (
	CreateAPIKeyDataTypeEnumAPIKeys CreateAPIKeyDataTypeEnum = "api_keys"
)

type CreateAPIKeyData struct {
	Attributes *CreateAPIKeyDataAttributes `json:"attributes,omitempty"`
	ID         *string                     `json:"id,omitempty"`
	Type       CreateAPIKeyDataTypeEnum    `json:"type"`
}

type CreateAPIKey struct {
	Data *CreateAPIKeyData `json:"data,omitempty"`
}
