package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetPlansFilterStockLevelEnum string

const (
	GetPlansFilterStockLevelEnumUnavailable GetPlansFilterStockLevelEnum = "Unavailable"
	GetPlansFilterStockLevelEnumLow         GetPlansFilterStockLevelEnum = "Low"
	GetPlansFilterStockLevelEnumMedium      GetPlansFilterStockLevelEnum = "Medium"
	GetPlansFilterStockLevelEnumHigh        GetPlansFilterStockLevelEnum = "High"
	GetPlansFilterStockLevelEnumUnique      GetPlansFilterStockLevelEnum = "Unique"
)

type GetPlansQueryParams struct {
	FilterInStock    *bool                         `queryParam:"style=form,explode=true,name=filter[in_stock]"`
	FilterLocation   *string                       `queryParam:"style=form,explode=true,name=filter[location]"`
	FilterName       *string                       `queryParam:"style=form,explode=true,name=filter[name]"`
	FilterSlug       *string                       `queryParam:"style=form,explode=true,name=filter[slug]"`
	FilterStockLevel *GetPlansFilterStockLevelEnum `queryParam:"style=form,explode=true,name=filter[stock_level]"`
}

type GetPlansSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetPlansRequest struct {
	QueryParams GetPlansQueryParams
	Security    GetPlansSecurity
}

type GetPlans200ApplicationJSON struct {
	Data []shared.PlanData `json:"data,omitempty"`
}

type GetPlansResponse struct {
	ContentType                      string
	StatusCode                       int
	GetPlans200ApplicationJSONObject *GetPlans200ApplicationJSON
}
