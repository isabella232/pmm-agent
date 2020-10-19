// Code generated by go-swagger; DO NOT EDIT.

package agent_local

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

// NewStatus2Params creates a new Status2Params object
// with the default values initialized.
func NewStatus2Params() *Status2Params {
	var ()
	return &Status2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewStatus2ParamsWithTimeout creates a new Status2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatus2ParamsWithTimeout(timeout time.Duration) *Status2Params {
	var ()
	return &Status2Params{

		timeout: timeout,
	}
}

// NewStatus2ParamsWithContext creates a new Status2Params object
// with the default values initialized, and the ability to set a context for a request
func NewStatus2ParamsWithContext(ctx context.Context) *Status2Params {
	var ()
	return &Status2Params{

		Context: ctx,
	}
}

// NewStatus2ParamsWithHTTPClient creates a new Status2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStatus2ParamsWithHTTPClient(client *http.Client) *Status2Params {
	var ()
	return &Status2Params{
		HTTPClient: client,
	}
}

/*Status2Params contains all the parameters to send to the API endpoint
for the status2 operation typically these are written to a http.Request
*/
type Status2Params struct {

	/*GetNetworkInfo
	  Returns network info (latency and clock_drift) if true.

	*/
	GetNetworkInfo *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the status2 params
func (o *Status2Params) WithTimeout(timeout time.Duration) *Status2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status2 params
func (o *Status2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status2 params
func (o *Status2Params) WithContext(ctx context.Context) *Status2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status2 params
func (o *Status2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status2 params
func (o *Status2Params) WithHTTPClient(client *http.Client) *Status2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status2 params
func (o *Status2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGetNetworkInfo adds the getNetworkInfo to the status2 params
func (o *Status2Params) WithGetNetworkInfo(getNetworkInfo *bool) *Status2Params {
	o.SetGetNetworkInfo(getNetworkInfo)
	return o
}

// SetGetNetworkInfo adds the getNetworkInfo to the status2 params
func (o *Status2Params) SetGetNetworkInfo(getNetworkInfo *bool) {
	o.GetNetworkInfo = getNetworkInfo
}

// WriteToRequest writes these params to a swagger request
func (o *Status2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GetNetworkInfo != nil {

		// query param get_network_info
		var qrGetNetworkInfo bool
		if o.GetNetworkInfo != nil {
			qrGetNetworkInfo = *o.GetNetworkInfo
		}
		qGetNetworkInfo := swag.FormatBool(qrGetNetworkInfo)
		if qGetNetworkInfo != "" {
			if err := r.SetQueryParam("get_network_info", qGetNetworkInfo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}