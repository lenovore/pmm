// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// StartPTSummaryActionReader is a Reader for the StartPTSummaryAction structure.
type StartPTSummaryActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPTSummaryActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStartPTSummaryActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewStartPTSummaryActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartPTSummaryActionOK creates a StartPTSummaryActionOK with default headers values
func NewStartPTSummaryActionOK() *StartPTSummaryActionOK {
	return &StartPTSummaryActionOK{}
}

/*StartPTSummaryActionOK handles this case with default header values.

A successful response.
*/
type StartPTSummaryActionOK struct {
	Payload *StartPTSummaryActionOKBody
}

func (o *StartPTSummaryActionOK) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/StartPTSummary][%d] startPTSummaryActionOk  %+v", 200, o.Payload)
}

func (o *StartPTSummaryActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPTSummaryActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPTSummaryActionDefault creates a StartPTSummaryActionDefault with default headers values
func NewStartPTSummaryActionDefault(code int) *StartPTSummaryActionDefault {
	return &StartPTSummaryActionDefault{
		_statusCode: code,
	}
}

/*StartPTSummaryActionDefault handles this case with default header values.

An error response.
*/
type StartPTSummaryActionDefault struct {
	_statusCode int

	Payload *StartPTSummaryActionDefaultBody
}

// Code gets the status code for the start p t summary action default response
func (o *StartPTSummaryActionDefault) Code() int {
	return o._statusCode
}

func (o *StartPTSummaryActionDefault) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/StartPTSummary][%d] StartPTSummaryAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartPTSummaryActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPTSummaryActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartPTSummaryActionBody start p t summary action body
swagger:model StartPTSummaryActionBody
*/
type StartPTSummaryActionBody struct {

	// Node ID for this Action.
	NodeID string `json:"node_id,omitempty"`

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start p t summary action body
func (o *StartPTSummaryActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTSummaryActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTSummaryActionBody) UnmarshalBinary(b []byte) error {
	var res StartPTSummaryActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTSummaryActionDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model StartPTSummaryActionDefaultBody
*/
type StartPTSummaryActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this start p t summary action default body
func (o *StartPTSummaryActionDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTSummaryActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTSummaryActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartPTSummaryActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTSummaryActionOKBody start p t summary action OK body
swagger:model StartPTSummaryActionOKBody
*/
type StartPTSummaryActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start p t summary action OK body
func (o *StartPTSummaryActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTSummaryActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTSummaryActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartPTSummaryActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}