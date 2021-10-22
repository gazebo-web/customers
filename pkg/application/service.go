package application

import (
	"context"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/api"
	"io"
	"log"
)

// service contains the business logic to manage customers.
type service struct {
	logger *log.Logger
}

// GetCustomer returns the customer of a certain application.
func (s *service) GetCustomer(ctx context.Context, req api.GetCustomerRequest) (api.GetCustomerResponse, error) {
	panic("implement me")
}

// CreateCustomer creates a new customer for a certain application.
func (s *service) CreateCustomer(ctx context.Context, req api.CreateCustomerRequest) (api.CreateCustomerResponse, error) {
	panic("implement me")
}

// Service holds the methods of the service in charge of managing user service.
type Service interface {
	api.CustomersV1
}

// NewService initializes a new api.CustomersV1 service implementation.
func NewService(logger *log.Logger) Service {
	if logger == nil {
		logger = log.New(io.Discard, "", log.LstdFlags)
	}
	return &service{
		logger: logger,
	}
}
