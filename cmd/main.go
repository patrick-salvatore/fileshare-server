package main

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/unrolled/render"

	application "github.com/patrick-salvatore/fileshare-service/application"
	v "github.com/patrick-salvatore/fileshare-service/internal/version"
	log "github.com/sirupsen/logrus"
)

// First init here is meant to load environment variables
// very crucial to keep this init function where this is
func init() {
	args := os.Args[1:]

	var LOADED = false
	for i := 0; i < len(args); i++ {
		res := strings.Split(args[i], "=")

		if res[0] == "-env" {
			err := godotenv.Load("../env." + strings.ToLower(res[0]))
			LOADED = true

			if err != nil {
				log.Fatal("Error loading .env file")
			}

		}
	}

	// default to loading local env var because i'm a lazy developer
	if !LOADED {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("Error loading LOCAL .env file")
		}
	}

	println(os.Getenv("ENV"))
}

func init() {
	if "LOCAL" == strings.ToUpper(os.Getenv("ENV")) {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	// Grab environment variables

	var (
		env     = strings.ToUpper(os.Getenv("ENV")) // LOCAL, DEV, STG, PRD
		port    = os.Getenv("PORT")                 // server traffic on this port
		version = os.Getenv("VERSION")              // path to VERSION file
	)
	// Read version information
	version, err := v.ParseVersionFile(version)
	if err != nil {
		log.WithFields(log.Fields{
			"env":  env,
			"err":  err,
			"path": os.Getenv("VERSION"),
		}).Fatal("Can't find a VERSION file")
		return
	}
	log.WithFields(log.Fields{
		"env":     env,
		"path":    os.Getenv("VERSION"),
		"version": version,
		"db":      os.Getenv("MONGO_URL"),
	}).Info("Loaded VERSION file")

	// Initialise application context
	appEnv := application.AppEnv{
		// dbClient: dbClient,
		Render:  render.New(),
		Version: version,
		Env:     env,
		Port:    port,
	}

	// Start application
	application.StartServer(appEnv)
}
