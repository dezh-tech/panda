package service

import (
	"context"

	schema "github.com/dezh-tech/panda/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

func (d Domain) GetAll(ctx context.Context, filter interface{}) (*[]schema.Domain, error) {
	domains, err := d.repo.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	return domains, nil
}

func (d Domain) GetAllWithoutFilter(ctx context.Context) (*[]schema.Domain, error) {
	domains, err := d.GetAll(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	return domains, nil
}
