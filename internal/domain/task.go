package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Code string             `json:"code" bson:"_id"`
}
