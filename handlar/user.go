package handlar

import (
	"github.com/gin-gonic/gin"
	"github.com/katsumeshi/tweetn-backend/model"
	"github.com/katsumeshi/tweetn-backend/service"
	"github.com/gin-contrib/sessions"
	"net/http"
)

type User struct {
	userService service.User
}

func InitUserHandler(userService service.User) *User {
	return &User{userService: userService}
}

func  (u User) Login(c *gin.Context) {
	var loginUser model.User
	c.Bind(&loginUser)

	user, err := u.userService.FindFirst(loginUser.Username)
	isNotFoundUser := err != nil

	if isNotFoundUser {
		c.JSON(200, model.Error{3, "Not found user"})
	} else {
		loginUser = user
		session := sessions.Default(c)
		v := session.Get("userId")
		var userId int
		if v == nil {
			userId = loginUser.Id
			session.Set("userId", userId)
			session.Save()
		} else {
			userId = v.(int)
		}
		c.Redirect(http.StatusMovedPermanently, "/v1/tweets")
	}
}