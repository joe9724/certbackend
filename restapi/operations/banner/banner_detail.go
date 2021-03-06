// Code generated by go-swagger; DO NOT EDIT.

package banner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// BannerDetailHandlerFunc turns a function with the right signature into a banner detail handler
type BannerDetailHandlerFunc func(BannerDetailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BannerDetailHandlerFunc) Handle(params BannerDetailParams) middleware.Responder {
	return fn(params)
}

// BannerDetailHandler interface for that can handle valid banner detail params
type BannerDetailHandler interface {
	Handle(BannerDetailParams) middleware.Responder
}

// NewBannerDetail creates a new http.Handler for the banner detail operation
func NewBannerDetail(ctx *middleware.Context, handler BannerDetailHandler) *BannerDetail {
	return &BannerDetail{Context: ctx, Handler: handler}
}

/*BannerDetail swagger:route GET /banner/detail Banner bannerDetail

Banner详情

Banner详情

*/
type BannerDetail struct {
	Context *middleware.Context
	Handler BannerDetailHandler
}

func (o *BannerDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBannerDetailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
