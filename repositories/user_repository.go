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

type User struct {
	*Base
}

func NewUserRepository(client *mongo.Client, dbName string, timeout time.Duration) *User {
	return &User{
		Base: NewBaseRepository(client, dbName, schema.UserSchemaName, timeout),
	}
}

func (r *User) Add(ctx context.Context, d *schema.User) (*mongo.InsertOneResult, error) {
	d.ID = primitive.NewObjectID()
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	return r.InsertOne(ctx, d)
}

func (r *User) GetByField(ctx context.Context,
	fieldName string, value interface{},
) (*schema.User, error) {
	result := new(schema.User)
	err := r.FindOne(ctx, bson.M{fieldName: value}, result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}

func (r *User) GetAll(ctx context.Context, filter interface{}) (*[]schema.User, error) {
	results := new([]schema.User)
	err := r.FindAll(ctx, filter, results)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &[]schema.User{}, nil
		}

		return nil, err
	}

	return results, nil
}

func (r *User) Update(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	return r.UpdateOne(ctx, filter, update)
}

func (r *User) Delete(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return r.DeleteOne(ctx, filter)
}
