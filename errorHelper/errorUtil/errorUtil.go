package errorUtil

import (
	"github.com/Kasparund/Go-Action-Test-Overload/errorHelper"
	"github.com/pkg/errors"
)

type errorUtil struct {
}

func NewErrorUtil() errorHelper.Helper {
	return &errorUtil{}
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func (e errorUtil) WithStack(err error) error {
	return errors.WithStack(err)
}

// New returns an error with the supplied message. New also records the stack trace at the point it was called.
func (e errorUtil) New(message string) error {
	return errors.New(message)
}
