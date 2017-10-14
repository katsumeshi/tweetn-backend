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
	db := gormConnect()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/users/show/:username", func(c *gin.Context) {
		username := c.Param("username")
		users := []User{}

		db.Find(&users, "username=?", username)
		if 0 < len(users) {
			json := convertToJson(users[0])
			m := convertToMap(json)
			c.HTML(http.StatusOK, "show.tmpl", m)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "can't find the user"})
		}

	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", nil)
	})

	var loginUser User
	router.POST("/login", func(c *gin.Context) {
		c.Bind(&loginUser)

		users := []User{}
		db.Find(&users, "username=?", loginUser.Username)

		if 0 < len(users) {
			loginUser = users[0]
		}
		//		tweet.UserId = users[0].Id
		//		db.Create(&tweet)
		//
		//		c.JSON(http.StatusOK, gin.H{"message": tweet.Content, "user": tweet.UserId})
	})

	router.GET("/tweets/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new.tmpl", nil)
	})

	router.POST("/tweets/new", func(c *gin.Context) {
		var tweet Tweet
		c.Bind(&tweet)

		tweet.UserId = loginUser.Id
		db.Create(&tweet)

		c.JSON(http.StatusOK, gin.H{"message": tweet.Content, "user": tweet.UserId})
	})

	visitCount := 0

	router.GET("/tweets/lists", func(c *gin.Context) {
		tweets := []Tweet{}
		db.Preload("User").Find(&tweets)
		c.HTML(http.StatusOK, "lists.tmpl", gin.H{"tweets": tweets})
		fmt.Println("%d", visitCount)
		visitCount++
	})

	router.Run(":8080")
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
