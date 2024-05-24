# ScaleOutRelay
(*ScaleOutRelay*)

## Overview

**Requires that the `scale-out-relay` plugin is installed on the server.**

Manage relays.

### Available Operations

* [DemoteToNode](#demotetonode) - Demote a relay to simple node
* [PromoteToRelay](#promotetorelay) - Promote a node to relay

## DemoteToNode

Demote a relay to a simple node.

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
    var nodeID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"
    ctx := context.Background()
    res, err := s.ScaleOutRelay.DemoteToNode(ctx, nodeID)
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
| `nodeID`                                              | *string*                                              | :heavy_check_mark:                                    | Id of the target node                                 | 9a1773c9-0889-40b6-be89-f6504443ac1b                  |


### Response

**[*operations.DemoteToNodeResponse](../../models/operations/demotetonoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PromoteToRelay

Promote a node to relay

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
    var nodeID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"
    ctx := context.Background()
    res, err := s.ScaleOutRelay.PromoteToRelay(ctx, nodeID)
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
| `nodeID`                                              | *string*                                              | :heavy_check_mark:                                    | Id of the target node                                 | 9a1773c9-0889-40b6-be89-f6504443ac1b                  |


### Response

**[*operations.PromoteToRelayResponse](../../models/operations/promotetorelayresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
