package mongo

import (
	"context"
	"errors"
	. "github.com/odair-pedro/gobbit/retry/domain"
	"github.com/odair-pedro/gobbit/retry/infrastructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION string = "events"

type eventRepository struct {
	collection *mongo.Collection
}

func NewEventRepository(ctx context.Context, db *infrastructure.Database) (*eventRepository, error) {
	opt := options.Client().ApplyURI(db.ConnectionString)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, err
	}

	collection := client.Database(db.Name).Collection(COLLECTION)
	return &eventRepository{collection}, nil
}

func (repository *eventRepository) Create(ctx context.Context, event *Event) error {
	r, err := repository.collection.InsertOne(ctx, event)
	if r == nil {
		return errors.New("")
	}
	event.Id = ConvertObjectIDToString(r.InsertedID.(primitive.ObjectID))
	return err
}

func (repository *eventRepository) Delete(ctx context.Context, id string) error {
	r := repository.collection.FindOneAndDelete(ctx, FilterById(id))
	var event *Event
	return r.Decode(event)
}

func (repository *eventRepository) FindById(ctx context.Context, id string) *Event {
	var result *Event
	_ = repository.collection.FindOne(ctx, FilterById(id)).Decode(&result)
	return result
}

func (repository *eventRepository) Update(ctx context.Context, event *Event) error {
	_, err := repository.collection.ReplaceOne(ctx, FilterById(event.Id), event)
	return err
}
