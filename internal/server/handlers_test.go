package server

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"gitlab.com/ignitionrobotics/billing/customers/internal/conf"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/api"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/application"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/models"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/persistence"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type handlersTestSuite struct {
	suite.Suite
	Logger           *log.Logger
	DB               *gorm.DB
	Service          application.Service
	Server           *Server
	Handler          http.HandlerFunc
	CustomerA        models.Customer
	CustomerB        models.Customer
	CustomerC        models.Customer
	ResponseRecorder *httptest.ResponseRecorder
}

func TestHandlers(t *testing.T) {
	suite.Run(t, new(handlersTestSuite))
}

func (s *handlersTestSuite) SetupSuite() {
	s.Logger = log.New(os.Stdout, "[TestHandlers] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)

	var c conf.Config
	s.Require().NoError(c.Parse())

	var err error
	s.DB, err = persistence.OpenConn(c.Database)
	s.Require().NoError(err)

	s.Service = application.NewService(s.DB, s.Logger)

	s.Server = NewServer(Options{
		config:    c,
		customers: s.Service,
		logger:    s.Logger,
	})
}

func (s *handlersTestSuite) SetupTest() {
	s.Require().NoError(persistence.MigrateTables(s.DB))

	var err error
	s.CustomerA = models.Customer{
		Handle:      "test1",
		Application: "fuel",
		Service:     "stripe",
		CustomerID:  "customer_test_a",
	}
	s.CustomerA, err = persistence.CreateCustomer(s.DB, s.CustomerA)
	s.Require().NoError(err)

	s.CustomerB = models.Customer{
		Handle:      "test2",
		Application: "fuel",
		Service:     "stripe",
		CustomerID:  "customer_test_b",
	}
	s.CustomerB, err = persistence.CreateCustomer(s.DB, s.CustomerB)
	s.Require().NoError(err)

	s.CustomerC = models.Customer{
		Handle:      "test3",
		Application: "fuel",
		Service:     "stripe",
		CustomerID:  "customer_test_c",
	}
	s.CustomerC, err = persistence.CreateCustomer(s.DB, s.CustomerC)
	s.Require().NoError(err)

	s.ResponseRecorder = httptest.NewRecorder()
}

func (s *handlersTestSuite) TearDownTest() {
	s.Require().NoError(persistence.DropTables(s.DB))
}

func (s *handlersTestSuite) TestGetCustomerByHandle_OK() {
	s.Handler = s.Server.GetCustomerByHandle

	in := api.GetCustomerByHandleRequest{
		Handle:      "test1",
		Service:     "stripe",
		Application: "fuel",
	}

	request := s.setupRequest(in, http.MethodPost)

	s.Handler.ServeHTTP(s.ResponseRecorder, request)
	s.Require().Equal(http.StatusOK, s.ResponseRecorder.Code)

	var out api.CustomerResponse
	s.parseResponseJSON(&out)
	s.Assert().Equal("customer_test_a", out.ID)
}

func (s *handlersTestSuite) TestGetCustomerByID_OK() {
	s.Handler = s.Server.GetCustomerByID

	in := api.GetCustomerByIDRequest{
		ID:          "customer_test_a",
		Service:     "stripe",
		Application: "fuel",
	}

	request := s.setupRequest(in, http.MethodPost)

	s.Handler.ServeHTTP(s.ResponseRecorder, request)
	s.Require().Equal(http.StatusOK, s.ResponseRecorder.Code)

	var out api.CustomerResponse
	s.parseResponseJSON(&out)
	s.Assert().Equal("test1", out.Handle)
}

func (s *handlersTestSuite) TestCreateCustomer_OK() {
	s.Handler = s.Server.CreateCustomer

	_, err := persistence.GetCustomerByUsername(s.DB, "fuel", "stripe", "test4")
	s.Require().Error(err)

	in := api.CreateCustomerRequest{
		ID:          "customer_test_d",
		Handle:      "test4",
		Service:     "stripe",
		Application: "fuel",
	}

	request := s.setupRequest(in, http.MethodPost)

	s.Handler.ServeHTTP(s.ResponseRecorder, request)
	s.Require().Equal(http.StatusOK, s.ResponseRecorder.Code)

	var out api.CustomerResponse
	s.parseResponseJSON(&out)

	cus, err := persistence.GetCustomerByUsername(s.DB, "fuel", "stripe", "test4")
	s.Require().NoError(err)

	var expected api.CustomerResponse
	expected.FromCustomer(cus)
	s.Assert().Equal(expected, out)
}

func (s *handlersTestSuite) setupRequest(in interface{}, method string) *http.Request {
	body, err := json.Marshal(in)
	s.Require().NoError(err)

	buff := bytes.NewBuffer(body)

	request, err := http.NewRequest(method, "/", buff)
	s.Require().NoError(err)

	return request
}

func (s *handlersTestSuite) parseResponseJSON(out interface{}) {
	body, err := io.ReadAll(s.ResponseRecorder.Body)
	s.Require().NoError(err)
	s.Require().NoError(json.Unmarshal(body, out))
}
