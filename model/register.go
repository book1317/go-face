package model

type Register struct {
	Profile Profile `json:"profile" bson:"profile"`
	Account Account `json:"account" bson:"account"`
}
