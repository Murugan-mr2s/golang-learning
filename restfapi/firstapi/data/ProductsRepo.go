package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"_"`
	UpdatedOn   string  `json:"_"`
	DeletedOn   string  `json:"_"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productsList = []*Product{

	&Product{
		ID:          1,
		Name:        "iphone",
		Description: " mobile phone",
		Price:       345.0,
		SKU:         "Electronics",
		CreatedOn:   time.Now().UTC().GoString(),
		UpdatedOn:   time.Now().UTC().GoString(),
	},

	&Product{
		ID:          1,
		Name:        "samsung",
		Description: " mobile phone",
		Price:       145.0,
		SKU:         "Electronics",
		CreatedOn:   time.Now().UTC().GoString(),
		UpdatedOn:   time.Now().UTC().GoString(),
	},
}

func GetProducts() Products {
	return productsList
}
