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
		users := []user{}

		db.Find(&users, "username=?", username)
		if 0 < len(users) {
			json := convertToJson(users[0])
			m := convertToMap(json)
			c.HTML(http.StatusOK, "show.tmpl", m)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "can't find the user"})
		}

	})
	router.Run(":8080")
}

func convertToMap(jsonString string) gin.H {
	m := make(gin.H)
	err := json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return m
}

func convertToJson(user user) string {
	j, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(j)
}

type user struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username"`
	Location string `json:"location"`
	About    string `json:"about"`
}

func gormConnect() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/development")
	if err != nil {
		fmt.Printf("can't connect db")
		panic("failed to connect database")
	}
	return db
}
