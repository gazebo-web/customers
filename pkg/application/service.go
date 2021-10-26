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

// GetIdentity returns the customer or a user of a certain application.
func (s *service) GetIdentity(ctx context.Context, req api.GetIdentityRequest) (api.GetIdentityResponse, error) {
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
