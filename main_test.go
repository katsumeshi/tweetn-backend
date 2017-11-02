package main

//import (
//	"testing"
//
//	"github.com/stretchr/testify/mock"
//)

//var Server *httptest.Server
//
//func TestInit(t *testing.T) {
//	Server = httptest.NewServer(GetMainEngine())
//	fmt.Printf("%v", Server.URL)
//}
//
//func TestGetLogin(t *testing.T) {
//
//	res, err := http.Get(Server.URL + "/v1/login")
//	if err != nil {
//		log.Fatal(err)
//	}
//	if res.StatusCode != 200 {
//		t.Fatalf("%v", res.StatusCode)
//	}
//}

/*
  Test objects
*/

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
//type MyMockedObject struct {
//	mock.Mock
//}
//
//// DoSomething is a method on MyMockedObject that implements some interface
//// and just records the activity, and returns what the Mock object tells it to.
////
//// In the real object, this method would do something useful, but since this
//// is a mocked object - we're just going to stub it out.
////
//// NOTE: This method is not being tested here, code that uses this object is.
//func (m *MyMockedObject) DoSomething(number int) (bool, error) {
//
//	args := m.Called(number)
//	return args.Bool(0), args.Error(1)
//
//}

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
//func TestSomething(t *testing.T) {
//
//	// create an instance of our test object
//	testObj := new(MyMockedObject)
//
//	// setup expectations
//	testObj.On("DoSomething", 123).Return(true, nil)
//
//	// call the code we are testing
//	targetFuncThatDoesSomethingWithObj(testObj)
//
//	// assert that the expectations were met
//	testObj.AssertExpectations(t)
//
//}

//func TestPostLoginNotFound(t *testing.T) {
//
//	res, err := http.PostForm(Server.URL+"/v1/login", url.Values{"username": {""}})
//	defer res.Body.Close()
//	var body Error
//	json.NewDecoder(res.Body).Decode(&body)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	if res.StatusCode != 200 {
//		t.Fatalf("%v", res.StatusCode)
//	}
//	if body.Code != 3 {
//		t.Fatalf("%v", body.Code)
//	}
//}

//func TestGetLogin2(t *testing.T) {
//
//	ts := httptest.NewServer(GetMainEngine())
//	defer ts.Close()
//
//	res, err := http.Get(ts.URL + "/v1/login")
//	if err != nil {
//		log.Fatal(err)
//	}
//	if res.StatusCode != 200 {
//		t.Fatalf("%v", res.StatusCode)
//	}
//}

//func TestLogin(t *testing.T) {
//	req, _ := http.NewRequest("POST", "/login", nil)
//	w := httptest.NewRecorder()
//
//	r := gin.Default()
//
//	r.ServeHTTP(w, req)
//
//	if !strings.Contains(w.HeaderMap.Get("Content-Type"), "text/html") {
//		t.Errorf("Content-Type should be text/html, was %s", w.HeaderMap.Get("Content-Type"))
//	}
//}
