package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/localghost/my-clippings/app"
	"log"
	"os"
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
		input, _ := ctx.MultipartForm.File["hela"][0].Open()
		defer input.Close()

		fmt.Println("INPUT OPENED")

		clippings, err := New().Parse(input)
		if err != nil {
			fmt.Println(err)
			return ctx.OK(&app.ClippingsMedia{})
		}
		fmt.Println("PARSED")

		output, _ := os.Create(fmt.Sprintf("/tmp/my-clippings/%s.json", ctx.MultipartForm.File["hela"][0].Filename))
		defer output.Close()
		err = json.NewEncoder(output).Encode(clippings)
		if err != nil {
			fmt.Println(err)
			return ctx.OK(&app.ClippingsMedia{})
		}

		filename := fmt.Sprintf("%s.json", ctx.MultipartForm.File["hela"][0].Filename)
		return ctx.OK(&app.ClippingsMedia{ID: &filename})
	}
	//buffer := &bytes.Buffer{}
	//io.Copy(buffer, ctx.Body)
	//ctx.Body.Close()
	//log.Println(buffer)

	// ClippingsController_Upload: end_implement
	return ctx.OK(&app.ClippingsMedia{})
}
