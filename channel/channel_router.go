package channel

import (
	"net/http"
	"wetube/auth/middleware"
	"wetube/channel/controller"
)

func RegisterRoutes() {
	http.Handle("/channel", middleware.NewAuthMiddleware(controller.NewCreateHandler()))
}
