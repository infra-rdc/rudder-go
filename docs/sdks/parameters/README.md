# Parameters
(*Parameters*)

## Overview

Global properties

### Available Operations

* [ListParameters](#listparameters) - List all global properties
* [ParameterDetails](#parameterdetails) - Get the value of a global property
* [UpdateParameter](#updateparameter) - Update a global property's value
* [DeleteParameter](#deleteparameter) - Delete a global parameter

## ListParameters

Get the current value of all the global properties (a.k.a. global parameters)

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

    ctx := context.Background()
    res, err := s.Parameters.ListParameters(ctx)
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

**[*operations.ListParametersResponse](../../models/operations/listparametersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ParameterDetails

Get the current value of a given global property (a.k.a. global parameter)

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
    var parameterID string = "rudder_file_edit_header"
    ctx := context.Background()
    res, err := s.Parameters.ParameterDetails(ctx, parameterID)
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
| `parameterID`                                         | *string*                                              | :heavy_check_mark:                                    | Id of the parameter to manage                         | rudder_file_edit_header                               |


### Response

**[*operations.ParameterDetailsResponse](../../models/operations/parameterdetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateParameter

Update the details of a global property

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
    var parameterID string = "rudder_file_edit_header"
    ctx := context.Background()
    res, err := s.Parameters.UpdateParameter(ctx, parameterID)
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
| `parameterID`                                         | *string*                                              | :heavy_check_mark:                                    | Id of the parameter to manage                         | rudder_file_edit_header                               |


### Response

**[*operations.UpdateParameterResponse](../../models/operations/updateparameterresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteParameter

Delete an existing global parameter

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
    var parameterID string = "rudder_file_edit_header"
    ctx := context.Background()
    res, err := s.Parameters.DeleteParameter(ctx, parameterID)
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
| `parameterID`                                         | *string*                                              | :heavy_check_mark:                                    | Id of the parameter to manage                         | rudder_file_edit_header                               |


### Response

**[*operations.DeleteParameterResponse](../../models/operations/deleteparameterresponse.md), error**
| Error Object                          | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.DeleteParameterResponseBody | 500                                   | application/json                      |
| sdkerrors.SDKError                    | 4xx-5xx                               | */*                                   |
