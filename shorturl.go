package shorturl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type cleanuri_struct struct {
	Result_url string `json:"result_url"`
}

// Get short link from cleanuri.com (only 2 Urls in 1 second for one ip )
func cleanuri(url string) (string, error) {
	body := strings.NewReader("url=" + url)
	req, err := http.NewRequest("POST", "https://cleanuri.com/api/v1/shorten", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	a := cleanuri_struct{}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	json.Unmarshal(byt, &a)
	if a.Result_url == "" {
		return "", errors.New(string(byt))
	}
	return strings.Replace(a.Result_url, "\\", "", -1), nil
}
