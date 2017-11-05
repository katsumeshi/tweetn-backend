package service

import (
	"github.com/katsumeshi/tweetn-backend/model"
	"github.com/katsumeshi/tweetn-backend/dao"
)

type User interface {
	FindFirst(userName string) (model.User, error)
}

type UserImpl struct {
	userDao dao.User
}

func InitUserService(userDao dao.User) *UserImpl {
	return &UserImpl{userDao: userDao}
}

func (u UserImpl) FindFirst(userName string) (model.User, error) {
	user, _  := u.userDao.FindFirst(userName)
	return user, nil
}