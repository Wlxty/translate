package libretranslate

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

func TestFetchLanguages(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// our database of articles
	languages := make([]map[string]interface{}, 0)

	// mock to list out the articles
	httpmock.RegisterResponder("GET", "https://libretranslate:5000/languages",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, languages)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}

func TestTranslateWord(t *testing.T) {
	httpmock.RegisterResponder("POST", "https://libretranslate:5000/translate?source=pl&target=en&q=szpital",
		func(req *http.Request) (*http.Response, error) {
			translates := make([]map[string]interface{}, 0)
			translate := make(map[string]interface{})
			if err := json.NewDecoder(req.Body).Decode(&translate); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}

			translates = append(translates, translate)

			resp, err := httpmock.NewJsonResponse(200, translate)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}
