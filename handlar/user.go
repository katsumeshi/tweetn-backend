package handlar

import (
	"github.com/gin-gonic/gin"
	"github.com/katsumeshi/tweetn-backend/model"
	"github.com/katsumeshi/tweetn-backend/service"
	"github.com/katsumeshi/tweetn-backend/session"

	"net/http"
	"github.com/gin-contrib/sessions"
)

type User struct {
	userService service.User
	sessions session.Session
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

		u.sessions = session.InitSession(sessions.Default(c));
		u.sessions.SaveUserId(user.Id)

		c.Redirect(http.StatusMovedPermanently, "/v1/tweets")
	}
}