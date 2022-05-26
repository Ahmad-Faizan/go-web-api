package handlers

import (
	"context"
	"net/http"

	"github.com/Ahmad-Faizan/go-web-api/data"
)

// MiddlewareProductValidator validates the request and calls the next handler
func (p *Product) MiddlewareProductValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// read product from request body
		prod := &data.Product{}
		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] cannot read the request body as product", err)
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		// Validate the product using validator package
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] error in validating the product", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return
		}

		// add the product to the request context
		ctx := context.WithValue(r.Context(), KeyProduct{}, *prod)
		r = r.WithContext(ctx)

		// call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
