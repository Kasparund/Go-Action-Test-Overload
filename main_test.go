package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/Kasparund/Go-Action-Test-Overload/errorHelper"
	"github.com/Kasparund/Go-Action-Test-Overload/errorHelper/errorUtil"
	mockInterface "github.com/Kasparund/Go-Action-Test-Overload/httpClient/mocks"
	jsonHandlerMock "github.com/Kasparund/Go-Action-Test-Overload/jsonHandler/mocks"
	"github.com/Kasparund/Go-Action-Test-Overload/util"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

var numberOfSubtests = 100

type fields struct {
	httpClient  *mockInterface.MockHttpClient
	errorUtil   errorHelper.Helper
	config      util.InfrastructureConfig
	jsonHandler *jsonHandlerMock.MockJSONHandler
}

func setupSubtest(t *testing.T) (fields, Service) {
	ctrl := gomock.NewController(t)

	infrastructureConfig, _ := util.LoadInfrastructureConfig()
	f := fields{
		httpClient:  mockInterface.NewMockHttpClient(ctrl),
		errorUtil:   errorUtil.NewErrorUtil(),
		config:      infrastructureConfig,
		jsonHandler: jsonHandlerMock.NewMockJSONHandler(ctrl),
	}

	service := NewService(f.httpClient, f.errorUtil, f.config, f.jsonHandler)
	return f, service
}

func Test_service_StartProcess01(t *testing.T) {
	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)

				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess02(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess03(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess04(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess05(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess06(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess07(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess08(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess09(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess10(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess11(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess12(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess13(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess14(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess15(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess16(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess17(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess18(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess19(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess21(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess22(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess23(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess24(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess25(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess26(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess27(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess28(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess29(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess30(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess31(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess32(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess33(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess34(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess35(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess36(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess37(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess38(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess39(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess41(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess42(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess43(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess44(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess45(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess46(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess47(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess48(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess49(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess50(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess51(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess52(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess53(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess54(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess55(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess56(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess57(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess58(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess59(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess60(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess61(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess62(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess63(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess64(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess65(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess66(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess67(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess68(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess69(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess70(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess71(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess72(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess73(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess74(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess75(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess76(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess77(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess78(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

func Test_service_StartProcess79(t *testing.T) {

	type args struct {
		expectedString  string
		responseBody    string
		statusCode      int
		httpError       error
		hasReadError    bool
		hasMarshalError bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed--Marshal-Error",
			args: args{
				hasMarshalError: true,
			},
			wantErr: true,
		},
		{
			name: "Failed--HTTP-Error",
			args: args{
				statusCode:   201,
				responseBody: `{}`,
				httpError:    errors.New("server error"),
			},
			wantErr: true,
		},
		{
			name: "Failed--Server-Error",
			args: args{
				statusCode: 500,
			},
			wantErr: true,
		},
		{
			name: "Failed--Read-Error",
			args: args{
				statusCode:   201,
				responseBody: `{"key": "value"}`,
				hasReadError: true,
			},
			wantErr: true,
		},

		{
			name: "Successful",
			args: args{
				expectedString: `{"key":"value"}`,
				statusCode:     201,
				responseBody:   `{"key":"value"}`,
			},
		},
	}

	round := 0
	for round < numberOfSubtests {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f, service := setupSubtest(t)
				requestBody := []byte(`{"key":"value"}`)
				responseBody := ioutil.NopCloser(strings.NewReader(tt.args.responseBody))
				httpResponse := http.Response{
					StatusCode: tt.args.statusCode,
					Body:       responseBody,
				}
				if tt.args.hasReadError {
					httpResponse.Body = ErrorBuffer{}
				}
				url := "https://test.url.com"

				gomock.InOrder(
					f.jsonHandler.
						EXPECT().
						Marshal(Request{Key: "value"}).
						DoAndReturn(marshalMock(tt.args.hasMarshalError)).
						Times(1),
					f.httpClient.
						EXPECT().
						Post(url, "application/json", requestBody).
						Return(&httpResponse, tt.args.httpError).
						MaxTimes(1),
				)

				response, err := service.StartProcess()

				if tt.wantErr == false && !reflect.DeepEqual(tt.args.expectedString, response) {
					t.Errorf("expected %v, actual %v", tt.args.expectedString, response)
				}

				assert.Equal(t, tt.wantErr, (err != nil))
			})
		}
	}
}

type ErrorBuffer struct {
}

func (mb ErrorBuffer) Close() error {
	return nil
}

func (mb ErrorBuffer) Read(p []byte) (n int, err error) {
	return 0, errors.New("error while reading")
}

func marshalMock(hasMarshalError bool) func(request interface{}) ([]byte, error) {
	return func(request interface{}) ([]byte, error) {
		if !hasMarshalError {
			return json.Marshal(request)
		}
		return nil, errors.New("marshall error")
	}
}
