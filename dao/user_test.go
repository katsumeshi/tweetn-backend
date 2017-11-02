package dao

import (
	"github.com/jinzhu/gorm"
)

var (
	dbMap   *gorm.DB
	userDao User
)

//
//func TestUserFindList(t *testing.T) {
//	// userDao.FindFist
//}
