package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type CreateServerReinstallPathParams struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

type CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum string

const (
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumWindowsServer2019StdV1 CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "windows_server_2019_std_v1"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumWindowsServer2019DcV1  CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "windows_server_2019_dc_v1"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumWindowsServer2016StdV1 CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "windows_server_2016_std_v1"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumWindowsServer2016DcV1  CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "windows_server_2016_dc_v1"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumUbuntu2204X64Lts       CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "ubuntu_22_04_x64_lts"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumDebian11               CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "debian_11"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumRockylinux8            CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "rockylinux_8"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumDebian10               CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "debian_10"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumFlatcarStable          CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "flatcar_stable"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumRhel8                  CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "rhel8"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumCentos74X64            CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "centos_7_4_x64"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumCentos8X64             CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "centos_8_x64"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumUbuntu1804X64Lts       CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "ubuntu_18_04_x64_lts"
	CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnumUbuntu2004X64Lts       CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum = "ubuntu_20_04_x64_lts"
)

type CreateServerReinstallRequestBodyDataAttributes struct {
	Hostname        *string                                                            `json:"hostname,omitempty"`
	OperatingSystem *CreateServerReinstallRequestBodyDataAttributesOperatingSystemEnum `json:"operating_system,omitempty"`
	Raid            *int64                                                             `json:"raid,omitempty"`
	SSHKeys         []string                                                           `json:"ssh_keys,omitempty"`
	UserData        *int64                                                             `json:"user_data,omitempty"`
}

type CreateServerReinstallRequestBodyDataTypeEnum string

const (
	CreateServerReinstallRequestBodyDataTypeEnumReinstalls CreateServerReinstallRequestBodyDataTypeEnum = "reinstalls"
)

type CreateServerReinstallRequestBodyData struct {
	Attributes *CreateServerReinstallRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       CreateServerReinstallRequestBodyDataTypeEnum    `json:"type"`
}

type CreateServerReinstallRequestBody struct {
	Data CreateServerReinstallRequestBodyData `json:"data"`
}

type CreateServerReinstallSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type CreateServerReinstallRequest struct {
	PathParams CreateServerReinstallPathParams
	Request    *CreateServerReinstallRequestBody `request:"mediaType=application/json"`
	Security   CreateServerReinstallSecurity
}

type CreateServerReinstallResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
}
