package router

import (
	"net/http"

	"github.com/KannelRetry/controllers/healthcheck"
	"github.com/gorilla/mux"
)

// NewRouter returns a router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range internalRoutes {
		switch route.AuthMiddleware {
		case 0:
			router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
			break

		}
	}
	return router
}

// NewHealthcheckRouter -
func NewHealthcheckRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").
		Path("/healthstatus").
		Name("Healthcheck api").
		Handler(http.HandlerFunc(healthcheck.HealthCheck))
	return router
}
