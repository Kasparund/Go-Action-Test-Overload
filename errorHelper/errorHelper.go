package errorHelper

//go:generate mockgen -source ../errorHelper/errorHelper.go -destination=mocks/mock_errorHelper.go -package=mocks

type Helper interface {
	WithStack(err error) error
	New(message string) error
}
