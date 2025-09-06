package users

import (
	"net/http"
	authMiddleware "wetube/auth/middleware"
	"wetube/role"
	"wetube/users/controller"
	"wetube/utils"
)

func RegisterRoutes() {
	userHandler := utils.MethodHandler{
		http.MethodGet:    controller.NewGetByIdHandler(),
		http.MethodDelete: controller.NewDeleteByIdHandler(),
		http.MethodPut:    controller.NewUpdateByIdHandler(),
	}
	http.Handle("/users/{userId}", authMiddleware.NewAuthMiddleware(userHandler))

	restoreHandler := controller.NewRestoreHandler()
	restoreHandler = role.NewRoleMiddleware([]string{"admin"}, restoreHandler)
	restoreHandler = authMiddleware.NewAuthMiddleware(restoreHandler)
	http.Handle("/users/{userId}/restore", restoreHandler)
}
