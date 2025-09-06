package users

import (
	"net/http"
	"wetube/auth/middleware"
	"wetube/users/controller"
	"wetube/utils"
)

func RegisterRoutes() {
	userHandler := utils.MethodHandler{
		http.MethodGet:    controller.GetById,
		http.MethodDelete: controller.DeleteUser,
		http.MethodPut:    controller.Update,
	}
	http.Handle("/users/{userId}", middleware.NewAuthMiddleware(userHandler))
	http.HandleFunc("/users/{userId}/restore", middleware.AuthFunc(controller.Restore))
}
