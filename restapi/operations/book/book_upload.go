// Code generated by go-swagger; DO NOT EDIT.

package book

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

// BookUploadHandlerFunc turns a function with the right signature into a book upload handler
type BookUploadHandlerFunc func(BookUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BookUploadHandlerFunc) Handle(params BookUploadParams) middleware.Responder {
	return fn(params)
}

// BookUploadHandler interface for that can handle valid book upload params
type BookUploadHandler interface {
	Handle(BookUploadParams) middleware.Responder
}

// NewBookUpload creates a new http.Handler for the book upload operation
func NewBookUpload(ctx *middleware.Context, handler BookUploadHandler) *BookUpload {
	return &BookUpload{Context: ctx, Handler: handler}
}

/*BookUpload swagger:route POST /book/upload Book bookUpload

添加一本书

*/
type BookUpload struct {
	Context *middleware.Context
	Handler BookUploadHandler
}

func (o *BookUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBookUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok BookUploadOK
	var response models.InlineResponse20016
	var status models.Response
	var book models.Book
	var msg string
	var code int64

	msg = "ok"
	code = 200

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	if(Params.IconUrl != ""){
		fmt.Println(Params.IconUrl)
		book.Icon = Params.IconUrl
	}

	//tt:= int64(-1)
	fmt.Println("Params.BookId=",*(Params.BookId))
	if(*(Params.BookId) == -1){ //新建
		fmt.Println("new")
		fmt.Println("Params.Summary is",Params.Summary)
		book.Summary = *(Params.Summary)
		book.Name = Params.Title
		book.AuthorName = Params.AuthorName
		fmt.Println("author is",Params.AuthorName)
		//t := int64(-1)
		//book.book_Id = &t
		book.Status = *(Params.Status)
		if(Params.IconUrl != ""){
			book.Icon = Params.IconUrl
		}
		//book.User_id = *(Params.MemberID)
		db.Table("books").Create(&book)
	}else{ //更新
		//fmt.Println("edit")
		//db.Table("sub_book_items").Where("id=?",*(Params.BookId)).Last(&book)
		if(Params.IconUrl != ""){
			fmt.Println("1",Params.IconUrl)
			db.Exec("update books set name=?,status=?,summary=?,icon=?,author_name=? where id=?",Params.Title,0,*(Params.Summary),Params.IconUrl,Params.AuthorName,&(Params.BookId))
		}else{
			fmt.Println("2",Params.IconUrl,*(Params.Summary))
			summary := *(Params.Summary)
			db.Exec("update books set name=?,status=?,summary=?,author_name=? where id=?",Params.Title,0,summary,Params.AuthorName,&(Params.BookId))
		}

	}

	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Return = &status
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
