package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"desc"`
	SKU         string `json:"sku" validate:"required,sku"`
	Price       int    `json:"price" validate:"gt=0"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

var ErrProductNotFound = fmt.Errorf("product does not exist")

type Products []*Product

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNewID()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()

	productList = append(productList, p)
}

func UpdateProduct(p *Product) error {
	_, idx, err := fetchProductByID(p.ID)
	if err != nil {
		return err
	}

	productList[idx] = p
	return nil
}

func (p *Products) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

func (p *Product) Validate() error {
	vd := validator.New()
	vd.RegisterValidation("sku", validateSKU)
	return vd.Struct(p)
}

func validateSKU(flvl validator.FieldLevel) bool {
	// SKU must be of the format 000-000-000
	regex := regexp.MustCompile(`[0-9]+-[0-9]+-[0-9]+`)
	matches := regex.FindAllString(flvl.Field().String(), -1)

	return len(matches) == 1
}

func getNewID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func fetchProductByID(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

// This is a static list of products acting as a dummy database
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
