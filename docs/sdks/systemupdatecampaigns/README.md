# SystemUpdateCampaigns
(*SystemUpdateCampaigns*)

## Overview

**Requires that the `system update` plugin is installed on the server.**

Fetch System update campaigns results.

### Available Operations

* [GetCampaignResults](#getcampaignresults) - Get a campaign result history
* [GetCampaignEventResult](#getcampaigneventresult) - Get a campaign event result
* [GetSystemUpdateResultForNode](#getsystemupdateresultfornode) - Get detailed campaign event result for a Node

## GetCampaignResults

Get a campaign result history

### Example Usage

```go
package main

import(
	rudder "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/operations"
	"context"
	"log"
)

func main() {
    s := rudder.New(
        rudder.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    request := operations.GetCampaignResultsRequest{
        ID: "0076a379-f32d-4732-9e91-33ab219d8fde",
    }
    ctx := context.Background()
    res, err := s.SystemUpdateCampaigns.GetCampaignResults(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetCampaignResultsRequest](../../models/operations/getcampaignresultsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetCampaignResultsResponse](../../models/operations/getcampaignresultsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCampaignEventResult

Get a campaign event result

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
    var id string = "0076a379-f32d-4732-9e91-33ab219d8fde"
    ctx := context.Background()
    res, err := s.SystemUpdateCampaigns.GetCampaignEventResult(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Id of the campaign event                              | 0076a379-f32d-4732-9e91-33ab219d8fde                  |


### Response

**[*operations.GetCampaignEventResultResponse](../../models/operations/getcampaigneventresultresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSystemUpdateResultForNode

Get detailed campaign event result for a Node

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
    var id string = "0076a379-f32d-4732-9e91-33ab219d8fde"

    var nodeID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"
    ctx := context.Background()
    res, err := s.SystemUpdateCampaigns.GetSystemUpdateResultForNode(ctx, id, nodeID)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Id of the campaign event                              | 0076a379-f32d-4732-9e91-33ab219d8fde                  |
| `nodeID`                                              | *string*                                              | :heavy_check_mark:                                    | Id of the target node                                 | 9a1773c9-0889-40b6-be89-f6504443ac1b                  |


### Response

**[*operations.GetSystemUpdateResultForNodeResponse](../../models/operations/getsystemupdateresultfornoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
