# System
(*System*)

## Overview

Internal components and administration

### Available Operations

* [ListArchives](#listarchives) - List archives
* [CreateArchive](#createarchive) - Create an archive
* [RestoreArchive](#restorearchive) - Restore an archive
* [GetZipArchive](#getziparchive) - Get an archive as a ZIP
* [GetHealthcheckResult](#gethealthcheckresult) - Get healthcheck
* [GetSystemInfo](#getsysteminfo) - Get server information
* [PurgeSoftware](#purgesoftware) - Trigger batch for cleaning unreferenced software
* [RegeneratePolicies](#regeneratepolicies) - Trigger a new policy generation
* [ReloadAll](#reloadall) - Reload both techniques and dynamic groups
* [ReloadGroups](#reloadgroups) - Reload dynamic groups
* [ReloadTechniques](#reloadtechniques) - Reload techniques
* [GetStatus](#getstatus) - Get server status
* [UpdatePolicies](#updatepolicies) - Trigger update of policies

## ListArchives

List configuration archives

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
    var archiveKind components.ArchiveKind = components.ArchiveKindFull
    ctx := context.Background()
    res, err := s.System.ListArchives(ctx, archiveKind)
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
| `archiveKind`                                                    | [components.ArchiveKind](../../models/components/archivekind.md) | :heavy_check_mark:                                               | Type of archive to make                                          | full                                                             |


### Response

**[*operations.ListArchivesResponse](../../models/operations/listarchivesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateArchive

Create new archive of the given kind

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
    var archiveKind components.ArchiveKind = components.ArchiveKindFull
    ctx := context.Background()
    res, err := s.System.CreateArchive(ctx, archiveKind)
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
| `archiveKind`                                                    | [components.ArchiveKind](../../models/components/archivekind.md) | :heavy_check_mark:                                               | Type of archive to make                                          | full                                                             |


### Response

**[*operations.CreateArchiveResponse](../../models/operations/createarchiveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RestoreArchive

Restore an archive of the given kind for the given moment

### Example Usage

```go
package main

import(
	rudder "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/models/operations"
	"context"
	"log"
)

func main() {
    s := rudder.New(
        rudder.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var archiveKind components.ArchiveKind = components.ArchiveKindFull

    var archiveRestoreKind operations.ArchiveRestoreKind = operations.ArchiveRestoreKindLatestCommit
    ctx := context.Background()
    res, err := s.System.RestoreArchive(ctx, archiveKind, archiveRestoreKind)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                       | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     | Example                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                           | :heavy_check_mark:                                                                                                              | The context to use for the request.                                                                                             |                                                                                                                                 |
| `archiveKind`                                                                                                                   | [components.ArchiveKind](../../models/components/archivekind.md)                                                                | :heavy_check_mark:                                                                                                              | Type of archive to make                                                                                                         | full                                                                                                                            |
| `archiveRestoreKind`                                                                                                            | [operations.ArchiveRestoreKind](../../models/operations/archiverestorekind.md)                                                  | :heavy_check_mark:                                                                                                              | What archive to restore (latest archive, latest commit in configuration repository, or archive with ID as given by listArchive) | latestCommit                                                                                                                    |


### Response

**[*operations.RestoreArchiveResponse](../../models/operations/restorearchiveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetZipArchive

Get an archive of the given kind as a zip

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
    var archiveKind components.ArchiveKind = components.ArchiveKindFull

    var commitID string = "<value>"
    ctx := context.Background()
    res, err := s.System.GetZipArchive(ctx, archiveKind, commitID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `archiveKind`                                                    | [components.ArchiveKind](../../models/components/archivekind.md) | :heavy_check_mark:                                               | Type of archive to make                                          | full                                                             |
| `commitID`                                                       | *string*                                                         | :heavy_check_mark:                                               | commit ID of the archive to get as a ZIP file                    |                                                                  |


### Response

**[*operations.GetZipArchiveResponse](../../models/operations/getziparchiveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetHealthcheckResult

Run and get the result of all checks

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
    res, err := s.System.GetHealthcheckResult(ctx)
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

**[*operations.GetHealthcheckResultResponse](../../models/operations/gethealthcheckresultresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSystemInfo

Get information about the server version

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
    res, err := s.System.GetSystemInfo(ctx)
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

**[*operations.GetSystemInfoResponse](../../models/operations/getsysteminforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PurgeSoftware

Start the software cleaning batch asynchronously.

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
    res, err := s.System.PurgeSoftware(ctx)
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

**[*operations.PurgeSoftwareResponse](../../models/operations/purgesoftwareresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RegeneratePolicies

Trigger a full policy generation

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
    res, err := s.System.RegeneratePolicies(ctx)
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

**[*operations.RegeneratePoliciesResponse](../../models/operations/regeneratepoliciesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadAll

Reload both techniques and dynamic groups

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
    res, err := s.System.ReloadAll(ctx)
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

**[*operations.ReloadAllResponse](../../models/operations/reloadallresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadGroups

Reload dynamic groups

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
    res, err := s.System.ReloadGroups(ctx)
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

**[*operations.ReloadGroupsResponse](../../models/operations/reloadgroupsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReloadTechniques

Reload techniques from local technique library

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
    res, err := s.System.ReloadTechniques(ctx)
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

**[*operations.ReloadTechniquesResponse](../../models/operations/reloadtechniquesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetStatus

Get information about current server status

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
    res, err := s.System.GetStatus(ctx)
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

**[*operations.GetStatusResponse](../../models/operations/getstatusresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdatePolicies

Update configuration policies if needed

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
    res, err := s.System.UpdatePolicies(ctx)
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

**[*operations.UpdatePoliciesResponse](../../models/operations/updatepoliciesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
