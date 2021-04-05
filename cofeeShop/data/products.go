package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedON   string
	UpdatedON   string
	DeletedON   string
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Forthy and milky",
		Price:       2.34,
		SKU:         "abc233",
		CreatedON:   time.Now().UTC().String(),
		UpdatedON:   time.Now().UTC().String(),
	},
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Forthy and milky",
		Price:       2.34,
		SKU:         "abc233",
		CreatedON:   time.Now().UTC().String(),
		UpdatedON:   time.Now().UTC().String(),
	},
}
