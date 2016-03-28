package env

import (
	"os"
)


// Environment is an interface for getting objects for the environment.
type Environment interface {
	// GetDBHost returns the hostname to MongoDB.
	GetDBHost() string

	// GetDBName returns the database name.
	GetDBName() string
}

var env Environment

func Init() {
	switch os.Getenv("GOSEARCH_ENV") {
	case "test":
		env = &TestEnvironment{}
	case "develop":
		env = &DevelopEnvironment{}
	default:
		env = &DevelopEnvironment{}
	}
}

func GetDBHost() string {
	return env.GetDBHost()
}

func GetDBName() string {
	return env.GetDBName()
}
