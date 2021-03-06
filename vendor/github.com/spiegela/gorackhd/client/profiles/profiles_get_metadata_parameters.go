package profiles

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

// NewProfilesGetMetadataParams creates a new ProfilesGetMetadataParams object
// with the default values initialized.
func NewProfilesGetMetadataParams() *ProfilesGetMetadataParams {

	return &ProfilesGetMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProfilesGetMetadataParamsWithTimeout creates a new ProfilesGetMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProfilesGetMetadataParamsWithTimeout(timeout time.Duration) *ProfilesGetMetadataParams {

	return &ProfilesGetMetadataParams{

		timeout: timeout,
	}
}

// NewProfilesGetMetadataParamsWithContext creates a new ProfilesGetMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewProfilesGetMetadataParamsWithContext(ctx context.Context) *ProfilesGetMetadataParams {

	return &ProfilesGetMetadataParams{

		Context: ctx,
	}
}

// NewProfilesGetMetadataParamsWithHTTPClient creates a new ProfilesGetMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProfilesGetMetadataParamsWithHTTPClient(client *http.Client) *ProfilesGetMetadataParams {

	return &ProfilesGetMetadataParams{
		HTTPClient: client,
	}
}

/*ProfilesGetMetadataParams contains all the parameters to send to the API endpoint
for the profiles get metadata operation typically these are written to a http.Request
*/
type ProfilesGetMetadataParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the profiles get metadata params
func (o *ProfilesGetMetadataParams) WithTimeout(timeout time.Duration) *ProfilesGetMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the profiles get metadata params
func (o *ProfilesGetMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the profiles get metadata params
func (o *ProfilesGetMetadataParams) WithContext(ctx context.Context) *ProfilesGetMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the profiles get metadata params
func (o *ProfilesGetMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the profiles get metadata params
func (o *ProfilesGetMetadataParams) WithHTTPClient(client *http.Client) *ProfilesGetMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the profiles get metadata params
func (o *ProfilesGetMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ProfilesGetMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
