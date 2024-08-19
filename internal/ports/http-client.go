package ports

import "net/http"

type HttpClient interface {
	POST(url string, body string, headers map[string]string) (*http.Response, error)
}
