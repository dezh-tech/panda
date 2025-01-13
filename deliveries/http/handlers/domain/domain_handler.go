package handlers

import service "github.com/dezh-tech/panda/services/domain"

type Domain struct {
	service service.Domain
}

func NewDomainService(domainSvc service.Domain) Domain {
	return Domain{
		service: domainSvc,
	}
}
