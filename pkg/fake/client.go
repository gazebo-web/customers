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

// GetIdentity mocks a call to the Customers API.
func (f *Fake) GetIdentity(ctx context.Context, req api.GetIdentityRequest) (api.GetIdentityResponse, error) {
	args := f.Called(ctx, req)
	res := args.Get(0).(api.GetIdentityResponse)
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
