package api

import (
	"context"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/models"
)

// CustomersV1 holds the methods that allow managing customers.
type CustomersV1 interface {
	// GetCustomerByHandle returns customer information based on the customer's application handle.
	GetCustomerByHandle(ctx context.Context, req GetCustomerByHandleRequest) (GetCustomerResponse, error)

	// GetCustomerByHandle returns customer information based on the customer's external service identity.
	GetCustomerByID(ctx context.Context, req GetCustomerByIDRequest) (GetCustomerResponse, error)

	// CreateCustomer creates a new customer for a certain application.
	CreateCustomer(ctx context.Context, req CreateCustomerRequest) (CreateCustomerResponse, error)
}

// GetCustomerByHandleRequest is the input of the CustomersV1.GetCustomerByHandle method.
type GetCustomerByHandleRequest struct {
	// Handle is the customer identity in the context of a certain application.
	// E.g. application username, application organization name.
	Handle string

	// Service is the payment service provider the customer is registered in.
	// E.g. Stripe, PayPal
	Service string

	// Application is the application that originated the creation of the customer.
	Application string
}

// GetCustomerByIDRequest is the input of the CustomersV1.GetCustomerByID method.
type GetCustomerByIDRequest struct {
	// ID is the customer identity in the context of an external service.
	ID string

	// Service is the service provider the customer is registered in.
	// E.g. Stripe, PayPal
	Service string

	// Application is the application that originated the creation of the customer.
	Application string
}

// GetCustomerResponse is the output from the CustomersV1.GetCustomerByHandle and the CustomersV1.GetCustomerByID methods.
type GetCustomerResponse struct {
	// ID is the customer identity in the context of an external service.
	ID string

	// Handle is the customer identity in the context of a certain application.
	// E.g. application username, application organization name.
	Handle string

	// Service is the service provider the customer is registered in.
	// E.g. Stripe, PayPal
	Service string

	// Application is the application that originated the creation of the customer.
	Application string
}

// FromCustomer fills the current response with data from the given models.Customer.
func (res *GetCustomerResponse) FromCustomer(customer models.Customer) GetCustomerResponse {
	res.ID = customer.CustomerID
	res.Handle = customer.Handle
	res.Service = customer.Service
	res.Application = customer.Application
	return *res
}

// CreateCustomerRequest is the input for the CustomersV1.CreateCustomer operation.
type CreateCustomerRequest struct{}

// CreateCustomerResponse is the output of the CustomersV1.CreateCustomer operation.
type CreateCustomerResponse struct{}
