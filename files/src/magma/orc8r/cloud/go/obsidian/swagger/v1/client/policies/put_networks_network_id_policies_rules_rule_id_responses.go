// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutNetworksNetworkIDPoliciesRulesRuleIDReader is a Reader for the PutNetworksNetworkIDPoliciesRulesRuleID structure.
type PutNetworksNetworkIDPoliciesRulesRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDPoliciesRulesRuleIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDPoliciesRulesRuleIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDPoliciesRulesRuleIDNoContent creates a PutNetworksNetworkIDPoliciesRulesRuleIDNoContent with default headers values
func NewPutNetworksNetworkIDPoliciesRulesRuleIDNoContent() *PutNetworksNetworkIDPoliciesRulesRuleIDNoContent {
	return &PutNetworksNetworkIDPoliciesRulesRuleIDNoContent{}
}

/*PutNetworksNetworkIDPoliciesRulesRuleIDNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDPoliciesRulesRuleIDNoContent struct {
}

func (o *PutNetworksNetworkIDPoliciesRulesRuleIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/policies/rules/{rule_id}][%d] putNetworksNetworkIdPoliciesRulesRuleIdNoContent ", 204)
}

func (o *PutNetworksNetworkIDPoliciesRulesRuleIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDPoliciesRulesRuleIDDefault creates a PutNetworksNetworkIDPoliciesRulesRuleIDDefault with default headers values
func NewPutNetworksNetworkIDPoliciesRulesRuleIDDefault(code int) *PutNetworksNetworkIDPoliciesRulesRuleIDDefault {
	return &PutNetworksNetworkIDPoliciesRulesRuleIDDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDPoliciesRulesRuleIDDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDPoliciesRulesRuleIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID policies rules rule ID default response
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDPoliciesRulesRuleIDDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/policies/rules/{rule_id}][%d] PutNetworksNetworkIDPoliciesRulesRuleID default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDPoliciesRulesRuleIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDPoliciesRulesRuleIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
