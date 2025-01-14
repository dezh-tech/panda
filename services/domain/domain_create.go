package service

import (
	"context"

	schema "github.com/dezh-tech/panda/schemas"
)

func (s Domain) Create(ctx context.Context, domain, status string,
	basePricePerIdentifier uint, defaultTTL uint32,
) (interface{}, error) {
	// Check if the domain already exists
	d, err := s.repo.GetByField(ctx, "Domain", domain)
	if err != nil {
		return nil, err
	}

	if d != nil {
		return nil, ErrIsExist
	}

	id, err := s.repo.Add(ctx, &schema.Domain{
		Domain:                 domain,
		BasePricePerIdentifier: basePricePerIdentifier,
		DefaultTTL:             defaultTTL,
		Status:                 status,
	})
	if err != nil {
		return nil, err
	}

	return id.InsertedID, nil
}
