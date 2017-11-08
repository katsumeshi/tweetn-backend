package session

import "testing"

func TestSession(t *testing.T) {
	loginUser = user
	session := sessions.Default(c)
	v := session.Get("userId")
	var userId int
	if v == nil {
		userId = loginUser.Id
		session.Set("userId", userId)
		session.Save()
	} else {
		userId = v.(int)
	}
}