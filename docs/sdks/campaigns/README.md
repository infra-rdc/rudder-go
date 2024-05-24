# Campaigns
(*Campaigns*)

## Overview

Campaigns

### Available Operations

* [GetAllCampaignEvents](#getallcampaignevents) - Get all campaign events
* [GetCampaignEvent](#getcampaignevent) - Get a campaign event details
* [SaveCampaignEvent](#savecampaignevent) - Update an existing event
* [DeleteCampaignEvent](#deletecampaignevent) - Delete a campaign event details
* [GetEventsCampaign](#geteventscampaign) - Get campaign events for a campaign
* [ScheduleCampaign](#schedulecampaign) - Schedule a campaign event for a campaign

## GetAllCampaignEvents

Get all campaign events

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/models/operations"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    request := operations.GetAllCampaignEventsRequest{
        CampaignType: components.CampaignTypeSystemUpdate.ToPointer(),
        State: components.CampaignEventStatusRunning.ToPointer(),
        CampaignID: components.CampaignIDSystemUpdate.ToPointer(),
    }
    ctx := context.Background()
    res, err := s.Campaigns.GetAllCampaignEvents(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAllCampaignEventsRequest](../../models/operations/getallcampaigneventsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetAllCampaignEventsResponse](../../models/operations/getallcampaigneventsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCampaignEvent

Get a campaign event details

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var id string = "0076a379-f32d-4732-9e91-33ab219d8fde"
    ctx := context.Background()
    res, err := s.Campaigns.GetCampaignEvent(ctx, id)
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

**[*operations.GetCampaignEventResponse](../../models/operations/getcampaigneventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SaveCampaignEvent

Update an existing event

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var id string = "0076a379-f32d-4732-9e91-33ab219d8fde"
    ctx := context.Background()
    res, err := s.Campaigns.SaveCampaignEvent(ctx, id)
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

**[*operations.SaveCampaignEventResponse](../../models/operations/savecampaigneventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteCampaignEvent

Delete a campaign event details

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var id string = "0076a379-f32d-4732-9e91-33ab219d8fde"
    ctx := context.Background()
    res, err := s.Campaigns.DeleteCampaignEvent(ctx, id)
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

**[*operations.DeleteCampaignEventResponse](../../models/operations/deletecampaigneventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetEventsCampaign

Get campaign events for a campaign

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/models/operations"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    request := operations.GetEventsCampaignRequest{
        ID: "0076a379-f32d-4732-9e91-33ab219d8fde",
        CampaignType: components.CampaignTypeSystemUpdate.ToPointer(),
        State: components.CampaignEventStatusScheduled.ToPointer(),
    }
    ctx := context.Background()
    res, err := s.Campaigns.GetEventsCampaign(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetEventsCampaignRequest](../../models/operations/geteventscampaignrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetEventsCampaignResponse](../../models/operations/geteventscampaignresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ScheduleCampaign

Schedule a campaign event for a campaign

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var id string = "0076a379-f32d-4732-9e91-33ab219d8fde"
    ctx := context.Background()
    res, err := s.Campaigns.ScheduleCampaign(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Id of the campaign                                    | 0076a379-f32d-4732-9e91-33ab219d8fde                  |


### Response

**[*operations.ScheduleCampaignResponse](../../models/operations/schedulecampaignresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
