# DataSources
(*DataSources*)

## Overview

**Requires that the `datasources` plugin is installed on the server.**

Data sources plugin configuration.

### Available Operations

* [GetAllDataSources](#getalldatasources) - List all data sources
* [ReloadAllDatasourcesAllNodes](#reloadalldatasourcesallnodes) - Update properties from data sources
* [ReloadOneDatasourceAllNodes](#reloadonedatasourceallnodes) - Update properties from data sources
* [GetDataSource](#getdatasource) - Get data source configuration
* [DeleteDataSource](#deletedatasource) - Delete a data source
* [ReloadAllDatasourcesOneNode](#reloadalldatasourcesonenode) - Update properties for one node from all data sources
* [ReloadOneDatasourceOneNode](#reloadonedatasourceonenode) - Update properties for one node from a data source

## GetAllDataSources

Get the configuration of all present data sources

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
    res, err := s.DataSources.GetAllDataSources(ctx)
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

**[*operations.GetAllDataSourcesResponse](../../models/operations/getalldatasourcesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadAllDatasourcesAllNodes

Update properties from all data source on all nodes. The call is asynchronous.

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
    res, err := s.DataSources.ReloadAllDatasourcesAllNodes(ctx)
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

**[*operations.ReloadAllDatasourcesAllNodesResponse](../../models/operations/reloadalldatasourcesallnodesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadOneDatasourceAllNodes

Update properties from all data source on all nodes. The call is asynchronous.

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
    var datasourceID string = "test-data-source"
    ctx := context.Background()
    res, err := s.DataSources.ReloadOneDatasourceAllNodes(ctx, datasourceID)
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
| `datasourceID`                                        | *string*                                              | :heavy_check_mark:                                    | Id of the data source                                 | test-data-source                                      |


### Response

**[*operations.ReloadOneDatasourceAllNodesResponse](../../models/operations/reloadonedatasourceallnodesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDataSource

Get the configuration of a data source

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
    var datasourceID string = "test-data-source"
    ctx := context.Background()
    res, err := s.DataSources.GetDataSource(ctx, datasourceID)
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
| `datasourceID`                                        | *string*                                              | :heavy_check_mark:                                    | Id of the data source                                 | test-data-source                                      |


### Response

**[*operations.GetDataSourceResponse](../../models/operations/getdatasourceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteDataSource

Delete a data source configuration

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
    var datasourceID string = "test-data-source"
    ctx := context.Background()
    res, err := s.DataSources.DeleteDataSource(ctx, datasourceID)
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
| `datasourceID`                                        | *string*                                              | :heavy_check_mark:                                    | Id of the data source                                 | test-data-source                                      |


### Response

**[*operations.DeleteDataSourceResponse](../../models/operations/deletedatasourceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadAllDatasourcesOneNode

Update properties from all data sources on one nodes. The call is asynchronous.

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
    res, err := s.DataSources.ReloadAllDatasourcesOneNode(ctx, nodeID)
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

**[*operations.ReloadAllDatasourcesOneNodeResponse](../../models/operations/reloadalldatasourcesonenoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadOneDatasourceOneNode

Update properties from a data source on one nodes. The call is asynchronous.

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

    var datasourceID string = "test-data-source"
    ctx := context.Background()
    res, err := s.DataSources.ReloadOneDatasourceOneNode(ctx, nodeID, datasourceID)
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
| `datasourceID`                                        | *string*                                              | :heavy_check_mark:                                    | Id of the data source                                 | test-data-source                                      |


### Response

**[*operations.ReloadOneDatasourceOneNodeResponse](../../models/operations/reloadonedatasourceonenoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
