package server

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/languages", nil)
	w := httptest.NewRecorder()
	var server Server

	server.LanguagePageHandler(w, req)
	res := w.Result()

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	assert.NotEqual(t, 0, len(data))
}

func TestPostRequest(t *testing.T) {
	parm := url.Values{}
	parm.Add("source", "polish")
	parm.Add("target", "english")
	parm.Add("word", "szpital")

	req := httptest.NewRequest(http.MethodPost, "/translate?source="+parm["source"][0]+"&target="+parm["target"][0]+"&word="+parm["word"][0], nil)
	w := httptest.NewRecorder()

	var server Server
	server.TranslatePageHandler(w, req)
	res := w.Result()

	defer res.Body.Close()
}
