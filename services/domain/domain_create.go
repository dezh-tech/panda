package service

import (
	"context"

	schema "github.com/dezh-tech/panda/schemas"
)

func (d Domain) Create(ctx context.Context, newDomain, status string,
	basePricePerIdentifier uint, defaultTTL uint32,
) (interface{}, error) {
	// Check if the domain already exists
	domain, err := d.repo.GetByField(ctx, "Domain", newDomain)
	if err != nil {
		return nil, err
	}

	if domain != nil {
		return nil, ErrIsExist
	}

	id, err := d.repo.Add(ctx, &schema.Domain{
		Domain:                 newDomain,
		BasePricePerIdentifier: basePricePerIdentifier,
		DefaultTTL:             defaultTTL,
		Status:                 status,
	})
	if err != nil {
		return nil, err
	}

	return id.InsertedID, nil
}
