package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ahmad-Faizan/go-web-api/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// create a global logger
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// get a product handler
	ph := handlers.NewProduct(l)

	// define a new server multiplexer
	mux := mux.NewRouter()

	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidator)

	putRouter := mux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidator)

	// define the server
	srv := http.Server{
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		Addr:         "127.0.0.1:9090",
		Handler:      mux,
		ErrorLog:     l,
	}

	// start the server in a new goroutine
	go func() {
		l.Println("Starting server on", fmt.Sprint(srv.Addr))
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			l.Fatalf("error in starting server : %s", err)
		}
	}()

	// handle graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	sig := <-shutdown
	l.Printf("%s received, shutting down gracefully", sig)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	srv.Shutdown(ctx)
}
