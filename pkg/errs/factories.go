package errs

import "fmt"

type OptionFunc func(*Error)

func New(opts ...OptionFunc) Error {
	var err Error
	for _, applyOpt := range opts {
		applyOpt(&err)
	}

	return err
}

func Reason(msg string) OptionFunc {
	return func(err *Error) {
		err.err = fmt.Errorf(msg)
		err.Reason = err.Error()
	}
}

func Template(messageTemplate string) OptionFunc {
	return func(err *Error) {
		err.template = messageTemplate
		err.Reason = messageTemplate
		err.err = fmt.Errorf(messageTemplate)
	}
}

func Code(code string) OptionFunc {
	return func(err *Error) {
		err.Code = code
	}
}

func withType(t string) OptionFunc {
	return func(err *Error) {
		err.Type = t
	}
}

var (
	Internal     = withType(InternalType)
	Conflict     = withType(ConflictType)
	NotFound     = withType(NotFoundType)
	BusinessRule = withType(BusinessRuleType)
	Validation   = withType(ValidationType)
)
