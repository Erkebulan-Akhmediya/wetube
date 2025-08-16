package role

import (
	"net/http"
	"wetube/auth/middleware"
	"wetube/role/controller"
)

func RegisterRoutes() {
	http.HandleFunc("/roles", middleware.Auth(controller.GetAll))
}
