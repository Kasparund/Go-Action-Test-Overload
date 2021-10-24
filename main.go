package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	httpclient "github.com/Kasparund/Go-Action-Test-Overload/httpClient"
	netclient "github.com/Kasparund/Go-Action-Test-Overload/httpClient/netHTTP"
	jsonHandler "github.com/Kasparund/Go-Action-Test-Overload/jsonHandler"
	json "github.com/Kasparund/Go-Action-Test-Overload/jsonHandler/json"
	"github.com/Kasparund/Go-Action-Test-Overload/util"
)

func main() {
	httpClient := netclient.NewNetHttpClient()
	jsonHandler := json.NewJSONHandler()
	config, _ := util.LoadInfrastructureConfig()
	service := NewService(httpClient, config, jsonHandler)

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
	httpClient  httpclient.HttpClient
	jsonHandler jsonHandler.JSONHandler
}

func NewService(httpClient httpclient.HttpClient, config util.InfrastructureConfig, jsonHandler jsonHandler.JSONHandler) Service {
	fmt.Println(config.ConfigName)
	return &service{httpClient, jsonHandler}
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
