package data

import (
	"fmt"
	"time"
)

// ErrProductNotFound is an error when the product is unavailable in the database
var ErrProductNotFound = fmt.Errorf("product does not exist")

// Products define a slice of Product
type Products []*Product

// GetProducts returns all products from the database
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product matching by ID
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	_, idx, err := fetchProductByID(id)
	if err != nil {
		return nil, err
	}

	return productList[idx], nil
}

// AddProduct adds a new product to the database
func AddProduct(p Product) {
	p.ID = getNewID()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()

	productList = append(productList, &p)
}

// UpdateProduct updates an existing product
// If a product is not found this function returns a ProductNotFound error
func UpdateProduct(p Product) error {
	_, idx, err := fetchProductByID(p.ID)
	if err != nil {
		return err
	}

	productList[idx] = &p
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	_, idx, err := fetchProductByID(id)
	if err != nil {
		return ErrProductNotFound
	}

	productList = append(productList[:idx], productList[idx+1:]...)
	return nil
}

// getNewID returns a new ID for a product
func getNewID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// fetchProductByID returns a product if the id matches in the database
// If a product is not found, it returns -1 as the index and a ProductNotFound error
func fetchProductByID(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}
