package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/wchen777/go-microservices/handlers"
)

func main() {

	// create instance of new handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := handlers.NewHello(l)
	ph := handlers.NewProducts(l)
	// gh := handlers.NewGoodbye(l)

	// // create new servemux for handler

	// USING GO STANDARD
	// sm := http.NewServeMux()
	// sm.Handle("/s", ph)
	// sm.Handle("/goodbye", gh)

	// USING GORILLA
	sm := mux.NewRouter()

	// GORILLA TO SEPARATE GET
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/}", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	// start server with our serve mux
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// go routine to run the server and catch error and won't block
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// signal channel
	sigChan := make(chan os.Signal)

	// listen for interrupt and kill signals
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// graceful shutdown to finish work before cutting off after 30 seconds
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
