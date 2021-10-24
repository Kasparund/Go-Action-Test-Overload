package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	mockInterface "github.com/Kasparund/Go-Action-Test-Overload/httpClient/mocks"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func Test_service_StartProcess1(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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

func Test_service_StartProcess2(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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

func Test_service_StartProcess3(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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

func Test_service_StartProcess4(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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

func Test_service_StartProcess5(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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

func Test_service_StartProcess6(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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

func Test_service_StartProcess7(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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

func Test_service_StartProcess8(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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

func Test_service_StartProcess9(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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

func Test_service_StartProcess40(t *testing.T) {
	type fields struct {
		httpClient *mockInterface.MockHttpClient
	}
	type args struct {
		expectedString string
		responseBody   string
		statusCode     int
		httpError      error
		hasReadError   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
	for round < 1000 {
		round++
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				f := fields{
					httpClient: mockInterface.NewMockHttpClient(ctrl),
				}

				service := NewService(f.httpClient)
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
