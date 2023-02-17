package shared

type UserDataPropertiesAttributes struct {
	Content     *string         `json:"content,omitempty"`
	CreatedAt   *string         `json:"created_at,omitempty"`
	Description *string         `json:"description,omitempty"`
	Projects    *ProjectInclude `json:"projects,omitempty"`
	UpdatedAt   *string         `json:"updated_at,omitempty"`
	User        *UserInclude    `json:"user,omitempty"`
}

type UserDataPropertiesRelationships struct {
	Projects map[string]interface{} `json:"projects,omitempty"`
	User     map[string]interface{} `json:"user,omitempty"`
}

type UserDataPropertiesTypeEnum string

const (
	UserDataPropertiesTypeEnumUserData UserDataPropertiesTypeEnum = "user_data"
)

type UserDataProperties struct {
	Attributes    *UserDataPropertiesAttributes    `json:"attributes,omitempty"`
	ID            *string                          `json:"id,omitempty"`
	Relationships *UserDataPropertiesRelationships `json:"relationships,omitempty"`
	Type          UserDataPropertiesTypeEnum       `json:"type"`
}
