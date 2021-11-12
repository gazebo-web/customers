package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"gitlab.com/ignitionrobotics/billing/customers/internal/conf"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/application"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/domain/persistence"
	"log"
	"net/http"
)

// Setup initializes the conf.Config to run the web server.
func Setup(logger *log.Logger) (conf.Config, error) {
	logger.Println("Parsing config")
	var c conf.Config
	if err := c.Parse(); err != nil {
		logger.Println("Failed to parse config:", err)
		return conf.Config{}, err
	}
	return c, nil
}

// Run runs the web server using the given config.
func Run(config conf.Config, logger *log.Logger) error {
	db, err := persistence.OpenConn(config.Database)
	if err != nil {
		return err
	}

	if err = persistence.MigrateTables(db); err != nil {
		return err
	}

	logger.Println("Initializing Customers service")
	cs := application.NewService(db, logger)

	logger.Println("Initializing HTTP server")
	s := NewServer(Options{
		config:    config,
		customers: cs,
		logger:    logger,
	})

	if err = s.ListenAndServe(); err != nil {
		logger.Println("Error while running HTTP server:", err)
		return err
	}
	return nil
}

// Options contains a set of components to be used when initializing a web server.
type Options struct {
	// config is the config used to set the web server and its components up
	config conf.Config

	// customers contains an application.Service implementation.
	customers application.Service

	// logger is used for logging important messages when the server is running
	logger *log.Logger
}

// Server is an HTTP web server used to expose api.CustomersV1 endpoints. It prepares the input for each
// service operation and returns a serialized JSON response from each operation output.
type Server struct {
	// customers contains an implementation of application.Service
	customers application.Service

	// logger contains the logger used to print debug information.
	logger *log.Logger

	// router is used to route requests in the HTTP server.
	router chi.Router

	// port is the HTTP port used to listen for incoming requests.
	port uint

	// httpServer is used to serve the router with fine-grained control of ListenAndServe and Shutdown operations.
	httpServer http.Server
}

// NewServer initializes a new web server that will serve api.CreditsV1 methods.
func NewServer(opts Options) *Server {
	s := Server{
		customers: opts.customers,
		logger:    opts.logger,
		port:      opts.config.Port,
	}

	s.router = chi.NewRouter()

	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.AllowContentType("application/json"))
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Route("/customers", func(r chi.Router) {
		r.Post("/search/handle", s.GetCustomerByHandle)
		r.Post("/search/id", s.GetCustomerByID)
		r.Post("/", s.CreateCustomer)
	})

	s.httpServer = http.Server{
		Addr:    s.getAddress(),
		Handler: s.router,
	}
	return &s
}

// ListenAndServe starts listening in the port defined on conf.Config. It's in charge of serving the different endpoints.
func (s *Server) ListenAndServe() error {
	s.logger.Println("Listening on", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Shutdown shuts the web server down.
func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

// getAddress returns a valid address (host:port) representation that the server will listen to.
func (s *Server) getAddress() string {
	return fmt.Sprintf(":%d", s.port)
}
