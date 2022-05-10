package application

// Route is the model for the router setup
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

// Routes are the main setup for our Router
type Routes []Route

var AppRoutes = Routes{
	Route{"Healthcheck", "GET", "/healthcheck", HealthcheckHandler},
	Route{"Healthcheck", "GET", "/files", GetFilesHandler},
	Route{"Healthcheck", "GET", "/files/{id}", GetFileHandler},
	Route{"Healthcheck", "POST", "/files", PostFileHandler},
}
