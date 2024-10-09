package models

type User struct {
	Id       int    `json:"-"`
	Login    string `json:"login" binding: "required"`
	Password string `json:"password" binding: "required"`
	Role     int    `json:"role" binding: "required"`
}
