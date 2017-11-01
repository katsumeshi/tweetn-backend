package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"tweetn-background/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var isHeroku = false
var db = gormConnect()

func main() {
	GetMainEngine().Run()
}

func GetMainEngine() *gin.Engine {

	b, err := strconv.ParseBool(os.Getenv("IS_HEROKU"))
	if err == nil {
		isHeroku = b
	}

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

	r.LoadHTMLGlob("templates/*")
	r.GET("/", Entrance)

	v1 := r.Group("/v1")
	{
		v1.GET("/login", GetLoginView)
		v1.POST("/login", Login)
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
	c.HTML(http.StatusOK, "account.tmpl", Error{})
}

func CreateUser(c *gin.Context) {
	user := model.User{}
	c.Bind(&user)
	if 0 < len(user.Username) {
		if canCreateUser(user.Username) {
			db.Create(&user)
			c.Redirect(http.StatusMovedPermanently, "/v1/users/"+user.Username)
		} else {
			c.HTML(http.StatusOK, "account.tmpl", Error{0, "Your account is already exisited"})
		}
	} else {
		c.HTML(http.StatusOK, "account.tmpl", Error{1, "At least you need to fill in username"})
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

func Login(c *gin.Context) {
	var loginUser model.User
	c.Bind(&loginUser)

	users := []model.User{}
	db.Find(&users, "username=?", loginUser.Username)

	isNotFoundUser := 0 == len(users)
	if isNotFoundUser {
		c.JSON(200, Error{3, "Not found user"})
	} else {
		loginUser = users[0]
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
		connectInfo = u.User.String() + "@tcp(" + u.Host + ")" + u.Path
	}
	db, err := gorm.Open(scheme, connectInfo)
	if err != nil {
		fmt.Printf(connectInfo + "\n")
		fmt.Printf("can't connect db")
		panic("failed to connect database")
	}
	return db
}

// Model ---------------------------------------------

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
