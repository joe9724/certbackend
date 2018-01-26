// Code generated by go-swagger; DO NOT EDIT.

package chapter

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

// ChapterDetailHandlerFunc turns a function with the right signature into a chapter detail handler
type ChapterDetailHandlerFunc func(ChapterDetailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ChapterDetailHandlerFunc) Handle(params ChapterDetailParams) middleware.Responder {
	return fn(params)
}

// ChapterDetailHandler interface for that can handle valid chapter detail params
type ChapterDetailHandler interface {
	Handle(ChapterDetailParams) middleware.Responder
}

// NewChapterDetail creates a new http.Handler for the chapter detail operation
func NewChapterDetail(ctx *middleware.Context, handler ChapterDetailHandler) *ChapterDetail {
	return &ChapterDetail{Context: ctx, Handler: handler}
}

/*ChapterDetail swagger:route GET /chapter/detail Chapter chapterDetail

章节详情

章节详情

*/
type ChapterDetail struct {
	Context *middleware.Context
	Handler ChapterDetailHandler
}

func (o *ChapterDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewChapterDetailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok ChapterDetailOK
	var response models.InlineResponse20025

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	var chapter models.Chapter
	db.Table("chapters").Where("id=?",Params.ChapterID).First(&chapter)
	//data
	response.Data = &chapter


	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
