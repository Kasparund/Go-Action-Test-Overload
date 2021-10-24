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

func Test_service_StartProcess(t *testing.T) {
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

type ErrorBuffer struct {
}

func (mb ErrorBuffer) Close() error {
	return nil
}

func (mb ErrorBuffer) Read(p []byte) (n int, err error) {
	return 0, errors.New("error while reading")
}
