package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Category    []string           `json:"category,omitempty" bson:"category,omitempty"`
}
