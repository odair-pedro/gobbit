package infrastructure

import (
	"os"
	"testing"
)

func TestNewDatabase_ShouldReturn_DefaultConnectionString(t *testing.T) {
	const defaultConnectionString string = "mongodb://localhost:27017"
	db := NewDatabase()
	if db.ConnectionString != defaultConnectionString {
		t.Fail()
	}
}

func TestNewDatabase_ShouldReturn_EnvConnectionString(t *testing.T) {
	const envConnectionString string = "test-connectionstring"
	_ = os.Setenv("GOBBIT_DB_CONNECTIONSTRING", envConnectionString)
	db := NewDatabase()
	if db.ConnectionString != envConnectionString {
		t.Fail()
	}
}

func TestNewDatabase_ShouldReturn_DefaultName(t *testing.T) {
	const defaultName string = "gobbit"
	db := NewDatabase()
	if db.Name != defaultName {
		t.Fail()
	}
}

func TestNewDatabase_ShouldReturn_EnvName(t *testing.T) {
	const envName string = "test-name"
	_ = os.Setenv("GOBBIT_DB_NAME", envName)
	db := NewDatabase()
	if db.Name != envName {
		t.Fail()
	}
}
