package routes

import (
	"authentication/handlers"
	"authentication/middleware"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/signup", handlers.Signup)
	mux.HandleFunc("/login", handlers.Login)

	mux.Handle("/api/profile", middleware.AuthMiddleware(http.HandlerFunc(handlers.ProfileHandler)))
}
