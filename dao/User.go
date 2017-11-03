package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/katsumeshi/tweetn-backend/model"
)

type User interface {
	FindFirst(userName string) (model.User, error)
}

type UserImpl struct {
	db *gorm.DB
}

func (u UserImpl) FindFirst(userName string) (model.User, error) {
	users := []model.User{}
	u.db.Limit(1).Find(&users, "username=?", userName)
	return users[0], nil
}

// can be more efficient
//func (u UserImpl) IsExist(string userName) (bool, error) {
//	isNotFoundUser := 0 == len(FindFirst(userName))
//	return isNotFoundUser, nil
//}
