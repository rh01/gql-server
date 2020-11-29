package store

import "errors"

var (
	errBadCredentials  = errors.New("Email/password combination don't work")
	errUnauthenticated = errors.New("Unauthenticated")
	errUnknown         = errors.New("Something went wrong")
	errRecordNotFound  = "record not found"
	errCreateRecord = errors.New("record creation failed")
	errNotImplemented = errors.New("not implemented")
)
