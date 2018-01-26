// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingbackend/models"
)

// BookChapterListRelationEditOKCode is the HTTP code returned for type BookChapterListRelationEditOK
const BookChapterListRelationEditOKCode int = 200

/*BookChapterListRelationEditOK 操作成功，返回操作信息

swagger:response bookChapterListRelationEditOK
*/
type BookChapterListRelationEditOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20018 `json:"body,omitempty"`
}

// NewBookChapterListRelationEditOK creates BookChapterListRelationEditOK with default headers values
func NewBookChapterListRelationEditOK() *BookChapterListRelationEditOK {
	return &BookChapterListRelationEditOK{}
}

// WithPayload adds the payload to the book chapter list relation edit o k response
func (o *BookChapterListRelationEditOK) WithPayload(payload *models.InlineResponse20018) *BookChapterListRelationEditOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the book chapter list relation edit o k response
func (o *BookChapterListRelationEditOK) SetPayload(payload *models.InlineResponse20018) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BookChapterListRelationEditOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*BookChapterListRelationEditDefault error

swagger:response bookChapterListRelationEditDefault
*/
type BookChapterListRelationEditDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBookChapterListRelationEditDefault creates BookChapterListRelationEditDefault with default headers values
func NewBookChapterListRelationEditDefault(code int) *BookChapterListRelationEditDefault {
	if code <= 0 {
		code = 500
	}

	return &BookChapterListRelationEditDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the book chapter list relation edit default response
func (o *BookChapterListRelationEditDefault) WithStatusCode(code int) *BookChapterListRelationEditDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the book chapter list relation edit default response
func (o *BookChapterListRelationEditDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the book chapter list relation edit default response
func (o *BookChapterListRelationEditDefault) WithPayload(payload *models.Error) *BookChapterListRelationEditDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the book chapter list relation edit default response
func (o *BookChapterListRelationEditDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BookChapterListRelationEditDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}