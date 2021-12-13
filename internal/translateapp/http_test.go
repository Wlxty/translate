package translateapp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"translateapp/internal/server"
)

func TestGetRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/languages", nil)
	w := httptest.NewRecorder()
	var server server.Server

	service := Service{server: &server}
	service.LanguagePageHandler(w, req)
	res := w.Result()

	defer res.Body.Close()
	
	data, _ := ioutil.ReadAll(res.Body)
	if err := json.NewEncoder(w).Encode(server.Languages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	assert.NotEqual(t, 0, len(data))
}

func TestPostRequest(t *testing.T) {
	parm := url.Values{}
	parm.Add("source", "polish")
	parm.Add("target", "english")
	parm.Add("word", "szpital")

	req := httptest.NewRequest(http.MethodPost, "/translate?source="+parm["source"][0]+"&target="+parm["target"][0]+"&word="+parm["word"][0], nil)
	w := httptest.NewRecorder()

	var server server.Server
	service := Service{server: &server}
	service.TranslatePageHandler(w, req)
	res := w.Result()

	defer res.Body.Close()
	if err := json.NewEncoder(w).Encode(server.Languages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
