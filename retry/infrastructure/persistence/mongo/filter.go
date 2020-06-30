package mongo

import "gopkg.in/mgo.v2/bson"

func FilterById(id string) interface{} {
	return bson.M{"_id": ConvertStringToObjectID(id)}
}
