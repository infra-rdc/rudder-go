# SecretManagement
(*SecretManagement*)

## Overview

**Requires that the `secret-management` plugin is installed on the server.**

Manage secrets variables.

### Available Operations

* [GetAllSecrets](#getallsecrets) - List all secrets
* [UpdateSecret](#updatesecret) - Update a secret
* [AddSecret](#addsecret) - Create a secret
* [GetSecret](#getsecret) - Get one secret
* [DeleteSecret](#deletesecret) - Delete a secret

## GetAllSecrets

Get the list of all secrets without their value

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
    res, err := s.SecretManagement.GetAllSecrets(ctx)
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

**[*operations.GetAllSecretsResponse](../../models/operations/getallsecretsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateSecret

Update a secret and override the value, the name cannot be overridden

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
    request := components.Secrets{
        Name: ruddergo.String("secret-password"),
        Description: ruddergo.String("Password of my super secret user account"),
        Value: ruddergo.String("nj-k;EO32!kFWewn2Nk,u"),
    }
    ctx := context.Background()
    res, err := s.SecretManagement.UpdateSecret(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Secrets](../../models/components/secrets.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.UpdateSecretResponse](../../models/operations/updatesecretresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AddSecret

Add a secret

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
    request := components.Secrets{
        Name: ruddergo.String("secret-password"),
        Description: ruddergo.String("Password of my super secret user account"),
        Value: ruddergo.String("nj-k;EO32!kFWewn2Nk,u"),
    }
    ctx := context.Background()
    res, err := s.SecretManagement.AddSecret(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Secrets](../../models/components/secrets.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.AddSecretResponse](../../models/operations/addsecretresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSecret

Get one secret by its unique name

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
    var name string = "<value>"
    ctx := context.Background()
    res, err := s.SecretManagement.GetSecret(ctx, name)
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
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | Unique name of the secret                             |


### Response

**[*operations.GetSecretResponse](../../models/operations/getsecretresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteSecret

Remove the secret by its unique name

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
    var name string = "<value>"
    ctx := context.Background()
    res, err := s.SecretManagement.DeleteSecret(ctx, name)
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
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | Unique name of the secret                             |


### Response

**[*operations.DeleteSecretResponse](../../models/operations/deletesecretresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
