package api

import "context"

// CustomersV1 holds the methods that allow managing customers.
type CustomersV1 interface {
	// GetIdentity returns the customer or a user of a certain application.
	GetIdentity(ctx context.Context, req GetIdentityRequest) (GetIdentityResponse, error)

	// CreateCustomer creates a new customer for a certain application.
	CreateCustomer(ctx context.Context, req CreateCustomerRequest) (CreateCustomerResponse, error)
}

// GetIdentityRequest is used to get the identity of a certain customer or user.
// If User is passed, it returns the Customer.
// If Customer is passed, it returns the User.
// If both values are passed, it returns an error.
type GetIdentityRequest struct {
	// User is the username of a certain application that will be returned.
	// Mutually exclusive with Customer.
	User string
	// Customer is the customer id of the user that will be returned.
	// Mutually exclusive with User.
	Customer string
	// Service is the payment service used to register the customer.
	Service string
	// Application is the application that originated the creation of the customer.
	Application string
}

// GetIdentityResponse is used to return a user or a customer depending on the parameters set in GetIdentityRequest.
type GetIdentityResponse struct {
	// User contains the username.
	User *string
	// Customer contains the customer id.
	Customer *string
	// Service is the payment service used to register the customer.
	Service string
	// Application is the application that originated the creation of the customer.
	Application string
}

type CreateCustomerRequest struct{}

type CreateCustomerResponse struct{}
