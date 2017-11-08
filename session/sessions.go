package session

import (
	"github.com/gin-contrib/sessions"
)

type Session struct {
	session sessions.Session
}

func InitSession(session sessions.Session) Session {
	return Session{session: session}
}

func (s Session) SaveUserId(userId int) int {
	v := s.GetUserId()
	if v == -1 {
		s.session.Set("userId", userId)
		s.session.Save()
	}
	return v
}

func (s Session) GetUserId() int {
	v := s.session.Get("userId")
	if v != nil {
		return v.(int)
	}
	return -1
}


