package controller

import (
	"time"
	"wetube/users/service"
)

type userDto struct {
	Id        int    `json:"id,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"createdAt"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

func newUserDto(user *service.User) *userDto {
	dto := userDto{
		Id:        user.Id,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format(time.DateOnly),
	}
	if user.DeletedAt.Valid {
		dto.DeletedAt = user.DeletedAt.Time.Format(time.DateOnly)
	}
	return &dto
}
