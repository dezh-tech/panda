package domainhandler

import domainService "github.com/dezh-tech/panda/services/domain"

type Handler struct {
	domainSvc domainService.DomainService
}

func New(domainSvc domainService.DomainService) Handler {
	return Handler{
		domainSvc: domainSvc,
	}
}
