package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetLogin(t *testing.T) {
	req, _ := http.NewRequest("GET", "/login", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	type TestData struct{ Name string }

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", nil)
	})

	r.ServeHTTP(w, req)

	if !strings.Contains(w.HeaderMap.Get("Content-Type"), "text/html") {
		t.Errorf("Content-Type should be text/html, was %s", w.HeaderMap.Get("Content-Type"))
	}
}
