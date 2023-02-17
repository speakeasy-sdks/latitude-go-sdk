<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/latitude-go-sdk"
    "github.com/speakeasy-sdks/latitude-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/latitude-go-sdk/pkg/models/operations"
)

func main() {
    s := latitude.New()
    
    req := operations.DeleteAPIKeyRequest{
        Security: operations.DeleteAPIKeySecurity{
            Bearer: shared.SchemeBearer{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
        PathParams: operations.DeleteAPIKeyPathParams{
            ID: "unde",
        },
    }
    
    res, err := s.APIKeys.DeleteAPIKey(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
```
<!-- End SDK Example Usage -->