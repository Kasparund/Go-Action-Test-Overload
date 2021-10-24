package nethttp

import (
	"bytes"
	"net/http"

	httpclient "github.com/Kasparund/Go-Action-Test-Overload/httpClient"
)

type netHttpClient struct {
	Client *http.Client
}

func NewNetHttpClient() httpclient.HttpClient {
	return &netHttpClient{
		&http.Client{},
	}
}

func (c *netHttpClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func (c *netHttpClient) Post(url string, contentType string, body []byte) (*http.Response, error) {
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	return c.Client.Do(req)
}

func (c *netHttpClient) Shutdown() {
	// Closing Idle Connections
	c.Client.CloseIdleConnections()
}
