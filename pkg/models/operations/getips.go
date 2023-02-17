package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetIpsFilterFamilyEnum string

const (
	GetIpsFilterFamilyEnumIPv4 GetIpsFilterFamilyEnum = "IPv4"
	GetIpsFilterFamilyEnumIPv6 GetIpsFilterFamilyEnum = "IPv6"
)

type GetIpsFilterTypeEnum string

const (
	GetIpsFilterTypeEnumPrivate GetIpsFilterTypeEnum = "private"
	GetIpsFilterTypeEnumPublic  GetIpsFilterTypeEnum = "public"
)

type GetIpsQueryParams struct {
	ExtraFieldsIPAddresses *string                 `queryParam:"style=form,explode=true,name=extra_fields[ip_addresses]"`
	FilterAddress          *string                 `queryParam:"style=form,explode=true,name=filter[address]"`
	FilterFamily           *GetIpsFilterFamilyEnum `queryParam:"style=form,explode=true,name=filter[family]"`
	FilterLocation         *string                 `queryParam:"style=form,explode=true,name=filter[location]"`
	FilterProject          *string                 `queryParam:"style=form,explode=true,name=filter[project]"`
	FilterServer           *string                 `queryParam:"style=form,explode=true,name=filter[server]"`
	FilterType             *GetIpsFilterTypeEnum   `queryParam:"style=form,explode=true,name=filter[type]"`
	Include                *string                 `queryParam:"style=form,explode=true,name=include"`
}

type GetIpsSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetIpsRequest struct {
	QueryParams GetIpsQueryParams
	Security    GetIpsSecurity
}

type GetIpsResponse struct {
	ContentType string
	StatusCode  int
	ErrorObject *shared.ErrorObject
	IPAddresses *shared.IPAddresses
}
