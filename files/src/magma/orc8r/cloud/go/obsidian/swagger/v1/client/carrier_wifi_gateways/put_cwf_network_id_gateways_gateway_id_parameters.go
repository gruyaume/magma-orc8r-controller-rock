// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

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

// NewPutCwfNetworkIDGatewaysGatewayIDParams creates a new PutCwfNetworkIDGatewaysGatewayIDParams object
// with the default values initialized.
func NewPutCwfNetworkIDGatewaysGatewayIDParams() *PutCwfNetworkIDGatewaysGatewayIDParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDParamsWithTimeout creates a new PutCwfNetworkIDGatewaysGatewayIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCwfNetworkIDGatewaysGatewayIDParamsWithTimeout(timeout time.Duration) *PutCwfNetworkIDGatewaysGatewayIDParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDParams{

		timeout: timeout,
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDParamsWithContext creates a new PutCwfNetworkIDGatewaysGatewayIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCwfNetworkIDGatewaysGatewayIDParamsWithContext(ctx context.Context) *PutCwfNetworkIDGatewaysGatewayIDParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDParams{

		Context: ctx,
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDParamsWithHTTPClient creates a new PutCwfNetworkIDGatewaysGatewayIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCwfNetworkIDGatewaysGatewayIDParamsWithHTTPClient(client *http.Client) *PutCwfNetworkIDGatewaysGatewayIDParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDParams{
		HTTPClient: client,
	}
}

/*PutCwfNetworkIDGatewaysGatewayIDParams contains all the parameters to send to the API endpoint
for the put cwf network ID gateways gateway ID operation typically these are written to a http.Request
*/
type PutCwfNetworkIDGatewaysGatewayIDParams struct {

	/*Gateway
	  Full desired configuration of the gateway

	*/
	Gateway *models.MutableCwfGateway
	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) WithTimeout(timeout time.Duration) *PutCwfNetworkIDGatewaysGatewayIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) WithContext(ctx context.Context) *PutCwfNetworkIDGatewaysGatewayIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) WithHTTPClient(client *http.Client) *PutCwfNetworkIDGatewaysGatewayIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGateway adds the gateway to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) WithGateway(gateway *models.MutableCwfGateway) *PutCwfNetworkIDGatewaysGatewayIDParams {
	o.SetGateway(gateway)
	return o
}

// SetGateway adds the gateway to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) SetGateway(gateway *models.MutableCwfGateway) {
	o.Gateway = gateway
}

// WithGatewayID adds the gatewayID to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) WithGatewayID(gatewayID string) *PutCwfNetworkIDGatewaysGatewayIDParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) WithNetworkID(networkID string) *PutCwfNetworkIDGatewaysGatewayIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put cwf network ID gateways gateway ID params
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCwfNetworkIDGatewaysGatewayIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Gateway != nil {
		if err := r.SetBodyParam(o.Gateway); err != nil {
			return err
		}
	}

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
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
