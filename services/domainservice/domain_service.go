package domainservice

import (
	"context"

	"github.com/dezh-tech/panda/pkg/validator"
	schema "github.com/dezh-tech/panda/schemas"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Add(ctx context.Context, entity *schema.Domain) (*mongo.InsertOneResult, error)
	GetByField(ctx context.Context, fieldName string, value interface{}) (*schema.Domain, error)
	GetAll(ctx context.Context, filter interface{}) (*[]schema.Domain, error)
	Update(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
}

type DomainService struct {
	repo      Repository
	validator *validator.Validator
}

func New(repo Repository) DomainService {
	return DomainService{repo: repo, validator: validator.NewValidator()}
}
