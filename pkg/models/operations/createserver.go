package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type CreateServerRequestBodyDataAttributesOperatingSystemEnum string

const (
	CreateServerRequestBodyDataAttributesOperatingSystemEnumWindowsServer2019StdV1 CreateServerRequestBodyDataAttributesOperatingSystemEnum = "windows_server_2019_std_v1"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumWindowsServer2019DcV1  CreateServerRequestBodyDataAttributesOperatingSystemEnum = "windows_server_2019_dc_v1"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumWindowsServer2016StdV1 CreateServerRequestBodyDataAttributesOperatingSystemEnum = "windows_server_2016_std_v1"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumWindowsServer2016DcV1  CreateServerRequestBodyDataAttributesOperatingSystemEnum = "windows_server_2016_dc_v1"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumUbuntu2204X64Lts       CreateServerRequestBodyDataAttributesOperatingSystemEnum = "ubuntu_22_04_x64_lts"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumDebian11               CreateServerRequestBodyDataAttributesOperatingSystemEnum = "debian_11"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumRockylinux8            CreateServerRequestBodyDataAttributesOperatingSystemEnum = "rockylinux_8"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumDebian10               CreateServerRequestBodyDataAttributesOperatingSystemEnum = "debian_10"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumFlatcarStable          CreateServerRequestBodyDataAttributesOperatingSystemEnum = "flatcar_stable"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumRhel8                  CreateServerRequestBodyDataAttributesOperatingSystemEnum = "rhel8"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumCentos74X64            CreateServerRequestBodyDataAttributesOperatingSystemEnum = "centos_7_4_x64"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumCentos8X64             CreateServerRequestBodyDataAttributesOperatingSystemEnum = "centos_8_x64"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumUbuntu1804X64Lts       CreateServerRequestBodyDataAttributesOperatingSystemEnum = "ubuntu_18_04_x64_lts"
	CreateServerRequestBodyDataAttributesOperatingSystemEnumUbuntu2004X64Lts       CreateServerRequestBodyDataAttributesOperatingSystemEnum = "ubuntu_20_04_x64_lts"
)

type CreateServerRequestBodyDataAttributesPlanEnum string

const (
	CreateServerRequestBodyDataAttributesPlanEnumC1LargeX86  CreateServerRequestBodyDataAttributesPlanEnum = "c1-large-x86"
	CreateServerRequestBodyDataAttributesPlanEnumC1MediumX86 CreateServerRequestBodyDataAttributesPlanEnum = "c1-medium-x86"
	CreateServerRequestBodyDataAttributesPlanEnumC1SmallX86  CreateServerRequestBodyDataAttributesPlanEnum = "c1-small-x86"
	CreateServerRequestBodyDataAttributesPlanEnumC1TinyX86   CreateServerRequestBodyDataAttributesPlanEnum = "c1-tiny-x86"
	CreateServerRequestBodyDataAttributesPlanEnumC2LargeArm  CreateServerRequestBodyDataAttributesPlanEnum = "c2-large-arm"
	CreateServerRequestBodyDataAttributesPlanEnumC2LargeX86  CreateServerRequestBodyDataAttributesPlanEnum = "c2-large-x86"
	CreateServerRequestBodyDataAttributesPlanEnumC2MediumX86 CreateServerRequestBodyDataAttributesPlanEnum = "c2-medium-x86"
	CreateServerRequestBodyDataAttributesPlanEnumC2SmallX86  CreateServerRequestBodyDataAttributesPlanEnum = "c2-small-x86"
	CreateServerRequestBodyDataAttributesPlanEnumT1Spot1     CreateServerRequestBodyDataAttributesPlanEnum = "t1-spot1"
)

type CreateServerRequestBodyDataAttributesRaidEnum string

const (
	CreateServerRequestBodyDataAttributesRaidEnumRaid0 CreateServerRequestBodyDataAttributesRaidEnum = "raid-0"
	CreateServerRequestBodyDataAttributesRaidEnumRaid1 CreateServerRequestBodyDataAttributesRaidEnum = "raid-1"
	CreateServerRequestBodyDataAttributesRaidEnumNull  CreateServerRequestBodyDataAttributesRaidEnum = "null"
)

type CreateServerRequestBodyDataAttributesSiteEnum string

const (
	CreateServerRequestBodyDataAttributesSiteEnumCh1  CreateServerRequestBodyDataAttributesSiteEnum = "CH1"
	CreateServerRequestBodyDataAttributesSiteEnumBgt  CreateServerRequestBodyDataAttributesSiteEnum = "BGT"
	CreateServerRequestBodyDataAttributesSiteEnumBue  CreateServerRequestBodyDataAttributesSiteEnum = "BUE"
	CreateServerRequestBodyDataAttributesSiteEnumAsh  CreateServerRequestBodyDataAttributesSiteEnum = "ASH"
	CreateServerRequestBodyDataAttributesSiteEnumDal2 CreateServerRequestBodyDataAttributesSiteEnum = "DAL2"
	CreateServerRequestBodyDataAttributesSiteEnumLa2  CreateServerRequestBodyDataAttributesSiteEnum = "LA2"
	CreateServerRequestBodyDataAttributesSiteEnumMh1  CreateServerRequestBodyDataAttributesSiteEnum = "MH1"
	CreateServerRequestBodyDataAttributesSiteEnumMi1  CreateServerRequestBodyDataAttributesSiteEnum = "MI1"
	CreateServerRequestBodyDataAttributesSiteEnumNy2  CreateServerRequestBodyDataAttributesSiteEnum = "NY2"
	CreateServerRequestBodyDataAttributesSiteEnumSan  CreateServerRequestBodyDataAttributesSiteEnum = "SAN"
	CreateServerRequestBodyDataAttributesSiteEnumSp1  CreateServerRequestBodyDataAttributesSiteEnum = "SP1"
	CreateServerRequestBodyDataAttributesSiteEnumSp2  CreateServerRequestBodyDataAttributesSiteEnum = "SP2"
	CreateServerRequestBodyDataAttributesSiteEnumSyd  CreateServerRequestBodyDataAttributesSiteEnum = "SYD"
	CreateServerRequestBodyDataAttributesSiteEnumTy6  CreateServerRequestBodyDataAttributesSiteEnum = "TY6"
	CreateServerRequestBodyDataAttributesSiteEnumTy8  CreateServerRequestBodyDataAttributesSiteEnum = "TY8"
)

type CreateServerRequestBodyDataAttributes struct {
	Hostname        *string                                                   `json:"hostname,omitempty"`
	OperatingSystem *CreateServerRequestBodyDataAttributesOperatingSystemEnum `json:"operating_system,omitempty"`
	Plan            *CreateServerRequestBodyDataAttributesPlanEnum            `json:"plan,omitempty"`
	Project         *string                                                   `json:"project,omitempty"`
	Raid            *CreateServerRequestBodyDataAttributesRaidEnum            `json:"raid,omitempty"`
	Site            *CreateServerRequestBodyDataAttributesSiteEnum            `json:"site,omitempty"`
	SSHKeys         []string                                                  `json:"ssh_keys,omitempty"`
	UserData        *int64                                                    `json:"user_data,omitempty"`
}

type CreateServerRequestBodyDataTypeEnum string

const (
	CreateServerRequestBodyDataTypeEnumServers CreateServerRequestBodyDataTypeEnum = "servers"
)

type CreateServerRequestBodyData struct {
	Attributes *CreateServerRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       CreateServerRequestBodyDataTypeEnum    `json:"type"`
}

type CreateServerRequestBody struct {
	Data *CreateServerRequestBodyData `json:"data,omitempty"`
}

type CreateServerSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type CreateServerRequest struct {
	Request  *CreateServerRequestBody `request:"mediaType=application/json"`
	Security CreateServerSecurity
}

type CreateServerResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
	Server      *shared.Server
}
