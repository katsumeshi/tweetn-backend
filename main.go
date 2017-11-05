package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/katsumeshi/tweetn-backend/dao"
	"github.com/katsumeshi/tweetn-backend/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/katsumeshi/tweetn-backend/service"
	"github.com/katsumeshi/tweetn-backend/handlar"
)

var isHeroku = false
var db *gorm.DB

func main() {
	GetMainEngine().Run()
}

func GetMainEngine() *gin.Engine {

	b, err := strconv.ParseBool(os.Getenv("IS_HEROKU"))
	if err == nil {
		isHeroku = b
	}
	db = gormConnect()

	r := gin.Default()

	if isHeroku {
		u, err := url.Parse(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err)
		}
		p, _ := u.User.Password()
		store, _ := sessions.NewRedisStore(10, "tcp", u.Host, p, []byte("secret"))
		r.Use(sessions.Sessions("session", store))
	} else {
		store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
		r.Use(sessions.Sessions("session", store))
	}

	userDao := dao.InitUserDao(db);
	userService := service.InitUserService(userDao);
	userHandler := handlar.InitUserHandler(userService);



	r.LoadHTMLGlob("templates/*")
	r.GET("/", Entrance)
	v1 := r.Group("/v1")
	{
		v1.GET("/login", GetLoginView)
		v1.POST("/login", userHandler.Login)
		v1.POST("/logout", Logout)

		v1.GET("/users", GetUserCreationView)
		v1.POST("/users", CreateUser)
		v1.GET("/users/:username", ShowUser)

		v1.POST("/tweets", PostTweets)
		v1.GET("/tweets", ShowTweets)
	}

	return r
}

func Entrance(c *gin.Context) {
	userId := getSessionUserId(c)
	fmt.Printf("%v", userId)
	if 0 <= userId {
		c.Redirect(302, "/v1/tweets")
	} else {
		c.Redirect(302, "/v1/login")
	}
}

func ShowUser(c *gin.Context) {
	username := c.Param("username")
	user := model.User{}
	if db.First(&user, "username = ?", username).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"message": "can't find the user"})
	} else {
		c.HTML(http.StatusOK, "show.tmpl", user)
	}
}

func GetUserCreationView(c *gin.Context) {
	c.HTML(http.StatusOK, "account.tmpl", model.Error{})
}

func CreateUser(c *gin.Context) {
	user := model.User{}
	c.Bind(&user)
	if 0 < len(user.Username) {
		if canCreateUser(user.Username) {
			db.Create(&user)
			c.Redirect(http.StatusMovedPermanently, "/v1/users/"+user.Username)
		} else {
			c.HTML(http.StatusOK, "account.tmpl", model.Error{0, "Your account is already exisited"})
		}
	} else {
		c.HTML(http.StatusOK, "account.tmpl",model. Error{1, "At least you need to fill in username"})
	}
}

func canCreateUser(username string) bool {
	user := model.User{}
	if db.First(&user, "username = ?", username).RecordNotFound() {
		return true
	}
	return false
}

func GetLoginView(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}


func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userId")
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/v1/login")
}

func getSessionUserId(c *gin.Context) int {
	session := sessions.Default(c)
	v := session.Get("userId")
	fmt.Printf("%v", v)
	if v == nil {
		return -1
	} else {
		return v.(int)
	}
}

func GetTweetView(c *gin.Context) {
	c.HTML(http.StatusOK, "new.tmpl", nil)
}

func PostTweets(c *gin.Context) {
	var tweet model.Tweet
	c.Bind(&tweet)

	userId := getSessionUserId(c)
	tweet.UserId = userId
	db.Create(&tweet)

	c.Redirect(http.StatusMovedPermanently, "/v1/tweets")
}

func ShowTweets(c *gin.Context) {

	session := sessions.Default(c)
	userId := session.Get("userId")
	if userId == nil {
		c.Redirect(http.StatusMovedPermanently, "/v1/login")
	}

	tweets := []model.Tweet{}
	db.Preload("User").Order("Id desc").Find(&tweets)

	c.HTML(http.StatusOK, "lists.tmpl", gin.H{"tweets": tweets})
}

// DB ---------------------------------------------
func gormConnect() *gorm.DB {

	dbUrl := os.Getenv("CLEARDB_DATABASE_URL")
	u, err := url.Parse(dbUrl)
	if err != nil {
		panic(err)
	}
	var scheme string = "mysql"
	var connectInfo string = "root:@tcp(127.0.0.1:3306)/development"
	if isHeroku {
		scheme = u.Scheme
		connectInfo = u.User.String() + "@tcp(" + u.Host + ":3306)" + u.Path
	}
	gormDB, err := gorm.Open(scheme, connectInfo)
	if err != nil {
		fmt.Printf(scheme + "\n")
		fmt.Printf(connectInfo + "\n")
		fmt.Printf("can't connect db")
		panic("failed to connect database")
	}
	return gormDB
}

