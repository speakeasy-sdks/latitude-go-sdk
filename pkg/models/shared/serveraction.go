package shared

type ServerActionDataAttributes struct {
	Status *string `json:"status,omitempty"`
}

type ServerActionDataTypeEnum string

const (
	ServerActionDataTypeEnumActions ServerActionDataTypeEnum = "actions"
)

type ServerActionData struct {
	Attributes *ServerActionDataAttributes `json:"attributes,omitempty"`
	ID         *string                     `json:"id,omitempty"`
	Type       *ServerActionDataTypeEnum   `json:"type,omitempty"`
}

type ServerAction struct {
	Data *ServerActionData      `json:"data,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}
