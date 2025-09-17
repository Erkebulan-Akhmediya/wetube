package controller

import "mime/multipart"

type videoDto struct {
	Name        string
	Description string
	File        multipart.File
	FileHeader  *multipart.FileHeader
}
