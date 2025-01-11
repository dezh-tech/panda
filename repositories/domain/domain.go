package domain

import (
	"context"
	"errors"
	"time"

	"github.com/dezh-tech/panda/infrastructures/database"
	schema "github.com/dezh-tech/panda/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Domain struct {
	db *database.Database
}

const SCHEMA_NAME = "users"

func New(db *database.Database) Domain {
	return Domain{
		db: db,
	}
}

func (u Domain) Add(d schema.Domain) (*mongo.InsertOneResult, error) {
	collection := u.db.Client.Database(u.db.DBName).Collection(SCHEMA_NAME)

	ctx, cancel := context.WithTimeout(context.Background(), u.db.QueryTimeout)
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.M{
		"Domain":                 d.Domain,
		"BasePricePerIdentifier": d.BasePricePerIdentifier,
		"DefaultTTL":             d.DefaultTTL,
		"Status":                 d.Status,

		"CreatedAt": time.Now(),
		"UpdatedAt": time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u Domain) GetByDomain(domain string) (*schema.Domain, error) {
	collection := u.db.Client.Database(u.db.DBName).Collection(SCHEMA_NAME)

	ctx, cancel := context.WithTimeout(context.Background(), u.db.QueryTimeout)
	defer cancel()

	var d *schema.Domain
	err := collection.FindOne(ctx, bson.M{"Domain": domain}).Decode(&d)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return d, nil
		}

		return d, err
	}

	return d, nil
}
