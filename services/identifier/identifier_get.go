package service

import (
	"context"

	schema "github.com/dezh-tech/panda/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

func (d Identifier) GetAll(ctx context.Context, filter interface{}) (*[]schema.Identifier, error) {
	identifiers, err := d.repo.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	return identifiers, nil
}

func (d Identifier) GetAllWithoutFilter(ctx context.Context) (*[]schema.Identifier, error) {
	identifiers, err := d.GetAll(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	return identifiers, nil
}

func (s Identifier) GetByField(ctx context.Context, fieldName string, value interface{}) (*schema.Identifier, error) {
	identifier, err := s.repo.GetByField(ctx, fieldName, value)
	if err != nil {
		return nil, err
	}

	return identifier, nil
}

func (d Identifier) GetAllByPubKey(ctx context.Context, pubKey string) (*[]schema.Identifier, error) {
	identifiers, err := d.GetAll(ctx, bson.M{"pubkey": pubKey})
	if err != nil {
		return nil, err
	}

	return identifiers, nil
}
