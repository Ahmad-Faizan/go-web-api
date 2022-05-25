package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Ahmad-Faizan/go-web-api/data"
	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

type KeyProduct struct{}

// NewProduct creates a new product handler with the specified logger
func NewProduct(l *log.Logger) *Product {
	return &Product{l: l}
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")

	// Fetch the data from data access layer
	pl := data.GetProducts()

	// return the response in JSON format
	err := pl.ToJSON(w)
	if err != nil {
		p.l.Println(err)
		http.Error(w, "Unable to return products", http.StatusInternalServerError)
	}
}

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	// read product from the request body
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(prod)
}

func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")

	// read product id from the request url
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("unable to parse ID from request", err)
		http.Error(w, "Unable to parse ID", http.StatusBadRequest)
		return
	}

	// read product from the request body
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	prod.ID = id
	prod.UpdatedOn = time.Now().UTC().String()
	err = data.UpdateProduct(prod)

	if err == data.ErrProductNotFound {
		p.l.Println("product not found", err)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		p.l.Println("product not found", err)
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

func (p Product) MiddlewareProductValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// read product from request body
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("cannot read the request body as product", err)
			http.Error(w, "Unable to parse request body", http.StatusBadRequest)
			return
		}

		// add the product to the request context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
