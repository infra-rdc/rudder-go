# Groups
(*Groups*)

## Overview

Groups management

### Available Operations

* [ListGroups](#listgroups) - List all groups
* [CreateGroup](#creategroup) - Create a group
* [GetGroupCategoryDetails](#getgroupcategorydetails) - Get group category details
* [DeleteGroupCategory](#deletegroupcategory) - Delete group category
* [UpdateGroupCategory](#updategroupcategory) - Update group category details
* [GetGroupTree](#getgrouptree) - Get groups tree
* [GroupDetails](#groupdetails) - Get group details
* [UpdateGroup](#updategroup) - Update group details
* [DeleteGroup](#deletegroup) - Delete a group
* [ReloadGroup](#reloadgroup) - Reload a group

## ListGroups

List all groups

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
    res, err := s.Groups.ListGroups(ctx)
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

**[*operations.ListGroupsResponse](../../models/operations/listgroupsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateGroup

Create a new group based in provided parameters

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
    var request *components.GroupNew = &components.GroupNew{
        Source: ruddergo.String("b9f6d98a-28bc-4d80-90f7-d2f14269e215"),
        Category: "e17ecf6a-a9f2-44de-a97c-116d24d30ff4",
        ID: ruddergo.String("32d013f7-b6d8-46c8-99d3-016307fa66c0"),
        DisplayName: "Ubuntu 18.04 nodes",
        Description: ruddergo.String("Documentation for the group"),
        Query: &components.GroupNewQuery{
            Composition: components.GroupNewCompositionAnd.ToPointer(),
            Where: []components.GroupNewWhere{
                components.GroupNewWhere{
                    ObjectType: ruddergo.String("node"),
                    Attribute: ruddergo.String("OS"),
                    Comparator: ruddergo.String("eq"),
                    Value: ruddergo.String("Linux"),
                },
            },
        },
        Properties: []components.GroupNewProperties{
            components.GroupNewProperties{
                Name: "os",
                Value: "debian10",
            },
        },
    }
    ctx := context.Background()
    res, err := s.Groups.CreateGroup(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [components.GroupNew](../../models/components/groupnew.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.CreateGroupResponse](../../models/operations/creategroupresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetGroupCategoryDetails

Get detailed information about a group category

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
    var groupCategoryID string = "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90"
    ctx := context.Background()
    res, err := s.Groups.GetGroupCategoryDetails(ctx, groupCategoryID)
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
| `groupCategoryID`                                     | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | e0a311fa-f7b2-4f9e-89a9-db517b9c6b90                  |


### Response

**[*operations.GetGroupCategoryDetailsResponse](../../models/operations/getgroupcategorydetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteGroupCategory

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
    var groupCategoryID string = "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90"
    ctx := context.Background()
    res, err := s.Groups.DeleteGroupCategory(ctx, groupCategoryID)
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
| `groupCategoryID`                                     | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | e0a311fa-f7b2-4f9e-89a9-db517b9c6b90                  |


### Response

**[*operations.DeleteGroupCategoryResponse](../../models/operations/deletegroupcategoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateGroupCategory

Update detailed information about a group category

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
    var groupCategoryID string = "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90"

    groupCategoryUpdate := components.GroupCategoryUpdate{
        Parent: "b9f6d98a-28bc-4d80-90f7-d2f14269e215",
        Name: "Hardware groups",
        Description: ruddergo.String("Nodes by hardware provider"),
    }
    ctx := context.Background()
    res, err := s.Groups.UpdateGroupCategory(ctx, groupCategoryID, groupCategoryUpdate)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `groupCategoryID`                                                                | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              | e0a311fa-f7b2-4f9e-89a9-db517b9c6b90                                             |
| `groupCategoryUpdate`                                                            | [components.GroupCategoryUpdate](../../models/components/groupcategoryupdate.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |


### Response

**[*operations.UpdateGroupCategoryResponse](../../models/operations/updategroupcategoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetGroupTree

Get all available groups and their categories in a tree

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
    res, err := s.Groups.GetGroupTree(ctx)
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

**[*operations.GetGroupTreeResponse](../../models/operations/getgrouptreeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GroupDetails

Get detailed information about a group

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
    var groupID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"
    ctx := context.Background()
    res, err := s.Groups.GroupDetails(ctx, groupID)
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
| `groupID`                                             | *string*                                              | :heavy_check_mark:                                    | Id of the group                                       | 9a1773c9-0889-40b6-be89-f6504443ac1b                  |


### Response

**[*operations.GroupDetailsResponse](../../models/operations/groupdetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateGroup

Update detailed information about a group

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
    var groupID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"

    groupUpdate := components.GroupUpdate{
        Category: ruddergo.String("e17ecf6a-a9f2-44de-a97c-116d24d30ff4"),
        DisplayName: ruddergo.String("Ubuntu 18.04 nodes"),
        Description: ruddergo.String("Documentation for the group"),
        Query: &components.GroupUpdateQuery{
            Composition: components.GroupUpdateCompositionAnd.ToPointer(),
            Where: []components.GroupUpdateWhere{
                components.GroupUpdateWhere{
                    ObjectType: ruddergo.String("node"),
                    Attribute: ruddergo.String("OS"),
                    Comparator: ruddergo.String("eq"),
                    Value: ruddergo.String("Linux"),
                },
            },
        },
    }
    ctx := context.Background()
    res, err := s.Groups.UpdateGroup(ctx, groupID, groupUpdate)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `groupID`                                                        | *string*                                                         | :heavy_check_mark:                                               | Id of the group                                                  | 9a1773c9-0889-40b6-be89-f6504443ac1b                             |
| `groupUpdate`                                                    | [components.GroupUpdate](../../models/components/groupupdate.md) | :heavy_check_mark:                                               | N/A                                                              |                                                                  |


### Response

**[*operations.UpdateGroupResponse](../../models/operations/updategroupresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteGroup

Update detailed information about a group

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
    var groupID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"
    ctx := context.Background()
    res, err := s.Groups.DeleteGroup(ctx, groupID)
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
| `groupID`                                             | *string*                                              | :heavy_check_mark:                                    | Id of the group                                       | 9a1773c9-0889-40b6-be89-f6504443ac1b                  |


### Response

**[*operations.DeleteGroupResponse](../../models/operations/deletegroupresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadGroup

Recompute the content of a group

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
    var groupID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"
    ctx := context.Background()
    res, err := s.Groups.ReloadGroup(ctx, groupID)
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
| `groupID`                                             | *string*                                              | :heavy_check_mark:                                    | Id of the group                                       | 9a1773c9-0889-40b6-be89-f6504443ac1b                  |


### Response

**[*operations.ReloadGroupResponse](../../models/operations/reloadgroupresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
