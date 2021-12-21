package libretranslate

import "go.uber.org/zap"

type Client struct {
	host   string
	logger *zap.SugaredLogger
}

func NewClient(host string, logger *zap.SugaredLogger) *Client {
	return &Client{
		host:   host,
		logger: logger,
	}
}

func (c *Client) Languages() ([]Language, error) {
	// var languages translateapp.Language
	// output := languages.Languages()
	// jsonify, err := json.Marshal(output)
	// if err != nil {
	// 	fmt.Fprintf(writer, "Error: %s", err.Error())
	// 	return nil
	// }

	// fmt.Fprintf(writer, string(jsonify))
	return nil, nil
}

func (c *Client) Translate(text, fromLang, toLang string) (string, error) {
	// resp, err := http.PostForm("http://libretranslate:5000//translate", data)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var res map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&res)
	// jsonify, err := json.Marshal(res)
	// if err != nil {
	// 	fmt.Fprintf(writer, "Error: %s", err.Error())
	// 	fmt.Fprintf(writer, "Error: %w", err)
	// 	return nil
	// }

	// fmt.Fprintf(writer, string(jsonify))
	return "", nil
}
