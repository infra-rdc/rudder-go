# Rules
(*Rules*)

## Overview

Rules management

### Available Operations

* [GetRuleCategoryDetails](#getrulecategorydetails) - Get rule category details
* [DeleteRuleCategory](#deleterulecategory) - Delete group category
* [UpdateRuleCategory](#updaterulecategory) - Update rule category details
* [GetRuleTree](#getruletree) - Get rules tree

## GetRuleCategoryDetails

Get detailed information about a rule category

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
    var ruleCategoryID string = "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90"
    ctx := context.Background()
    res, err := s.Rules.GetRuleCategoryDetails(ctx, ruleCategoryID)
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
| `ruleCategoryID`                                      | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | e0a311fa-f7b2-4f9e-89a9-db517b9c6b90                  |


### Response

**[*operations.GetRuleCategoryDetailsResponse](../../models/operations/getrulecategorydetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteRuleCategory

Delete a group category. It must have no child groups and no children categories.

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
    var ruleCategoryID string = "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90"
    ctx := context.Background()
    res, err := s.Rules.DeleteRuleCategory(ctx, ruleCategoryID)
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
| `ruleCategoryID`                                      | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | e0a311fa-f7b2-4f9e-89a9-db517b9c6b90                  |


### Response

**[*operations.DeleteRuleCategoryResponse](../../models/operations/deleterulecategoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateRuleCategory

Update detailed information about a rule category

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/components"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var ruleCategoryID string = "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90"

    ruleCategoryUpdate := components.RuleCategoryUpdate{
        Parent: "b9f6d98a-28bc-4d80-90f7-d2f14269e215",
        Name: "Security policies",
        Description: ruddergo.String("Baseline applying CIS guidelines"),
    }
    ctx := context.Background()
    res, err := s.Rules.UpdateRuleCategory(ctx, ruleCategoryID, ruleCategoryUpdate)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `ruleCategoryID`                                                               | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            | e0a311fa-f7b2-4f9e-89a9-db517b9c6b90                                           |
| `ruleCategoryUpdate`                                                           | [components.RuleCategoryUpdate](../../models/components/rulecategoryupdate.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |


### Response

**[*operations.UpdateRuleCategoryResponse](../../models/operations/updaterulecategoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRuleTree

Get all available rules and their categories in a tree

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
    res, err := s.Rules.GetRuleTree(ctx)
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

**[*operations.GetRuleTreeResponse](../../models/operations/getruletreeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
