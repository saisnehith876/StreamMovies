package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type user struct {
	ID        bson.ObjectID
	UserId    string
	FirstName string
	LastNmae  string
	Email     string
	Password  string // need to hash it
	Role      string //only 2 roles- admin or user
	CreatedAt time.Time
	UpdatedAt time.Time
	Token     string

	//we r going to use JWT/ Json web Tokens for authentication and  authorization purposes in this application
	//we'll persist the values with revlant tokens within the documents tat is represented in code by this user struct
	// a JWT is comoact URL safe token used to securely transmit info between parties as JSON object
	//header -Payload -Signature 3 parts of JWT
}

