// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// UserListOKCode is the HTTP code returned for type UserListOK
const UserListOKCode int = 200

/*UserListOK 操作成功，返回登录信息

swagger:response userListOK
*/
type UserListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20030 `json:"body,omitempty"`
}

// NewUserListOK creates UserListOK with default headers values
func NewUserListOK() *UserListOK {
	return &UserListOK{}
}

// WithPayload adds the payload to the user list o k response
func (o *UserListOK) WithPayload(payload *models.InlineResponse20030) *UserListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user list o k response
func (o *UserListOK) SetPayload(payload *models.InlineResponse20030) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrUserListDefault error

swagger:response userListDefault
*/
type NrUserListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrUserListDefault creates NrUserListDefault with default headers values
func NewNrUserListDefault(code int) *NrUserListDefault {
	if code <= 0 {
		code = 500
	}

	return &NrUserListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user list default response
func (o *NrUserListDefault) WithStatusCode(code int) *NrUserListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user list default response
func (o *NrUserListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user list default response
func (o *NrUserListDefault) WithPayload(payload *models.Error) *NrUserListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user list default response
func (o *NrUserListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrUserListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
