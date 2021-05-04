package handlers

import (
	"log"
	"net/http"

	"github.com/wchen777/go-microservices/data"
)

type Products struct {
	l *log.Logger
}

// instantiate instance of this handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// handler for GET request to this endpoint
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	// // encode products list into json w/ marshal
	// d, err := json.Marshal(lp)

	// use encoder instead of marshal for performance
	err := lp.ToJSON(rw)

	// error check
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// rw.Write(d)
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle GET
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	// handle PUT

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
