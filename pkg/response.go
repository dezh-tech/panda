package pkg

import "github.com/dezh-tech/panda/pkg/validator"

type ResponseDto struct {
	Success bool             `json:"success"`
	Error   validator.Varror `json:"error"`
	Message string           `json:"message"`
	Data    interface{}      `json:"data"`
}
