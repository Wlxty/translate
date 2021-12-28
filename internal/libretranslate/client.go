package libretranslate

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	Logger *zap.SugaredLogger
	Host   string
}

type Libre interface {
	Translate(q string, source string, target string) (string, error)
	Languages() (string, error)
}

// Constructor for libretranslate client
func NewClient(Logger *zap.SugaredLogger, Host string) *Client {
	client := Client{Logger, Host}
	return &client
}

// PostForm to get translation
func (client *Client) Translate(q string, source string, target string) (string, error) {
	data := url.Values{
		"q":      {q},
		"source": {source},
		"target": {target},
	}
	resp, err := http.PostForm(client.Host+"translate", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	j := json.NewDecoder(resp.Body).Decode(&res)
	if j != nil {
		client.Logger.Debug("Service Languages: Not valid Json")
	}
	client.Logger.Debug("Service Languages works fine")

	jsonify, err := json.Marshal(res)

	return string(jsonify), err
}

// Get request, read all languages in libretranslate server
func (client *Client) Languages() (string, error) {
	resp, err := http.Get(client.Host + "languages")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	return sb, err
}
