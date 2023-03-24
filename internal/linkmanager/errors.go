package linkmanager

import "github.com/italorfeitosa/go-shorten/pkg/errs"

var ErrWrongOwnerID = errs.New(
	errs.BusinessRule,
	errs.Code("WrongOwnerID"),
	errs.Reason("resource doens't belong to given owner"),
)

var ErrMissingOwnerID = errs.New(
	errs.Validation,
	errs.Code("MissingOwnerIDHeader"),
	errs.Reason("missing x-owner-id in header"),
)
