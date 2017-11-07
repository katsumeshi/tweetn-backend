package dao

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/katsumeshi/tweetn-backend/model"
	"time"
)

var (
	userDao User
)

func TestInitUserDao(t *testing.T) {
	var connectInfo string = "root:@tcp(127.0.0.1:3306)/development"
	db, err := gorm.Open("mysql", connectInfo)
	userDao = InitUserDao(db)
	if err != nil {
		panic("failed to connect database")
	}
	userDao = InitUserDao(db)

	if (!db.HasTable("users")) {
		db.AutoMigrate(&model.User{})
	}

	user := model.User{0, "unko", "test", "chiba", "ware"}
	db.Create(&user)

}

func TestUserFindList(t *testing.T) {
	user, _ := userDao.FindFirst("test")
	assert.Equal(t, "unko", user.Name, "user name should be unko")
}
