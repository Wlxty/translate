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
	GetHost() string
	GetLogger() *zap.SugaredLogger
	GetLibre() *Client
}

func (c *Client) GetLibre() *Client {
	return c
}
func (c *Client) GetHost() string {
	return c.Host
}

func (c *Client) GetLogger() *zap.SugaredLogger {
	return c.Logger
}

// Constructor for libretranslate client
func NewClient(Logger *zap.SugaredLogger, Host string) *Client {
	client := Client{Logger, Host}
	return &client
}

// PostForm to get translation
func (client *Client) Translate(q string, source string, target string) (string, error) {
	input := url.Values{
		"q":      {q},
		"source": {source},
		"target": {target},
	}
	data, err := http.PostForm(client.GetHost()+"translate", input)

	if err != nil {
		log.Fatal(err)
	}

	var response map[string]interface{}

	j := json.NewDecoder(data.Body).Decode(&response)
	if j != nil {
		client.Logger.Debug("Service Languages: Not valid Json")
	}
	client.Logger.Debug("Service Languages works fine")

	jsonify, err := json.Marshal(response)

	return string(jsonify), err
}

// Get request, read all languages in libretranslate server
func (client *Client) Languages() (string, error) {
	data, err := http.Get(client.GetHost() + "languages")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	response := string(body)
	return response, err
}
