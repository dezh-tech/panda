package domainhandler

import domainService "github.com/dezh-tech/panda/services/domainservice"

type Handler struct {
	domainService domainService.DomainService
}

func New(domainSvc domainService.DomainService) Handler {
	return Handler{
		domainService: domainSvc,
	}
}
