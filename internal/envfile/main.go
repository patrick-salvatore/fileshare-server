package envfile

import (
	"log"
	"os"
	"strings"
)

var LOADED = false

func LoadEnvFile(filePath string) {
	// default to loading local env var because i'm a lazy developer
	if filePath == "" {
		filePath = "LOCAL"
	}

	if LOADED {
		log.Fatal("why are you loading env again?")
		return
	}

	f, err := os.ReadFile(".env." + strings.ToLower(filePath))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lines := strings.Split(string(string(f)), "\n")

	for _, line := range lines {
		split := strings.Split(line, "=")
		key, val := split[0], split[1]

		os.Setenv(key, val)
	}

	LOADED = true
}
