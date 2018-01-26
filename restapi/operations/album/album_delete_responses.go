// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "certbackend/models"
)

// AlbumDeleteOKCode is the HTTP code returned for type AlbumDeleteOK
const AlbumDeleteOKCode int = 200

/*AlbumDeleteOK 专辑详情

swagger:response albumDeleteOK
*/
type AlbumDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20018 `json:"body,omitempty"`
}

// NewAlbumDeleteOK creates AlbumDeleteOK with default headers values
func NewAlbumDeleteOK() *AlbumDeleteOK {
	return &AlbumDeleteOK{}
}

// WithPayload adds the payload to the album delete o k response
func (o *AlbumDeleteOK) WithPayload(payload *models.InlineResponse20018) *AlbumDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album delete o k response
func (o *AlbumDeleteOK) SetPayload(payload *models.InlineResponse20018) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AlbumDeleteDefault error

swagger:response albumDeleteDefault
*/
type AlbumDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAlbumDeleteDefault creates AlbumDeleteDefault with default headers values
func NewAlbumDeleteDefault(code int) *AlbumDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &AlbumDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the album delete default response
func (o *AlbumDeleteDefault) WithStatusCode(code int) *AlbumDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the album delete default response
func (o *AlbumDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the album delete default response
func (o *AlbumDeleteDefault) WithPayload(payload *models.Error) *AlbumDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album delete default response
func (o *AlbumDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
