package domain

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Config struct {
	Id             bson.ObjectId `bson:"_id"`
	CreatedAt      time.Time     `bson:"created_at"`
	UpdatedAt      time.Time     `bson:"updated_at"`
	HandlerId      string        `bson:"handler_id"`
	ExecutionLimit uint          `bson:"execution_limit"`
	Delay          uint          `bson:"delay"`
}
