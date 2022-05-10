package application

import (
	"net/http"

	"github.com/patrick-salvatore/fileshare-service/dao/files"
	"github.com/patrick-salvatore/fileshare-service/dao/health"
)

// HandlerFunc is a custom implementation of the http.HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request, AppEnv)

// MakeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func MakeHandler(appEnv AppEnv, fn func(http.ResponseWriter, *http.Request, AppEnv)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Terry Pratchett tribute
		w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Vary", "Access-Control-Request-Method")
		w.Header().Set("Vary", "Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// return function with AppEnv
		fn(w, r, appEnv)
	}
}

// HealthcheckHandler returns useful info about the app
func HealthcheckHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	check := health.Check{
		AppName: "bulwark_backend_service",
		Version: appEnv.Version,
	}
	appEnv.Render.JSON(w, http.StatusOK, check)
}

// HealthcheckHandler returns useful info about the app
func GetFilesHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	files, err := files.GetFiles()

	if err != nil {
		appEnv.Render.JSON(w, http.StatusInternalServerError, err)
	} else {
		appEnv.Render.JSON(w, http.StatusOK, files)
	}

}

// HealthcheckHandler returns useful info about the app
func GetFileHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	files, _ := files.GetFile(req)
	appEnv.Render.JSON(w, http.StatusOK, files)
}

// HealthcheckHandler returns useful info about the app
func PostFileHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	files, _ := files.PostFiles(req)
	appEnv.Render.JSON(w, http.StatusOK, files)
}
