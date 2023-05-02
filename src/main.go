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
	controller2 "github.com/scrapes/haw-cloudwp-openapi/src/controller"
	"github.com/scrapes/haw-cloudwp-openapi/src/service"
	openapi "github.com/scrapes/haw-cloudwp-openapi/src/v1/go"
	"log"
	"net/http"
	"os"
	"strings"
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
		corsOrigins = "http://localhost:3000,https://app.cloudwp.anwski.de,https://api.cloudwp.anwski.de"
	}

	s := new(service.V1Service)
	controller := new(controller2.V1Controller).Init(s)
	allowedOrigins := strings.Split(corsOrigins, ",")

	router := openapi.NewRouter(controller)

	router.Use(handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
	))

	//router.Use(middleware.EnsureValidToken(auth0Domain, auth0Audience))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
