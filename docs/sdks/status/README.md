# Status
(*Status*)

## Overview

Is alive check

### Available Operations

* [None](#none) - Check if Rudder is alive

## None

An unauthenticated API to check if Rudder web application is up and running. Be careful: this API does not follow other Rudder's API convention: 
- it is not versioned, so the path is just `/api/status`;
- it returns a plain text message.

### Example Usage

```go
package main

import(
	rudder "github.com/infra-rdc/rudder-go"
	"context"
	"log"
)

func main() {
    s := rudder.New()

    ctx := context.Background()
    res, err := s.Status.None(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.NoneResponse](../../models/operations/noneresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
