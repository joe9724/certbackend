// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "certbackend/models"
)

// CategoryDeleteOKCode is the HTTP code returned for type CategoryDeleteOK
const CategoryDeleteOKCode int = 200

/*CategoryDeleteOK 分类详情

swagger:response categoryDeleteOK
*/
type CategoryDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20018 `json:"body,omitempty"`
}

// NewCategoryDeleteOK creates CategoryDeleteOK with default headers values
func NewCategoryDeleteOK() *CategoryDeleteOK {
	return &CategoryDeleteOK{}
}

// WithPayload adds the payload to the category delete o k response
func (o *CategoryDeleteOK) WithPayload(payload *models.InlineResponse20018) *CategoryDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the category delete o k response
func (o *CategoryDeleteOK) SetPayload(payload *models.InlineResponse20018) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CategoryDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CategoryDeleteDefault error

swagger:response categoryDeleteDefault
*/
type CategoryDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCategoryDeleteDefault creates CategoryDeleteDefault with default headers values
func NewCategoryDeleteDefault(code int) *CategoryDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &CategoryDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the category delete default response
func (o *CategoryDeleteDefault) WithStatusCode(code int) *CategoryDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the category delete default response
func (o *CategoryDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the category delete default response
func (o *CategoryDeleteDefault) WithPayload(payload *models.Error) *CategoryDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the category delete default response
func (o *CategoryDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CategoryDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
