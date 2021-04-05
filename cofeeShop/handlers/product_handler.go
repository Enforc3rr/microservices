package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nicholasjackson/building-microservices-youtube/product-api/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServerHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to unmarshal Json", http.StatusInternalServerError)
	}
}
