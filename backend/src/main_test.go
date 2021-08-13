package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)



func Test_DbConnection(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/todolist", todoList())
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/todolist")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}


func Test_todoList(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/todolist", todoList())
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/todolist")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func Test_todoCreate(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/todo", todoCreate())
	ts := httptest.NewServer(r)
	defer ts.Close()
	reader := strings.NewReader("number=2")
	t.Run("create todo", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/todo", reader) //BTW check for error
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	})

}
