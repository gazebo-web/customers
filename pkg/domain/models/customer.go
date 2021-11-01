package models

import (
	"gorm.io/gorm"
)

// Customer links a user from a certain application to a specific payment system identity.
// 	Example: A Fuel user with a Stripe customer.
type Customer struct {
	gorm.Model

	// Handle is the customer identity in the context of a certain application.
	// E.g. application username, application organization name.
	Handle string

	// Application is where the user is registered.
	Application string

	// Service is the service provider the customer is registered in.
	// E.g. Stripe, PayPal
	Service string

	// ID is the customer identity in the context of an external service.
	CustomerID string
}
