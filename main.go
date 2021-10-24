package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	httpclient "github.com/Kasparund/Go-Action-Test-Overload/httpClient"
	netclient "github.com/Kasparund/Go-Action-Test-Overload/httpClient/netHTTP"
)

func main() {
	httpClient := netclient.NewNetHttpClient()
	service := NewService(httpClient)

	response, err := service.StartProcess()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

type Service interface {
	StartProcess() (response string, err error)
}

type service struct {
	httpClient httpclient.HttpClient
}

func NewService(httpClient httpclient.HttpClient) Service {
	return &service{httpClient}
}

func (of *service) StartProcess() (response string, err error) {
	var requestBody = []byte(`{"key":"value"}`)
	resp, err := of.httpClient.Post("https://test.url.com", "application/json", requestBody)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		err = errors.New("unexpected status code from server")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body), nil
}
