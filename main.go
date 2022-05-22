package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// index handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello handler")

		// read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("error in reading body of the request", err)
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// write the body of the request to the response
		fmt.Fprintf(w, "Hello %s", body)
	})

	// goodbye handler
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Running in Goodbye handler")
	})

	log.Println("Starting server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatalf("error in starting server : %s", err)

}
