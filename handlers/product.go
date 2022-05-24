package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

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

	// Handle the POST verb on the endpoint
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	// Handle the PUT verb on the endpoint
	if r.Method == http.MethodPut {
		p.updateProduct(w, r)
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

func (p *Product) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	// read product from the request body
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		p.l.Println("cannot read the request body as product", err)
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)
}

func (p *Product) updateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")

	// read product id from the request url
	id, err := parseID(r.URL.Path)
	if err != nil {
		p.l.Println("unable to parse ID from request", err)
		http.Error(w, "Unable to parse ID", http.StatusBadRequest)
		return
	}

	// read product from request body
	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		p.l.Println("cannot read the request body as product", err)
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

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

func parseID(path string) (int, error) {
	pt := regexp.MustCompile(`/([0-9]+)`)
	mch := pt.FindAllStringSubmatch(path, -1)

	if len(mch) != 1 {
		return -1, fmt.Errorf("invalid URI, more than one ID")
	}

	if len(mch[0]) != 2 {
		return -1, fmt.Errorf("invalid URI, more than one capture group")
	}

	idRaw := mch[0][1]
	return strconv.Atoi(idRaw)
}
