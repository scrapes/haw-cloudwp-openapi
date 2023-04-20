/*
 * cloudwpss23-openapi-cyan
 *
 * OpenAPI Reference für das CloudWP der HAW Hamburg für das SommerSemster 2023
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"os"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	GetApiExternal(http.ResponseWriter, *http.Request)
	GetFiles(http.ResponseWriter, *http.Request)
	GetFilesName(http.ResponseWriter, *http.Request)
	GetUsersUserId(http.ResponseWriter, *http.Request)
	OptionsApiExternal(http.ResponseWriter, *http.Request)
	OptionsFileUpload(http.ResponseWriter, *http.Request)
	OptionsFiles(http.ResponseWriter, *http.Request)
	OptionsFilesName(http.ResponseWriter, *http.Request)
	OptionsUser(http.ResponseWriter, *http.Request)
	OptionsUsersUserId(http.ResponseWriter, *http.Request)
	PatchUsersUserId(http.ResponseWriter, *http.Request)
	PostUser(http.ResponseWriter, *http.Request)
	PutFileUpload(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	GetApiExternal(context.Context) (ImplResponse, error)
	GetFiles(context.Context) (ImplResponse, error)
	GetFilesName(context.Context, string) (ImplResponse, error)
	GetUsersUserId(context.Context, int32) (ImplResponse, error)
	OptionsApiExternal(context.Context) (ImplResponse, error)
	OptionsFileUpload(context.Context) (ImplResponse, error)
	OptionsFiles(context.Context) (ImplResponse, error)
	OptionsFilesName(context.Context, string) (ImplResponse, error)
	OptionsUser(context.Context) (ImplResponse, error)
	OptionsUsersUserId(context.Context, int32) (ImplResponse, error)
	PatchUsersUserId(context.Context, int32, PatchUsersUserIdRequest) (ImplResponse, error)
	PostUser(context.Context, PostUserRequest) (ImplResponse, error)
	PutFileUpload(context.Context, string, *os.File) (ImplResponse, error)
}
