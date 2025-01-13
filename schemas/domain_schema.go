package schema

const DomainSchemaName = "domains"

type Domain struct {
	Domain                 string `bson:"Domain"`
	BasePricePerIdentifier uint   `bson:"BasePricePerIdentifier"`
	DefaultTTL             uint32 `bson:"DefaultTTL"`
	Status                 string `bson:"Status"`

	Base
}
