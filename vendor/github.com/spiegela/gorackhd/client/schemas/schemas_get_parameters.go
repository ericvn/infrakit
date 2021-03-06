package schemas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSchemasGetParams creates a new SchemasGetParams object
// with the default values initialized.
func NewSchemasGetParams() *SchemasGetParams {

	return &SchemasGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemasGetParamsWithTimeout creates a new SchemasGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemasGetParamsWithTimeout(timeout time.Duration) *SchemasGetParams {

	return &SchemasGetParams{

		timeout: timeout,
	}
}

// NewSchemasGetParamsWithContext creates a new SchemasGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemasGetParamsWithContext(ctx context.Context) *SchemasGetParams {

	return &SchemasGetParams{

		Context: ctx,
	}
}

// NewSchemasGetParamsWithHTTPClient creates a new SchemasGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemasGetParamsWithHTTPClient(client *http.Client) *SchemasGetParams {

	return &SchemasGetParams{
		HTTPClient: client,
	}
}

/*SchemasGetParams contains all the parameters to send to the API endpoint
for the schemas get operation typically these are written to a http.Request
*/
type SchemasGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schemas get params
func (o *SchemasGetParams) WithTimeout(timeout time.Duration) *SchemasGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schemas get params
func (o *SchemasGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schemas get params
func (o *SchemasGetParams) WithContext(ctx context.Context) *SchemasGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schemas get params
func (o *SchemasGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schemas get params
func (o *SchemasGetParams) WithHTTPClient(client *http.Client) *SchemasGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schemas get params
func (o *SchemasGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SchemasGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
