package domain

import "time"

type Event struct {
	Id        string    `bson:"_id,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Count     int       `bson:"count"`
}

//func getRepository() (repository.EventRepository, error) {
//
//	rep, err := mongo.GetRepositoryInstance()
//	if err != nil {
//		return nil, err
//	}
//	return rep, nil
//}
