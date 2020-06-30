package repository

import (
	"github.com/odair-pedro/gobbit/retry/domain"
	"gopkg.in/mgo.v2/bson"
)

type EventRepository interface {
	Create(event *domain.Event) error
	Delete(id bson.ObjectId) error
	FindById(id bson.ObjectId) *domain.Event
	Update(event *domain.Event) error
}
