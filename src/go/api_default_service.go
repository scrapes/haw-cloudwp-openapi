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
	"errors"
	"net/http"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// GetApiExternal - Your GET endpoint
func (s *DefaultApiService) GetApiExternal(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetApiExternal with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetApiExternal200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetApiExternal200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetApiExternal method not implemented")
}

// GetUsersUserId - Get User Info by User ID
func (s *DefaultApiService) GetUsersUserId(ctx context.Context, userId interface{}) (ImplResponse, error) {
	// TODO - update GetUsersUserId with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	//return Response(200, User{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUsersUserId method not implemented")
}

// OptionsApiExternal -
func (s *DefaultApiService) OptionsApiExternal(ctx context.Context) (ImplResponse, error) {
	// TODO - update OptionsApiExternal with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("OptionsApiExternal method not implemented")
}

// OptionsUser -
func (s *DefaultApiService) OptionsUser(ctx context.Context) (ImplResponse, error) {
	// TODO - update OptionsUser with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("OptionsUser method not implemented")
}

// OptionsUsersUserId -
func (s *DefaultApiService) OptionsUsersUserId(ctx context.Context, userId interface{}) (ImplResponse, error) {
	// TODO - update OptionsUsersUserId with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("OptionsUsersUserId method not implemented")
}

// PatchUsersUserId - Update User Information
func (s *DefaultApiService) PatchUsersUserId(ctx context.Context, userId interface{}, patchUsersUserIdRequest PatchUsersUserIdRequest) (ImplResponse, error) {
	// TODO - update PatchUsersUserId with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	//return Response(200, User{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(409, {}) or use other options such as http.Ok ...
	//return Response(409, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("PatchUsersUserId method not implemented")
}

// PostUser - Create New User
func (s *DefaultApiService) PostUser(ctx context.Context, postUserRequest PostUserRequest) (ImplResponse, error) {
	// TODO - update PostUser with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	//return Response(200, User{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(409, {}) or use other options such as http.Ok ...
	//return Response(409, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostUser method not implemented")
}
