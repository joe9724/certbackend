// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "certbackend/models"
)

// UserLoginOKCode is the HTTP code returned for type UserLoginOK
const UserLoginOKCode int = 200

/*UserLoginOK 登录成功，返回登录信息

swagger:response userLoginOK
*/
type UserLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20029 `json:"body,omitempty"`
}

// NewUserLoginOK creates UserLoginOK with default headers values
func NewUserLoginOK() *UserLoginOK {
	return &UserLoginOK{}
}

// WithPayload adds the payload to the user login o k response
func (o *UserLoginOK) WithPayload(payload *models.InlineResponse20029) *UserLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user login o k response
func (o *UserLoginOK) SetPayload(payload *models.InlineResponse20029) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrUserLoginDefault error

swagger:response userLoginDefault
*/
type NrUserLoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrUserLoginDefault creates NrUserLoginDefault with default headers values
func NewNrUserLoginDefault(code int) *NrUserLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &NrUserLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user login default response
func (o *NrUserLoginDefault) WithStatusCode(code int) *NrUserLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user login default response
func (o *NrUserLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user login default response
func (o *NrUserLoginDefault) WithPayload(payload *models.Error) *NrUserLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user login default response
func (o *NrUserLoginDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrUserLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
