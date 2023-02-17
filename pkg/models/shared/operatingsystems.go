package shared

type OperatingSystemsAttributesFeatures struct {
	Raid     *bool `json:"raid,omitempty"`
	SSHKeys  *bool `json:"ssh_keys,omitempty"`
	UserData *bool `json:"user_data,omitempty"`
}

type OperatingSystemsAttributes struct {
	Distro   *string                             `json:"distro,omitempty"`
	Features *OperatingSystemsAttributesFeatures `json:"features,omitempty"`
	Name     *string                             `json:"name,omitempty"`
	Slug     *string                             `json:"slug,omitempty"`
	User     *string                             `json:"user,omitempty"`
	Version  *string                             `json:"version,omitempty"`
}

type OperatingSystems struct {
	Attributes *OperatingSystemsAttributes `json:"attributes,omitempty"`
	ID         *string                     `json:"id,omitempty"`
}
