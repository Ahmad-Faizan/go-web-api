package handlers

import (
	"net/http"
	"time"

	"github.com/Ahmad-Faizan/go-web-api/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Update the details of a product
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	// read id from URL
	id := getProductID(r)

	// read product from the request body
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	prod.ID = id
	prod.UpdatedOn = time.Now().UTC().String()

	p.l.Println("[DEBUG] updating product id", prod.ID)
	w.Header().Add("Content-Type", "application/json")

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
