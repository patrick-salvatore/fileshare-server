package application

import (
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/mongo"
)

// AppEnv holds application configuration data
type AppEnv struct {
	dbClient *mongo.Client
	Render   *render.Render
	Version  string
	Env      string
	Port     string
}

// CreateContextForTestSetup initialises an application context struct
// for testing purposes
func CreateContextForTestSetup() AppEnv {
	testVersion := "0.0.0"
	appEnv := AppEnv{
		Render:  render.New(),
		Version: testVersion,
		Env:     "LOCAL",
		Port:    "3001",
	}
	return appEnv
}
