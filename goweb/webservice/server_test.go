package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", HandleRequest)
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	r := httptest.NewRequest("GET", "/post/?id=2", nil)
	mux.ServeHTTP(writer, r)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 2 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	r := httptest.NewRequest("PUT", "/post/2", json)
	mux.ServeHTTP(writer, r)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
