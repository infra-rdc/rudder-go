# ChangeRequests
(*ChangeRequests*)

## Overview

**Requires that the `changes-validation` plugin is installed on the server.**

Manage change requests.

### Available Operations

* [ListChangeRequests](#listchangerequests) - List all change requests
* [ChangeRequestDetails](#changerequestdetails) - Get a change request details
* [DeclineChangeRequest](#declinechangerequest) - Decline a request details
* [UpdateChangeRequest](#updatechangerequest) - Update a request details
* [AcceptChangeRequest](#acceptchangerequest) - Accept a request details
* [ListUsers](#listusers) - List user
* [SaveWorkflowUser](#saveworkflowuser) - Update validated user list
* [RemoveValidatedUser](#removevalidateduser) - Remove an user from validated user list

## ListChangeRequests

List all change requests

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
    res, err := s.ChangeRequests.ListChangeRequests(ctx)
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

**[*operations.ListChangeRequestsResponse](../../models/operations/listchangerequestsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ChangeRequestDetails

Get a change request details

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
    var changeRequestID int64 = 37
    ctx := context.Background()
    res, err := s.ChangeRequests.ChangeRequestDetails(ctx, changeRequestID)
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
| `changeRequestID`                                     | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   | 37                                                    |


### Response

**[*operations.ChangeRequestDetailsResponse](../../models/operations/changerequestdetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeclineChangeRequest

Refuse a change request

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
    var changeRequestID int64 = 37
    ctx := context.Background()
    res, err := s.ChangeRequests.DeclineChangeRequest(ctx, changeRequestID)
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
| `changeRequestID`                                     | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   | 37                                                    |


### Response

**[*operations.DeclineChangeRequestResponse](../../models/operations/declinechangerequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateChangeRequest

Update a change request

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
    var changeRequestID int64 = 37

    requestBody := operations.UpdateChangeRequestRequestBody{}
    ctx := context.Background()
    res, err := s.ChangeRequests.UpdateChangeRequest(ctx, changeRequestID, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `changeRequestID`                                                                                      | *int64*                                                                                                | :heavy_check_mark:                                                                                     | N/A                                                                                                    | 37                                                                                                     |
| `requestBody`                                                                                          | [operations.UpdateChangeRequestRequestBody](../../models/operations/updatechangerequestrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |


### Response

**[*operations.UpdateChangeRequestResponse](../../models/operations/updatechangerequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AcceptChangeRequest

Accept a change request

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
    var changeRequestID int64 = 37

    requestBody := operations.AcceptChangeRequestRequestBody{
        Status: operations.StatusDeployed.ToPointer(),
    }
    ctx := context.Background()
    res, err := s.ChangeRequests.AcceptChangeRequest(ctx, changeRequestID, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `changeRequestID`                                                                                      | *int64*                                                                                                | :heavy_check_mark:                                                                                     | N/A                                                                                                    | 37                                                                                                     |
| `requestBody`                                                                                          | [operations.AcceptChangeRequestRequestBody](../../models/operations/acceptchangerequestrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |


### Response

**[*operations.AcceptChangeRequestResponse](../../models/operations/acceptchangerequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUsers

List all validated and unvalidated users

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
    res, err := s.ChangeRequests.ListUsers(ctx)
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

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SaveWorkflowUser

Add and remove user from validated users

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
    request := operations.SaveWorkflowUserRequestBody{
        ValidatedUsers: []string{
            "John Do",
        },
    }
    ctx := context.Background()
    res, err := s.ChangeRequests.SaveWorkflowUser(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.SaveWorkflowUserRequestBody](../../models/operations/saveworkflowuserrequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.SaveWorkflowUserResponse](../../models/operations/saveworkflowuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveValidatedUser

The user is again subject to workflow validation

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
    var username string = "JaneDoe"
    ctx := context.Background()
    res, err := s.ChangeRequests.RemoveValidatedUser(ctx, username)
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
| `username`                                            | *string*                                              | :heavy_check_mark:                                    | Username of an user (unique)                          | JaneDoe                                               |


### Response

**[*operations.RemoveValidatedUserResponse](../../models/operations/removevalidateduserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
