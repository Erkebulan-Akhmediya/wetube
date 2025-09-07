package controller

type authDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwtDto struct {
	Token string `json:"token"`
	Id    int    `json:"id"`
}
