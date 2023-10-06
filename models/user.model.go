package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId       string             `json:"userId" bson:"userId"`
	FirstName    string             `json:"firstName" bson:"firstName"`
	LastName     string             `json:"lastName" bson:"lastName"`
	Email        string             `json:"email" bson:"email"`
	MobileNumber string             `json:"mobileNumber" bson:"mobileNumber"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
}
