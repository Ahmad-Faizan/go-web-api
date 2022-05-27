package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/Ahmad-Faizan/go-web-api/client"
	"github.com/Ahmad-Faizan/go-web-api/client/products"
)

func TestClient(t *testing.T) {
	tcfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, tcfg)

	params := products.NewGetProductsParams()
	prod, err := c.Products.GetProducts(params)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
}
