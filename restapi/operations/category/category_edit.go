// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
	"fmt"
	_"os"
	"tingtingbackend/models"
	"tingtingbackend/var"
)

// CategoryEditHandlerFunc turns a function with the right signature into a category edit handler
type CategoryEditHandlerFunc func(CategoryEditParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CategoryEditHandlerFunc) Handle(params CategoryEditParams) middleware.Responder {
	return fn(params)
}

// CategoryEditHandler interface for that can handle valid category edit params
type CategoryEditHandler interface {
	Handle(CategoryEditParams) middleware.Responder
}

// NewCategoryEdit creates a new http.Handler for the category edit operation
func NewCategoryEdit(ctx *middleware.Context, handler CategoryEditHandler) *CategoryEdit {
	return &CategoryEdit{Context: ctx, Handler: handler}
}

/*CategoryEdit swagger:route POST /category/edit Category categoryEdit

编辑大类资料

*/
type CategoryEdit struct {
	Context *middleware.Context
	Handler CategoryEditHandler
}

func (o *CategoryEdit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCategoryEditParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok CategoryUploadOK
	var response models.InlineResponse20016
	var status models.Response
	var category models.Category
	var msg string
	var code int64
	msg = "ok"
	code = 200

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}

	//db.Table("sub_category_items").Where("id=?",Params.CategoryId).First(&category)
	//db.Raw("select * from sub_category_items where id=?",Params.CategoryId).First(&category)

	//如果有icon

	category.Name = &(Params.Title)
	t := int64(-1)
	category.Category_Id = &t
	category.Status = Params.Status
	//album.User_id = *(Params.MemberID)
	if(Params.IconUrl != ""){
		fmt.Println(Params.IconUrl)
		category.Icon = &(Params.IconUrl)
	}
		// status := *(Params.Status)
		categoryId := *(Params.CategoryId)
		db.Exec("update sub_category_items set name=?,status=? where id=?",Params.Title,0,categoryId)

	//db.Table("sub_category_items").Save(&category)
	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Return = &status
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
