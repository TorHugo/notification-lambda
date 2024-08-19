package http

import (
	"bytes"
	"fmt"
	"net/http"
)

type Client struct {
}

func NewHttpClient() *Client {
	return &Client{}
}

func (hc *Client) POST(url string, body string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, fmt.Errorf("error to create request! error: %w", err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error to send request! error: %w", err)
	}
	return resp, nil
}
