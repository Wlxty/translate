package server

import (
        "io/ioutil"
        "net/http"
        "testing"
        "time"
        "github.com/stretchr/testify/assert"
	"strings"
	"net/url"
)

func TestLanuchHandleRequests(t *testing.T){
	go HandleRequests(":8080")
}

func TestPostRequest(t *testing.T) {
	client := &http.Client{
                Timeout: 1 * time.Second,
        }

        parm := url.Values{}
        parm.Add("source", "polish")
        parm.Add("target", "english")
        parm.Add("word", "szpital")
        r, _ := http.NewRequest("POST", "http://localhost:8080/translate",strings.NewReader(parm.Encode()))

        resp, err := client.Do(r)
        if err != nil {
                panic(err)
        }
        assert.Equal(t, http.StatusOK, resp.StatusCode)
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                panic(err)
        }
        assert.NotEqual(t, 0, body)
}

func TestGetRequest(t *testing.T) {
        client := &http.Client{
                Timeout: 1 * time.Second,
        }

        r, _ := http.NewRequest("GET", "http://localhost:8080/languages", nil)

        resp, err := client.Do(r)
        if err != nil {
                panic(err)
        }
        assert.Equal(t, http.StatusOK, resp.StatusCode)
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                panic(err)
        }
        assert.NotEqual(t, 0, len(body))
}
