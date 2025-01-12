package repositories

import (
	"context"
	"errors"
	"time"

	schema "github.com/dezh-tech/panda/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DomainRepository struct {
	*BaseRepository
}

func NewDoaminRepository(client *mongo.Client, dbName string, timeout time.Duration) *DomainRepository {
	return &DomainRepository{
		BaseRepository: NewBaseRepository(client, dbName, schema.DOMAIN_SCHEMA_NAME, timeout),
	}
}

func (r *DomainRepository) Add(ctx context.Context, d *schema.Domain) (*mongo.InsertOneResult, error) {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	return r.InsertOne(ctx, d)
}

func (r *DomainRepository) GetByField(ctx context.Context, fieldName string, value interface{}) (*schema.Domain, error) {
	var result *schema.Domain
	err := r.FindOne(ctx, bson.M{fieldName: value}, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *DomainRepository) GetAll(ctx context.Context, filter interface{}) (*[]schema.Domain, error) {
	results := new([]schema.Domain)
	err := r.FindAll(ctx, filter, results)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &[]schema.Domain{}, nil
		}

		return nil, err
	}
	
	return results, nil
}

func (r *DomainRepository) Update(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	return r.UpdateOne(ctx, filter, update)
}

func (r *DomainRepository) Delete(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return r.DeleteOne(ctx, filter)
}
