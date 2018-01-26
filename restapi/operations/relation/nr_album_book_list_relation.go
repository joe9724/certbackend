// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrAlbumBookListRelationHandlerFunc turns a function with the right signature into a album book list relation handler
type NrAlbumBookListRelationHandlerFunc func(NrAlbumBookListRelationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrAlbumBookListRelationHandlerFunc) Handle(params NrAlbumBookListRelationParams) middleware.Responder {
	return fn(params)
}

// NrAlbumBookListRelationHandler interface for that can handle valid album book list relation params
type NrAlbumBookListRelationHandler interface {
	Handle(NrAlbumBookListRelationParams) middleware.Responder
}

// NewNrAlbumBookListRelation creates a new http.Handler for the album book list relation operation
func NewNrAlbumBookListRelation(ctx *middleware.Context, handler NrAlbumBookListRelationHandler) *NrAlbumBookListRelation {
	return &NrAlbumBookListRelation{Context: ctx, Handler: handler}
}

/*NrAlbumBookListRelation swagger:route GET /relation/album/bookList Relation albumBookListRelation

获取专辑下书本集合

获取专辑下书本集合

*/
type NrAlbumBookListRelation struct {
	Context *middleware.Context
	Handler NrAlbumBookListRelationHandler
}

func (o *NrAlbumBookListRelation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrAlbumBookListRelationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}