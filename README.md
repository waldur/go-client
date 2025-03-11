# go-client
Go client for Waldur MasterMind generated from OpenAPI schema

Example usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/waldur/go-client"
)

func main() {
	auth, err := client.NewTokenAuth("API_TOKEN")
	if err != nil {
		log.Fatal(err)
	}

	hc := http.Client{}

	c, err := client.NewClientWithResponses("API_URL", client.WithHTTPClient(&hc), client.WithRequestEditorFn(auth.Intercept))

	ctx := context.Background()

	resp, err := c.UsersMeRetrieveWithResponse(ctx, &client.UsersMeRetrieveParams{})
	if err != nil {
		log.Fatalf("Error calling API: %v", err)
	}

	fmt.Printf("Response: %+v\n", *resp.JSON200.FirstName)

}
```