// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "tingtingbackend/models"
)

// NewNrRelationAlbumBooklistEditParams creates a new NrRelationAlbumBooklistEditParams object
// with the default values initialized.
func NewNrRelationAlbumBooklistEditParams() NrRelationAlbumBooklistEditParams {
	var ()
	return NrRelationAlbumBooklistEditParams{}
}

// NrRelationAlbumBooklistEditParams contains all the bound params for the relation album booklist edit operation
// typically these are obtained from a http.Request
//
// swagger:parameters /relation/album/booklist/edit
type NrRelationAlbumBooklistEditParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*/relation/album/booklist/edit
	  In: body
	*/
	Body *models.Body1
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrRelationAlbumBooklistEditParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Body1
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}