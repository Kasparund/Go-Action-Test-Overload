package httpclient

import (
	"net/http"
)

//go:generate mockgen -source ../httpClient/httpClient.go -destination=mocks/mock_httpClient.go -package=mocks

type HttpClient interface {
	Get(url string) (*http.Response, error)
	Post(url string, contentType string, body []byte) (*http.Response, error)
	Shutdown()
}
