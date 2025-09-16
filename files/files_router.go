package files

import (
	"net/http"
	"wetube/files/controller"
)

func RegisterRoutes() {
	registerGetHandler()
}

func registerGetHandler() {
	getByNameHandler := controller.NewGetByNameHandler()
	http.Handle("/files/{name}", getByNameHandler)
}
