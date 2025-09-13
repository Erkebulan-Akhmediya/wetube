package users

import (
	"net/http"
	authMiddleware "wetube/auth/middleware"
	"wetube/role"
	"wetube/users/controller"
	"wetube/users/middleware"
	"wetube/utils"
)

func RegisterRoutes() {
	registerDetailHandlers()
	registerRestoreHandler()
}

func registerDetailHandlers() {
	deleteByIdHandler := controller.NewDeleteByIdHandler()
	updateByIdHandler := controller.NewUpdateByIdHandler()

	var userHandler http.Handler = utils.MethodHandler{
		http.MethodGet:    controller.NewGetByIdHandler(),
		http.MethodDelete: middleware.NewSelfOrAdminMiddleware(deleteByIdHandler),
		http.MethodPut:    middleware.NewSelfOrAdminMiddleware(updateByIdHandler),
	}
	userHandler = middleware.NewURLUserMiddleware(userHandler)
	userHandler = authMiddleware.NewAuthMiddleware(userHandler)
	http.Handle("/users/{userId}", userHandler)
}

func registerRestoreHandler() {
	restoreHandler := controller.NewRestoreHandler()
	restoreHandler = role.NewRoleMiddleware([]string{"admin"}, restoreHandler)
	restoreHandler = authMiddleware.NewAuthMiddleware(restoreHandler)
	restoreHandler = utils.MethodHandler{
		http.MethodPatch: restoreHandler,
	}
	http.Handle("/users/{userId}/restore", restoreHandler)
}
