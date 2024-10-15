package handlers

import (
	"log"
	"net/http"

	"pixels.com/firstapi/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter) {
	lp := data.GetProducts()
	//data, err := json.Marshal(products)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unble to marshal json", http.StatusInternalServerError)
	}

}
