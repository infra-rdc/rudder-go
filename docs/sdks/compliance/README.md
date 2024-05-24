# Compliance
(*Compliance*)

## Overview

Access compliance data

### Available Operations

* [GetGlobalCompliance](#getglobalcompliance) - Global compliance
* [GetDirectivesCompliance](#getdirectivescompliance) - Compliance details for all directives
* [GetDirectiveComplianceID](#getdirectivecomplianceid) - Compliance details by directive
* [GetNodesCompliance](#getnodescompliance) - Compliance details for all nodes
* [GetNodeCompliance](#getnodecompliance) - Compliance details by node
* [GetRulesCompliance](#getrulescompliance) - Compliance details for all rules
* [GetRuleCompliance](#getrulecompliance) - Compliance details by rule

## GetGlobalCompliance

Get current global compliance of a Rudder server

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
    var precision *int64 = ruddergo.Int64(0)
    ctx := context.Background()
    res, err := s.Compliance.GetGlobalCompliance(ctx, precision)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `precision`                                                | **int64*                                                   | :heavy_minus_sign:                                         | Number of digits after comma in compliance percent figures | 0                                                          |


### Response

**[*operations.GetGlobalComplianceResponse](../../models/operations/getglobalcomplianceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDirectivesCompliance

Get current compliance of all the nodes of a Rudder server

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
    res, err := s.Compliance.GetDirectivesCompliance(ctx)
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

**[*operations.GetDirectivesComplianceResponse](../../models/operations/getdirectivescomplianceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDirectiveComplianceID

Get current compliance of a directive of a Rudder server

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
    var directiveID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"

    var format *string = ruddergo.String("<value>")
    ctx := context.Background()
    res, err := s.Compliance.GetDirectiveComplianceID(ctx, directiveID, format)
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `directiveID`                                         | *string*                                              | :heavy_check_mark:                                    | Id of the directive                                   | 9a1773c9-0889-40b6-be89-f6504443ac1b                  |
| `format`                                              | **string*                                             | :heavy_minus_sign:                                    | format of export                                      |                                                       |


### Response

**[*operations.GetDirectiveComplianceIDResponse](../../models/operations/getdirectivecomplianceidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetNodesCompliance

Get current compliance of all the nodes of a Rudder server

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
    var level *int64 = ruddergo.Int64(4)

    var precision *int64 = ruddergo.Int64(0)
    ctx := context.Background()
    res, err := s.Compliance.GetNodesCompliance(ctx, level, precision)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                | Example                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |                                                                                                                            |
| `level`                                                                                                                    | **int64*                                                                                                                   | :heavy_minus_sign:                                                                                                         | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) | 4                                                                                                                          |
| `precision`                                                                                                                | **int64*                                                                                                                   | :heavy_minus_sign:                                                                                                         | Number of digits after comma in compliance percent figures                                                                 | 0                                                                                                                          |


### Response

**[*operations.GetNodesComplianceResponse](../../models/operations/getnodescomplianceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetNodeCompliance

Get current compliance of a node of a Rudder server

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

    var level *int64 = ruddergo.Int64(4)

    var precision *int64 = ruddergo.Int64(0)
    ctx := context.Background()
    res, err := s.Compliance.GetNodeCompliance(ctx, nodeID, level, precision)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                | Example                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |                                                                                                                            |
| `nodeID`                                                                                                                   | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | Id of the target node                                                                                                      | 9a1773c9-0889-40b6-be89-f6504443ac1b                                                                                       |
| `level`                                                                                                                    | **int64*                                                                                                                   | :heavy_minus_sign:                                                                                                         | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) | 4                                                                                                                          |
| `precision`                                                                                                                | **int64*                                                                                                                   | :heavy_minus_sign:                                                                                                         | Number of digits after comma in compliance percent figures                                                                 | 0                                                                                                                          |


### Response

**[*operations.GetNodeComplianceResponse](../../models/operations/getnodecomplianceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRulesCompliance

Get current compliance of all the rules of a Rudder server

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
    var level *int64 = ruddergo.Int64(4)

    var precision *int64 = ruddergo.Int64(0)
    ctx := context.Background()
    res, err := s.Compliance.GetRulesCompliance(ctx, level, precision)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                | Example                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |                                                                                                                            |
| `level`                                                                                                                    | **int64*                                                                                                                   | :heavy_minus_sign:                                                                                                         | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) | 4                                                                                                                          |
| `precision`                                                                                                                | **int64*                                                                                                                   | :heavy_minus_sign:                                                                                                         | Number of digits after comma in compliance percent figures                                                                 | 0                                                                                                                          |


### Response

**[*operations.GetRulesComplianceResponse](../../models/operations/getrulescomplianceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRuleCompliance

Get current compliance of a rule of a Rudder server

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
    var ruleID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"

    var level *int64 = ruddergo.Int64(4)

    var precision *int64 = ruddergo.Int64(0)
    ctx := context.Background()
    res, err := s.Compliance.GetRuleCompliance(ctx, ruleID, level, precision)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                | Example                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |                                                                                                                            |
| `ruleID`                                                                                                                   | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | Id of the target rule                                                                                                      | 9a1773c9-0889-40b6-be89-f6504443ac1b                                                                                       |
| `level`                                                                                                                    | **int64*                                                                                                                   | :heavy_minus_sign:                                                                                                         | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) | 4                                                                                                                          |
| `precision`                                                                                                                | **int64*                                                                                                                   | :heavy_minus_sign:                                                                                                         | Number of digits after comma in compliance percent figures                                                                 | 0                                                                                                                          |


### Response

**[*operations.GetRuleComplianceResponse](../../models/operations/getrulecomplianceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
