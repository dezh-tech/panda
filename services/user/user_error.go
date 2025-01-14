package service

import "errors"

var (
	Err         = errors.New("something went wrong")
	ErrNotFound = errors.New("user not found")
	ErrIsExist  = errors.New("user exist")
)
