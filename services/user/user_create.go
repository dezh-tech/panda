package service

import (
	"context"

	schema "github.com/dezh-tech/panda/schemas"
)

func (u User) Create(ctx context.Context, pubKey string) (interface{}, error) {
	// Check if the user already exists
	d, err := u.repo.GetByField(ctx, "pubkey", pubKey)
	if err != nil {
		return nil, err
	}

	if d != nil {
		return nil, ErrIsExist
	}

	id, err := u.repo.Add(ctx, &schema.User{
		PubKey: pubKey,
	})
	if err != nil {
		return nil, err
	}

	return id.InsertedID, nil
}
