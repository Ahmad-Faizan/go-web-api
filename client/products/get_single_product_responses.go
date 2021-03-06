// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Ahmad-Faizan/go-web-api/models"
)

// GetSingleProductReader is a Reader for the GetSingleProduct structure.
type GetSingleProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSingleProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSingleProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSingleProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSingleProductOK creates a GetSingleProductOK with default headers values
func NewGetSingleProductOK() *GetSingleProductOK {
	return &GetSingleProductOK{}
}

/* GetSingleProductOK describes a response with status code 200, with default header values.

Data structure representing a single product
*/
type GetSingleProductOK struct {
	Payload *models.Product
}

func (o *GetSingleProductOK) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getSingleProductOK  %+v", 200, o.Payload)
}
func (o *GetSingleProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *GetSingleProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleProductNotFound creates a GetSingleProductNotFound with default headers values
func NewGetSingleProductNotFound() *GetSingleProductNotFound {
	return &GetSingleProductNotFound{}
}

/* GetSingleProductNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type GetSingleProductNotFound struct {
	Payload *models.GenericError
}

func (o *GetSingleProductNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getSingleProductNotFound  %+v", 404, o.Payload)
}
func (o *GetSingleProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSingleProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
