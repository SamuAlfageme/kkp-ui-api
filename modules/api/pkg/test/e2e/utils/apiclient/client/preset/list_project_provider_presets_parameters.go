// Code generated by go-swagger; DO NOT EDIT.

package preset

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
	"github.com/go-openapi/swag"
)

// NewListProjectProviderPresetsParams creates a new ListProjectProviderPresetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectProviderPresetsParams() *ListProjectProviderPresetsParams {
	return &ListProjectProviderPresetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectProviderPresetsParamsWithTimeout creates a new ListProjectProviderPresetsParams object
// with the ability to set a timeout on a request.
func NewListProjectProviderPresetsParamsWithTimeout(timeout time.Duration) *ListProjectProviderPresetsParams {
	return &ListProjectProviderPresetsParams{
		timeout: timeout,
	}
}

// NewListProjectProviderPresetsParamsWithContext creates a new ListProjectProviderPresetsParams object
// with the ability to set a context for a request.
func NewListProjectProviderPresetsParamsWithContext(ctx context.Context) *ListProjectProviderPresetsParams {
	return &ListProjectProviderPresetsParams{
		Context: ctx,
	}
}

// NewListProjectProviderPresetsParamsWithHTTPClient creates a new ListProjectProviderPresetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectProviderPresetsParamsWithHTTPClient(client *http.Client) *ListProjectProviderPresetsParams {
	return &ListProjectProviderPresetsParams{
		HTTPClient: client,
	}
}

/*
ListProjectProviderPresetsParams contains all the parameters to send to the API endpoint

	for the list project provider presets operation.

	Typically these are written to a http.Request.
*/
type ListProjectProviderPresetsParams struct {

	// Datacenter.
	Datacenter *string

	// Disabled.
	Disabled *bool

	// Name.
	Name *string

	// ProjectID.
	ProjectID string

	// ProviderName.
	ProviderName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project provider presets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectProviderPresetsParams) WithDefaults() *ListProjectProviderPresetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project provider presets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectProviderPresetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project provider presets params
func (o *ListProjectProviderPresetsParams) WithTimeout(timeout time.Duration) *ListProjectProviderPresetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project provider presets params
func (o *ListProjectProviderPresetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project provider presets params
func (o *ListProjectProviderPresetsParams) WithContext(ctx context.Context) *ListProjectProviderPresetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project provider presets params
func (o *ListProjectProviderPresetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project provider presets params
func (o *ListProjectProviderPresetsParams) WithHTTPClient(client *http.Client) *ListProjectProviderPresetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project provider presets params
func (o *ListProjectProviderPresetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatacenter adds the datacenter to the list project provider presets params
func (o *ListProjectProviderPresetsParams) WithDatacenter(datacenter *string) *ListProjectProviderPresetsParams {
	o.SetDatacenter(datacenter)
	return o
}

// SetDatacenter adds the datacenter to the list project provider presets params
func (o *ListProjectProviderPresetsParams) SetDatacenter(datacenter *string) {
	o.Datacenter = datacenter
}

// WithDisabled adds the disabled to the list project provider presets params
func (o *ListProjectProviderPresetsParams) WithDisabled(disabled *bool) *ListProjectProviderPresetsParams {
	o.SetDisabled(disabled)
	return o
}

// SetDisabled adds the disabled to the list project provider presets params
func (o *ListProjectProviderPresetsParams) SetDisabled(disabled *bool) {
	o.Disabled = disabled
}

// WithName adds the name to the list project provider presets params
func (o *ListProjectProviderPresetsParams) WithName(name *string) *ListProjectProviderPresetsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list project provider presets params
func (o *ListProjectProviderPresetsParams) SetName(name *string) {
	o.Name = name
}

// WithProjectID adds the projectID to the list project provider presets params
func (o *ListProjectProviderPresetsParams) WithProjectID(projectID string) *ListProjectProviderPresetsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project provider presets params
func (o *ListProjectProviderPresetsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithProviderName adds the providerName to the list project provider presets params
func (o *ListProjectProviderPresetsParams) WithProviderName(providerName string) *ListProjectProviderPresetsParams {
	o.SetProviderName(providerName)
	return o
}

// SetProviderName adds the providerName to the list project provider presets params
func (o *ListProjectProviderPresetsParams) SetProviderName(providerName string) {
	o.ProviderName = providerName
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectProviderPresetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Datacenter != nil {

		// query param datacenter
		var qrDatacenter string

		if o.Datacenter != nil {
			qrDatacenter = *o.Datacenter
		}
		qDatacenter := qrDatacenter
		if qDatacenter != "" {

			if err := r.SetQueryParam("datacenter", qDatacenter); err != nil {
				return err
			}
		}
	}

	if o.Disabled != nil {

		// query param disabled
		var qrDisabled bool

		if o.Disabled != nil {
			qrDisabled = *o.Disabled
		}
		qDisabled := swag.FormatBool(qrDisabled)
		if qDisabled != "" {

			if err := r.SetQueryParam("disabled", qDisabled); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param provider_name
	if err := r.SetPathParam("provider_name", o.ProviderName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
