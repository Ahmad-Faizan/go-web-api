package handlers

import (
	"log"
	"net/http"
)

// Goodbye is a simple handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye creates a new goodbye handler with the specified logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{
		l: l,
	}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Running in Goodbye handler")
}
