// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	middleware "github.com/go-openapi/runtime/middleware"
	"certbackend/models"
	"fmt"
	"certbackend/var"
)

// MemberListHandlerFunc turns a function with the right signature into a member list handler
type MemberListHandlerFunc func(MemberListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MemberListHandlerFunc) Handle(params MemberListParams) middleware.Responder {
	return fn(params)
}

// MemberListHandler interface for that can handle valid member list params
type MemberListHandler interface {
	Handle(MemberListParams) middleware.Responder
}

// NewMemberList creates a new http.Handler for the member list operation
func NewMemberList(ctx *middleware.Context, handler MemberListHandler) *MemberList {
	return &MemberList{Context: ctx, Handler: handler}
}

/*MemberList swagger:route POST /member/list Member memberList

获取会员列表(含条件检索)

获取会员列表

*/
type MemberList struct {
	Context *middleware.Context
	Handler MemberListHandler
}

func (o *MemberList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMemberListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok MemberListOK
	var response models.InlineResponse2001
	var memberList models.InlineResponse2001Orders

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	db.Where(map[string]interface{}{"status":0}).Find(&memberList).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
	//data
	response.Orders = memberList

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
