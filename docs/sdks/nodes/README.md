# Nodes
(*Nodes*)

## Overview

Nodes management

### Available Operations

* [ListAcceptedNodes](#listacceptednodes) - List managed nodes
* [CreateNodes](#createnodes) - Create one or several new nodes
* [ApplyPolicyAllNodes](#applypolicyallnodes) - Trigger an agent run on all nodes
* [ListPendingNodes](#listpendingnodes) - List pending nodes
* [ChangePendingNodeStatus](#changependingnodestatus) - Update pending Node status
* [GetNodesStatus](#getnodesstatus) - Get nodes acceptation status
* [NodeDetails](#nodedetails) - Get information about a node
* [UpdateNode](#updatenode) - Update node settings and properties
* [DeleteNode](#deletenode) - Delete a node
* [ApplyPolicy](#applypolicy) - Trigger an agent run
* [NodeInheritedProperties](#nodeinheritedproperties) - Get inherited node properties for a node

## ListAcceptedNodes

Get information about the nodes managed by the target server

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/models/operations"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    request := operations.ListAcceptedNodesRequest{
        Include: ruddergo.String("minimal"),
        Query: &components.NodeQuery{
            Composition: components.CompositionAnd.ToPointer(),
            Where: []components.Where{
                components.Where{
                    ObjectType: ruddergo.String("node"),
                    Attribute: ruddergo.String("OS"),
                    Comparator: ruddergo.String("eq"),
                    Value: ruddergo.String("Linux"),
                },
            },
        },
        Where: []components.NodeWhere{
            components.NodeWhere{
                ObjectType: ruddergo.String("node"),
                Attribute: ruddergo.String("OS"),
                Comparator: ruddergo.String("eq"),
                Value: ruddergo.String("Linux"),
            },
        },
        Composition: components.NodeCompositionAnd.ToPointer(),
    }
    ctx := context.Background()
    res, err := s.Nodes.ListAcceptedNodes(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListAcceptedNodesRequest](../../models/operations/listacceptednodesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListAcceptedNodesResponse](../../models/operations/listacceptednodesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateNodes

Use the provided array of node information to create new nodes

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
    request := []components.NodeAdd{
        components.NodeAdd{
            ID: "378740d3-c4a9-4474-8485-478e7e52db52",
            Hostname: "my.node.hostname.local",
            Status: components.NodeAddStatusPending,
            Os: components.Os{
                Type: components.OsTypeFreebsd,
                Name: components.NameKali,
                Version: "9.5",
                FullName: "Debian GNU/Linux 9 (stretch)",
            },
            PolicyServerID: ruddergo.String("root"),
            Machine: &components.NodeAddMachine{
                Type: components.NodeAddTypeVirtual,
                Manufacturer: ruddergo.String("corp inc."),
                SerialNumber: ruddergo.String("ece12459-2b90-49c9-ab1e-72e38f797421"),
            },
            AgentKey: &components.AgentKey{
                Value: "-----BEGIN CERTIFICATE-----
            MIIFqDCC[...]3tALNn
            -----END CERTIFICATE-----",
            },
            Properties: []components.NodeAddProperties{
                components.NodeAddProperties{
                    Name: "datacenter",
                    Value: "AMS2",
                },
            },
            IPAddresses: []string{
                "192.168.180.90",
            },
            Timezone: &components.Timezone{
                Name: "CEST",
                Offset: "+0200",
            },
        },
    }
    ctx := context.Background()
    res, err := s.Nodes.CreateNodes(ctx, request)
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
| `request`                                             | [[]components.NodeAdd](../../.md)                     | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateNodesResponse](../../models/operations/createnodesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ApplyPolicyAllNodes

This API allows to trigger an agent run on all nodes. Response contains a json stating if agent could be started on each node, but not if the run went fine and do not display any output from it. You can see the result of the run in Rudder web interface or in the each agent logs.

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
    res, err := s.Nodes.ApplyPolicyAllNodes(ctx)
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

**[*operations.ApplyPolicyAllNodesResponse](../../models/operations/applypolicyallnodesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListPendingNodes

Get information about the nodes pending acceptation

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/models/operations"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    request := operations.ListPendingNodesRequest{
        Include: ruddergo.String("minimal"),
        Query: &components.NodeQuery{
            Composition: components.CompositionAnd.ToPointer(),
            Where: []components.Where{
                components.Where{
                    ObjectType: ruddergo.String("node"),
                    Attribute: ruddergo.String("OS"),
                    Comparator: ruddergo.String("eq"),
                    Value: ruddergo.String("Linux"),
                },
            },
        },
        Where: []components.NodeWhere{
            components.NodeWhere{
                ObjectType: ruddergo.String("node"),
                Attribute: ruddergo.String("OS"),
                Comparator: ruddergo.String("eq"),
                Value: ruddergo.String("Linux"),
            },
        },
        Composition: components.NodeCompositionAnd.ToPointer(),
    }
    ctx := context.Background()
    res, err := s.Nodes.ListPendingNodes(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListPendingNodesRequest](../../models/operations/listpendingnodesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListPendingNodesResponse](../../models/operations/listpendingnodesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ChangePendingNodeStatus

Accept or refuse a pending node

### Example Usage

```go
package main

import(
	ruddergo "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/operations"
	"context"
	"log"
)

func main() {
    s := ruddergo.New(
        ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var nodeID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"

    var requestBody *operations.ChangePendingNodeStatusRequestBody = &operations.ChangePendingNodeStatusRequestBody{
        Status: operations.ChangePendingNodeStatusStatusAccepted.ToPointer(),
    }
    ctx := context.Background()
    res, err := s.Nodes.ChangePendingNodeStatus(ctx, nodeID, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                       | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     | Example                                                                                                         |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                           | :heavy_check_mark:                                                                                              | The context to use for the request.                                                                             |                                                                                                                 |
| `nodeID`                                                                                                        | *string*                                                                                                        | :heavy_check_mark:                                                                                              | Id of the target node                                                                                           | 9a1773c9-0889-40b6-be89-f6504443ac1b                                                                            |
| `requestBody`                                                                                                   | [*operations.ChangePendingNodeStatusRequestBody](../../models/operations/changependingnodestatusrequestbody.md) | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |                                                                                                                 |


### Response

**[*operations.ChangePendingNodeStatusResponse](../../models/operations/changependingnodestatusresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetNodesStatus

Get acceptation status (pending, accepted, deleted, unknown) of a list of nodes

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
    var ids string = "8403353b-42c4-46f5-a08d-bc77a1a0bad9,root"
    ctx := context.Background()
    res, err := s.Nodes.GetNodesStatus(ctx, ids)
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
| `ids`                                                 | *string*                                              | :heavy_check_mark:                                    | Comma separated list of node Ids                      | 8403353b-42c4-46f5-a08d-bc77a1a0bad9,root             |


### Response

**[*operations.GetNodesStatusResponse](../../models/operations/getnodesstatusresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## NodeDetails

Get details about a given node

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

    var include *string = ruddergo.String("minimal")
    ctx := context.Background()
    res, err := s.Nodes.NodeDetails(ctx, nodeID, include)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `nodeID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Id of the target node                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | 9a1773c9-0889-40b6-be89-f6504443ac1b                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| `include`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don't provide a base level, they will be added to `default` level, so if you only want os details, use `minimal,os` as the value for this parameter.<br/>* **minimal** includes: `id`, `hostname` and `status`<br/>* **default** includes **minimal** plus `architectureDescription`, `description`, `ipAddresses`, `lastRunDate`, `lastInventoryDate`, `machine`, `os`, `managementTechnology`, `policyServerId`, `properties` (be careful! Only node own properties, if you also need inherited properties, look at the dedicated `/nodes/{id}/inheritedProperties` endpoint), `policyMode `, `ram` and `timezone`<br/>* **full** includes: **default** plus `accounts`, `bios`, `controllers`, `environmentVariables`, `fileSystems`, `managementTechnologyDetails`, `memories`, `networkInterfaces`, `ports`, `processes`, `processors`, `slots`, `software`, `softwareUpdate`, `sound`, `storage`, `videos` and `virtualMachines` | minimal                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |


### Response

**[*operations.NodeDetailsResponse](../../models/operations/nodedetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateNode

Update node settings and properties

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
    var nodeID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"

    var nodeSettings *components.NodeSettings = &components.NodeSettings{
        Properties: []components.NodeSettingsProperties{
            components.NodeSettingsProperties{
                Name: "datacenter",
                Value: "AMS2",
            },
        },
        PolicyMode: components.NodeSettingsPolicyModeAudit.ToPointer(),
        State: components.NodeSettingsStateEnabled.ToPointer(),
        AgentKey: &components.AgentKey{
            Value: "-----BEGIN CERTIFICATE-----
        MIIFqDCC[...]3tALNn
        -----END CERTIFICATE-----",
        },
    }
    ctx := context.Background()
    res, err := s.Nodes.UpdateNode(ctx, nodeID, nodeSettings)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |                                                                     |
| `nodeID`                                                            | *string*                                                            | :heavy_check_mark:                                                  | Id of the target node                                               | 9a1773c9-0889-40b6-be89-f6504443ac1b                                |
| `nodeSettings`                                                      | [*components.NodeSettings](../../models/components/nodesettings.md) | :heavy_minus_sign:                                                  | N/A                                                                 |                                                                     |


### Response

**[*operations.UpdateNodeResponse](../../models/operations/updatenoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteNode

Remove a node from the Rudder server. It won't be managed anymore.

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
    var nodeID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"

    var mode *components.NodeDeleteMode = components.NodeDeleteModeMove.ToPointer()
    ctx := context.Background()
    res, err := s.Nodes.DeleteNode(ctx, nodeID, mode)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |                                                                                 |
| `nodeID`                                                                        | *string*                                                                        | :heavy_check_mark:                                                              | Id of the target node                                                           | 9a1773c9-0889-40b6-be89-f6504443ac1b                                            |
| `mode`                                                                          | [*components.NodeDeleteMode](../../models/components/nodedeletemode.md)         | :heavy_minus_sign:                                                              | Deletion mode to use, either move to trash ('move', default) or erase ('erase') | move                                                                            |


### Response

**[*operations.DeleteNodeResponse](../../models/operations/deletenoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ApplyPolicy

This API allows to trigger an agent run on the target node. Response is a stream of the actual agent run on the node.

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
    res, err := s.Nodes.ApplyPolicy(ctx, nodeID)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentOutput != nil {
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

**[*operations.ApplyPolicyResponse](../../models/operations/applypolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## NodeInheritedProperties

This API returns all node properties for a node, including group inherited ones.

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
    res, err := s.Nodes.NodeInheritedProperties(ctx, nodeID)
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

**[*operations.NodeInheritedPropertiesResponse](../../models/operations/nodeinheritedpropertiesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
