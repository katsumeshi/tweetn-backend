package model

type User struct {
	Id       int    `json:"id"`
	Name     string `form:"name" json:"name"`
	Username string `form:"username" json:"username"`
	Location string `form:"location" json:"location"`
	About    string `form:"about" json:"about"`
}
