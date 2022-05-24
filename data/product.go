package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	SKU         string `json:"sku"`
	Price       int    `json:"price"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

type Products []*Product

func GetProducts() Products {
	return productList
}

func (p *Products) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

// This is a static list of products acting as a dummy database
var productList = []*Product{
	{
		ID:          100,
		Name:        "Latte",
		Description: "A savoury drink with a dash of coffee on milk",
		SKU:         "123-456",
		Price:       5,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          101,
		Name:        "Capuccino",
		Description: "A delicious dose of coffee and milk",
		SKU:         "789-456",
		Price:       8,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          102,
		Name:        "Espresso",
		Description: "A sleepless drink with coffee without milk",
		SKU:         "456-789",
		Price:       4,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
