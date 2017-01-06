package utils

import (
	"github.com/YanshuoH/youkonger/consts"
	"github.com/pkg/errors"
)

type CommonError struct {
	Code        string
	Description string
	Detail      string
	Err         error
}

// A helper function aimed to:
// Return a wrapped error with code, description and more details
// It uses pkg/errors for error wrapping
func NewCommonError(code string, err error, optionals ...string) *CommonError {
	e := &CommonError{
		Code: code,
		Description: consts.Messenger.Get(code),
	}

	// override description if necessary
	if len(optionals) > 0 {
		e.Description = optionals[0]
	}

	if len(optionals) > 1 {
		e.Detail = optionals[1]
	}

	if err == nil {
		e.Err = errors.New(e.Description)
	} else {
		e.Err = errors.Wrap(err, e.Description)
	}

	return e
}
