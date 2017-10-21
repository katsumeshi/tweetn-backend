package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := gin.Default()
	//	store := sessions.NewCookieStore([]byte("secret"))

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("session", store))

	r.LoadHTMLGlob("templates/*")

	r.GET("/user/:username", ShowUser)

	r.GET("/login", GetLoginView)
	r.POST("/login", Login)
	r.POST("/logout", Logout)

	// r.GET("/tweets/new", GetTweetView)
	r.POST("/tweets/new", PostTweet)

	r.GET("/account", GetAccountView)
	r.POST("/account", CreateAccount)

	r.GET("/tweets/lists", TweetsList)
	r.Run(":8080")
}

func ShowUser(c *gin.Context) {
	username := c.Param("username")

	user := User{}
	db := gormConnect()
	if db.First(&user, "username = ?", username).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"message": "can't find the user"})
	} else {
		c.HTML(http.StatusOK, "show.tmpl", user)
	}
}

func GetAccountView(c *gin.Context) {
	c.HTML(http.StatusOK, "account.tmpl", Error{})
}

func CreateAccount(c *gin.Context) {
	user := User{}
	c.Bind(&user)
	if 0 < len(user.Username) {
		db := gormConnect()
		if canCreateUser(user.Username) {
			db.Create(&user)
			c.Redirect(http.StatusMovedPermanently, "/user/"+user.Username)
		} else {
			c.HTML(http.StatusOK, "account.tmpl", Error{0, "Your account is already exisited"})
		}
	} else {
		c.HTML(http.StatusOK, "account.tmpl", Error{1, "At least you need to fill in username"})
	}
}

func canCreateUser(username string) bool {
	db := gormConnect()
	user := User{}
	if db.First(&user, "username = ?", username).RecordNotFound() {
		return true
	}
	return false
}

func GetLoginView(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func Login(c *gin.Context) {
	var loginUser User
	c.Bind(&loginUser)

	users := []User{}
	db := gormConnect()
	db.Find(&users, "username=?", loginUser.Username)

	isNotFoundUser := 0 == len(users)
	if isNotFoundUser {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"error": true})
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
		c.Redirect(http.StatusMovedPermanently, "/tweets/lists")
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userId")
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/login")
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

func PostTweet(c *gin.Context) {
	var tweet Tweet
	c.Bind(&tweet)

	userId := getSessionUserId(c)
	tweet.UserId = userId
	db := gormConnect()
	db.Create(&tweet)

	// c.JSON(http.StatusOK, gin.H{"message": tweet.Content, "user": tweet.UserId})
	c.Redirect(http.StatusMovedPermanently, "/tweets/lists")
}

func TweetsList(c *gin.Context) {

	session := sessions.Default(c)
	userId := session.Get("userId")
	if userId == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	tweets := []Tweet{}
	db := gormConnect()
	db.Preload("User").Order("Id desc").Find(&tweets)

	c.HTML(http.StatusOK, "lists.tmpl", gin.H{"tweets": tweets})
}

// DB ---------------------------------------------
func gormConnect() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/development")
	if err != nil {
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

type User struct {
	Id       int    `json:"id"`
	Name     string `form:"name" json:"name"`
	Username string `form:"username" json:"username"`
	Location string `form:"location" json:"location"`
	About    string `form:"about" json:"about"`
}

type Tweet struct {
	Id      int
	Content string `form:"content" json:"content"`
	UserId  int    `json:"user_id"`
	User    User
}
