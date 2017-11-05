// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/katsumeshi/tweetn-backend/model"

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// FindFirst provides a mock function with given fields: userName
func (_m *User) FindFirst(userName string) (model.User, error) {
	ret := _m.Called(userName)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(string) model.User); ok {
		r0 = rf(userName)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
