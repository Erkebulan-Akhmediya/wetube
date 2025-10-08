package dto

import "mime/multipart"

type UploadVideoDto struct {
	ChannelId   int
	Name        string
	Description string
	File        multipart.File
	FileHeader  *multipart.FileHeader
}

type GetVideoDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	File        string `json:"file"`
	ChannelId   int    `json:"channelId"`
}
