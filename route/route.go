package route

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/riankrishandi/dans/controller"
	"github.com/riankrishandi/dans/middleware"
)

func NewAPI(c *controller.Controller, m *middleware.Middleware) (http.Handler, error) {
	router := mux.NewRouter().StrictSlash(true)

	// API routes.
	{
		v1 := router.PathPrefix("/api/v1").Subrouter()

		// Publics.
		v1.Handle("/login", c.HandleLogin()).Methods(http.MethodPost)

		// Privates.
		v1.Handle("/job", m.Auth(c.HandleGetJobList())).Methods(http.MethodGet)
		v1.Handle("/job/{jobID}", m.Auth(c.HandleGetJobDetail())).Methods(http.MethodGet)
	}

	return handlers.CORS(handlers.AllowCredentials())(router), nil
}
