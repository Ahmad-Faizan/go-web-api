package handlers

import (
	"net/http"

	"github.com/Ahmad-Faizan/go-web-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product by id
//
// responses:
//	204: noContentResponse
//  404: errorResponse
//  500: errorResponse

// Delete handles DELETE requests and removes products from the database
func (p *Product) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] record id does not exist")

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
