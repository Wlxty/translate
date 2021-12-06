package server

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "net/url"
)

func TestGetRequest(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/languages", nil)
    w := httptest.NewRecorder()
    var server Server
    server.LanguagePageHandler(w, req)
    res := w.Result()
    defer res.Body.Close()
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("expected error to be nil got %v", err)
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
    var server Server
    server.TranslatePageHandler(w, req)
    res := w.Result()
    defer res.Body.Close()
    _, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("expected error to be nil got %v", err)
    }
    if w.Code != http.StatusOK {
        t.Errorf("Home page didn't return %v", http.StatusOK)
    }
}
