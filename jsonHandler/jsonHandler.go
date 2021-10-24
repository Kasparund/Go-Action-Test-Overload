package jsonHandler

//go:generate mockgen -source ../jsonHandler/jsonHandler.go -destination=mocks/mock_jsonHandler.go -package=mocks

type JSONHandler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}
