// Code generated by go-swagger; DO NOT EDIT.

package todo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	modules "github.com/chinajuanbob/helloworld/pkg/gen/modules"
)

// NewUpdateTodoParams creates a new UpdateTodoParams object
// with the default values initialized.
func NewUpdateTodoParams() *UpdateTodoParams {
	var ()
	return &UpdateTodoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTodoParamsWithTimeout creates a new UpdateTodoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTodoParamsWithTimeout(timeout time.Duration) *UpdateTodoParams {
	var ()
	return &UpdateTodoParams{

		timeout: timeout,
	}
}

// NewUpdateTodoParamsWithContext creates a new UpdateTodoParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTodoParamsWithContext(ctx context.Context) *UpdateTodoParams {
	var ()
	return &UpdateTodoParams{

		Context: ctx,
	}
}

// NewUpdateTodoParamsWithHTTPClient creates a new UpdateTodoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTodoParamsWithHTTPClient(client *http.Client) *UpdateTodoParams {
	var ()
	return &UpdateTodoParams{
		HTTPClient: client,
	}
}

/*UpdateTodoParams contains all the parameters to send to the API endpoint
for the update todo operation typically these are written to a http.Request
*/
type UpdateTodoParams struct {

	/*Body
	  new content

	*/
	Body *modules.PbUpdateTodoRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update todo params
func (o *UpdateTodoParams) WithTimeout(timeout time.Duration) *UpdateTodoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update todo params
func (o *UpdateTodoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update todo params
func (o *UpdateTodoParams) WithContext(ctx context.Context) *UpdateTodoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update todo params
func (o *UpdateTodoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update todo params
func (o *UpdateTodoParams) WithHTTPClient(client *http.Client) *UpdateTodoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update todo params
func (o *UpdateTodoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update todo params
func (o *UpdateTodoParams) WithBody(body *modules.PbUpdateTodoRequest) *UpdateTodoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update todo params
func (o *UpdateTodoParams) SetBody(body *modules.PbUpdateTodoRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTodoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
