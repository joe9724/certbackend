// Code generated by go-swagger; DO NOT EDIT.

package chapter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ChapterEditHandlerFunc turns a function with the right signature into a chapter edit handler
type ChapterEditHandlerFunc func(ChapterEditParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ChapterEditHandlerFunc) Handle(params ChapterEditParams) middleware.Responder {
	return fn(params)
}

// ChapterEditHandler interface for that can handle valid chapter edit params
type ChapterEditHandler interface {
	Handle(ChapterEditParams) middleware.Responder
}

// NewChapterEdit creates a new http.Handler for the chapter edit operation
func NewChapterEdit(ctx *middleware.Context, handler ChapterEditHandler) *ChapterEdit {
	return &ChapterEdit{Context: ctx, Handler: handler}
}

/*ChapterEdit swagger:route POST /chapter/edit Chapter chapterEdit

编辑一个章节

*/
type ChapterEdit struct {
	Context *middleware.Context
	Handler ChapterEditHandler
}

func (o *ChapterEdit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewChapterEditParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
