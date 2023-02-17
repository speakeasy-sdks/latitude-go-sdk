package shared

type IPAddressAttributesFamilyEnum string

const (
	IPAddressAttributesFamilyEnumIPv4 IPAddressAttributesFamilyEnum = "IPv4"
	IPAddressAttributesFamilyEnumIPv6 IPAddressAttributesFamilyEnum = "IPv6"
)

type IPAddressAttributesRegionSite struct {
	Facility *string `json:"facility,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Slug     *string `json:"slug,omitempty"`
}

// IPAddressAttributesRegion
// NOTE: `region` is provided as an extra attribute that is lazy loaded. To request it, just add a `?extra_fields[ip_addresses]=region` in the query string
type IPAddressAttributesRegion struct {
	City    *string                        `json:"city,omitempty"`
	Country *string                        `json:"country,omitempty"`
	Site    *IPAddressAttributesRegionSite `json:"site,omitempty"`
}

type IPAddressAttributesRole struct {
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

type IPAddressAttributesTypeEnum string

const (
	IPAddressAttributesTypeEnumPublic  IPAddressAttributesTypeEnum = "Public"
	IPAddressAttributesTypeEnumPrivate IPAddressAttributesTypeEnum = "Private"
)

type IPAddressAttributes struct {
	Address *string                        `json:"address,omitempty"`
	Family  *IPAddressAttributesFamilyEnum `json:"family,omitempty"`
	Project *ProjectInclude                `json:"project,omitempty"`
	Region  *IPAddressAttributesRegion     `json:"region,omitempty"`
	Role    *IPAddressAttributesRole       `json:"role,omitempty"`
	Type    *IPAddressAttributesTypeEnum   `json:"type,omitempty"`
}

type IPAddressRelationshipsProjectMeta struct {
	Included *bool `json:"included,omitempty"`
}

type IPAddressRelationshipsProject struct {
	Meta *IPAddressRelationshipsProjectMeta `json:"meta,omitempty"`
}

type IPAddressRelationships struct {
	Project *IPAddressRelationshipsProject `json:"project,omitempty"`
}

type IPAddress struct {
	Attributes    *IPAddressAttributes    `json:"attributes,omitempty"`
	ID            *string                 `json:"id,omitempty"`
	Relationships *IPAddressRelationships `json:"relationships,omitempty"`
}
