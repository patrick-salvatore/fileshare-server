package application

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrick-salvatore/fileshare-service/db"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
	"github.com/urfave/negroni"
)

// StartServer Wraps the mux Router and uses the Negroni Middleware
func StartServer(appEnv AppEnv) {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range AppRoutes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(MakeHandler(appEnv, route.HandlerFunc))
	}

	// security
	var isDevelopment = false
	if appEnv.Env == "LOCAL" {
		isDevelopment = true
	}
	secureMiddleware := secure.New(secure.Options{
		// This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains
		// options to be ignored during development. When deploying to production,
		// be sure to set this to false.
		IsDevelopment: isDevelopment,
		// AllowedHosts is a list of fully qualified domain names that are allowed (CORS)
		AllowedHosts: []string{},
		// If ContentTypeNosniff is true, adds the X-Content-Type-Options header
		// with the value `nosniff`. Default is false.
		ContentTypeNosniff: true,
		// If BrowserXssFilter is true, adds the X-XSS-Protection header with the
		// value `1; mode=block`. Default is false.
		BrowserXssFilter: true,
	})

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.UseHandler(c.Handler(router))

	// start now
	var Port string
	if appEnv.Env == "LOCAL" {
		Port = "localhost:" + appEnv.Port
	} else {
		Port = ":" + appEnv.Port
	}

	srv := &http.Server{
		Addr:           Port,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// initalize db
		db.DatabaseInstance.Setup()
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// print to std output to represent service has started
	startupMessage := "===================> Starting app (v" + appEnv.Version + ")"
	startupMessage = startupMessage + " on port " + appEnv.Port
	startupMessage = startupMessage + " in " + appEnv.Env + " mode."
	log.Println(startupMessage)

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		var disconnectErr error
		if disconnectErr = db.DatabaseInstance.Shutdown(); disconnectErr != nil {
			panic(disconnectErr)
		}
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
