package application

import (
	"context"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/api"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/models"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/persistence"
	"gorm.io/gorm"
	"io"
	"log"
)

// service contains business logic to manage customers.
type service struct {
	logger *log.Logger
	db     *gorm.DB
}

// GetCustomerByHandle returns a models.Customer based on the models.Customer's Handle field.
func (s *service) GetCustomerByHandle(ctx context.Context, req api.GetCustomerByHandleRequest) (api.CustomerResponse, error) {
	if len(req.Application) == 0 {
		return api.CustomerResponse{}, api.ErrIdentityMissingApplication
	}
	if len(req.Service) == 0 {
		return api.CustomerResponse{}, api.ErrIdentityMissingService
	}
	if len(req.Handle) == 0 {
		return api.CustomerResponse{}, api.ErrCustomerMissingIdentityValue
	}

	c, err := persistence.GetCustomerByUsername(s.db, req.Application, req.Service, req.Handle)
	if err != nil {
		if persistence.IsErrorNotFound(err) {
			return api.CustomerResponse{}, api.ErrCustomerNotFound
		}
		return api.CustomerResponse{}, err
	}

	var res api.CustomerResponse
	return res.FromCustomer(c), nil
}

// GetCustomerByID returns a models.Customer based on the models.Customer's CustomerID field.
func (s *service) GetCustomerByID(ctx context.Context, req api.GetCustomerByIDRequest) (api.CustomerResponse, error) {
	if len(req.Application) == 0 {
		return api.CustomerResponse{}, api.ErrIdentityMissingApplication
	}
	if len(req.Service) == 0 {
		return api.CustomerResponse{}, api.ErrIdentityMissingService
	}
	if len(req.ID) == 0 {
		return api.CustomerResponse{}, api.ErrCustomerMissingIdentityValue
	}

	c, err := persistence.GetCustomerByCustomerID(s.db, req.Application, req.Service, req.ID)
	if err != nil {
		if persistence.IsErrorNotFound(err) {
			return api.CustomerResponse{}, api.ErrCustomerNotFound
		}
		return api.CustomerResponse{}, err
	}

	var res api.CustomerResponse
	return res.FromCustomer(c), nil
}

// CreateCustomer creates a new customer for a certain application.
func (s *service) CreateCustomer(ctx context.Context, req api.CreateCustomerRequest) (api.CustomerResponse, error) {
	if len(req.Application) == 0 {
		return api.CustomerResponse{}, api.ErrIdentityMissingApplication
	}
	if len(req.Service) == 0 {
		return api.CustomerResponse{}, api.ErrIdentityMissingService
	}
	if len(req.ID) == 0 || len(req.Handle) == 0 {
		return api.CustomerResponse{}, api.ErrCustomerMissingIdentityValue
	}

	c, err := persistence.CreateCustomer(s.db, models.Customer{
		CustomerID:  req.ID,
		Handle:      req.Handle,
		Application: req.Application,
		Service:     req.Service,
	})
	if err != nil {
		return api.CustomerResponse{}, err
	}

	var res api.CustomerResponse
	return res.FromCustomer(c), nil
}

// Service holds the methods of the service in charge of managing users.
type Service interface {
	api.CustomersV1
}

// NewService initializes a new api.CustomersV1 service implementation.
func NewService(db *gorm.DB, logger *log.Logger) Service {
	if logger == nil {
		logger = log.New(io.Discard, "", log.LstdFlags)
	}
	return &service{
		db:     db,
		logger: logger,
	}
}
