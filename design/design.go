package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("my-clippings", func() {
	Description("This API exposes an image resource that allows uploading and downloading images")
	BasePath("/api")
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		Headers("Accept", "Content-Type")
		Expose("Content-Type", "Origin")
		MaxAge(600)
	})
})

var _ = Resource("clippings", func() {
	BasePath("/clippings")

	Action("upload", func() {
		Routing(POST("/"))
		Description("Upload My Clippings file")
		Response(OK, ClippingsMedia)
	})

	Files("/download/*filename", "/tmp/my-clippings/") // Serve files from the "images" directory
})

var ClippingsMedia = MediaType("application/vnd.upload.clipping", func() {
	Description("My Clippings metadata")
	TypeName("ClippingsMedia")
	Attributes(func() {
		Attribute("id", String, "File ID")
	})
	View("default", func() {
		Attribute("id")
	})
})
