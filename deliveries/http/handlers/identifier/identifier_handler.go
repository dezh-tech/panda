package handlers

import service "github.com/dezh-tech/panda/services/identifier"

type Identifier struct {
	service service.Identifier
}

func NewIdentifierService(identifierSvc service.Identifier) Identifier {
	return Identifier{
		service: identifierSvc,
	}
}
