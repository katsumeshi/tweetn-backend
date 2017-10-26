package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestPostLoginNotFound(t *testing.T) {

	res, err := http.PostForm(Server.URL+"/v1/login", url.Values{"username": {""}})
	defer res.Body.Close()
	var body Error
	json.NewDecoder(res.Body).Decode(&body)

	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("%v", res.StatusCode)
	}
	if body.Code != 3 {
		t.Fatalf("%v", body.Code)
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
