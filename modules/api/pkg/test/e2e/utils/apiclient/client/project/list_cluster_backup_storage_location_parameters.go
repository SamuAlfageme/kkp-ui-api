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

// NewListClusterBackupStorageLocationParams creates a new ListClusterBackupStorageLocationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListClusterBackupStorageLocationParams() *ListClusterBackupStorageLocationParams {
	return &ListClusterBackupStorageLocationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListClusterBackupStorageLocationParamsWithTimeout creates a new ListClusterBackupStorageLocationParams object
// with the ability to set a timeout on a request.
func NewListClusterBackupStorageLocationParamsWithTimeout(timeout time.Duration) *ListClusterBackupStorageLocationParams {
	return &ListClusterBackupStorageLocationParams{
		timeout: timeout,
	}
}

// NewListClusterBackupStorageLocationParamsWithContext creates a new ListClusterBackupStorageLocationParams object
// with the ability to set a context for a request.
func NewListClusterBackupStorageLocationParamsWithContext(ctx context.Context) *ListClusterBackupStorageLocationParams {
	return &ListClusterBackupStorageLocationParams{
		Context: ctx,
	}
}

// NewListClusterBackupStorageLocationParamsWithHTTPClient creates a new ListClusterBackupStorageLocationParams object
// with the ability to set a custom HTTPClient for a request.
func NewListClusterBackupStorageLocationParamsWithHTTPClient(client *http.Client) *ListClusterBackupStorageLocationParams {
	return &ListClusterBackupStorageLocationParams{
		HTTPClient: client,
	}
}

/*
ListClusterBackupStorageLocationParams contains all the parameters to send to the API endpoint

	for the list cluster backup storage location operation.

	Typically these are written to a http.Request.
*/
type ListClusterBackupStorageLocationParams struct {

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list cluster backup storage location params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClusterBackupStorageLocationParams) WithDefaults() *ListClusterBackupStorageLocationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list cluster backup storage location params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClusterBackupStorageLocationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list cluster backup storage location params
func (o *ListClusterBackupStorageLocationParams) WithTimeout(timeout time.Duration) *ListClusterBackupStorageLocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list cluster backup storage location params
func (o *ListClusterBackupStorageLocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list cluster backup storage location params
func (o *ListClusterBackupStorageLocationParams) WithContext(ctx context.Context) *ListClusterBackupStorageLocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list cluster backup storage location params
func (o *ListClusterBackupStorageLocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list cluster backup storage location params
func (o *ListClusterBackupStorageLocationParams) WithHTTPClient(client *http.Client) *ListClusterBackupStorageLocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list cluster backup storage location params
func (o *ListClusterBackupStorageLocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the list cluster backup storage location params
func (o *ListClusterBackupStorageLocationParams) WithProjectID(projectID string) *ListClusterBackupStorageLocationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list cluster backup storage location params
func (o *ListClusterBackupStorageLocationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListClusterBackupStorageLocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}