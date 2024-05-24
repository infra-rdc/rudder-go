# OpenSCAP
(*OpenSCAP*)

## Overview

**Requires that the `openscap` plugin is installed on the server.**

Get OpenSCAP reports for nodes.

### Available Operations

* [OpenscapReport](#openscapreport) - Get an OpenSCAP report

## OpenscapReport

Get latest OpenSCAP report for the given node

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
    res, err := s.OpenSCAP.OpenscapReport(ctx, nodeID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
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

**[*operations.OpenscapReportResponse](../../models/operations/openscapreportresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
