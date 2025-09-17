package service

import (
	"wetube/database"
	fileService "wetube/files/service"
	"wetube/video/controller/dto"
)

type Video struct {
	Id          int
	Name        string
	Description string
	File        string
	ChannelId   int
}

func Create(dto *dto.VideoDto) error {
	fileName, err := fileService.Upload(dto.File, dto.FileHeader)
	if err != nil {
		return err
	}
	video := Video{
		Name:        dto.Name,
		Description: dto.Description,
		File:        fileName,
		ChannelId:   dto.ChannelId,
	}
	return create(&video)
}

func create(video *Video) error {
	query := `INSERT INTO video (name, description, channel_id, file) VALUES ($1, $2, $3, $4)`
	_, err := database.Db().Exec(query, video.Name, video.Description, video.ChannelId, video.File)
	return err
}
