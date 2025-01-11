package domainService

import (
	schema "github.com/dezh-tech/panda/schemas"
	"github.com/dezh-tech/panda/pkg/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Add(usr schema.Domain) (*mongo.InsertOneResult, error)
	GetByDomain(domain string) (*schema.Domain, error)
}

type DomainService struct {
	repo      Repository
	validator *validator.Validator
}

func New(repo Repository) DomainService {
	return DomainService{repo: repo, validator: validator.NewValidator()}
}
