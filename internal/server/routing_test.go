package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"io/ioutil"
	"fmt"
	"github.com/gorilla/mux"
)

func TestingHomePage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := mux.NewRouter()
		server := NewTaskServer()

		r.HandleFunc("/", server.HomePageHandler).Methods("GET")
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(ts.URL + "/")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})
}


func TestingTranslatePage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := mux.NewRouter()
		server := NewTaskServer()

		r.HandleFunc("/translate", server.TranslatePageHandler).Methods("POST")
		ts := httptest.NewServer(r)
		defer ts.Close()
		variables := strings.NewReader(`{
		"word" : "drzwi",
		"source" : "polish",
		"target" : "english"
	}`)
		res, err := http.NewRequest("POST", "/translate", variables) //BTW check for error
		res.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
			return
		}else{
			fmt.Println(string(body))
		}
	})
}