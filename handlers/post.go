package handlers

import (
	"net/http"

	"github.com/Ahmad-Faizan/go-web-api/data"
)

// swagger:route POST /products products createProduct
// Creates a new product
//
// responses:
//	201: noContentResponse
//  422: errorValidation
//  500: errorResponse

// AddProduct handles POST requests to add new products
func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {

	// read product from the request body
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)

	data.AddProduct(prod)
}
