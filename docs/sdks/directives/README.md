# Directives
(*Directives*)

## Overview

Directives management

### Available Operations

* [ListDirectives](#listdirectives) - List all directives
* [CreateDirective](#createdirective) - Create a directive
* [DirectiveDetails](#directivedetails) - Get directive details
* [DeleteDirective](#deletedirective) - Delete a directive

## ListDirectives

List all directives

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
    res, err := s.Directives.ListDirectives(ctx)
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

**[*operations.ListDirectivesResponse](../../models/operations/listdirectivesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateDirective

Create a new directive from provided parameters. You can specify a source directive to clone it.

### Example Usage

```go
package main

import(
	rudder "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/components"
	"context"
	"log"
)

func main() {
    s := rudder.New(
        rudder.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var request *components.DirectiveNew = &components.DirectiveNew{
        Source: rudder.String("b9f6d98a-28bc-4d80-90f7-d2f14269e215"),
        ID: rudder.String("91252ea2-feb2-412d-8599-c6945fee02c4"),
        DisplayName: rudder.String("91252ea2-feb2-412d-8599-c6945fee02c4"),
        ShortDescription: rudder.String("91252ea2-feb2-412d-8599-c6945fee02c4"),
        LongDescription: rudder.String("# Documentation
    * [Ticket link](https://tickets.example.com/issues/3456)"),
        TechniqueName: rudder.String("userManagement"),
        TechniqueVersion: rudder.String("8.0"),
        Priority: rudder.Int64(5),
        Enabled: rudder.Bool(true),
        System: rudder.Bool(false),
        Tags: []components.DirectiveNewTags{
            components.DirectiveNewTags{
                Name: rudder.String("value"),
            },
        },
        Parameters: &components.DirectiveNewParameters{},
    }
    ctx := context.Background()
    res, err := s.Directives.CreateDirective(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.DirectiveNew](../../models/components/directivenew.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.CreateDirectiveResponse](../../models/operations/createdirectiveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DirectiveDetails

Get all information about a given directive

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
    var directiveID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"
    ctx := context.Background()
    res, err := s.Directives.DirectiveDetails(ctx, directiveID)
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
| `directiveID`                                         | *string*                                              | :heavy_check_mark:                                    | Id of the directive                                   | 9a1773c9-0889-40b6-be89-f6504443ac1b                  |


### Response

**[*operations.DirectiveDetailsResponse](../../models/operations/directivedetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteDirective

Delete a directive

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
    var directiveID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"
    ctx := context.Background()
    res, err := s.Directives.DeleteDirective(ctx, directiveID)
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
| `directiveID`                                         | *string*                                              | :heavy_check_mark:                                    | Id of the directive                                   | 9a1773c9-0889-40b6-be89-f6504443ac1b                  |


### Response

**[*operations.DeleteDirectiveResponse](../../models/operations/deletedirectiveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
