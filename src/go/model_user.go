/*
 * cloudwpss23-openapi-cyan
 *
 * OpenAPI Reference für das CloudWP der HAW Hamburg für das SommerSemster 2023
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// User - 
type User struct {

	// Unique identifier for the given user.
	Id *interface{} `json:"id"`

	FirstName *interface{} `json:"firstName"`

	LastName *interface{} `json:"lastName"`

	Email *interface{} `json:"email"`

	DateOfBirth *interface{} `json:"dateOfBirth,omitempty"`

	// Set to true if the user's email has been verified.
	EmailVerified *interface{} `json:"emailVerified"`

	// The date that the user was created.
	CreateDate *interface{} `json:"createDate,omitempty"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"firstName": obj.FirstName,
		"lastName": obj.LastName,
		"email": obj.Email,
		"emailVerified": obj.EmailVerified,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUserRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of User (e.g. [][]User), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUser, ok := obj.(User)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserRequired(aUser)
	})
}
