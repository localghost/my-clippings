package main

import (
	"bytes"
	"github.com/goadesign/goa"
	"github.com/localghost/my-clippings/app"
	"io"
	"log"
)

// ClippingsController implements the clippings resource.
type ClippingsController struct {
	*goa.Controller
}

// NewClippingsController creates a clippings controller.
func NewClippingsController(service *goa.Service) *ClippingsController {
	return &ClippingsController{Controller: service.NewController("ClippingsController")}
}

// Upload runs the upload action.
func (c *ClippingsController) Upload(ctx *app.UploadClippingsContext) error {
	// ClippingsController_Upload: start_implement

	// Put your logic here
	err := ctx.ParseMultipartForm(1024)
	//err := ctx.ParseForm()
	if err != nil {
		log.Printf("error: %s", err)
	} else {
		log.Println(ctx.MultipartForm.Value)
		file, _ := ctx.MultipartForm.File["hela"][0].Open()
		defer file.Close()
		buffer := &bytes.Buffer{}
		io.Copy(buffer, file)
		log.Println(buffer)
	}
	//buffer := &bytes.Buffer{}
	//io.Copy(buffer, ctx.Body)
	//ctx.Body.Close()
	//log.Println(buffer)

	// ClippingsController_Upload: end_implement
	res := &app.ClippingsMedia{}
	return ctx.OK(res)
}
