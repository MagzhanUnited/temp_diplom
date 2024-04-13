package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person Model
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required,alpha"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required,alpha"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,alpha"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,alpha"`
}

type Residential struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title              string             `json:"title,omitempty" bson:"title,omitempty"`
	Description        string             `json:"description,omitempty" bson:"description,omitempty"`
	Price              string             `json:"price,omitempty" bson:"price,omitempty"`
	DiscountPercentage float64            `json:"discountPercentage,omitempty" bson:"discountPercentage,omitempty"`
	Rating             float64            `json:"rating,omitempty" bson:"rating,omitempty"`
	Category           string             `json:"category,omitempty" bson:"category,omitempty"`
	Thumbnail          []string           `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
	Builder            string             `json:"builder,omitempty" bson:"builder,omitempty"`
	Schemas            []string           `json:"schemas,omitempty" bson:"schemas,omitempty"`
}

type Order struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	City   string             `json:"city,omitempty" bson:"city,omitempty"`
	Number string             `json:"number,omitempty" bson:"number,omitempty"`
}
