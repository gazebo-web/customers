package fake

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/api"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/client"
)

var _ client.Client = (*Fake)(nil)

// Fake is a fake client.Client implementation.
type Fake struct {
	mock.Mock
}

// GetCustomerByHandle mocks a call to the Customers API.
func (f *Fake) GetCustomerByHandle(ctx context.Context, req api.GetCustomerByHandleRequest) (api.GetCustomerResponse, error) {
	args := f.Called(ctx, req)
	res := args.Get(0).(api.GetCustomerResponse)
	return res, args.Error(1)
}

// GetCustomerByID mocks a call to the Customers API.
func (f *Fake) GetCustomerByID(ctx context.Context, req api.GetCustomerByIDRequest) (api.GetCustomerResponse, error) {
	args := f.Called(ctx, req)
	res := args.Get(0).(api.GetCustomerResponse)
	return res, args.Error(1)
}

// CreateCustomer mocks a call to the Customers API.
func (f *Fake) CreateCustomer(ctx context.Context, req api.CreateCustomerRequest) (api.CreateCustomerResponse, error) {
	args := f.Called(ctx, req)
	res := args.Get(0).(api.CreateCustomerResponse)
	return res, args.Error(1)
}

// NewClient initializes a fake client.Client implementation.
func NewClient() *Fake {
	return &Fake{}
}
