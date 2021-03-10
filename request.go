package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const URI = "https://jsonplaceholder.typicode.com/"

type Post struct {
	UserID 	int `json:"userId"`
	ID 		int `json:"id"`
	Title 	string `json:"title"`
	Body 	string `json:"body"`
}

type Posts []Post

func get(path string, param url.Values)([]byte, error) {
	uri, err := url.Parse(URI + path)
	if err != nil {
		return nil, err
	}

	if param != nil {
		uri.RawQuery = param.Encode()
	}

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func post(path string, body interface{}) ([]byte, error) {
	uri, err := url.Parse(URI + path)
	if err != nil {
		return nil, err
	}

	postBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri.String(), bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return results, nil
}