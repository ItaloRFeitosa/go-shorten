package errs

import (
	"errors"
	"fmt"
)

type Error struct {
	Type     string `json:"type"`
	Code     string `json:"code"`
	Reason   string `json:"reason"`
	template string
	err      error
}

func (e Error) Error() string {
	if e.err == nil {
		return e.Reason
	}

	return e.err.Error()
}

func (e Error) Unwrap() error {
	return errors.Unwrap(e.err)
}

func (e Error) Is(target error) bool {
	if target == nil {
		return false
	}

	if err, ok := target.(Error); ok {
		return e.Equals(err)
	}

	return errors.Is(e.err, target)
}

func (e Error) Equals(target Error) bool {
	return target.Code == e.Code ||
		target.Reason == e.Reason ||
		target.template == e.template
}

func (e Error) WithArgs(opts ...any) Error {
	e.err = fmt.Errorf(e.template, opts...)
	e.Reason = e.err.Error()

	return e
}
