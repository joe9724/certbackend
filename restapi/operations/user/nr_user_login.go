// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NrUserLoginHandlerFunc turns a function with the right signature into a user login handler
type NrUserLoginHandlerFunc func(NrUserLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserLoginHandlerFunc) Handle(params NrUserLoginParams) middleware.Responder {
	return fn(params)
}

// NrUserLoginHandler interface for that can handle valid user login params
type NrUserLoginHandler interface {
	Handle(NrUserLoginParams) middleware.Responder
}

// NewNrUserLogin creates a new http.Handler for the user login operation
func NewNrUserLogin(ctx *middleware.Context, handler NrUserLoginHandler) *NrUserLogin {
	return &NrUserLogin{Context: ctx, Handler: handler}
}

/*NrUserLogin swagger:route GET /user/login User userLogin

登录接口

登录接口

*/
type NrUserLogin struct {
	Context *middleware.Context
	Handler NrUserLoginHandler
}

func (o *NrUserLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserLoginParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok UserLoginOK
	var response models.InlineResponse20029

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	var user models.User
	db.Where("name=?","admin").First(&user)
	//data
	response.Data = &user

	//routers
	/*var routers []models.Router
	//routers = append(routers, models.Router{1,"用户列表","/user/list"})
	//routers = append(routers, models.Router{1,"章节列表","/chapter/list"})
	var temp models.Router
	id := int64(1)
	name := "章节列表"
	routerUrl := "/chapter/list"
	temp.Name = &name
	temp.ID = &id
	temp.RouterURL = &routerUrl
	routers = append(routers, temp)
	//
	id1 := int64(2)
	name1 := "管理员列表"
	routerUrl1 := "/user/list"
	temp.Name = &name1
	temp.ID = &id1
	temp.RouterURL = &routerUrl1
	routers = append(routers, temp)
	//
	id2 := int64(3)
	name2 := "书本列表"
	routerUrl2 := "/book/list"
	temp.Name = &name2
	temp.ID = &id2
	temp.RouterURL = &routerUrl2
	routers = append(routers, temp)
	//
	id3 := int64(4)
	name3 := "专辑列表"
	routerUrl3 := "/album/list"
	temp.Name = &name3
	temp.ID = &id3
	temp.RouterURL = &routerUrl3
	routers = append(routers, temp)

	//
	id4 := int64(5)
	name4 := "会员列表"
	routerUrl4 := "/member/list"
	temp.Name = &name4
	temp.ID = &id4
	temp.RouterURL = &routerUrl4
	routers = append(routers, temp)

	//
	id5 := int64(6)
	name5 := "订单列表"
	routerUrl5 := "/order/list"
	temp.Name = &name5
	temp.ID = &id5
	temp.RouterURL = &routerUrl5
	routers = append(routers, temp)

	//
	id6 := int64(7)
	name6 := "子类列表"
	routerUrl6 := "/subCategory/list"
	temp.Name = &name6
	temp.ID = &id6
	temp.RouterURL = &routerUrl6
	routers = append(routers, temp)

	//
	id7 := int64(8)
	name7 := "大类列表"
	routerUrl7 := "/category/list"
	temp.Name = &name7
	temp.ID = &id7
	temp.RouterURL = &routerUrl7
	routers = append(routers, temp)

	response.Routers = routers
*/
	//get routers from sql
	db1,err1 := _var.OpenConnection()
	if err1!=nil{
		fmt.Println(err1.Error())
	}
	//query
	var dynamicRouters [] models.DynamicRouter
	db1.Table("routers").Where(map[string]interface{}{"status":0}).Find(&dynamicRouters)

	for i:=0; i<len(dynamicRouters); i++  {
		var childrenRouters [] models.ChildrenRouter
		db1.Table("routers").Where(map[string]interface{}{"status":0}).Where("parent_id=?",dynamicRouters[i].ID).Find(&childrenRouters)
		dynamicRouters[i].Children = &childrenRouters
	}

	//data
	response.DynamicRouters = dynamicRouters


	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status


	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}