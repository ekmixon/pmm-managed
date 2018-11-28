// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

	models "github.com/percona/pmm/api/json/models"
)

// NewAddAWSRDSNodeParams creates a new AddAWSRDSNodeParams object
// with the default values initialized.
func NewAddAWSRDSNodeParams() *AddAWSRDSNodeParams {
	var ()
	return &AddAWSRDSNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAWSRDSNodeParamsWithTimeout creates a new AddAWSRDSNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAWSRDSNodeParamsWithTimeout(timeout time.Duration) *AddAWSRDSNodeParams {
	var ()
	return &AddAWSRDSNodeParams{

		timeout: timeout,
	}
}

// NewAddAWSRDSNodeParamsWithContext creates a new AddAWSRDSNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAWSRDSNodeParamsWithContext(ctx context.Context) *AddAWSRDSNodeParams {
	var ()
	return &AddAWSRDSNodeParams{

		Context: ctx,
	}
}

// NewAddAWSRDSNodeParamsWithHTTPClient creates a new AddAWSRDSNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAWSRDSNodeParamsWithHTTPClient(client *http.Client) *AddAWSRDSNodeParams {
	var ()
	return &AddAWSRDSNodeParams{
		HTTPClient: client,
	}
}

/*AddAWSRDSNodeParams contains all the parameters to send to the API endpoint
for the add AWS RDS node operation typically these are written to a http.Request
*/
type AddAWSRDSNodeParams struct {

	/*Body*/
	Body *models.InventoryAddAWSRDSNodeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add AWS RDS node params
func (o *AddAWSRDSNodeParams) WithTimeout(timeout time.Duration) *AddAWSRDSNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add AWS RDS node params
func (o *AddAWSRDSNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add AWS RDS node params
func (o *AddAWSRDSNodeParams) WithContext(ctx context.Context) *AddAWSRDSNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add AWS RDS node params
func (o *AddAWSRDSNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add AWS RDS node params
func (o *AddAWSRDSNodeParams) WithHTTPClient(client *http.Client) *AddAWSRDSNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add AWS RDS node params
func (o *AddAWSRDSNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add AWS RDS node params
func (o *AddAWSRDSNodeParams) WithBody(body *models.InventoryAddAWSRDSNodeRequest) *AddAWSRDSNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add AWS RDS node params
func (o *AddAWSRDSNodeParams) SetBody(body *models.InventoryAddAWSRDSNodeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAWSRDSNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
