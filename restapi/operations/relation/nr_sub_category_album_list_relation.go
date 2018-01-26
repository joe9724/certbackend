// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrSubCategoryAlbumListRelationHandlerFunc turns a function with the right signature into a sub category album list relation handler
type NrSubCategoryAlbumListRelationHandlerFunc func(NrSubCategoryAlbumListRelationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrSubCategoryAlbumListRelationHandlerFunc) Handle(params NrSubCategoryAlbumListRelationParams) middleware.Responder {
	return fn(params)
}

// NrSubCategoryAlbumListRelationHandler interface for that can handle valid sub category album list relation params
type NrSubCategoryAlbumListRelationHandler interface {
	Handle(NrSubCategoryAlbumListRelationParams) middleware.Responder
}

// NewNrSubCategoryAlbumListRelation creates a new http.Handler for the sub category album list relation operation
func NewNrSubCategoryAlbumListRelation(ctx *middleware.Context, handler NrSubCategoryAlbumListRelationHandler) *NrSubCategoryAlbumListRelation {
	return &NrSubCategoryAlbumListRelation{Context: ctx, Handler: handler}
}

/*NrSubCategoryAlbumListRelation swagger:route GET /relation/subCategory/albumList Relation subCategoryAlbumListRelation

获取子类下专辑集合

获取子类下专辑集合

*/
type NrSubCategoryAlbumListRelation struct {
	Context *middleware.Context
	Handler NrSubCategoryAlbumListRelationHandler
}

func (o *NrSubCategoryAlbumListRelation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrSubCategoryAlbumListRelationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}