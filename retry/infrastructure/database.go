package infrastructure

import (
	"os"
	"reflect"
)

type Database struct {
	ConnectionString string `env:"GOBBIT_DB_CONNECTIONSTRING"`
	Name             string `env:"GOBBIT_DB_NAME"`
}

func NewDatabase() *Database {
	connectionString := getValueFromEnvWithDefault("ConnectionString", "mongodb://localhost:27017")
	name := getValueFromEnvWithDefault("Name", "gobbit")
	return &Database{connectionString, name}
}

func getValueFromEnvWithDefault(fieldName string, def string) string {
	t := reflect.TypeOf(Database{})
	field, _ := t.FieldByName(fieldName)
	value := field.Tag.Get("env")
	return getEnvWithDefault(value, def)
}

func getEnvWithDefault(key string, def string) string {
	env := os.Getenv(key)
	if len(env) > 0 {
		return env
	}
	return def
}
