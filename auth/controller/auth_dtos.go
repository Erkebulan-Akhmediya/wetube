package controller

import (
	"mime/multipart"
)

type signInDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwtDto struct {
	Token string `json:"token"`
	Id    int    `json:"id"`
}

type signUpDto struct {
	username  string
	password  string
	pfp       multipart.File
	pfpHeader *multipart.FileHeader
}
