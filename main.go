//go:generate goagen bootstrap -d github.com/localghost/my-clippings/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/localghost/my-clippings/app"
)

func main() {
	// Create service
	service := goa.New("my-clippings")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	//app.MountCategoryController(service, NewCategoryController(service))
	//app.MountPaymentController(service, NewPaymentController(service))
	//app.MountTransferController(service, NewTransferController(service))
	//app.UseJWTMiddleware(service, jwt.New(jwt.NewSimpleResolver([]jwt.Key{"jwt"}), nil, app.NewJWTSecurity()))
	app.MountClippingsController(service, NewClippingsController(service))

	// Start service
	if err := service.ListenAndServe(":8111"); err != nil {
		service.LogError("startup", "err", err)
	}
}
