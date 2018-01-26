// Code generated by go-swagger; DO NOT EDIT.

package recharge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	middleware "github.com/go-openapi/runtime/middleware"
	"tingtingbackend/models"
	"fmt"
	"tingtingbackend/var"
)

// RechargeListHandlerFunc turns a function with the right signature into a recharge list handler
type RechargeListHandlerFunc func(RechargeListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RechargeListHandlerFunc) Handle(params RechargeListParams) middleware.Responder {
	return fn(params)
}

// RechargeListHandler interface for that can handle valid recharge list params
type RechargeListHandler interface {
	Handle(RechargeListParams) middleware.Responder
}

// NewRechargeList creates a new http.Handler for the recharge list operation
func NewRechargeList(ctx *middleware.Context, handler RechargeListHandler) *RechargeList {
	return &RechargeList{Context: ctx, Handler: handler}
}

/*RechargeList swagger:route POST /recharge/list Recharge rechargeList

获取充值流水列表(含条件检索)

获取订单列表

*/
type RechargeList struct {
	Context *middleware.Context
	Handler RechargeListHandler
}

func (o *RechargeList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRechargeListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok RechargeListOK
	var response models.InlineResponse2003
	var rechargeList models.InlineResponse2003Orders

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	//db.Table("recharge").Where(map[string]interface{}{"status":0}).Find(&rechargeList).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
	db.Table("recharge").Select("recharge.id, recharge.memberId,recharge.type,recharge.order_no,recharge.status,recharge.time,recharge.value,members.name").Joins("left join members on recharge.memberId = members.id").Find(&rechargeList)
	//data
	response.Orders = rechargeList
	//fmt.Println("haspushed is",albumList[0].HasPushed)
	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}