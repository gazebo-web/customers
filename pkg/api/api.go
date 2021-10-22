package api

import "context"

// CustomersV1 holds the methods that allow managing customers.
type CustomersV1 interface {
	// GetCustomer returns the customer of a certain application.
	GetCustomer(ctx context.Context, req GetCustomerRequest) (GetCustomerResponse, error)

	// CreateCustomer creates a new customer for a certain application.
	CreateCustomer(ctx context.Context, req CreateCustomerRequest) (CreateCustomerResponse, error)
}

type GetCustomerRequest struct{}

type GetCustomerResponse struct{}

type CreateCustomerRequest struct{}

type CreateCustomerResponse struct{}
