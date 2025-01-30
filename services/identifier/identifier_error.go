package service

import "errors"

var (
	Err         = errors.New("something went wrong")
	ErrNotFound = errors.New("identifier not found")
	ErrIsExist  = errors.New("identifier exist")
)
