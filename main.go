package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wchen777/go-microservices/handlers"
)

func main() {

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("hello")
	// 	d, err := ioutil.ReadAll((r.Body))

	// 	// err check w/ response code
	// 	if err != nil {
	// 		http.Error(rw, "error caught", http.StatusBadRequest)
	// 		// rw.WriteHeader(http.StatusBadRequest)
	// 		// rw.Write([]byte("error caught"))
	// 		return
	// 	}

	// 	log.Printf("Data %s\n", d)
	// 	fmt.Fprintf(rw, "Hello %s", d)
	// })

	// create instance of new handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// create new servemux for handler
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// start server with our serve mux
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()

	// http.ListenAndServe(":9090", sm)

}
