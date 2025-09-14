package service

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
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

func GenerateUniqueName(file io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	name := fileHeader.Filename
	ext := filepath.Ext(name)
	var fileBytes bytes.Buffer
	if _, err := io.Copy(&fileBytes, file); err != nil {
		return "", err
	}
	return uuid.NewSHA1(uuid.NameSpaceDNS, fileBytes.Bytes()).String() + ext, nil
}
