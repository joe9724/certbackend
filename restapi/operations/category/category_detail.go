// Code generated by go-swagger; DO NOT EDIT.

package category

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

// CategoryDetailHandlerFunc turns a function with the right signature into a category detail handler
type CategoryDetailHandlerFunc func(CategoryDetailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CategoryDetailHandlerFunc) Handle(params CategoryDetailParams) middleware.Responder {
	return fn(params)
}

// CategoryDetailHandler interface for that can handle valid category detail params
type CategoryDetailHandler interface {
	Handle(CategoryDetailParams) middleware.Responder
}

// NewCategoryDetail creates a new http.Handler for the category detail operation
func NewCategoryDetail(ctx *middleware.Context, handler CategoryDetailHandler) *CategoryDetail {
	return &CategoryDetail{Context: ctx, Handler: handler}
}

/*CategoryDetail swagger:route GET /category/detail Category categoryDetail

分类详情

分类详情

*/
type CategoryDetail struct {
	Context *middleware.Context
	Handler CategoryDetailHandler
}

func (o *CategoryDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCategoryDetailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok CategoryDetailOK
	var response models.InlineResponse20017

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	var category models.Category
	db.Table("sub_category_items").Where("id=?",Params.CategoryID).First(&category)
	//data
	response.Data = &category

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
