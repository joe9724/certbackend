// Code generated by go-swagger; DO NOT EDIT.

package msg

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// MsgUnsendListHandlerFunc turns a function with the right signature into a msg unsend list handler
type MsgUnsendListHandlerFunc func(MsgUnsendListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MsgUnsendListHandlerFunc) Handle(params MsgUnsendListParams) middleware.Responder {
	return fn(params)
}

// MsgUnsendListHandler interface for that can handle valid msg unsend list params
type MsgUnsendListHandler interface {
	Handle(MsgUnsendListParams) middleware.Responder
}

// NewMsgUnsendList creates a new http.Handler for the msg unsend list operation
func NewMsgUnsendList(ctx *middleware.Context, handler MsgUnsendListHandler) *MsgUnsendList {
	return &MsgUnsendList{Context: ctx, Handler: handler}
}

/*MsgUnsendList swagger:route GET /msg/unsend/list Msg msgUnsendList

未发送消息列表

未发送消息列表

*/
type MsgUnsendList struct {
	Context *middleware.Context
	Handler MsgUnsendListHandler
}

func (o *MsgUnsendList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMsgUnsendListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
