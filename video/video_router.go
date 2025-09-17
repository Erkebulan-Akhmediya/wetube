package video

import (
	"net/http"
	authMiddleware "wetube/auth/middleware"
	channelMiddleware "wetube/channel/middleware"
	"wetube/utils"
	"wetube/video/controller"
)

func RegisterRoutes() {
	registerVideoHandlers()
}

func registerVideoHandlers() {
	uploadHandler := controller.NewUploadHandler()
	uploadHandler = channelMiddleware.NewIsOwnerMiddleware(uploadHandler)
	uploadHandler = authMiddleware.NewAuthMiddleware(uploadHandler)
	videoHandlers := utils.MethodHandler{
		http.MethodPost: uploadHandler,
	}
	http.Handle("/channel/{channelId}/video", videoHandlers)
}
