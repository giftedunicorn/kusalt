package router

import (
	"net/http"

	"github.com/giftedunicorn/kusalt/middlewares"
	"github.com/giftedunicorn/kusalt/router/v1"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		next.ServeHTTP(w, r)
	})
}

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(corsMiddleware)

	for _, route := range routesv1.Routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middlewares.Logger(handler, route.Name)

		if route.Auth {
			handler = middlewares.Authenticate(handler)
		}

		if route.Verify {
			handler = middlewares.VerifySign(handler)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	return c.Handler(router)
}
