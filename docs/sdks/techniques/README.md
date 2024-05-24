# Techniques
(*Techniques*)

## Overview

Techniques management

### Available Operations

* [Methods](#methods) - List methods
* [ReloadMethods](#reloadmethods) - Reload methods
* [ListTechniques](#listtechniques) - List all techniques
* [Techniques](#techniques) - Reload techniques
* [ListTechniquesVersions](#listtechniquesversions) - List versions
* [GetTechniqueAllVersion](#gettechniqueallversion) - Technique metadata by ID
* [ListTechniquesDirectives](#listtechniquesdirectives) - List all directives based on a technique
* [DeleteTechnique](#deletetechnique) - Delete technique
* [GetTechniqueAllVersionID](#gettechniqueallversionid) - Technique metadata by version and ID
* [ListTechniqueVersionDirectives](#listtechniqueversiondirectives) - List all directives based on a version of a technique
* [GetTechniquesResources](#gettechniquesresources) - Technique's resources
* [TechniqueRevisions](#techniquerevisions) - Technique's revisions

## Methods

Get all generic methods metadata

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
    res, err := s.Techniques.Methods(ctx)
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

**[*operations.MethodsResponse](../../models/operations/methodsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadMethods

Reload methods metadata from file system

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
    res, err := s.Techniques.ReloadMethods(ctx)
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

**[*operations.ReloadMethodsResponse](../../models/operations/reloadmethodsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListTechniques

List all technique with their versions

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
    res, err := s.Techniques.ListTechniques(ctx)
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

**[*operations.ListTechniquesResponse](../../models/operations/listtechniquesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Techniques

Reload all techniques metadata from file system

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
    res, err := s.Techniques.Techniques(ctx)
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

**[*operations.TechniquesResponse](../../models/operations/techniquesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListTechniquesVersions

List all techniques version

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
    res, err := s.Techniques.ListTechniquesVersions(ctx)
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

**[*operations.ListTechniquesVersionsResponse](../../models/operations/listtechniquesversionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTechniqueAllVersion

Get each Technique's versions with their metadata by ID

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
    var techniqueID string = "userManagement"
    ctx := context.Background()
    res, err := s.Techniques.GetTechniqueAllVersion(ctx, techniqueID)
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
| `techniqueID`                                         | *string*                                              | :heavy_check_mark:                                    | Technique ID                                          | userManagement                                        |


### Response

**[*operations.GetTechniqueAllVersionResponse](../../models/operations/gettechniqueallversionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListTechniquesDirectives

List all directives based on all version of a technique

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
    var techniqueID string = "userManagement"
    ctx := context.Background()
    res, err := s.Techniques.ListTechniquesDirectives(ctx, techniqueID)
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
| `techniqueID`                                         | *string*                                              | :heavy_check_mark:                                    | Technique ID                                          | userManagement                                        |


### Response

**[*operations.ListTechniquesDirectivesResponse](../../models/operations/listtechniquesdirectivesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteTechnique

Delete a technique from technique editor

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
    var techniqueID string = "userManagement"

    var techniqueVersion string = "6.0"
    ctx := context.Background()
    res, err := s.Techniques.DeleteTechnique(ctx, techniqueID, techniqueVersion)
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
| `techniqueID`                                         | *string*                                              | :heavy_check_mark:                                    | Technique ID                                          | userManagement                                        |
| `techniqueVersion`                                    | *string*                                              | :heavy_check_mark:                                    | Technique version                                     | 6.0                                                   |


### Response

**[*operations.DeleteTechniqueResponse](../../models/operations/deletetechniqueresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTechniqueAllVersionID

Get Technique metadata

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
    var techniqueID string = "userManagement"

    var techniqueVersion string = "6.0"
    ctx := context.Background()
    res, err := s.Techniques.GetTechniqueAllVersionID(ctx, techniqueID, techniqueVersion)
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
| `techniqueID`                                         | *string*                                              | :heavy_check_mark:                                    | Technique ID                                          | userManagement                                        |
| `techniqueVersion`                                    | *string*                                              | :heavy_check_mark:                                    | Technique version                                     | 6.0                                                   |


### Response

**[*operations.GetTechniqueAllVersionIDResponse](../../models/operations/gettechniqueallversionidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListTechniqueVersionDirectives

List all directives based on a version of a technique

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
    var techniqueID string = "userManagement"

    var techniqueVersion string = "6.0"
    ctx := context.Background()
    res, err := s.Techniques.ListTechniqueVersionDirectives(ctx, techniqueID, techniqueVersion)
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
| `techniqueID`                                         | *string*                                              | :heavy_check_mark:                                    | Technique ID                                          | userManagement                                        |
| `techniqueVersion`                                    | *string*                                              | :heavy_check_mark:                                    | Technique version                                     | 6.0                                                   |


### Response

**[*operations.ListTechniqueVersionDirectivesResponse](../../models/operations/listtechniqueversiondirectivesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTechniquesResources

Get currently deployed resources of a technique

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
    var techniqueID string = "userManagement"

    var techniqueVersion string = "6.0"
    ctx := context.Background()
    res, err := s.Techniques.GetTechniquesResources(ctx, techniqueID, techniqueVersion)
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
| `techniqueID`                                         | *string*                                              | :heavy_check_mark:                                    | Technique ID                                          | userManagement                                        |
| `techniqueVersion`                                    | *string*                                              | :heavy_check_mark:                                    | Technique version                                     | 6.0                                                   |


### Response

**[*operations.GetTechniquesResourcesResponse](../../models/operations/gettechniquesresourcesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## TechniqueRevisions

Get revisions for given technique

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
    var techniqueID string = "userManagement"

    var techniqueVersion string = "6.0"
    ctx := context.Background()
    res, err := s.Techniques.TechniqueRevisions(ctx, techniqueID, techniqueVersion)
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
| `techniqueID`                                         | *string*                                              | :heavy_check_mark:                                    | Technique ID                                          | userManagement                                        |
| `techniqueVersion`                                    | *string*                                              | :heavy_check_mark:                                    | Technique version                                     | 6.0                                                   |


### Response

**[*operations.TechniqueRevisionsResponse](../../models/operations/techniquerevisionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
