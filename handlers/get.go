package handlers

import (
	"net/http"

	"github.com/Ahmad-Faizan/go-web-api/data"
)

// swagger:route GET /products products getProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// GetProducts handles GET requests and returns all current products
func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Get all products")
	w.Header().Add("Content-Type", "application/json")

	// Fetch the product list from database
	pl := data.GetProducts()

	// return the response in JSON format
	err := data.ToJSON(pl, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product ", err)
		return
	}
}

// swagger:route GET /products/{id} products getSingleProduct
// Return a single product from the database
// responses:
//	200: productResponse
//	404: errorResponse

// GetProduct handles GET requests for a single product
func (p *Product) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)
	w.Header().Add("Content-Type", "application/json")

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
		return
	}
}
