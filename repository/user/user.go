package user

import (
	"context"
	"errors"

	"github.com/dezh-tech/geb/entity"
	"github.com/dezh-tech/geb/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	db *database.Database
}

func New(db *database.Database) User {
	return User{
		db: db,
	}
}

func (u User) Add(usr entity.User) error {
	collection := u.db.Client.Database(u.db.DBName).Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), u.db.QueryTimeout)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{
		"name":   usr.Name,
		"pubkey": usr.Pubkey,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u User) GetByPubkey(pubkey string) (entity.User, error) {
	collection := u.db.Client.Database(u.db.DBName).Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), u.db.QueryTimeout)
	defer cancel()

	var usr entity.User
	err := collection.FindOne(ctx, bson.M{"pubkey": pubkey}).Decode(&usr)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return usr, nil
		}

		return usr, err
	}

	return usr, nil
}
