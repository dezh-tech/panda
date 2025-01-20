package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const IdentifierSchemaName = "identifier"

type Identifier struct {
	Name           string    `bson:"name"`
	Pubkey         string    `bson:"pubkey"`
	DomainID       string    `bson:"domain_id"`
	ExpiresAt      time.Time `bson:"expires_at"`
	FullIdentifier string    `bson:"full_identifier"`

	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
