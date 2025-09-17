package dto

import "mime/multipart"

type VideoDto struct {
	ChannelId   int
	Name        string
	Description string
	File        multipart.File
	FileHeader  *multipart.FileHeader
}
