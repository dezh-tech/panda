package domainService

import (
	"github.com/dezh-tech/panda/pkg/validator"
	schema "github.com/dezh-tech/panda/schemas"
)

type DomainInsertArgs struct {
	Domain                 string
	BasePricePerIdentifier uint
	DefaultTTL             uint32
	Status                 string
}

type DomainInsertRes struct {
	ID interface{}
}

func (s DomainService) Create(req DomainInsertArgs) (*DomainInsertRes, *validator.Varror) {
	// Check if the domain already exists
	domain, err := s.repo.GetByDomain(req.Domain)
	if err != nil {
		return nil, &validator.Varror{Error: err.Error()}
	}

	if domain != nil {
		return nil, &validator.Varror{Error: ErrIsExist.Error()}
	}

	// Add the domain to the repository
	id, err := s.repo.Add(schema.Domain{
		Domain:                 req.Domain,
		BasePricePerIdentifier: req.BasePricePerIdentifier,
		DefaultTTL:             req.DefaultTTL,
		Status:                 req.Status,
	})

	if err != nil {
		return nil, &validator.Varror{Error: err.Error()}
	}

	// Return the response
	return &DomainInsertRes{ID: id.InsertedID}, nil
}
