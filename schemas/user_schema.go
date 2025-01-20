package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserSchemaName = "users"

type User struct {
	PubKey string `bson:"pubkey"`

	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
