package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type CreateVirtualNetworkRequestBodyDataAttributes struct {
	Description *string `json:"description,omitempty"`
	Project     *string `json:"project,omitempty"`
	Site        *string `json:"site,omitempty"`
}

type CreateVirtualNetworkRequestBodyDataTypeEnum string

const (
	CreateVirtualNetworkRequestBodyDataTypeEnumVirtualNetwork CreateVirtualNetworkRequestBodyDataTypeEnum = "virtual_network"
)

type CreateVirtualNetworkRequestBodyData struct {
	Attributes *CreateVirtualNetworkRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       CreateVirtualNetworkRequestBodyDataTypeEnum    `json:"type"`
}

type CreateVirtualNetworkRequestBody struct {
	Data CreateVirtualNetworkRequestBodyData `json:"data"`
}

type CreateVirtualNetworkSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type CreateVirtualNetworkRequest struct {
	Request  *CreateVirtualNetworkRequestBody `request:"mediaType=application/json"`
	Security CreateVirtualNetworkSecurity
}

type CreateVirtualNetworkResponse struct {
	ContentType    string
	StatusCode     int
	ErrorObject    *shared.ErrorObject
	VirtualNetwork *shared.VirtualNetwork
}
