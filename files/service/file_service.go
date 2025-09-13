package service

import (
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client

func InitClient() {
	var err error
	minioHost := os.Getenv("MINIO_HOST")
	minioPort := os.Getenv("MINIO_PORT")
	minioAddr := fmt.Sprintf("%s:%s", minioHost, minioPort)
	minioUser := os.Getenv("MINIO_USER")
	minioPassword := os.Getenv("MINIO_PASSWORD")
	minioOptions := &minio.Options{
		Creds: credentials.NewStaticV4(minioUser, minioPassword, ""),
	}
	if client, err = minio.New(minioAddr, minioOptions); err != nil {
		log.Fatal(err)
	}
}

func Client() *minio.Client {
	return client
}
