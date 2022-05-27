package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ahmad-Faizan/go-web-api/data"
	_ "github.com/Ahmad-Faizan/go-web-api/docs"
	"github.com/Ahmad-Faizan/go-web-api/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {

	// create a global logger
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// create a global validator
	v := data.NewValidation()

	// get a product handler
	ph := handlers.NewProduct(l, v)

	// define a new server multiplexer
	mux := mux.NewRouter()

	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProduct)

	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidator)

	putRouter := mux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidator)

	deleteRouter := mux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	// handlers for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	swagh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", swagh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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
