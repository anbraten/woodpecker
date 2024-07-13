package woodpeckergo_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	woodpeckergo "go.woodpecker-ci.org/woodpecker/v2/woodpecker-go"
)

func TestClient(t *testing.T) {
	client, err := woodpeckergo.New("http://localhost:8080")
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestClient_canCall() {
	// custom HTTP client
	hc := http.Client{}

	// with a raw http.Response
	{
		c, err := client.NewClient("http://localhost:1234", client.WithHTTPClient(&hc))
		if err != nil {
			log.Fatal(err)
		}

		resp, err := c.GetClient(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Expected HTTP 200 but received %d", resp.StatusCode)
		}
	}

	// or to get a struct with the parsed response body
	{
		c, err := client.NewClientWithResponses("http://localhost:1234", client.WithHTTPClient(&hc))
		if err != nil {
			log.Fatal(err)
		}

		resp, err := c.GetClientWithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode() != http.StatusOK {
			log.Fatalf("Expected HTTP 200 but received %d", resp.StatusCode())
		}

		fmt.Printf("resp.JSON200: %v\n", resp.JSON200)
	}

}
