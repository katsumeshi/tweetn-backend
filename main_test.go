package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var Server *httptest.Server

func TestInit(t *testing.T) {
	Server = httptest.NewServer(GetMainEngine())
	fmt.Printf("%v", Server.URL)
}

func TestGetLogin(t *testing.T) {

	res, err := http.Get(Server.URL + "/v1/login")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("%v", res.StatusCode)
	}
}

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
