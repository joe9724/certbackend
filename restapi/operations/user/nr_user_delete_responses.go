// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "certbackend/models"
)

// UserDeleteOKCode is the HTTP code returned for type UserDeleteOK
const UserDeleteOKCode int = 200

/*UserDeleteOK 删除成功，返回成功信息

swagger:response userDeleteOK
*/
type UserDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20018 `json:"body,omitempty"`
}

// NewUserDeleteOK creates UserDeleteOK with default headers values
func NewUserDeleteOK() *UserDeleteOK {
	return &UserDeleteOK{}
}

// WithPayload adds the payload to the user delete o k response
func (o *UserDeleteOK) WithPayload(payload *models.InlineResponse20018) *UserDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user delete o k response
func (o *UserDeleteOK) SetPayload(payload *models.InlineResponse20018) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrUserDeleteDefault error

swagger:response userDeleteDefault
*/
type NrUserDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrUserDeleteDefault creates NrUserDeleteDefault with default headers values
func NewNrUserDeleteDefault(code int) *NrUserDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &NrUserDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user delete default response
func (o *NrUserDeleteDefault) WithStatusCode(code int) *NrUserDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user delete default response
func (o *NrUserDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user delete default response
func (o *NrUserDeleteDefault) WithPayload(payload *models.Error) *NrUserDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user delete default response
func (o *NrUserDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrUserDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
