/*
 * customerfacing
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"github.com/gorilla/handlers"
	"github.com/scrapes/haw-cloudwp-openapi/src/Controller"
	"github.com/scrapes/haw-cloudwp-openapi/src/middleware"
	"log"
	"net/http"
	"os"
	"strings"

	openapi "github.com/scrapes/haw-cloudwp-openapi/src/go"
)

func main() {
	log.Printf("Server started")

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("defaulting to port %s", port)
	}

	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	auth0Audience := os.Getenv("AUTH0_AUDIENCE")
	corsOrigins := os.Getenv("CORS_ORIGINS")

	if auth0Domain == "" {
		auth0Domain = "dev-5n5igzycxiz22p3w.us.auth0.com"
	}

	if auth0Audience == "" {
		auth0Audience = "http://localhost:3001"
	}

	if corsOrigins == "" {
		corsOrigins = "http://localhost:3001,https://app.cloudwp.anwski.de,https://api.cloudwp.anwski.de"
	}

	DefaultApiService := new(Controller.ApiController)
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)
	allowedOrigins := strings.Split(corsOrigins, ",")

	router := openapi.NewRouter(DefaultApiController)
	router.Use(handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"Authorization"})))

	router.Use(middleware.EnsureValidToken(auth0Domain, auth0Audience))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
