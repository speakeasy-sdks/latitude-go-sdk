package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type AssignServerVirtualNetworkRequestBodyDataAttributes struct {
	ServerID         int64 `json:"server_id"`
	VirtualNetworkID int64 `json:"virtual_network_id"`
}

type AssignServerVirtualNetworkRequestBodyDataTypeEnum string

const (
	AssignServerVirtualNetworkRequestBodyDataTypeEnumVirtualNetworkAssignment AssignServerVirtualNetworkRequestBodyDataTypeEnum = "virtual_network_assignment"
)

type AssignServerVirtualNetworkRequestBodyData struct {
	Attributes *AssignServerVirtualNetworkRequestBodyDataAttributes `json:"attributes,omitempty"`
	Type       AssignServerVirtualNetworkRequestBodyDataTypeEnum    `json:"type"`
}

type AssignServerVirtualNetworkRequestBody struct {
	Data *AssignServerVirtualNetworkRequestBodyData `json:"data,omitempty"`
}

type AssignServerVirtualNetworkSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type AssignServerVirtualNetworkRequest struct {
	Request  *AssignServerVirtualNetworkRequestBody `request:"mediaType=application/json"`
	Security AssignServerVirtualNetworkSecurity
}

type AssignServerVirtualNetworkResponse struct {
	ContentType              string
	StatusCode               int
	ErrorObject              *shared.ErrorObject
	VirtualNetworkAssignment *shared.VirtualNetworkAssignment
}
