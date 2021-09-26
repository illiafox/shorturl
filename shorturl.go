package shorturl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Get link "cleanuri.com/" (only 2 Url in 1 second for one ip): url - The long URL to shorten https://www.example.me
func Cleanuri(url string) (string, error) {
	req, err := http.NewRequest("POST", "https://cleanuri.com/api/v1/shorten", strings.NewReader("url="+url))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	a := struct {
		Result_url string `json:"result_url"`
	}{}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(byt, &a); err != nil {
		return "", err
	}
	if a.Result_url == "" {
		return "", errors.New(string(byt))
	}
	return strings.Replace(a.Result_url, "\\", "", -1), nil
}

// Short text or URL to 6-digit code (gotiny.cc/abc123): url - The long URL to shorten https://www.example.me OR text
func GoTiny(url string) (string, error) {
	req, err := http.NewRequest("POST", "https://gotiny.cc/api", strings.NewReader(fmt.Sprintf("{ \"input\" : \"%s\" }", url)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	a := struct {
		Long string `json:"long"`
		Code string `json:"code"`
	}{}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(byt[1:len(byt)-1], &a)
	if a.Code == "" || a.Long == "" {
		return "", errors.New(string(byt))
	}
	return "gotiny.cc/" + a.Code, nil
}

//  Get short link "1pt.co/": url - The long URL to shorten https://www.example.me , short(optional) - The part after 1pt.co/ that will redirect to your long URL. If this paramter is not provided or the requested short URL is already taken, it will return a random 5-letter string
func Pt1(url string, short string) (string, error) {
	body, _ := json.Marshal(struct {
		Long  string `json:"long"`
		Short string `json:"short"`
	}{url, short})
	req, err := http.NewRequest("GET", "https://api.1pt.co/addURL", strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	a := struct {
		Status  int    `json:"status"`
		Message string `json:"code"`
		Short   string `json:"short"`
		Long    string `json:"long"`
	}{}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(byt[1:len(byt)-1], &a)
	if a.Short == "" || a.Message != "Added!" || a.Status != 201 || a.Long == "" {
		return "", errors.New(string(byt))
	}
	return "1pt.co/" + a.Short, nil
}
