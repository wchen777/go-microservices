package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

// create instance of handler given a logger object
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("goodbye"))
}
