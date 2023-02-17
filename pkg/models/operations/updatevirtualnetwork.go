package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type UpdateVirtualNetworkPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateVirtualNetworkRequestBodyDataAttributes struct {
	Description *string `json:"description,omitempty"`
}

type UpdateVirtualNetworkRequestBodyDataTypeEnum string

const (
	UpdateVirtualNetworkRequestBodyDataTypeEnumVirtualNetwork UpdateVirtualNetworkRequestBodyDataTypeEnum = "virtual_network"
)

type UpdateVirtualNetworkRequestBodyData struct {
	Attributes *UpdateVirtualNetworkRequestBodyDataAttributes `json:"attributes,omitempty"`
	ID         string                                         `json:"id"`
	Type       UpdateVirtualNetworkRequestBodyDataTypeEnum    `json:"type"`
}

type UpdateVirtualNetworkRequestBody struct {
	Data UpdateVirtualNetworkRequestBodyData `json:"data"`
}

type UpdateVirtualNetworkSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateVirtualNetworkRequest struct {
	PathParams UpdateVirtualNetworkPathParams
	Request    *UpdateVirtualNetworkRequestBody `request:"mediaType=application/json"`
	Security   UpdateVirtualNetworkSecurity
}

type UpdateVirtualNetworkResponse struct {
	ContentType    string
	StatusCode     int
	VirtualNetwork *shared.VirtualNetwork
}
