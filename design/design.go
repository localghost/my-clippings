package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("my-clippings", func() {
	Description("This API exposes an image resource that allows uploading and downloading images")
	BasePath("/api")
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
		Attribute("id", Integer, "Image ID")
		Attribute("filename", String, "Image filename")
		Attribute("uploaded_at", DateTime, "Upload timestamp")
		Required("id", "filename", "uploaded_at")
	})
	View("default", func() {
		Attribute("id")
		Attribute("filename")
		Attribute("uploaded_at")
	})
})
