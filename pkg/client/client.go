package client

import (
	"context"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/api"
	"gitlab.com/ignitionrobotics/web/ign-go/encoders"
	"gitlab.com/ignitionrobotics/web/ign-go/net"
	"net/http"
	"net/url"
	"time"
)

// client contains the HTTP client to connect to a api.CustomersV1 service.
type client struct {
	client net.Client
}

// GetCustomerByHandle performs an HTTP request to get a customer based on its handle.
func (c *client) GetCustomerByHandle(ctx context.Context, in api.GetCustomerByHandleRequest) (api.CustomerResponse, error) {
	var out api.CustomerResponse
	if err := c.client.Call(ctx, "GetCustomerByHandle", &in, &out); err != nil {
		return api.CustomerResponse{}, err
	}
	return out, nil
}

// GetCustomerByID performs an HTTP request to get a customer based on its id.
func (c *client) GetCustomerByID(ctx context.Context, in api.GetCustomerByIDRequest) (api.CustomerResponse, error) {
	var out api.CustomerResponse
	if err := c.client.Call(ctx, "GetCustomerByID", &in, &out); err != nil {
		return api.CustomerResponse{}, err
	}
	return out, nil
}

// CreateCustomer performs an HTTP request to create a new customer for a certain application.
func (c *client) CreateCustomer(ctx context.Context, in api.CreateCustomerRequest) (api.CustomerResponse, error) {
	var out api.CustomerResponse
	if err := c.client.Call(ctx, "CreateCustomer", &in, &out); err != nil {
		return api.CustomerResponse{}, err
	}
	return out, nil
}

// Client holds methods to interact with the api.CustomersV1.
type Client interface {
	api.CustomersV1
}

// NewCustomersClientV1 initializes a new api.CustomersV1 client implementation using an HTTP client.
func NewCustomersClientV1(baseURL *url.URL, timeout time.Duration) Client {
	endpoints := map[string]net.EndpointHTTP{
		"GetCustomerByHandle": {
			Method: http.MethodPost,
			Path:   "/customers/search/handle",
		},
		"GetCustomerByID": {
			Method: http.MethodPost,
			Path:   "/customers/search/id",
		},
		"CreateCustomer": {
			Method: http.MethodPost,
			Path:   "/customers",
		},
	}
	return &client{
		client: net.NewClient(net.NewCallerHTTP(baseURL, endpoints, timeout), encoders.JSON),
	}
}
