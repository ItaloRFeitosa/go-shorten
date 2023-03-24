package errs

import (
	"net/http"
)

var statusCodeMap = map[string]int{
	InternalType:     http.StatusInternalServerError,
	ValidationType:   http.StatusBadRequest,
	NotFoundType:     http.StatusNotFound,
	BusinessRuleType: http.StatusUnprocessableEntity,
	ConflictType:     http.StatusConflict,
}

func StatusCode(err error) int {

	asError, ok := err.(Error)
	if !ok {
		return http.StatusInternalServerError
	}

	code, ok := statusCodeMap[asError.Type]
	if !ok {
		return http.StatusInternalServerError
	}

	return code
}
