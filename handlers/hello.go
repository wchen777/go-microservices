package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// create instance of handler given a logger object
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println()

	d, err := ioutil.ReadAll((r.Body))

	// err check w/ response code
	if err != nil {
		http.Error(rw, "error caught", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("error caught"))
		return
	}

	log.Printf("Data %s\n", d)
	fmt.Fprintf(rw, "Hello %s", d)
}
