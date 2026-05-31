package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type user struct {
	ID        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    string        `json:"user_id" bson:"user_id"`
	FirstName string        `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastNmae  string        `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Email     string        `json:"email" bson:"email" validate:"required,email"`
	Password  string        `json:"password" bson:"password" validate:"required,min=6"`
	// need to hash it
	Role string `json:"role" bson:"role" validate:"oneof= ADMIN USER"`
	//only 2 roles- admin or user
	CreatedAt time.Time `json:"created_at" bson:"created _at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated _at"`
	Token     string    `json:"token" bson:"token"`

	//we r going to use JWT/ Json web Tokens for authentication and  authorization purposes in this application
	//we'll persist the values with revlant tokens within the documents tat is represented in code by this user struct
	// a JWT is comoact URL safe token used to securely transmit info between parties as JSON object
	//header -Payload -Signature 3 parts of JWT

	RefreshToken   string  `json:"refresh_token" bson:"refresh_token"`
	FavouriteGenre []Genre `json:"favourite_genre" bson:"favourite_genre" vaidate:"required,dive"`
}
