package users

import (
	"net/http"
	"wetube/auth/middleware"
	"wetube/users/controller"
	"wetube/utils"
)

func RegisterRoutes() {
	userHandler := utils.MethodHandler{
		http.MethodGet:    controller.NewGetByIdHandler(),
		http.MethodDelete: controller.NewDeleteByIdHandler(),
		http.MethodPut:    controller.NewUpdateByIdHandler(),
	}
	http.Handle("/users/{userId}", middleware.NewAuthMiddleware(userHandler))
	http.Handle("/users/{userId}/restore", middleware.NewAuthMiddleware(controller.NewRestoreHandler()))
}
