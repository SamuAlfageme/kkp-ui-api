// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListVMwareCloudDirectorComputePoliciesNoCredentialsReader is a Reader for the ListVMwareCloudDirectorComputePoliciesNoCredentials structure.
type ListVMwareCloudDirectorComputePoliciesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVMwareCloudDirectorComputePoliciesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVMwareCloudDirectorComputePoliciesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVMwareCloudDirectorComputePoliciesNoCredentialsOK creates a ListVMwareCloudDirectorComputePoliciesNoCredentialsOK with default headers values
func NewListVMwareCloudDirectorComputePoliciesNoCredentialsOK() *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK {
	return &ListVMwareCloudDirectorComputePoliciesNoCredentialsOK{}
}

/*
ListVMwareCloudDirectorComputePoliciesNoCredentialsOK describes a response with status code 200, with default header values.

VMwareCloudDirectorComputePolicyList
*/
type ListVMwareCloudDirectorComputePoliciesNoCredentialsOK struct {
	Payload models.VMwareCloudDirectorComputePolicyList
}

// IsSuccess returns true when this list v mware cloud director compute policies no credentials o k response has a 2xx status code
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list v mware cloud director compute policies no credentials o k response has a 3xx status code
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list v mware cloud director compute policies no credentials o k response has a 4xx status code
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list v mware cloud director compute policies no credentials o k response has a 5xx status code
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list v mware cloud director compute policies no credentials o k response a status code equal to that given
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/computepolicies][%d] listVMwareCloudDirectorComputePoliciesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/computepolicies][%d] listVMwareCloudDirectorComputePoliciesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK) GetPayload() models.VMwareCloudDirectorComputePolicyList {
	return o.Payload
}

func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVMwareCloudDirectorComputePoliciesNoCredentialsDefault creates a ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault with default headers values
func NewListVMwareCloudDirectorComputePoliciesNoCredentialsDefault(code int) *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault {
	return &ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v mware cloud director compute policies no credentials default response
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list v mware cloud director compute policies no credentials default response has a 2xx status code
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list v mware cloud director compute policies no credentials default response has a 3xx status code
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list v mware cloud director compute policies no credentials default response has a 4xx status code
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list v mware cloud director compute policies no credentials default response has a 5xx status code
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list v mware cloud director compute policies no credentials default response a status code equal to that given
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/computepolicies][%d] listVMwareCloudDirectorComputePoliciesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/computepolicies][%d] listVMwareCloudDirectorComputePoliciesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}