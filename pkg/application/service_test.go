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

type testGetIdentitySuite struct {
	suite.Suite
	DB        *gorm.DB
	Service   Service
	Logger    *log.Logger
	CustomerA models.Customer
	CustomerB models.Customer
	CustomerC models.Customer
}

func TestGetIdentity(t *testing.T) {
	suite.Run(t, new(testGetIdentitySuite))
}

func (s *testGetIdentitySuite) SetupSuite() {
	s.Logger = log.New(os.Stdout, "[TestGetCustomer] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
}

func (s *testGetIdentitySuite) SetupTest() {
	var err error

	var c conf.Config
	s.Require().NoError(c.Parse())

	s.DB, err = persistence.OpenConn(c.Database)
	s.Require().NoError(err)

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

	s.CustomerA, err = persistence.CreateCustomer(s.DB, s.CustomerA)
	s.Require().NoError(err)

	s.CustomerB, err = persistence.CreateCustomer(s.DB, s.CustomerB)
	s.Require().NoError(err)

	s.CustomerC, err = persistence.CreateCustomer(s.DB, s.CustomerC)
	s.Require().NoError(err)
}

func (s *testGetIdentitySuite) TearDownTest() {
	s.Require().NoError(persistence.DropTables(s.DB))
}

func (s *testGetIdentitySuite) TestGetCustomerByHandle() {
	res, err := s.Service.GetCustomerByHandle(context.Background(), api.GetCustomerByHandleRequest{
		Handle:      "user1",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Require().NoError(err)
	s.Assert().Equal(res.ID, s.CustomerA.CustomerID)
}

func (s *testGetIdentitySuite) TestGetCustomerByID() {
	res, err := s.Service.GetCustomerByID(context.Background(), api.GetCustomerByIDRequest{
		ID:          "customer2",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Require().NoError(err)
	s.Assert().Equal(res.Handle, s.CustomerB.Handle)
}

func (s *testGetIdentitySuite) TestGetIdentityMissingIdentity() {
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

func (s *testGetIdentitySuite) TestGetIdentityMissingApplication() {
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

func (s *testGetIdentitySuite) TestGetIdentityMissingService() {
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

func (s *testGetIdentitySuite) TestGetIdentityNotFound() {
	_, err := s.Service.GetCustomerByHandle(context.Background(), api.GetCustomerByHandleRequest{
		Handle:      "user4",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Assert().Error(err)

	_, err = s.Service.GetCustomerByID(context.Background(), api.GetCustomerByIDRequest{
		ID:          "customer5",
		Service:     "stripe",
		Application: "fuel",
	})
	s.Assert().Error(err)
}
