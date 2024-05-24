# Cve
(*Cve*)

## Overview

**Requires that the `cve` plugin is installed on the server.**

Manage CVE plugins data and configuration.

### Available Operations

* [GetAllCve](#getallcve) - Get all CVE details
* [CheckCVE](#checkcve) - Trigger a CVE check
* [GetCVECheckConfiguration](#getcvecheckconfiguration) - Get CVE check config
* [UpdateCVECheckConfiguration](#updatecvecheckconfiguration) - Update cve check config
* [GetLastCVECheck](#getlastcvecheck) - Get last CVE check result
* [GetCVEList](#getcvelist) - Get a list of CVE details
* [UpdateCVE](#updatecve) - Update CVE database from remote source
* [ReadCVEfromFS](#readcvefromfs) - Update CVE database from file system
* [GetCve](#getcve) - Get a CVE details

## GetAllCve

Get all CVE details

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
    res, err := s.Cve.GetAllCve(ctx)
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

**[*operations.GetAllCveResponse](../../models/operations/getallcveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CheckCVE

Trigger a CVE check

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
    res, err := s.Cve.CheckCVE(ctx)
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

**[*operations.CheckCVEResponse](../../models/operations/checkcveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCVECheckConfiguration

Get CVE check config

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
    res, err := s.Cve.GetCVECheckConfiguration(ctx)
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

**[*operations.GetCVECheckConfigurationResponse](../../models/operations/getcvecheckconfigurationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCVECheckConfiguration

Update cve check config

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
    var request *operations.UpdateCVECheckConfigurationRequestBody = &operations.UpdateCVECheckConfigurationRequestBody{
        URL: rudder.String("https://api.rudder.io/cve/v1/"),
    }
    ctx := context.Background()
    res, err := s.Cve.UpdateCVECheckConfiguration(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpdateCVECheckConfigurationRequestBody](../../models/operations/updatecvecheckconfigurationrequestbody.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.UpdateCVECheckConfigurationResponse](../../models/operations/updatecvecheckconfigurationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetLastCVECheck

Get last CVE check result

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
    request := operations.GetLastCVECheckRequest{}
    ctx := context.Background()
    res, err := s.Cve.GetLastCVECheck(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLastCVECheckRequest](../../models/operations/getlastcvecheckrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetLastCVECheckResponse](../../models/operations/getlastcvecheckresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCVEList

Get CVE details, from a list passed as parameter

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
    var request *operations.GetCVEListRequestBody = &operations.GetCVEListRequestBody{
        CveIds: []string{
            "CVE-2019-5953",
        },
        OnlyScore: rudder.Bool(true),
        MinScore: rudder.String("7.5"),
        MaxScore: rudder.String("8.5"),
    }
    ctx := context.Background()
    res, err := s.Cve.GetCVEList(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetCVEListRequestBody](../../models/operations/getcvelistrequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetCVEListResponse](../../models/operations/getcvelistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCVE

Update CVE database from remote source

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
    var request *operations.UpdateCVERequestBody = &operations.UpdateCVERequestBody{
        URL: rudder.String("https://nvd.nist.gov/feeds/json/cve/1.1"),
        Years: []string{
            "2019",
        },
    }
    ctx := context.Background()
    res, err := s.Cve.UpdateCVE(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateCVERequestBody](../../models/operations/updatecverequestbody.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateCVEResponse](../../models/operations/updatecveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReadCVEfromFS

Update CVE database from file system

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
    res, err := s.Cve.ReadCVEfromFS(ctx)
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

**[*operations.ReadCVEfromFSResponse](../../models/operations/readcvefromfsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCve

Get a CVE details

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
    var cveID string = "CVE-2022-25235"
    ctx := context.Background()
    res, err := s.Cve.GetCve(ctx, cveID)
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
| `cveID`                                               | *string*                                              | :heavy_check_mark:                                    | Id of the CVE                                         | CVE-2022-25235                                        |


### Response

**[*operations.GetCveResponse](../../models/operations/getcveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
