package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Kasparund/Go-Action-Test-Overload/errorHelper"
	"github.com/Kasparund/Go-Action-Test-Overload/errorHelper/errorUtil"
	httpclient "github.com/Kasparund/Go-Action-Test-Overload/httpClient"
	netclient "github.com/Kasparund/Go-Action-Test-Overload/httpClient/netHTTP"
	jsonHandler "github.com/Kasparund/Go-Action-Test-Overload/jsonHandler"
	json "github.com/Kasparund/Go-Action-Test-Overload/jsonHandler/json"
	"github.com/Kasparund/Go-Action-Test-Overload/util"
)

func main() {
	httpClient := netclient.NewNetHttpClient()
	errorHandler := errorUtil.NewErrorUtil()
	jsonHandler := json.NewJSONHandler()
	config, _ := util.LoadInfrastructureConfig()
	service := NewService(httpClient, errorHandler, config, jsonHandler)

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
	errorUtil   errorHelper.Helper
	jsonHandler jsonHandler.JSONHandler
}

func NewService(httpClient httpclient.HttpClient, errorUtil errorHelper.Helper, config util.InfrastructureConfig, jsonHandler jsonHandler.JSONHandler) Service {
	fmt.Println(config.ConfigName)
	return &service{httpClient, errorUtil, jsonHandler}
}

func (of *service) StartProcess() (response string, err error) {
	request := Request{Key: "value"}
	requestBody, err := of.jsonHandler.Marshal(request)
	if err != nil {
		return
	}

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

type Request struct {
	Key string `json:"key"`
}
