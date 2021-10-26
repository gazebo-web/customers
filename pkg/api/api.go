package api

import "context"

// CustomersV1 holds the methods that allow managing customers.
type CustomersV1 interface {
	// GetIdentity returns the customer or a user of a certain application.
	GetIdentity(ctx context.Context, req GetIdentityRequest) (GetIdentityResponse, error)

	// CreateCustomer creates a new customer for a certain application.
	CreateCustomer(ctx context.Context, req CreateCustomerRequest) (CreateCustomerResponse, error)
}

type GetCustomerRequest struct{}

type GetCustomerResponse struct{}

type CreateCustomerRequest struct{}

type CreateCustomerResponse struct{}
