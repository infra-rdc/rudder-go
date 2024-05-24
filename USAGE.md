<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	ruddergo "github.com/infra-rdc/rudder-go"
	"log"
)

func main() {
	s := ruddergo.New(
		ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
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
<!-- End SDK Example Usage [usage] -->