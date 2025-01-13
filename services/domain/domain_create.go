package service

import (
	"context"

	"github.com/dezh-tech/panda/pkg/validator"
	"github.com/dezh-tech/panda/schemas"
)

func (s Domain) Create(ctx context.Context, domain, status string,
	basePricePerIdentifier uint, defaultTTL uint32,
) (interface{}, *validator.Varror) {
	// Check if the domain already exists
	d, err := s.repo.GetByField(ctx, "Domain", domain)
	if err != nil {
		return nil, &validator.Varror{Error: err.Error()}
	}

	if d != nil {
		return nil, &validator.Varror{Error: ErrIsExist.Error()}
	}

	id, err := s.repo.Add(ctx, &schema.Domain{
		Domain:                 domain,
		BasePricePerIdentifier: basePricePerIdentifier,
		DefaultTTL:             defaultTTL,
		Status:                 status,
	})
	if err != nil {
		return nil, &validator.Varror{Error: err.Error()}
	}

	return id.InsertedID, nil
}
