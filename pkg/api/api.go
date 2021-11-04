package api

import (
	"context"
	"errors"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/models"
)

var (
	// ErrCustomerMissingIdentityValue is returned when no user or customer is provided.
	ErrCustomerMissingIdentityValue = errors.New("no identity value provided")

	// ErrIdentityMissingApplication is returned when no application is provided.
	ErrIdentityMissingApplication = errors.New("no application provided")

	// ErrIdentityMissingService is returned when no payment service is provided.
	ErrIdentityMissingService = errors.New("no service provided")

	// ErrCustomerNotFound is returned when the customer has not been found.
	ErrCustomerNotFound = errors.New("customer not found")
)

// CustomersV1 holds the methods that allow managing customers.
type CustomersV1 interface {
	// GetCustomerByHandle returns customer information based on the customer's application handle.
	GetCustomerByHandle(ctx context.Context, req GetCustomerByHandleRequest) (CustomerResponse, error)

	// GetCustomerByID returns customer information based on the customer's external service identity.
	GetCustomerByID(ctx context.Context, req GetCustomerByIDRequest) (CustomerResponse, error)

	// CreateCustomer creates a new customer for a certain application.
	CreateCustomer(ctx context.Context, req CreateCustomerRequest) (CustomerResponse, error)
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

// CustomerResponse is the output from multiple operations that need to return a single customer information.
// CustomersV1.GetCustomerByHandle, CustomersV1.GetCustomerByID and CustomersV1.CreateCustomer are examples of methods that
// return this data structure.
type CustomerResponse struct {
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
func (res *CustomerResponse) FromCustomer(customer models.Customer) CustomerResponse {
	res.ID = customer.CustomerID
	res.Handle = customer.Handle
	res.Service = customer.Service
	res.Application = customer.Application
	return *res
}

// CreateCustomerRequest is the input for the CustomersV1.CreateCustomer operation.
type CreateCustomerRequest struct {
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
