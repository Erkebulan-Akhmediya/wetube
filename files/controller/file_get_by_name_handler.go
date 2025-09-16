package controller

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"wetube/files/service"

	"github.com/minio/minio-go/v7"
)

func NewGetByNameHandler() http.Handler {
	return &getByNameHandler{}
}

type getByNameHandler struct{}

func (gh *getByNameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	obj, err := service.Client().GetObject(
		context.Background(),
		os.Getenv("MINIO_BUCKET_NAME"),
		name,
		minio.GetObjectOptions{},
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get file", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err = obj.Close(); err != nil {
			log.Println(err)
		}
	}()

	info, err := obj.Stat()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", info.ContentType)
	w.Header().Set("Content-Length", strconv.FormatInt(info.Size, 10))
	w.Header().Set("Content-Disposition", "attachment; filename=\""+name+"\"")

	if _, err = io.Copy(w, obj); err != nil {
		log.Println(err)
		http.Error(w, "Error sending file", http.StatusInternalServerError)
		return
	}
}
