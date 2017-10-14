package main

import (
	"fmt"
	"net/http"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/users/show/:username", ShowUser)

	router.GET("/login", GetLoginView)
	router.POST("/login", Login)

	router.GET("/tweets/new", GetTweetView)
	router.POST("/tweets/new", PostTweet)

	router.GET("/tweets/lists", TweetsList)

	router.Run(":8080")
}

func ShowUser(c *gin.Context) {
	username := c.Param("username")
	users := []User{}

	db := gormConnect()
	db.Find(&users, "username=?", username)
	if 0 < len(users) {
		json := convertToJson(users[0])
		m := convertToMap(json)
		c.HTML(http.StatusOK, "show.tmpl", m)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "can't find the user"})
	}

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

	if 0 < len(users) {
		loginUser = users[0]
	}
}

func GetTweetView(c *gin.Context) {
	c.HTML(http.StatusOK, "new.tmpl", nil)
}

func PostTweet(c *gin.Context) {
	var tweet Tweet
	c.Bind(&tweet)

	tweet.UserId = 1
	db := gormConnect()
	db.Create(&tweet)

	c.JSON(http.StatusOK, gin.H{"message": tweet.Content, "user": tweet.UserId})
}

func TweetsList(c *gin.Context) {
	tweets := []Tweet{}
	db := gormConnect()
	db.Preload("User").Find(&tweets)
	c.HTML(http.StatusOK, "lists.tmpl", gin.H{"tweets": tweets})
}

// Util ---------------------------------------------

func convertToMap(jsonString string) gin.H {
	m := make(gin.H)
	err := json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return m
}

func convertToJson(user User) string {
	j, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(j)
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

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `form:"username" json:"username"`
	Location string `json:"location"`
	About    string `json:"about"`
}

type Tweet struct {
	Id      int
	Content string `form:"content" json:"content"`
	UserId  int    `json:"user_id"`
	User    User
}
