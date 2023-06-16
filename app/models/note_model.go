package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID primitive.ObjectID `bson:"_id"`
}
