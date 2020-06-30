package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertStringToObjectID(value string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(value)
	return id
}

func ConvertObjectIDToString(id primitive.ObjectID) string {
	return id.Hex()
}
