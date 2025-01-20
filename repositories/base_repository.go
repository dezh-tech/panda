package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Base struct {
	Client       *mongo.Client
	DBName       string
	Collection   string
	QueryTimeout time.Duration
}

// NewBaseRepository creates a new BaseRepository instance.
func NewBaseRepository(client *mongo.Client, dbName, collection string, timeout time.Duration) *Base {
	return &Base{
		Client:       client,
		DBName:       dbName,
		Collection:   collection,
		QueryTimeout: timeout,
	}
}

// InsertOne inserts a single document into the collection.
func (r *Base) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()

	return collection.InsertOne(ctx, document)
}

// FindByField finds a single document by a specific field and value.
func (r *Base) FindByField(ctx context.Context, field string, value, result interface{}) error {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()

	filter := bson.M{field: value}
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// FindOne finds a single document matching the filter.
func (r *Base) FindOne(ctx context.Context, filter, result interface{}) error {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// FindAll finds all documents matching the filter.
func (r *Base) FindAll(ctx context.Context, filter, results interface{}) error {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, results)
}

// UpdateOne updates a single document matching the filter.
func (r *Base) UpdateOne(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()

	return collection.UpdateOne(ctx, filter, update)
}

// DeleteOne deletes a single document matching the filter.
func (r *Base) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()

	return collection.DeleteOne(ctx, filter)
}

// CountDocuments counts documents matching the filter.
func (r *Base) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()

	return collection.CountDocuments(ctx, filter)
}
