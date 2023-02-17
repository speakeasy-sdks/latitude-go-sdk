package shared

type VpnSessionDataWithPasswordAttributesStatusEnum string

const (
	VpnSessionDataWithPasswordAttributesStatusEnumEnable  VpnSessionDataWithPasswordAttributesStatusEnum = "enable"
	VpnSessionDataWithPasswordAttributesStatusEnumDisable VpnSessionDataWithPasswordAttributesStatusEnum = "disable"
)

type VpnSessionDataWithPasswordAttributes struct {
	CreatedAt *string                                         `json:"created_at,omitempty"`
	ExpiresAt *string                                         `json:"expires_at,omitempty"`
	Host      *string                                         `json:"host,omitempty"`
	Password  *string                                         `json:"password,omitempty"`
	Port      *string                                         `json:"port,omitempty"`
	Region    *RegionResourceData                             `json:"region,omitempty"`
	Status    *VpnSessionDataWithPasswordAttributesStatusEnum `json:"status,omitempty"`
	UpdatedAt *string                                         `json:"updated_at,omitempty"`
	UserName  *string                                         `json:"user_name,omitempty"`
}

type VpnSessionDataWithPasswordTypeEnum string

const (
	VpnSessionDataWithPasswordTypeEnumVpnSessions VpnSessionDataWithPasswordTypeEnum = "vpn_sessions"
)

type VpnSessionDataWithPassword struct {
	Attributes *VpnSessionDataWithPasswordAttributes `json:"attributes,omitempty"`
	ID         *string                               `json:"id,omitempty"`
	Type       *VpnSessionDataWithPasswordTypeEnum   `json:"type,omitempty"`
}
