package service

import (
	"context"
	"time"

	schema "github.com/dezh-tech/panda/schemas"
	domainService "github.com/dezh-tech/panda/services/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i Identifier) Create(ctx context.Context, name, domainID, pubkey string) (interface{}, error) {

	d, err := i.checkDomain(ctx, domainID)
	if err != nil {
		return nil, err
	}

	if !i.checkUser(ctx, pubkey) {
		return nil, Err
	}

	// TODO ::: check method for fullIdentifier
	f := name + "@" + d

	fi, err := i.repo.GetByField(ctx, "full_identifier", f)
	if err != nil {
		return nil, err
	}

	if fi != nil {
		return nil, ErrIsExist
	}

	id, err := i.repo.Add(ctx, &schema.Identifier{
		Name:           name,
		Pubkey:         pubkey,
		DomainID:       domainID,
		ExpiresAt:      time.Time{},
		FullIdentifier: f,
	})
	if err != nil {
		return nil, err
	}

	return id.InsertedID, nil
}

func (i Identifier) checkDomain(ctx context.Context, domainID string) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(domainID)
	if err != nil {
		return "", err
	}
	d, err := i.domainService.GetByField(ctx, "_id", objectID)
	if err != nil {
		return "", err
	}

	if d == nil {
		return "", domainService.ErrNotFound
	}

	return d.Domain, nil
}

func (i Identifier) checkUser(ctx context.Context, pubkey string) bool {
	u, err := i.userService.GetByField(ctx, "pubkey", pubkey)
	if err != nil {
		return false
	}

	if u == nil {
		return false
	}

	return true
}
