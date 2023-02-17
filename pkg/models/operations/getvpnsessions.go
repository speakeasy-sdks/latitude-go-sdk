package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetVpnSessionsFilterSiteEnum string

const (
	GetVpnSessionsFilterSiteEnumMh1  GetVpnSessionsFilterSiteEnum = "MH1"
	GetVpnSessionsFilterSiteEnumSp1  GetVpnSessionsFilterSiteEnum = "SP1"
	GetVpnSessionsFilterSiteEnumSp4  GetVpnSessionsFilterSiteEnum = "SP4"
	GetVpnSessionsFilterSiteEnumSyd  GetVpnSessionsFilterSiteEnum = "SYD"
	GetVpnSessionsFilterSiteEnumCh1  GetVpnSessionsFilterSiteEnum = "CH1"
	GetVpnSessionsFilterSiteEnumDal2 GetVpnSessionsFilterSiteEnum = "DAL2"
	GetVpnSessionsFilterSiteEnumLa2  GetVpnSessionsFilterSiteEnum = "LA2"
	GetVpnSessionsFilterSiteEnumMi1  GetVpnSessionsFilterSiteEnum = "MI1"
	GetVpnSessionsFilterSiteEnumNy2  GetVpnSessionsFilterSiteEnum = "NY2"
	GetVpnSessionsFilterSiteEnumSan  GetVpnSessionsFilterSiteEnum = "SAN"
	GetVpnSessionsFilterSiteEnumTy6  GetVpnSessionsFilterSiteEnum = "TY6"
	GetVpnSessionsFilterSiteEnumTy8  GetVpnSessionsFilterSiteEnum = "TY8"
)

type GetVpnSessionsQueryParams struct {
	FilterSite *GetVpnSessionsFilterSiteEnum `queryParam:"style=form,explode=true,name=filter[site]"`
}

type GetVpnSessionsSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetVpnSessionsRequest struct {
	QueryParams GetVpnSessionsQueryParams
	Security    GetVpnSessionsSecurity
}

type GetVpnSessions200ApplicationJSON struct {
	Data []shared.VpnSessionDataWithPassword `json:"data,omitempty"`
	Meta map[string]interface{}              `json:"meta,omitempty"`
}

type GetVpnSessionsResponse struct {
	ContentType                            string
	StatusCode                             int
	GetVpnSessions200ApplicationJSONObject *GetVpnSessions200ApplicationJSON
}
