package server

import (
	"encoding/json"
	"fmt"
	"gitlab.com/ignitionrobotics/billing/customers/pkg/api"
	"io"
	"net/http"
)

// GetCustomerByHandle is an HTTP handler to call the api.CustomersV1's GetCustomerByHandle method.
func (s *Server) GetCustomerByHandle(w http.ResponseWriter, r *http.Request) {
	var in api.GetCustomerByHandleRequest
	if err := s.readBodyJSON(w, r, &in); err != nil {
		return
	}
	out, err := s.customers.GetCustomerByHandle(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, &out)
}

// GetCustomerByID is an HTTP handler to call the api.CustomersV1's GetCustomerByID method.
func (s *Server) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	var in api.GetCustomerByIDRequest
	if err := s.readBodyJSON(w, r, &in); err != nil {
		return
	}
	out, err := s.customers.GetCustomerByID(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, &out)
}

// CreateCustomer is an HTTP handler to call the api.CustomersV1's CreateCustomer method.
func (s *Server) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var in api.CreateCustomerRequest
	if err := s.readBodyJSON(w, r, &in); err != nil {
		return
	}
	out, err := s.customers.CreateCustomer(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, &out)
}

func (s *Server) writeResponse(w http.ResponseWriter, out interface{}) {
	body, err := json.Marshal(out)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s - %s", http.StatusText(http.StatusInternalServerError), "Failed to write JSON body"), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s - %s", http.StatusText(http.StatusInternalServerError), "Failed to write body"), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) readBodyJSON(w http.ResponseWriter, r *http.Request, in interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s - %s", http.StatusText(http.StatusBadRequest), "Failed to read body"), http.StatusBadRequest)
		return err
	}

	if err = json.Unmarshal(body, &in); err != nil {
		http.Error(w, fmt.Sprintf("%s - %s", http.StatusText(http.StatusInternalServerError), "Failed to read JSON body"), http.StatusInternalServerError)
		return err
	}

	return nil
}
