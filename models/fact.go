package models

import "gopkg.in/mgo.v2/bson"

type (
	// Fact represents the structure of our resource
	Fact struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Name   string        `json:"name" bson:"name"`
		Gender string        `json:"gender" bson:"gender"`
		Age    int           `json:"age" bson:"age"`
	}
)
