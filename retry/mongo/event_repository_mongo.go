package mongo

//import (
//	. "github.com/odair-pedro/gobbit/retry"
//	"github.com/odair-pedro/gobbit/retry/domain"
//	"github.com/odair-pedro/gobbit/retry/domain/repository"
//	"github.com/odair-pedro/gobbit/retry/infrastructure"
//	"gopkg.in/mgo.v2/bson"
//)
//
//var instance repository.EventRepository
//
//type eventRepositoryMongo struct {
//	db *infrastructure.Database
//}
//
//func newEventRepositoryMongo() (*eventRepositoryMongo, error) {
//	db, err := GetCurrentDatabase()
//	if err != nil {
//		return nil, err
//	}
//	return &eventRepositoryMongo{db}, nil
//}
//
//func GetRepositoryInstance() (repository.EventRepository, error) {
//	if instance == nil {
//		newInstance, err := newEventRepositoryMongo()
//		if err != nil {
//			return nil, err
//		}
//		instance = newInstance
//	}
//	return instance, nil
//}
//
//func (repository eventRepositoryMongo) Create(event *domain.Event) error {
//	return nil
//}
//
//func (repository eventRepositoryMongo) Delete(id bson.ObjectId) error {
//	return nil
//}
//
//func (repository eventRepositoryMongo) FindById(id bson.ObjectId) *domain.Event {
//	return nil
//}
//
//func (repository eventRepositoryMongo) Update(event *domain.Event) error {
//	return nil
//}
