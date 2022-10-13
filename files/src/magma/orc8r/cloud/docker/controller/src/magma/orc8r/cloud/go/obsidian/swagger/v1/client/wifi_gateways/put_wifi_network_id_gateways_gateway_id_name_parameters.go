// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

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

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// NewPutWifiNetworkIDGatewaysGatewayIDNameParams creates a new PutWifiNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized.
func NewPutWifiNetworkIDGatewaysGatewayIDNameParams() *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &PutWifiNetworkIDGatewaysGatewayIDNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDNameParamsWithTimeout creates a new PutWifiNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutWifiNetworkIDGatewaysGatewayIDNameParamsWithTimeout(timeout time.Duration) *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &PutWifiNetworkIDGatewaysGatewayIDNameParams{

		timeout: timeout,
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDNameParamsWithContext creates a new PutWifiNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutWifiNetworkIDGatewaysGatewayIDNameParamsWithContext(ctx context.Context) *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &PutWifiNetworkIDGatewaysGatewayIDNameParams{

		Context: ctx,
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDNameParamsWithHTTPClient creates a new PutWifiNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutWifiNetworkIDGatewaysGatewayIDNameParamsWithHTTPClient(client *http.Client) *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &PutWifiNetworkIDGatewaysGatewayIDNameParams{
		HTTPClient: client,
	}
}

/*PutWifiNetworkIDGatewaysGatewayIDNameParams contains all the parameters to send to the API endpoint
for the put wifi network ID gateways gateway ID name operation typically these are written to a http.Request
*/
type PutWifiNetworkIDGatewaysGatewayIDNameParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*Name*/
	Name models.GatewayName
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) WithTimeout(timeout time.Duration) *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) WithContext(ctx context.Context) *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) WithHTTPClient(client *http.Client) *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) WithGatewayID(gatewayID string) *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithName adds the name to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) WithName(name models.GatewayName) *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) SetName(name models.GatewayName) {
	o.Name = name
}

// WithNetworkID adds the networkID to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) WithNetworkID(networkID string) *PutWifiNetworkIDGatewaysGatewayIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put wifi network ID gateways gateway ID name params
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutWifiNetworkIDGatewaysGatewayIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Name); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
