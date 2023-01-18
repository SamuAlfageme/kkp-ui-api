// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListProjectAWSSubnetsParams creates a new ListProjectAWSSubnetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectAWSSubnetsParams() *ListProjectAWSSubnetsParams {
	return &ListProjectAWSSubnetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectAWSSubnetsParamsWithTimeout creates a new ListProjectAWSSubnetsParams object
// with the ability to set a timeout on a request.
func NewListProjectAWSSubnetsParamsWithTimeout(timeout time.Duration) *ListProjectAWSSubnetsParams {
	return &ListProjectAWSSubnetsParams{
		timeout: timeout,
	}
}

// NewListProjectAWSSubnetsParamsWithContext creates a new ListProjectAWSSubnetsParams object
// with the ability to set a context for a request.
func NewListProjectAWSSubnetsParamsWithContext(ctx context.Context) *ListProjectAWSSubnetsParams {
	return &ListProjectAWSSubnetsParams{
		Context: ctx,
	}
}

// NewListProjectAWSSubnetsParamsWithHTTPClient creates a new ListProjectAWSSubnetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectAWSSubnetsParamsWithHTTPClient(client *http.Client) *ListProjectAWSSubnetsParams {
	return &ListProjectAWSSubnetsParams{
		HTTPClient: client,
	}
}

/*
ListProjectAWSSubnetsParams contains all the parameters to send to the API endpoint

	for the list project a w s subnets operation.

	Typically these are written to a http.Request.
*/
type ListProjectAWSSubnetsParams struct {

	// AccessKeyID.
	AccessKeyID *string

	// AssumeRoleARN.
	AssumeRoleARN *string

	// AssumeRoleExternalID.
	AssumeRoleExternalID *string

	// Credential.
	Credential *string

	// SecretAccessKey.
	SecretAccessKey *string

	// VPC.
	VPC *string

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project a w s subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAWSSubnetsParams) WithDefaults() *ListProjectAWSSubnetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project a w s subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAWSSubnetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithTimeout(timeout time.Duration) *ListProjectAWSSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithContext(ctx context.Context) *ListProjectAWSSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithHTTPClient(client *http.Client) *ListProjectAWSSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKeyID adds the accessKeyID to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithAccessKeyID(accessKeyID *string) *ListProjectAWSSubnetsParams {
	o.SetAccessKeyID(accessKeyID)
	return o
}

// SetAccessKeyID adds the accessKeyId to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetAccessKeyID(accessKeyID *string) {
	o.AccessKeyID = accessKeyID
}

// WithAssumeRoleARN adds the assumeRoleARN to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithAssumeRoleARN(assumeRoleARN *string) *ListProjectAWSSubnetsParams {
	o.SetAssumeRoleARN(assumeRoleARN)
	return o
}

// SetAssumeRoleARN adds the assumeRoleARN to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetAssumeRoleARN(assumeRoleARN *string) {
	o.AssumeRoleARN = assumeRoleARN
}

// WithAssumeRoleExternalID adds the assumeRoleExternalID to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithAssumeRoleExternalID(assumeRoleExternalID *string) *ListProjectAWSSubnetsParams {
	o.SetAssumeRoleExternalID(assumeRoleExternalID)
	return o
}

// SetAssumeRoleExternalID adds the assumeRoleExternalId to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetAssumeRoleExternalID(assumeRoleExternalID *string) {
	o.AssumeRoleExternalID = assumeRoleExternalID
}

// WithCredential adds the credential to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithCredential(credential *string) *ListProjectAWSSubnetsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithSecretAccessKey adds the secretAccessKey to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithSecretAccessKey(secretAccessKey *string) *ListProjectAWSSubnetsParams {
	o.SetSecretAccessKey(secretAccessKey)
	return o
}

// SetSecretAccessKey adds the secretAccessKey to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetSecretAccessKey(secretAccessKey *string) {
	o.SecretAccessKey = secretAccessKey
}

// WithVPC adds the vPC to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithVPC(vPC *string) *ListProjectAWSSubnetsParams {
	o.SetVPC(vPC)
	return o
}

// SetVPC adds the vPC to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetVPC(vPC *string) {
	o.VPC = vPC
}

// WithDC adds the dc to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithDC(dc string) *ListProjectAWSSubnetsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) WithProjectID(projectID string) *ListProjectAWSSubnetsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project a w s subnets params
func (o *ListProjectAWSSubnetsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectAWSSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessKeyID != nil {

		// header param AccessKeyID
		if err := r.SetHeaderParam("AccessKeyID", *o.AccessKeyID); err != nil {
			return err
		}
	}

	if o.AssumeRoleARN != nil {

		// header param AssumeRoleARN
		if err := r.SetHeaderParam("AssumeRoleARN", *o.AssumeRoleARN); err != nil {
			return err
		}
	}

	if o.AssumeRoleExternalID != nil {

		// header param AssumeRoleExternalID
		if err := r.SetHeaderParam("AssumeRoleExternalID", *o.AssumeRoleExternalID); err != nil {
			return err
		}
	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.SecretAccessKey != nil {

		// header param SecretAccessKey
		if err := r.SetHeaderParam("SecretAccessKey", *o.SecretAccessKey); err != nil {
			return err
		}
	}

	if o.VPC != nil {

		// header param VPC
		if err := r.SetHeaderParam("VPC", *o.VPC); err != nil {
			return err
		}
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}