package client

import (
	"context"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/api"
)

// client contains the HTTP client to connect to the customers API.
type client struct{}

// GetIdentity performs an HTTP request to get a customer or a user of a certain application.
func (c *client) GetIdentity(ctx context.Context, req api.GetIdentityRequest) (api.GetIdentityResponse, error) {
	panic("implement me")
}

// CreateCustomer performs an HTTP request to create a new customer for a certain application.
func (c *client) CreateCustomer(ctx context.Context, req api.CreateCustomerRequest) (api.CreateCustomerResponse, error) {
	panic("implement me")
}

// Client holds methods to interact with the api.CustomersV1.
type Client interface {
	api.CustomersV1
}

// NewClient initializes a new api.CustomersV1 client implementation using an HTTP client.
func NewClient() Client {
	return &client{}
}
