// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
	"fmt"
	_"os"
	"certbackend/models"
	"certbackend/var"
)

// AlbumUploadHandlerFunc turns a function with the right signature into a album upload handler
type AlbumUploadHandlerFunc func(AlbumUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AlbumUploadHandlerFunc) Handle(params AlbumUploadParams) middleware.Responder {
	return fn(params)
}

// AlbumUploadHandler interface for that can handle valid album upload params
type AlbumUploadHandler interface {
	Handle(AlbumUploadParams) middleware.Responder
}

// NewAlbumUpload creates a new http.Handler for the album upload operation
func NewAlbumUpload(ctx *middleware.Context, handler AlbumUploadHandler) *AlbumUpload {
	return &AlbumUpload{Context: ctx, Handler: handler}
}

/*AlbumUpload swagger:route POST /album/upload Album albumUpload

添加一个专辑

*/
type AlbumUpload struct {
	Context *middleware.Context
	Handler AlbumUploadHandler
}

func (o *AlbumUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAlbumUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok AlbumUploadOK
	var response models.InlineResponse20016
	var status models.Response
	var album models.Album
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
		album.Icon = Params.IconUrl
	}

	//tt:= int64(-1)
	fmt.Println("Params.AlbumId=",*(Params.AlbumId))
	if(*(Params.AlbumId) == -1){ //新建
		fmt.Println("new")
		fmt.Println("Params.Summary is",Params.Summary)
		album.Summary = *(Params.Summary)
		album.Name = Params.Title
		//t := int64(-1)
		//album.album_Id = &t
		album.Status = *(Params.Status)
		if(Params.IconUrl != ""){
			album.Icon = Params.IconUrl
		}
		//album.User_id = *(Params.MemberID)
		db.Table("albums").Create(&album)
	}else{ //更新
		//fmt.Println("edit")
		//db.Table("sub_album_items").Where("id=?",*(Params.AlbumId)).Last(&album)
		if(Params.IconUrl != ""){
			fmt.Println("1",Params.IconUrl)
			db.Exec("update albums set name=?,status=?,summary=?,icon=? where id=?",Params.Title,0,*(Params.Summary),Params.IconUrl,&(Params.AlbumId))
		}else{
			fmt.Println("2",Params.IconUrl,*(Params.Summary))
			summary := *(Params.Summary)
			db.Exec("update albums set name=?,status=?,summary=? where id=?",Params.Title,0,summary,&(Params.AlbumId))
		}

	}

	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Return = &status
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)

}
