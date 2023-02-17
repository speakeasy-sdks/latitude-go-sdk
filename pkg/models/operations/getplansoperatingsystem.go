package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetPlansOperatingSystemSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetPlansOperatingSystemRequest struct {
	Security GetPlansOperatingSystemSecurity
}

type GetPlansOperatingSystem200ApplicationJSON struct {
	Data []shared.OperatingSystems `json:"data,omitempty"`
}

type GetPlansOperatingSystemResponse struct {
	ContentType                                     string
	StatusCode                                      int
	GetPlansOperatingSystem200ApplicationJSONObject *GetPlansOperatingSystem200ApplicationJSON
}
