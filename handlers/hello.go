package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the specified logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{
		l: l,
	}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Running Hello handler")

	// read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("error in reading body of the request", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// write the body of the request to the response
	fmt.Fprintf(w, "Hello %s", body)
}
