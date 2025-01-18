package service

import (
	"context"

	"github.com/dezh-tech/panda/pkg/validator"
	schema "github.com/dezh-tech/panda/schemas"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Add(ctx context.Context, schema *schema.User) (*mongo.InsertOneResult, error)
	GetByField(ctx context.Context, fieldName string, value interface{}) (*schema.User, error)
	GetAll(ctx context.Context, filter interface{}) (*[]schema.User, error)
	Update(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
}

type User struct {
	repo      Repository
	validator *validator.Validator
}

func NewUserService(repo Repository) User {
	return User{repo: repo, validator: validator.NewValidator()}
}
