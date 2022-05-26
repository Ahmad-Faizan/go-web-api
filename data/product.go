package data

import "time"

// Product defines the structure of a sample product
// swagger:model
type Product struct {
	// the id for this product
	//
	// required: false
	// min: 1
	ID int `json:"id"`

	// the name for this product
	//
	// required: true
	// max length : 255
	Name string `json:"name" validate:"required,max=255"`

	// the description for this poduct
	//
	// required: false
	// max length: 1000
	Description string `json:"desc" validate:"max=1000"`

	// the SKU for the product
	//
	// required: true
	// pattern: [0-9]+-[0-9]+-[0-9]+
	SKU string `json:"sku" validate:"required,sku"`

	// the price for the product
	//
	// required: true
	// min: 1
	Price int `json:"price" validate:"required,gt=0"`

	// internal fields for the database
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// productList is a slice of products acting as a dummy database
var productList = []*Product{
	{
		ID:          100,
		Name:        "Latte",
		Description: "A savoury drink with a dash of coffee on milk",
		SKU:         "123-456-001",
		Price:       5,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          101,
		Name:        "Capuccino",
		Description: "A delicious dose of coffee and milk",
		SKU:         "789-456-001",
		Price:       8,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          102,
		Name:        "Espresso",
		Description: "A sleepless drink with coffee without milk",
		SKU:         "456-789-001",
		Price:       4,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
