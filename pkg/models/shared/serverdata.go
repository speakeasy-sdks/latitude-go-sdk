package shared

type ServerDataAttributesIpmiStatusEnum string

const (
	ServerDataAttributesIpmiStatusEnumUnavailable  ServerDataAttributesIpmiStatusEnum = "Unavailable"
	ServerDataAttributesIpmiStatusEnumIntermittent ServerDataAttributesIpmiStatusEnum = "Intermittent"
	ServerDataAttributesIpmiStatusEnumNormal       ServerDataAttributesIpmiStatusEnum = "Normal"
)

type ServerDataAttributesOperatingSystemDistro struct {
	Name   *string `json:"name,omitempty"`
	Series *string `json:"series,omitempty"`
	Slug   *string `json:"slug,omitempty"`
}

type ServerDataAttributesOperatingSystemFeatures struct {
	Raid     *bool `json:"raid,omitempty"`
	SSHKeys  *bool `json:"ssh_keys,omitempty"`
	UserData *bool `json:"user_data,omitempty"`
}

type ServerDataAttributesOperatingSystem struct {
	Distro   *ServerDataAttributesOperatingSystemDistro   `json:"distro,omitempty"`
	Features *ServerDataAttributesOperatingSystemFeatures `json:"features,omitempty"`
	Name     *string                                      `json:"name,omitempty"`
	Slug     *string                                      `json:"slug,omitempty"`
	Version  *string                                      `json:"version,omitempty"`
}

type ServerDataAttributesPlan struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ServerDataAttributesSpecs struct {
	CPU  *string `json:"cpu,omitempty"`
	Disk *string `json:"disk,omitempty"`
	RAM  *string `json:"ram,omitempty"`
}

type ServerDataAttributesStatusEnum string

const (
	ServerDataAttributesStatusEnumOn                ServerDataAttributesStatusEnum = "on"
	ServerDataAttributesStatusEnumOff               ServerDataAttributesStatusEnum = "off"
	ServerDataAttributesStatusEnumUnknown           ServerDataAttributesStatusEnum = "unknown"
	ServerDataAttributesStatusEnumReady             ServerDataAttributesStatusEnum = "ready"
	ServerDataAttributesStatusEnumDiskErasing       ServerDataAttributesStatusEnum = "disk_erasing"
	ServerDataAttributesStatusEnumFailedDiskErasing ServerDataAttributesStatusEnum = "failed_disk_erasing"
	ServerDataAttributesStatusEnumDeploying         ServerDataAttributesStatusEnum = "deploying"
	ServerDataAttributesStatusEnumFailedDeployment  ServerDataAttributesStatusEnum = "failed_deployment"
)

type ServerDataAttributes struct {
	CreatedAt       *string                              `json:"created_at,omitempty"`
	Hostname        *string                              `json:"hostname,omitempty"`
	IpmiStatus      *ServerDataAttributesIpmiStatusEnum  `json:"ipmi_status,omitempty"`
	OperatingSystem *ServerDataAttributesOperatingSystem `json:"operating_system,omitempty"`
	Plan            *ServerDataAttributesPlan            `json:"plan,omitempty"`
	PrimaryIpv4     *string                              `json:"primary_ipv4,omitempty"`
	Project         *ProjectInclude                      `json:"project,omitempty"`
	Region          *RegionResourceData                  `json:"region,omitempty"`
	Role            *string                              `json:"role,omitempty"`
	Site            *string                              `json:"site,omitempty"`
	Specs           *ServerDataAttributesSpecs           `json:"specs,omitempty"`
	Status          *ServerDataAttributesStatusEnum      `json:"status,omitempty"`
	Team            *TeamInclude                         `json:"team,omitempty"`
}

type ServerData struct {
	Attributes *ServerDataAttributes `json:"attributes,omitempty"`
	ID         *string               `json:"id,omitempty"`
	Type       *string               `json:"type,omitempty"`
}
