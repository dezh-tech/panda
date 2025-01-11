package domainService

import "errors"

var (
	Err  = errors.New("something went wrong")
	ErrNotFound   = errors.New("domain not found")
	ErrIsExist   = errors.New("domain exist")
)