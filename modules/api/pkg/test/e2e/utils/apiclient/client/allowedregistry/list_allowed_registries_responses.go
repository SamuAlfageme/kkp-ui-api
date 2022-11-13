// Code generated by go-swagger; DO NOT EDIT.

package allowedregistry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAllowedRegistriesReader is a Reader for the ListAllowedRegistries structure.
type ListAllowedRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllowedRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllowedRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAllowedRegistriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAllowedRegistriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAllowedRegistriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAllowedRegistriesOK creates a ListAllowedRegistriesOK with default headers values
func NewListAllowedRegistriesOK() *ListAllowedRegistriesOK {
	return &ListAllowedRegistriesOK{}
}

/*
ListAllowedRegistriesOK describes a response with status code 200, with default header values.

AllowedRegistry
*/
type ListAllowedRegistriesOK struct {
	Payload []*models.AllowedRegistry
}

// IsSuccess returns true when this list allowed registries o k response has a 2xx status code
func (o *ListAllowedRegistriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list allowed registries o k response has a 3xx status code
func (o *ListAllowedRegistriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list allowed registries o k response has a 4xx status code
func (o *ListAllowedRegistriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list allowed registries o k response has a 5xx status code
func (o *ListAllowedRegistriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list allowed registries o k response a status code equal to that given
func (o *ListAllowedRegistriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAllowedRegistriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistriesOK  %+v", 200, o.Payload)
}

func (o *ListAllowedRegistriesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistriesOK  %+v", 200, o.Payload)
}

func (o *ListAllowedRegistriesOK) GetPayload() []*models.AllowedRegistry {
	return o.Payload
}

func (o *ListAllowedRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllowedRegistriesUnauthorized creates a ListAllowedRegistriesUnauthorized with default headers values
func NewListAllowedRegistriesUnauthorized() *ListAllowedRegistriesUnauthorized {
	return &ListAllowedRegistriesUnauthorized{}
}

/*
ListAllowedRegistriesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListAllowedRegistriesUnauthorized struct {
}

// IsSuccess returns true when this list allowed registries unauthorized response has a 2xx status code
func (o *ListAllowedRegistriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list allowed registries unauthorized response has a 3xx status code
func (o *ListAllowedRegistriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list allowed registries unauthorized response has a 4xx status code
func (o *ListAllowedRegistriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list allowed registries unauthorized response has a 5xx status code
func (o *ListAllowedRegistriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list allowed registries unauthorized response a status code equal to that given
func (o *ListAllowedRegistriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListAllowedRegistriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistriesUnauthorized ", 401)
}

func (o *ListAllowedRegistriesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistriesUnauthorized ", 401)
}

func (o *ListAllowedRegistriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAllowedRegistriesForbidden creates a ListAllowedRegistriesForbidden with default headers values
func NewListAllowedRegistriesForbidden() *ListAllowedRegistriesForbidden {
	return &ListAllowedRegistriesForbidden{}
}

/*
ListAllowedRegistriesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListAllowedRegistriesForbidden struct {
}

// IsSuccess returns true when this list allowed registries forbidden response has a 2xx status code
func (o *ListAllowedRegistriesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list allowed registries forbidden response has a 3xx status code
func (o *ListAllowedRegistriesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list allowed registries forbidden response has a 4xx status code
func (o *ListAllowedRegistriesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list allowed registries forbidden response has a 5xx status code
func (o *ListAllowedRegistriesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list allowed registries forbidden response a status code equal to that given
func (o *ListAllowedRegistriesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListAllowedRegistriesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistriesForbidden ", 403)
}

func (o *ListAllowedRegistriesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistriesForbidden ", 403)
}

func (o *ListAllowedRegistriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAllowedRegistriesDefault creates a ListAllowedRegistriesDefault with default headers values
func NewListAllowedRegistriesDefault(code int) *ListAllowedRegistriesDefault {
	return &ListAllowedRegistriesDefault{
		_statusCode: code,
	}
}

/*
ListAllowedRegistriesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAllowedRegistriesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list allowed registries default response
func (o *ListAllowedRegistriesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list allowed registries default response has a 2xx status code
func (o *ListAllowedRegistriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list allowed registries default response has a 3xx status code
func (o *ListAllowedRegistriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list allowed registries default response has a 4xx status code
func (o *ListAllowedRegistriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list allowed registries default response has a 5xx status code
func (o *ListAllowedRegistriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list allowed registries default response a status code equal to that given
func (o *ListAllowedRegistriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAllowedRegistriesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistries default  %+v", o._statusCode, o.Payload)
}

func (o *ListAllowedRegistriesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistries default  %+v", o._statusCode, o.Payload)
}

func (o *ListAllowedRegistriesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAllowedRegistriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}