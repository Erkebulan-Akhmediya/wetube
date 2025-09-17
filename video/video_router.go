package video

import (
	"net/http"
	authMiddleware "wetube/auth/middleware"
	"wetube/utils"
	"wetube/video/controller"
)

func RegisterRoutes() {
	registerVideoHandlers()
}

func registerVideoHandlers() {
	uploadHandler := controller.NewUploadHandler()
	uploadHandler = authMiddleware.NewAuthMiddleware(uploadHandler)
	videoHandlers := utils.MethodHandler{
		http.MethodPost: uploadHandler,
	}
	http.Handle("/video", videoHandlers)
}
