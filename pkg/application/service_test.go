package application

import (
	"context"
	"github.com/stretchr/testify/suite"
	"gitlab.com/ignitionrobotics/billing/customers/internal/conf"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/api"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/models"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/persistence"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

type serviceTestSuite struct {
	suite.Suite
	DB        *gorm.DB
	Service   Service
	Logger    *log.Logger
	CustomerA models.Customer
	CustomerB models.Customer
	CustomerC models.Customer
}

func TestGetIdentity(t *testing.T) {
	suite.Run(t, new(serviceTestSuite))
}

func (s *serviceTestSuite) SetupSuite() {
	s.Logger = log.New(os.Stdout, "[TestGetCustomer] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
	var c conf.Config
	s.Require().NoError(c.Parse())

	var err error
	s.DB, err = persistence.OpenConn(c.Database)
	s.Require().NoError(err)

	s.Require().NoError(persistence.DropTables(s.DB))
}

func (s *serviceTestSuite) SetupTest() {
	s.Require().NoError(persistence.MigrateTables(s.DB))

	s.Service = NewService(s.DB, s.Logger)

	s.CustomerA = models.Customer{
		Handle:      "user1",
		Application: "fuel",
		Service:     "stripe",
		CustomerID:  "customer1",
	}

	s.CustomerB = models.Customer{
		Handle:      "user2",
		Application: "fuel",
		Service:     "stripe",
		CustomerID:  "customer2",
	}

	s.CustomerC = models.Customer{
		Handle:      "user3",
		Application: "fuel",
		Service:     "stripe",
		CustomerID:  "customer3",
	}

	var err error
	s.CustomerA, err = persistence.CreateCustomer(s.DB, s.CustomerA)
	s.Require().NoError(err)

	s.CustomerB, err = persistence.CreateCustomer(s.DB, s.CustomerB)
	s.Require().NoError(err)

	s.CustomerC, err = persistence.CreateCustomer(s.DB, s.CustomerC)
	s.Require().NoError(err)
}

func (s *serviceTestSuite) TearDownTest() {
	s.Require().NoError(persistence.DropTables(s.DB))
}

func (s *serviceTestSuite) TestGetCustomerByHandle() {
	res, err := s.Service.GetCustomerByHandle(context.Background(), api.GetCustomerByHandleRequest{
		Handle:      "user1",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Require().NoError(err)
	s.Assert().Equal(res.ID, s.CustomerA.CustomerID)
}

func (s *serviceTestSuite) TestGetCustomerByID() {
	res, err := s.Service.GetCustomerByID(context.Background(), api.GetCustomerByIDRequest{
		ID:          "customer2",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Require().NoError(err)
	s.Assert().Equal(res.Handle, s.CustomerB.Handle)
}

func (s *serviceTestSuite) TestGetIdentityMissingIdentity() {
	_, err := s.Service.GetCustomerByHandle(context.Background(), api.GetCustomerByHandleRequest{
		Handle:      "",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Assert().Error(err)
	s.Assert().Equal(api.ErrCustomerMissingIdentityValue, err)

	_, err = s.Service.GetCustomerByID(context.Background(), api.GetCustomerByIDRequest{
		ID:          "",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Assert().Error(err)
	s.Assert().Equal(api.ErrCustomerMissingIdentityValue, err)
}

func (s *serviceTestSuite) TestGetIdentityMissingApplication() {
	_, err := s.Service.GetCustomerByHandle(context.Background(), api.GetCustomerByHandleRequest{
		Handle:      "user1",
		Service:     "stripe",
		Application: "",
	})
	s.Assert().Error(err)
	s.Assert().Equal(api.ErrIdentityMissingApplication, err)

	_, err = s.Service.GetCustomerByID(context.Background(), api.GetCustomerByIDRequest{
		ID:          "customer1",
		Service:     "stripe",
		Application: "",
	})
	s.Assert().Error(err)
	s.Assert().Equal(api.ErrIdentityMissingApplication, err)
}

func (s *serviceTestSuite) TestGetIdentityMissingService() {
	_, err := s.Service.GetCustomerByHandle(context.Background(), api.GetCustomerByHandleRequest{
		Handle:      "user1",
		Service:     "",
		Application: "fuel",
	})
	s.Assert().Error(err)
	s.Assert().Equal(api.ErrIdentityMissingService, err)

	_, err = s.Service.GetCustomerByID(context.Background(), api.GetCustomerByIDRequest{
		ID:          "customer1",
		Service:     "",
		Application: "fuel",
	})
	s.Assert().Error(err)
	s.Assert().Equal(api.ErrIdentityMissingService, err)
}

func (s *serviceTestSuite) TestGetIdentityNotFound() {
	_, err := s.Service.GetCustomerByHandle(context.Background(), api.GetCustomerByHandleRequest{
		Handle:      "user4",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Assert().Error(err)
	s.Assert().Equal(api.ErrCustomerNotFound, err)

	_, err = s.Service.GetCustomerByID(context.Background(), api.GetCustomerByIDRequest{
		ID:          "customer5",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Assert().Error(err)
	s.Assert().Equal(api.ErrCustomerNotFound, err)
}

func (s *serviceTestSuite) TestCreateCustomer() {
	_, err := persistence.GetCustomerByUsername(s.DB, "fuel", "stripe", "user4")
	s.Require().Error(err)

	res, err := s.Service.CreateCustomer(context.Background(), api.CreateCustomerRequest{
		ID:          "customer4",
		Handle:      "user4",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Require().NoError(err)

	s.Assert().Equal("customer4", res.ID)
	s.Assert().Equal("user4", res.Handle)
	s.Assert().Equal("stripe", res.Service)
	s.Assert().Equal("fuel", res.Application)

	_, err = persistence.GetCustomerByUsername(s.DB, "fuel", "stripe", "user4")
	s.Require().NoError(err)
}

func (s *serviceTestSuite) TestCreateCustomerMissingAttributes() {
	_, err := s.Service.CreateCustomer(context.Background(), api.CreateCustomerRequest{
		ID:          "",
		Handle:      "user4",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Assert().Error(err)

	_, err = s.Service.CreateCustomer(context.Background(), api.CreateCustomerRequest{
		ID:          "customer4",
		Handle:      "",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Assert().Error(err)

	_, err = s.Service.CreateCustomer(context.Background(), api.CreateCustomerRequest{
		ID:          "customer4",
		Handle:      "user4",
		Service:     "",
		Application: "fuel",
	})
	s.Assert().Error(err)

	_, err = s.Service.CreateCustomer(context.Background(), api.CreateCustomerRequest{
		ID:          "customer4",
		Handle:      "user4",
		Service:     "stripe",
		Application: "",
	})
	s.Assert().Error(err)
}
