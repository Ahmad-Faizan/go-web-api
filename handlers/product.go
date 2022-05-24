package handlers

import (
	"log"
	"net/http"

	"github.com/Ahmad-Faizan/go-web-api/data"
)

type Product struct {
	l *log.Logger
}

// NewProduct creates a new product handler with the specified logger
func NewProduct(l *log.Logger) *Product {
	return &Product{l: l}
}

// ServeHTTP satisfies the http.Handler interface for Products
func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle the GET verb on the endpoint
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// To catch all other HTTP verbs
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(w http.ResponseWriter, r *http.Request) {
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
