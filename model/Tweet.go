package model

type Tweet struct {
	Id      int
	Content string `form:"content" json:"content"`
	UserId  int    `json:"user_id"`
	User    User
}
