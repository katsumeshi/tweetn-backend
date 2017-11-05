package service

import (
	"testing"
	"github.com/katsumeshi/tweetn-backend/dao/mocks"
	"github.com/katsumeshi/tweetn-backend/model"
	"github.com/stretchr/testify/assert"
)

func TestFindFirst(t *testing.T) {
	userDaoMock := new(mocks.User)

	expectedResult := model.User{0, "unko", "test", "akasaka", "programmer"}
	userDaoMock.On("FindFirst", "test").Return(expectedResult, nil).Once()
	userService := InitUserService(userDaoMock);

	user, err := userService.FindFirst("test");
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expectedResult.Id, user.Id)
	userDaoMock.AssertExpectations(t)
}