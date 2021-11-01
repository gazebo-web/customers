package client

import (
	"context"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/api"
)

// client contains the HTTP client to connect to a api.CustomersV1 service.
type client struct{}

// GetCustomerByHandle performs an HTTP request to get a customer based on its handle.
func (c *client) GetCustomerByHandle(ctx context.Context, req api.GetCustomerByHandleRequest) (api.GetCustomerResponse, error) {
	panic("implement me")
}

// GetCustomerByID performs an HTTP request to get a customer based on its id.
func (c *client) GetCustomerByID(ctx context.Context, req api.GetCustomerByIDRequest) (api.GetCustomerResponse, error) {
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
