package domainService

import (
	"context"

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

func (s DomainService) Create(ctx context.Context, req DomainInsertArgs) (*DomainInsertRes, *validator.Varror) {
	// Check if the domain already exists
	domain, err := s.repo.GetByField(ctx, "Domain", req.Domain)
	if err != nil {
		return nil, &validator.Varror{Error: err.Error()}
	}

	if domain != nil {
		return nil, &validator.Varror{Error: ErrIsExist.Error()}
	}

	id, err := s.repo.Add(ctx, schema.Domain{
		Domain:                 req.Domain,
		BasePricePerIdentifier: req.BasePricePerIdentifier,
		DefaultTTL:             req.DefaultTTL,
		Status:                 req.Status,
	})
	if err != nil {
		return nil, &validator.Varror{Error: err.Error()}
	}

	return &DomainInsertRes{ID: id.InsertedID}, nil
}
