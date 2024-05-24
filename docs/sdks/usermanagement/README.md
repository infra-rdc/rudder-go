# UserManagement
(*UserManagement*)

## Overview

**Requires that the `user-management` plugin is installed on the server.**

Manage users settings and configuration file.

### Available Operations

* [AddUser](#adduser) - Add user
* [GetRole](#getrole) - List all roles
* [UpdateUser](#updateuser) - Update user's infos
* [GetUserInfo](#getuserinfo) - List all users
* [ReloadUserConf](#reloaduserconf) - Reload user
* [DeleteUser](#deleteuser) - Delete an user

## AddUser

Add a new user

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
    request := components.Users{
        Username: rudder.String("John Doe"),
        Password: rudder.String("passwdWillBeStoredHashed"),
        Role: []string{
            "user",
        },
    }
    ctx := context.Background()
    res, err := s.UserManagement.AddUser(ctx, request)
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
| `request`                                             | [components.Users](../../models/components/users.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddUserResponse](../../models/operations/adduserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRole

Get all available roles and their rights

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
    res, err := s.UserManagement.GetRole(ctx)
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

**[*operations.GetRoleResponse](../../models/operations/getroleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateUser

Rename, change password (pre-hashed or not) and change permission of an user. If a parameter is empty, it will be ignored.

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
    var username string = "JaneDoe"

    users := components.Users{
        Username: rudder.String("John Doe"),
        Password: rudder.String("passwdWillBeStoredHashed"),
        Role: []string{
            "user",
        },
    }
    ctx := context.Background()
    res, err := s.UserManagement.UpdateUser(ctx, username, users)
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
| `users`                                               | [components.Users](../../models/components/users.md)  | :heavy_check_mark:                                    | N/A                                                   |                                                       |


### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUserInfo

Get the list of all present users and their permissions

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
    res, err := s.UserManagement.GetUserInfo(ctx)
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

**[*operations.GetUserInfoResponse](../../models/operations/getuserinforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadUserConf

Reload the users from the file system, in the configuration file

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
    res, err := s.UserManagement.ReloadUserConf(ctx)
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

**[*operations.ReloadUserConfResponse](../../models/operations/reloaduserconfresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteUser

Delete the user and their permissions

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
    res, err := s.UserManagement.DeleteUser(ctx, username)
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

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
