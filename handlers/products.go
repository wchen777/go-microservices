package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

// handler for POST request to this endpoint
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	// decode json from body
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	// decode json from body
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
	}

}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle GET
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	// handle POST
	if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
		return
	}

	// handle PUT
	if r.Method == http.MethodPut {
		// expect the id in the URI manually
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 || len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("got id", id)

		p.UpdateProducts(id, rw, r)

	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
