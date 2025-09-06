package channel

import (
	"net/http"
	"wetube/auth/middleware"
	"wetube/channel/controller"
)

func RegisterRoutes() {
	http.HandleFunc("/channel", middleware.AuthFunc(controller.Create))
}
