# Settings
(*Settings*)

## Overview

Server configuration

### Available Operations

* [GetAllSettings](#getallsettings) - List all settings
* [GetAllowedNetworks](#getallowednetworks) - Get allowed networks for a policy server
* [SetAllowedNetworks](#setallowednetworks) - Set allowed networks for a policy server
* [ModifyAllowedNetworks](#modifyallowednetworks) - Modify allowed networks for a policy server
* [GetSetting](#getsetting) - Get the value of a setting
* [ModifySetting](#modifysetting) - Set the value of a setting

## GetAllSettings

Get the current value of all the settings

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
    res, err := s.Settings.GetAllSettings(ctx)
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

**[*operations.GetAllSettingsResponse](../../models/operations/getallsettingsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAllowedNetworks

Get the list of allowed networks for a policy server

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
    var nodeID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"
    ctx := context.Background()
    res, err := s.Settings.GetAllowedNetworks(ctx, nodeID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                       | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ctx`                                                           | [context.Context](https://pkg.go.dev/context#Context)           | :heavy_check_mark:                                              | The context to use for the request.                             |                                                                 |
| `nodeID`                                                        | *string*                                                        | :heavy_check_mark:                                              | Policy server ID for which you want to manage allowed networks. | 9a1773c9-0889-40b6-be89-f6504443ac1b                            |


### Response

**[*operations.GetAllowedNetworksResponse](../../models/operations/getallowednetworksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SetAllowedNetworks

Set the list of allowed networks for a policy server

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
    var nodeID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"

    requestBody := operations.SetAllowedNetworksRequestBody{
        Value: &operations.Value{},
    }
    ctx := context.Background()
    res, err := s.Settings.SetAllowedNetworks(ctx, nodeID, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `nodeID`                                                                                             | *string*                                                                                             | :heavy_check_mark:                                                                                   | Policy server ID for which you want to manage allowed networks.                                      | 9a1773c9-0889-40b6-be89-f6504443ac1b                                                                 |
| `requestBody`                                                                                        | [operations.SetAllowedNetworksRequestBody](../../models/operations/setallowednetworksrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |


### Response

**[*operations.SetAllowedNetworksResponse](../../models/operations/setallowednetworksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ModifyAllowedNetworks

Add or delete allowed networks for a policy server

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
    var nodeID string = "9a1773c9-0889-40b6-be89-f6504443ac1b"

    requestBody := operations.ModifyAllowedNetworksRequestBody{
        AllowedNetworks: &operations.AllowedNetworks{
            Add: []string{
                "192.168.2.0/24",
                "192.168.0.0/16",
            },
            Delete: []string{
                "162.168.1.0/24",
            },
        },
    }
    ctx := context.Background()
    res, err := s.Settings.ModifyAllowedNetworks(ctx, nodeID, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `nodeID`                                                                                                   | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Policy server ID for which you want to manage allowed networks.                                            | 9a1773c9-0889-40b6-be89-f6504443ac1b                                                                       |
| `requestBody`                                                                                              | [operations.ModifyAllowedNetworksRequestBody](../../models/operations/modifyallowednetworksrequestbody.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |


### Response

**[*operations.ModifyAllowedNetworksResponse](../../models/operations/modifyallowednetworksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSetting

Get the current value of a specific setting

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
    var settingID string = "global_policy_mode"
    ctx := context.Background()
    res, err := s.Settings.GetSetting(ctx, settingID)
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
| `settingID`                                           | *string*                                              | :heavy_check_mark:                                    | Id of the setting to set                              | global_policy_mode                                    |


### Response

**[*operations.GetSettingResponse](../../models/operations/getsettingresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ModifySetting

Set the current value of a specific setting

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
    var settingID string = "global_policy_mode"

    requestBody := operations.ModifySettingRequestBody{
        Value: rudder.String("enforce"),
    }
    ctx := context.Background()
    res, err := s.Settings.ModifySetting(ctx, settingID, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `settingID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | Id of the setting to set                                                                   | global_policy_mode                                                                         |
| `requestBody`                                                                              | [operations.ModifySettingRequestBody](../../models/operations/modifysettingrequestbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |


### Response

**[*operations.ModifySettingResponse](../../models/operations/modifysettingresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
