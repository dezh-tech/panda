package service

import (
	"context"

	schema "github.com/dezh-tech/panda/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

func (u User) GetAll(ctx context.Context, filter interface{}) (*[]schema.User, error) {
	Users, err := u.repo.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	return Users, nil
}

func (u User) GetAllWithoutFilter(ctx context.Context) (*[]schema.User, error) {
	Users, err := u.GetAll(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	return Users, nil
}

func (s User) GetByField(ctx context.Context, fieldName string, value interface{}) (*schema.User, error) {
	domain, err := s.repo.GetByField(ctx, fieldName, value)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
