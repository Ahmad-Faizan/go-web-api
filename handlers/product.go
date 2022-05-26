package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Ahmad-Faizan/go-web-api/data"
	"github.com/gorilla/mux"
)

// Product handler for operating on a product
type Product struct {
	l *log.Logger
	v *data.Validation
}

// KeyProduct is a key used for Product inside the request context
type KeyProduct struct{}

// NewProduct creates a new product handler with the specified logger
func NewProduct(l *log.Logger, v *data.Validation) *Product {
	return &Product{
		l: l,
		v: v,
	}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}
