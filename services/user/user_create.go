package service

import (
	"context"

	schema "github.com/dezh-tech/panda/schemas"
)

func (s User) Create(ctx context.Context, npub string) (interface{}, error) {
	// Check if the user already exists
	d, err := s.repo.GetByField(ctx, "npub", npub)
	if err != nil {
		return nil, err
	}

	if d != nil {
		return nil, err
	}

	id, err := s.repo.Add(ctx, &schema.User{
		Npub: npub,
	})
	if err != nil {
		return nil, err
	}

	return id.InsertedID, nil
}
