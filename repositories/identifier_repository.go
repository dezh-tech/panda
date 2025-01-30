package repositories

import (
	"context"
	"errors"
	"time"

	schema "github.com/dezh-tech/panda/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Identifier struct {
	*Base
}

func NewIdentifierRepository(client *mongo.Client, dbName string, timeout time.Duration) *Identifier {
	return &Identifier{
		Base: NewBaseRepository(client, dbName, schema.IdentifierSchemaName, timeout),
	}
}

func (r *Identifier) Add(ctx context.Context, d *schema.Identifier) (*mongo.InsertOneResult, error) {
	d.ID = primitive.NewObjectID()
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	return r.InsertOne(ctx, d)
}

func (r *Identifier) GetByField(ctx context.Context,
	fieldName string, value interface{},
) (*schema.Identifier, error) {
	result := new(schema.Identifier)
	err := r.FindOne(ctx, bson.M{fieldName: value}, result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}

func (r *Identifier) GetAll(ctx context.Context, filter interface{}) (*[]schema.Identifier, error) {
	results := new([]schema.Identifier)
	err := r.FindAll(ctx, filter, results)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &[]schema.Identifier{}, nil
		}

		return nil, err
	}

	return results, nil
}

func (r *Identifier) Update(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	return r.UpdateOne(ctx, filter, update)
}

func (r *Identifier) Delete(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return r.DeleteOne(ctx, filter)
}
