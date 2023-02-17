package operations

import (
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
)

type GetProjectsQueryParams struct {
	ExtraFieldsProjects *string `queryParam:"style=form,explode=true,name=extra_fields[projects]"`
	FilterBillingType   *string `queryParam:"style=form,explode=true,name=filter[billing_type]"`
	FilterDescription   *string `queryParam:"style=form,explode=true,name=filter[description]"`
	FilterEnvironment   *string `queryParam:"style=form,explode=true,name=filter[environment]"`
	FilterName          *string `queryParam:"style=form,explode=true,name=filter[name]"`
	FilterSlug          *string `queryParam:"style=form,explode=true,name=filter[slug]"`
	Include             *string `queryParam:"style=form,explode=true,name=include"`
}

type GetProjectsSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=apiKey,subtype=header"`
}

type GetProjectsRequest struct {
	QueryParams GetProjectsQueryParams
	Security    GetProjectsSecurity
}

type GetProjectsResponse struct {
	ContentType string
	StatusCode  int
	Projects    *shared.Projects
}
