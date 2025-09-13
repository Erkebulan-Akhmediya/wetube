package channel

import (
	"net/http"
	"wetube/auth/middleware"
	"wetube/channel/controller"
	"wetube/utils"
)

func RegisterRoutes() {
	createHandler := controller.NewCreateHandler()
	createHandler = middleware.NewAuthMiddleware(createHandler)
	channelHandler := utils.MethodHandler{
		http.MethodPost: createHandler,
	}
	http.Handle("/channel", channelHandler)
}
