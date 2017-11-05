package dao

//import
//(
//	"fmt"
//	"testing"
//
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//	"github.com/stretchr/testify/assert"
//)

//var (
//	userDao User
//)
//
//func TestMain(t *testing.T) {
//	var connectInfo string = "root:@tcp(127.0.0.1:3306)/development"
//	db, err := gorm.Open("mysql", connectInfo)
//	userDao = InitUserDao(db)
//	//	d, err :=
//	if err != nil {
//		fmt.Printf(connectInfo + "\n")
//		fmt.Printf("can't connect db")
//		panic("failed to connect database")
//	}
//	userDao = InitUserDao(db)
//}
//
////
//func TestUserFindList(t *testing.T) {
//	user, _ := userDao.FindFirst("test")
//	assert.Equal(t, "unko", user.Name, "user name should be unko")
//}
