// Code generated by go-swagger; DO NOT EDIT.

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PostChannelsReader is a Reader for the PostChannels structure.
type PostChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostChannelsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostChannelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostChannelsCreated creates a PostChannelsCreated with default headers values
func NewPostChannelsCreated() *PostChannelsCreated {
	return &PostChannelsCreated{}
}

/*PostChannelsCreated handles this case with default header values.

Success
*/
type PostChannelsCreated struct {
}

func (o *PostChannelsCreated) Error() string {
	return fmt.Sprintf("[POST /channels][%d] postChannelsCreated ", 201)
}

func (o *PostChannelsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChannelsDefault creates a PostChannelsDefault with default headers values
func NewPostChannelsDefault(code int) *PostChannelsDefault {
	return &PostChannelsDefault{
		_statusCode: code,
	}
}

/*PostChannelsDefault handles this case with default header values.

Unexpected Error
*/
type PostChannelsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post channels default response
func (o *PostChannelsDefault) Code() int {
	return o._statusCode
}

func (o *PostChannelsDefault) Error() string {
	return fmt.Sprintf("[POST /channels][%d] PostChannels default  %+v", o._statusCode, o.Payload)
}

func (o *PostChannelsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostChannelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
