package schema

import "time"

const DomainSchemaName = "domains"

type Domain struct {
	Domain                 string `bson:"domain"`
	BasePricePerIdentifier uint   `bson:"base_price_per_identifier"`
	DefaultTTL             uint32 `bson:"default_ttl"`
	Status                 string `bson:"status"`

	ID        interface{} `bson:"_id"`
	CreatedAt time.Time   `bson:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at"`
}
