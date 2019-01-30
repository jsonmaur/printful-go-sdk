package printful

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var apiKey string

func Register(key string) {
	apiKey = key
}

func Request(url string) ([]byte, error) {
	if apiKey == "" {
		return nil, errors.New("No Printful API key was registered")
	}

	client := &http.Client{
		Timeout: time.Second * 15,
	}

	endpoint := fmt.Sprintf("https://api.printful.com/%v", url)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	apiKeyEncoded := base64.StdEncoding.EncodeToString([]byte(apiKey))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", apiKeyEncoded))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var parsed Response
	if err = json.Unmarshal(body, &parsed); err != nil {
		return nil, errors.New("Could not parse response JSON")
	}

	if parsed.Code != 200 {
		return nil, errors.New("Received a non-200 response")
	}

	return body, nil
}
