# APIInfo
(*APIInfo*)

## Overview

Information about API endpoints and versions

### Available Operations

* [APIGeneralInformations](#apigeneralinformations) - List all endpoints
* [APIInformations](#apiinformations) - Get information about one API endpoint
* [APISubInformations](#apisubinformations) - Get information on endpoint in a section

## APIGeneralInformations

List all endpoints and their version

### Example Usage

```go
package main

import(
	rudder "github.com/infra-rdc/rudder-go"
	"context"
	"log"
)

func main() {
    s := rudder.New(
        rudder.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.APIInfo.APIGeneralInformations(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.APIGeneralInformationsResponse](../../models/operations/apigeneralinformationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## APIInformations

Get the description and the list of supported version for one API endpoint

### Example Usage

```go
package main

import(
	rudder "github.com/infra-rdc/rudder-go"
	"context"
	"log"
)

func main() {
    s := rudder.New(
        rudder.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var endpointName string = "listAcceptedNodes"
    ctx := context.Background()
    res, err := s.APIInfo.APIInformations(ctx, endpointName)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `endpointName`                                        | *string*                                              | :heavy_check_mark:                                    | Name of the endpoint for which one wants information  | listAcceptedNodes                                     |


### Response

**[*operations.APIInformationsResponse](../../models/operations/apiinformationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## APISubInformations

Get all endpoints in the given section with their supported version.

### Example Usage

```go
package main

import(
	rudder "github.com/infra-rdc/rudder-go"
	"context"
	"log"
)

func main() {
    s := rudder.New(
        rudder.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var sectionID string = "nodes"
    ctx := context.Background()
    res, err := s.APIInfo.APISubInformations(ctx, sectionID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `sectionID`                                           | *string*                                              | :heavy_check_mark:                                    | Id of the API section                                 | nodes                                                 |


### Response

**[*operations.APISubInformationsResponse](../../models/operations/apisubinformationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
