package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type UpdateServerDeployConfigPathParams struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

type UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum string

const (
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumWindowsServer2019StdV1 UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "windows_server_2019_std_v1"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumWindowsServer2019DcV1  UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "windows_server_2019_dc_v1"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumWindowsServer2016StdV1 UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "windows_server_2016_std_v1"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumWindowsServer2016DcV1  UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "windows_server_2016_dc_v1"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumUbuntu2204X64Lts       UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "ubuntu_22_04_x64_lts"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumDebian11               UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "debian_11"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumRockylinux8            UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "rockylinux_8"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumDebian10               UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "debian_10"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumFlatcarStable          UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "flatcar_stable"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumRhel8                  UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "rhel8"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumCentos74X64            UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "centos_7_4_x64"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumCentos8X64             UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "centos_8_x64"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumUbuntu1804X64Lts       UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "ubuntu_18_04_x64_lts"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumUbuntu2004X64Lts       UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "ubuntu_20_04_x64_lts"
	UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnumNull                   UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum = "null"
)

type UpdateServerDeployConfigRequestBodyAttributesRaidEnum string

const (
	UpdateServerDeployConfigRequestBodyAttributesRaidEnumRaid0 UpdateServerDeployConfigRequestBodyAttributesRaidEnum = "raid-0"
	UpdateServerDeployConfigRequestBodyAttributesRaidEnumRaid1 UpdateServerDeployConfigRequestBodyAttributesRaidEnum = "raid-1"
	UpdateServerDeployConfigRequestBodyAttributesRaidEnumNull  UpdateServerDeployConfigRequestBodyAttributesRaidEnum = "null"
)

type UpdateServerDeployConfigRequestBodyAttributes struct {
	Hostname        *string                                                           `json:"hostname,omitempty"`
	OperatingSystem *UpdateServerDeployConfigRequestBodyAttributesOperatingSystemEnum `json:"operating_system,omitempty"`
	Raid            *UpdateServerDeployConfigRequestBodyAttributesRaidEnum            `json:"raid,omitempty"`
	SSHKeys         []int64                                                           `json:"ssh_keys,omitempty"`
	UserData        *int64                                                            `json:"user_data,omitempty"`
}

type UpdateServerDeployConfigRequestBodyTypeEnum string

const (
	UpdateServerDeployConfigRequestBodyTypeEnumDeployConfig UpdateServerDeployConfigRequestBodyTypeEnum = "deploy_config"
)

type UpdateServerDeployConfigRequestBody struct {
	Attributes *UpdateServerDeployConfigRequestBodyAttributes `json:"attributes,omitempty"`
	Type       UpdateServerDeployConfigRequestBodyTypeEnum    `json:"type"`
}

type UpdateServerDeployConfigSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateServerDeployConfigRequest struct {
	PathParams UpdateServerDeployConfigPathParams
	Request    *UpdateServerDeployConfigRequestBody `request:"mediaType=application/json"`
	Security   UpdateServerDeployConfigSecurity
}

type UpdateServerDeployConfigResponse struct {
	ContentType  string
	StatusCode   int
	DeployConfig *shared.DeployConfig
	ErrorObject  *shared.ErrorObject
}
