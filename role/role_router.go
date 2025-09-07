package role

import (
	"net/http"
	authMiddleware "wetube/auth/middleware"
	"wetube/role/controller"
	"wetube/utils"
)

func RegisterRoutes() {
	registerUpdateRoute()
}

func registerUpdateRoute() {
	updateHandler := controller.NewUpdateByUserIdHandler()
	updateHandler = NewRoleMiddleware([]string{"admin"}, updateHandler)
	updateHandler = authMiddleware.NewAuthMiddleware(updateHandler)
	updateHandler = utils.MethodHandler{
		http.MethodPut: updateHandler,
	}
	http.Handle("/roles/{userId}", updateHandler)
}
