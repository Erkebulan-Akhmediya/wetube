package files

import (
	"net/http"
	"wetube/files/controller"
	"wetube/utils"
)

func RegisterRoutes() {
	registerGetHandler()
}

func registerGetHandler() {
	getByNameHandler := controller.NewGetByNameHandler()
	getByNameHandler = utils.MethodHandler{
		http.MethodGet: getByNameHandler,
	}
	http.Handle("/files/{name}", getByNameHandler)
}
