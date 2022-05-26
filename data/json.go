package data

import (
	"encoding/json"
	"io"
)

// ToJSON serializes the given interface to an io.Writer in JSON format
func ToJSON(i interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(i)
}

// FromJSON deserializes the JSON string into an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(i)
}
