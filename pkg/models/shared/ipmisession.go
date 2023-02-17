package shared

type IpmiSessionDataAttributes struct {
	IpmiAddress  *string `json:"ipmi_address,omitempty"`
	IpmiPassword *string `json:"ipmi_password,omitempty"`
	IpmiUsername *string `json:"ipmi_username,omitempty"`
}

type IpmiSessionDataTypeEnum string

const (
	IpmiSessionDataTypeEnumIpmiSessions IpmiSessionDataTypeEnum = "ipmi_sessions"
)

type IpmiSessionData struct {
	Attributes *IpmiSessionDataAttributes `json:"attributes,omitempty"`
	ID         *string                    `json:"id,omitempty"`
	Type       *IpmiSessionDataTypeEnum   `json:"type,omitempty"`
}

type IpmiSession struct {
	Data *IpmiSessionData `json:"data,omitempty"`
}
