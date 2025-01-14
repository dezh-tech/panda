package schema

import "time"

const UserSchemaName = "users"

type User struct {
	Npub string `bson:"npub"`

	ID        interface{} `bson:"_id"`
	CreatedAt time.Time   `bson:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at"`
}
