// Code generated by go-swagger; DO NOT EDIT.

package msg

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "certbackend/models"
)

// MsgDetailOKCode is the HTTP code returned for type MsgDetailOK
const MsgDetailOKCode int = 200

/*MsgDetailOK 消息详情

swagger:response msgDetailOK
*/
type MsgDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20012 `json:"body,omitempty"`
}

// NewMsgDetailOK creates MsgDetailOK with default headers values
func NewMsgDetailOK() *MsgDetailOK {
	return &MsgDetailOK{}
}

// WithPayload adds the payload to the msg detail o k response
func (o *MsgDetailOK) WithPayload(payload *models.InlineResponse20012) *MsgDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the msg detail o k response
func (o *MsgDetailOK) SetPayload(payload *models.InlineResponse20012) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MsgDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*MsgDetailDefault error

swagger:response msgDetailDefault
*/
type MsgDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMsgDetailDefault creates MsgDetailDefault with default headers values
func NewMsgDetailDefault(code int) *MsgDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &MsgDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the msg detail default response
func (o *MsgDetailDefault) WithStatusCode(code int) *MsgDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the msg detail default response
func (o *MsgDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the msg detail default response
func (o *MsgDetailDefault) WithPayload(payload *models.Error) *MsgDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the msg detail default response
func (o *MsgDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MsgDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
