package errors

import "errors"

var (
	ErrNoUser               = errors.New("no user in DB")
	ErrInsertUser           = errors.New("can't insert new user")
	ErrReadUser             = errors.New("can't read user")
	ErrDecode               = errors.New("failed to decode structure")
	ErrMethodNotImplemented = errors.New("method is not implemented")
)
