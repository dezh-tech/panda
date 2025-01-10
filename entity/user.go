package entity

type User struct {
	Name   string `bson:"name"`
	Pubkey string `bson:"pubkey"`
}
